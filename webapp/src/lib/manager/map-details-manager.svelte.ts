import type { CreateNode, Map, Node } from '$lib/models';
import type { CreateNodeSchema } from '$lib/schema';
import { MapsSDK } from '$lib/sdk/map';
import { NodeSDK } from '$lib/sdk/node';
import { MapLoading, NodeLoading, type ResourceLoading } from '$lib/utils/types/loading';
import { toast } from 'svelte-sonner';

export class MapDetailsManager {
	#mapsSDK: MapsSDK;
	#nodeSDK: NodeSDK;
	#mapId;
	map = $state<Map>();
	nodes = $state<Node[]>();
	loading = $state<ResourceLoading>({});

	loadSDK(token: string) {
		this.#mapsSDK.newAuthToken(token);
		this.#nodeSDK.newAuthToken(token);
	}

	constructor(mapId: string) {
		this.#mapId = mapId;
		this.#mapsSDK = new MapsSDK();
		this.#nodeSDK = new NodeSDK();
	}

	#formDataToCreateNode(nodeData: CreateNodeSchema): CreateNode {
		return {
			text: nodeData.nodeText,
			isSeedNode: false
		};
	}

	async createNewNode(nodeData: CreateNodeSchema) {
		try {
			this.loading[NodeLoading.CreatingNode] = true;
			const node = await this.#nodeSDK.createNode(
				this.#mapId,
				this.#formDataToCreateNode(nodeData)
			);

			this.nodes?.push(node);
		} catch {
			toast.error('Something went wrong while creating node');
		} finally {
			this.loading[NodeLoading.CreatingNode] = false;
		}
	}

	async loadingNodes() {
		try {
			this.loading[NodeLoading.FetchingNodes] = true;
			this.nodes = await this.#nodeSDK.getAllNodes(this.#mapId);
		} catch {
			toast.error('Something went wrong while loading this map', {
				action: {
					label: 'Retry',
					onClick: () => this.loadingNodes()
				}
			});
		} finally {
			this.loading[NodeLoading.FetchingNodes] = false;
		}
	}

	async loadMapDetails() {
		try {
			this.loading[MapLoading.FetchingMapById] = true;
			this.map = await this.#mapsSDK.getMapById(this.#mapId);
		} catch {
			toast.error('Something went wrong while loading this map', {
				action: {
					label: 'Retry',
					onClick: () => this.loadMapDetails()
				}
			});
		} finally {
			this.loading[MapLoading.FetchingMapById] = false;
		}
	}
}
