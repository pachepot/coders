package model

import (
	"errors"

	"gorm.io/gorm"
)

type Problem struct {
	ID           int          `json:"id" example:"1" format:"int64" gorm:"autoIncrement"`
	Title        string       `json:"title" example:"Problem title"`
	Class        string       `json:"class" example:"Problem class"`
	Description  string       `json:"description" example:"Problem description"`
	TimeLimit    int          `json:"timeLimit" example:"1000" format:"int64"`
	MemoryLimit  int          `json:"memoryLimit" example:"128" format:"int64"`
	ShortCircuit bool         `json:"shortCircuit" example:"false"`
	MemberID     int          `json:"memberID" example:"1" format:"int64"`
	Member       Member       `gorm:"ForeignKey:MemberID;"`
	Submissions  []Submission `gorm:"ForeignKey:ProblemID";constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type EditProblem struct {
	Title        string `json:"title" example:"Problem title"`
	Class        string `json:"class" example:"Problem class"`
	Description  string `json:"description" example:"Problem description"`
	TimeLimit    int    `json:"timeLimit" example:"1000" format:"int64"`
	MemoryLimit  int    `json:"memoryLimit" example:"1024" format:"int64"`
	ShortCircuit bool   `json:"shortCircuit" example:"false"`
	MemberID     int    `json:"memberID" example:"1" format:"int64"`
}

type PrintProblem struct {
	Title string
	Class string
}

var (
	ErrTitleInvalid       = errors.New("title is empty")
	ErrDescInvalid        = errors.New("description is empty")
	ErrMemberIDInvalid    = errors.New("member ID is empty")
	ErrTimeLimitInvalid   = errors.New("time limit is empty")
	ErrMemoryLimitInvalid = errors.New("memory limit is empty")
)

func (p EditProblem) ProblemValidation() error {
	switch {
	case len(p.Title) == 0:
		return ErrTitleInvalid
	case len(p.Description) == 0:
		return ErrDescInvalid
	case p.MemberID == 0:
		return ErrMemberIDInvalid
	case p.TimeLimit == 0:
		return ErrTimeLimitInvalid
	case p.MemoryLimit == 0:
		return ErrMemoryLimitInvalid
	default:
		return nil
	}
}

func ProblemAll(db *gorm.DB, num int, page int, mid int, search string, sort string) ([]PrintProblem, error) {
	var problem []PrintProblem
	var err error

	db = db.Model(&Problem{})

	if mid != 0 {
		db = db.Where("member_id = ?", mid)
	}
	if search != "" {
		db = db.Where("title LIKE ? OR class LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if sort != "" {
		db = db.Order(sort)
	}

	err = db.Limit(num).Offset(num*(page-1)).Select("title", "class").Find(&problem).Error

	return problem, err
}

func ProblemOne(db *gorm.DB, id int) (Problem, error) {
	var problem Problem
	err := db.Where("id = ?", id).First(&problem).Error
	return problem, err
}

func ProblemInsert(db *gorm.DB, problem Problem) (Problem, error) {
	err := db.Create(&problem).Error
	return problem, err
}

func ProblemUpdate(db *gorm.DB, problem Problem) (Problem, error) {
	err := db.Model(&Problem{}).Where("id = ?", problem.ID).Update("title", problem.Title).Error
	if err != nil {
		return problem, err
	}
	err = db.Model(&Problem{}).Where("id = ?", problem.ID).Update("class", problem.Class).Error
	if err != nil {
		return problem, err
	}
	err = db.Model(&Problem{}).Where("id = ?", problem.ID).Update("description", problem.Description).Error
	if err != nil {
		return problem, err
	}
	err = db.Model(&Problem{}).Where("id = ?", problem.ID).Update("time_limit", problem.TimeLimit).Error
	if err != nil {
		return problem, err
	}
	err = db.Model(&Problem{}).Where("id = ?", problem.ID).Update("memory_limit", problem.MemoryLimit).Error
	if err != nil {
		return problem, err
	}
	err = db.Model(&Problem{}).Where("id = ?", problem.ID).Update("short_circuit", problem.ShortCircuit).Error
	if err != nil {
		return problem, err
	}
	return problem, err
}

func ProblemDelete(db *gorm.DB, id int) error {
	var problem Problem
	err := db.Where("id=?", id).First(&problem).Error
	if err != nil {
		return err
	}
	err = db.Delete(&problem).Error
	return err
}
