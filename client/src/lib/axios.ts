import { PUBLIC_BACKEND_URL } from '$env/static/public';
import base, { isAxiosError } from 'axios';
import { englishKeys } from './translations/en';

export const axios = base.create({
	baseURL: PUBLIC_BACKEND_URL,
	withCredentials: true
});

export const getApiErrorMessage = (error: any): { message: string; data: any } => {
	if (isAxiosError(error)) {
		const axiosError = error;
		const responseData: any = axiosError.response?.data;
		const message = responseData.message;
		if (typeof message === 'string') {
			const translationKey = messageToTranslationKey(message);
			return { message: translationKey, data: responseData.data };
		}
		return { message: axiosError.message, data: undefined };
	}
	if (typeof error?.messsage === 'string') {
		return { message: error.message, data: undefined };
	}
	return { message: error, data: undefined };
};

const messageToTranslationKey = (message: string): string => {
	const key = message.toLowerCase().replace(/[^a-z]/g, '_');
	if (Object.keys(englishKeys.en.translation).includes(key)) return key;
	return message;
};
