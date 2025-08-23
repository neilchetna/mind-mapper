import type { Map } from '$lib/modles';

export async function getMaps(): Promise<Map[]> {
	const mock: Map[] = [
		{
			title: 'School project about climate change',
			id: '1'
		},
		{
			title: 'Starting a coffee business',
			id: '2'
		},
		{
			title: 'Best place to plan a trip to',
			id: '3'
		}
	];

	return new Promise((resolve) => {
		setTimeout(() => resolve(mock), 1000);
	});
}

export async function getMapById(id: string): Promise<Map> {
	const map: Map = {
		title: 'School project about climate change',
		id
	};

	return new Promise((resolve) => {
		setTimeout(() => resolve(map), 1000);
	});
}
