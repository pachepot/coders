import React, { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Grid, Button } from '@material-ui/core';
import { SelectForm, SearchInput, Pagination } from '..';
import Table from '../Table';

const head = [
	'채점 번호',
	'아이디',
	'메모리',
	'시간',
	'언어',
	'코드 길이',
	'제출한 시간',
];

const ProblemRank = ({ submissions, handleSubmissions }) => {
	const { id } = useParams();
	useEffect(() => {
		handleSubmissions(id);
	}, [id, submissions]);
	return (
		<Grid className="problemrank-container">
			<Grid className="problemrank-input" container>
				<Grid className="problemrank-search">
					<SearchInput label="아이디로 검색" />
				</Grid>
				<SelectForm
					defaultValue="모든 언어"
					values={['모든 언어', 'C++', 'Java', 'Python']}
				/>
				<Button size="small">↓ 메모리로 정렬</Button>
				<Button size="small">↓ 시간으로 정렬</Button>
			</Grid>
			<Table head={head} rows={submissions} />
			<Grid className="problemrank-pagination">
				<Pagination />
			</Grid>
		</Grid>
	);
};

export default ProblemRank;
