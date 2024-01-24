import env from '$env/static/public';
import base from 'axios';

export const axios = base.create({
	baseURL: env.PUBLIC_BACKEND_URL,
	withCredentials: true
});
