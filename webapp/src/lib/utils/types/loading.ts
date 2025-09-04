export enum MapLoading {
	FetchingMaps = 'fetching-maps',
	FetchingMapById = 'fetching-map-by-id',
	CreatingMap = 'creating-map'
}

export enum NodeLoading {
	FetchingNodes = 'fetching-nodes',
	CreatingNode = 'creating-node'
}

export enum EdgeLoading {
	FetchingEdges = 'fetching-edges'
}

type Resource = `${MapLoading}` | `${NodeLoading}` | `${EdgeLoading}`;

export type ResourceLoading = Partial<Record<Resource, boolean>>;
