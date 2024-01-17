import { useEffect, useState } from 'react';
import axios from 'axios';
import { getTicketingApiUrl } from '@/utilities/common';

const MyPage = () => {
	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await axios.get(`${getTicketingApiUrl()}/api/v1/ticket/user`, {
					withCredentials: true,
				});
				console.log(response.data);
			} catch (error) {
				console.error('Axios error:', error);
			}
		};

		fetchData(); // Call the async function to fetch data when the component mounts
	}, []); // Empty dependency array means this effect runs once on mount

	return (
		<div>
			<h1>API Response</h1>
		</div>
	);
};

export default MyPage;
