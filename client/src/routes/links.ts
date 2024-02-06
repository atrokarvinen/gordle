import { base } from '$app/paths';
import { DEFAULT_LIMIT, DEFAULT_PAGE } from './statistics/game-history/defaults';

export type LinkType = {
	href: string;
	iconClass: string;
	datatestid: string;
	name: string;
};

export const links: LinkType[] = [
	{
		name: 'history',
		href: `${base}/statistics/game-history?page=${DEFAULT_PAGE}&limit=${DEFAULT_LIMIT}`,
		iconClass: 'fas fa-history',
		datatestid: 'game-history'
	},
	{ name: 'settings', href: `${base}/settings`, iconClass: 'fas fa-cog', datatestid: 'settings' },
	{ name: 'info', href: `${base}/info`, iconClass: 'fas fa-info', datatestid: 'info' }
];
