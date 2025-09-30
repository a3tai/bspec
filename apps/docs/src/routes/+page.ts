import type { PageLoad } from './$types';
import homeData from '$lib/data/home.json';

export const load: PageLoad = async ({ parent }) => {
	const parentData = await parent();

	return {
		...homeData,
		specification: parentData.specification
	};
};