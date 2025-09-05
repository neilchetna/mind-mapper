import type { CreateNode } from '$lib/dto';
import type { Edge, Map, Node } from '$lib/models';
import type { CreateNodeSchema } from '$lib/schema';
import { EdgeSDK } from '$lib/sdk/edge';
import { MapsSDK } from '$lib/sdk/map';
import { NodeSDK } from '$lib/sdk/node';
import {
	EdgeLoading,
	MapLoading,
	NodeLoading,
	type ResourceLoading
} from '$lib/utils/types/loading';
import { toast } from 'svelte-sonner';

export class MapDetailsManager {
	#mapsSDK: MapsSDK;
	#nodeSDK: NodeSDK;
	#edgeSDK: EdgeSDK;
	#mapId;
	map = $state<Map>();
	nodes = $state<Node[]>([]);
	edges = $state<Edge[]>([]);
	loading = $state<ResourceLoading>({});
	selectedNode = $state<string>('');

	loadSDK(token: string) {
		this.#mapsSDK.newAuthToken(token);
		this.#nodeSDK.newAuthToken(token);
		this.#edgeSDK.newAuthToken(token);
	}

	constructor(mapId: string) {
		this.#mapId = mapId;
		this.#mapsSDK = new MapsSDK();
		this.#nodeSDK = new NodeSDK();
		this.#edgeSDK = new EdgeSDK();
	}

	setSelectedNode(nodeId: string): void {
		this.selectedNode = nodeId;
	}

	#formDataToCreateNode(nodeData: CreateNodeSchema, parentId: string): CreateNode {
		return {
			parentId,
			text: nodeData.nodeText,
			isSeedNode: false
		};
	}

	async loadingEdges() {
		try {
			this.loading[EdgeLoading.FetchingEdges] = true;
			this.edges = await this.#edgeSDK.getEdges(this.#mapId);
		} catch {
			toast.error('Something went wrong while loading this map', {
				action: {
					label: 'Retry',
					onClick: () => this.loadingEdges()
				}
			});
		} finally {
			this.loading[EdgeLoading.FetchingEdges] = false;
		}
	}

	async createNewNode(nodeData: CreateNodeSchema) {
		try {
			this.loading[NodeLoading.CreatingNode] = true;
			const res = await this.#nodeSDK.createNode(
				this.#mapId,
				this.#formDataToCreateNode(nodeData, this.selectedNode)
			);

			this.nodes = [...this.nodes, res.node];
			this.edges = [...this.edges, res.edge];
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
