import type { PageLoad } from './$types';
import domainsData from '$lib/data/domains.json';

export const load: PageLoad = async ({ parent }) => {
	const { specification } = await parent();

	return {
		...domainsData,
		specification
	};
};