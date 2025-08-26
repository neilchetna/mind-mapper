export enum MapLoading {
	FetchingMaps = 'fetching-maps',
	FetchingMapById = 'fetching-map-by-id',
	CreatingMap = 'creating-map'
}

type Resource = `${MapLoading}`;

export type ResourceLoading = Partial<Record<Resource, boolean>>;
