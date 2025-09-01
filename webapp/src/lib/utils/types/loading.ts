export enum MapLoading {
	FetchingMaps = 'fetching-maps',
	FetchingMapById = 'fetching-map-by-id',
	CreatingMap = 'creating-map'
}

export enum NodeLoading {
	FetchingNodes = 'fetching-nodes'
}

type Resource = `${MapLoading}` | `${NodeLoading}`;

export type ResourceLoading = Partial<Record<Resource, boolean>>;
