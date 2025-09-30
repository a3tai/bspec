import type { PageLoad } from './$types';
import domainsData from '$lib/data/domains.json';

export const load: PageLoad = async ({ params, parent }) => {
	const { specification } = await parent();
	const domain = domainsData.domains.find(d => d.slug === params.domain);

	if (!domain) {
		return {
			status: 404,
			error: new Error('Domain not found')
		};
	}

	return {
		domain,
		specification
	};
};