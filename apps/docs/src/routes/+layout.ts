import type { LayoutLoad } from './$types';
import homeData from '$lib/data/home.json';
import type { BSpecSpecification } from '$lib/bspec-types';

export const load: LayoutLoad = () => {
	// Transform home data into specification format
	const specification: BSpecSpecification = {
		version: homeData.version,
		domains: homeData.domains.map(domain => ({
			name: domain.name,
			slug: domain.slug,
			documents: domain.documentTypes?.map(dt => ({
				type: dt.type,
				name: dt.name,
				slug: dt.slug,
				path: `/domains/${domain.slug}/${dt.slug}`,
				frontmatter: {},
				content: ''
			})) || []
		}))
	};

	return {
		specification
	};
};