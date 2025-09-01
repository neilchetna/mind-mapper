import type { CreateMap, CreateNode, Map } from '$lib/models';
import type { CreateMapSchema } from '$lib/schema/create-map';
import { MapsSDK } from '$lib/sdk/map';
import { MapLoading, type ResourceLoading } from '$lib/utils/types/loading';
import { toast } from 'svelte-sonner';

export class MapsManager {
	#sdk: MapsSDK;
	maps: Map[] = $state<Map[]>([]);
	loading: ResourceLoading = $state<ResourceLoading>({});

	constructor() {
		this.#sdk = new MapsSDK();
	}

	loadSDK(token: string) {
		this.#sdk.newAuthToken(token);
	}

	async loadMaps() {
		try {
			this.loading[MapLoading.FetchingMaps] = true;
			this.maps = await this.#sdk.getMaps();
		} catch {
			toast.error('Something went wrong while getting your maps', {
				action: {
					label: 'Retry',
					onClick: () => this.loadMaps()
				}
			});
		} finally {
			this.loading[MapLoading.FetchingMaps] = false;
		}
	}

	#formData2CreateNode(seedNodeData: CreateMapSchema): CreateMap {
		const node: CreateNode = { isSeedNode: true, text: seedNodeData.seedNode };
		const res: CreateMap = { nodes: [node] };

		if (seedNodeData.explorationDetails) {
			res.explorationDetails = seedNodeData.explorationDetails;
		}

		return res;
	}

	async createMap(seedNodeData: CreateMapSchema) {
		try {
			const mapData = this.#formData2CreateNode(seedNodeData);
			this.loading[MapLoading.CreatingMap] = true;
			const map = await this.#sdk.createNewMap(mapData);
			this.maps = [...this.maps, map];
			return map;
		} catch {
			toast.error('Something went wrong while creating maps');
		} finally {
			this.loading[MapLoading.CreatingMap] = false;
		}
	}
}
