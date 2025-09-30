import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { findDocument } from '$lib/navigation';

export const load: PageLoad = async ({ params, parent }) => {
	const { specification } = await parent();

	if (!specification) {
		throw error(404, 'Specification not found');
	}

	const domain = params.domain;
	const doctype = params.doctype;

	// Find the document
	const document = findDocument(specification, domain, doctype);

	if (!document) {
		throw error(404, `Document ${doctype} not found in domain ${domain}`);
	}

	// Find the domain info
	const domainInfo = specification.domains.find(d => d.slug === domain);

	if (!domainInfo) {
		throw error(404, `Domain ${domain} not found`);
	}

	return {
		document,
		domain: domainInfo,
		specification
	};
};