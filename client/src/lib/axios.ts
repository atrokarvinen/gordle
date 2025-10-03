import { PUBLIC_BACKEND_URL } from '$env/static/public';
import base, { isAxiosError } from 'axios';
import { englishKeys } from './translations/en';

export const axios = base.create({
	baseURL: PUBLIC_BACKEND_URL,
	withCredentials: true
});

type ApiErrorMessage = {
	message?: string;
};

export const getApiErrorMessage = (error: unknown) => {
	const apiError = error as ApiErrorMessage;
	if (isAxiosError(error)) {
		const axiosError = error;
		const responseData = axiosError.response?.data;
		const message = responseData.message;
		if (typeof message === 'string') {
			const translationKey = messageToTranslationKey(message);
			return { message: translationKey, data: responseData.data };
		}
		return { message: axiosError.message, data: undefined };
	}
	if (typeof apiError?.message === 'string') {
		return { message: apiError.message, data: undefined };
	}
	return { message: error as string, data: undefined };
};

const messageToTranslationKey = (message: string): string => {
	const key = message.toLowerCase().replace(/[^a-z]/g, '_');
	if (Object.keys(englishKeys.en.translation).includes(key)) return key;
	return message;
};
