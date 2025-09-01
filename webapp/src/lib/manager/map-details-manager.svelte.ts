import type { Node, Map } from '$lib/models';
import { MapsSDK } from '$lib/sdk/map';
import { NodeSDK } from '$lib/sdk/node';
import { MapLoading, NodeLoading, type ResourceLoading } from '$lib/utils/types/loading';
import { toast } from 'svelte-sonner';

export class MapDetailsManager {
	#mapsSDK: MapsSDK;
	#nodeSDK: NodeSDK;
	map = $state<Map>();
	nodes = $state<Node[]>();
	loading = $state<ResourceLoading>({});

	loadSDK(token: string) {
		this.#mapsSDK.newAuthToken(token);
		this.#nodeSDK.newAuthToken(token);
	}

	constructor() {
		this.#mapsSDK = new MapsSDK();
		this.#nodeSDK = new NodeSDK();
	}

	async loadingNodes(mapId: string) {
		try {
			this.loading[NodeLoading.FetchingNodes] = true;
			this.nodes = await this.#nodeSDK.getAllNodes(mapId);
		} catch (error) {
			console.error(error);

			toast.error('Something went wrong while loading this map', {
				action: {
					label: 'Retry',
					onClick: () => this.loadingNodes(mapId)
				}
			});
		} finally {
			this.loading[NodeLoading.FetchingNodes] = false;
		}
	}

	async loadMapDetails(id: string) {
		try {
			this.loading[MapLoading.FetchingMapById] = true;
			this.map = await this.#mapsSDK.getMapById(id);
		} catch (error) {
			console.error(error);

			toast.error('Something went wrong while loading this map', {
				action: {
					label: 'Retry',
					onClick: () => this.loadMapDetails(id)
				}
			});
		} finally {
			this.loading[MapLoading.FetchingMapById] = false;
		}
	}
}
