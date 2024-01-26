import { PUBLIC_BACKEND_URL } from '$env/static/public';
import base, { AxiosError, isAxiosError } from 'axios';

export const axios = base.create({
	baseURL: PUBLIC_BACKEND_URL,
	withCredentials: true
});

export const getApiErrorMessage = (error: any) => {
	if (isAxiosError(error)) {
		const axiosError: AxiosError = error;
		const data: any = axiosError.response?.data;
		const message = data.message;
		if (typeof message === 'string') {
			return message;
		}
		return axiosError.message;
	}
	if (typeof error?.messsage === 'string') {
		return error.message;
	}
	return error;
};
