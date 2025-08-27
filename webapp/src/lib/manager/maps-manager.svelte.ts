import type { Map } from '$lib/modles';
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
		} catch (err) {
			console.error(err);
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

	async createEmptyMap() {
		try {
			this.loading[MapLoading.CreatingMap] = true;
			const map = await this.#sdk.createNewMap({});
			this.maps = [...this.maps, map];
			return map;
		} catch (error) {
			console.error(error);
			toast.error('Something went wrong while creating maps');
		} finally {
			this.loading[MapLoading.CreatingMap] = false;
		}
	}
}
