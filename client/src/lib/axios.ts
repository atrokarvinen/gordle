import { PUBLIC_BACKEND_URL } from '$env/static/public';
import base from 'axios';

export const axios = base.create({
	baseURL: PUBLIC_BACKEND_URL,
	withCredentials: true
});
