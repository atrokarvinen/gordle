import base from 'axios';

export const axios = base.create({
	baseURL: 'http://localhost:9000',
	withCredentials: true
});
