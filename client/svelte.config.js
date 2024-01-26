import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [vitePreprocess({})],
	kit: {
		adapter: adapter(),
		prerender: {
			entries: ['/game/-1']
		},
		paths: {
			base: process.env.NODE_ENV === 'production' ? '/gordle' : ''
		}
	}
};

export default config;
