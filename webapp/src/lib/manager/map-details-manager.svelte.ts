import type { Map } from '$lib/modles';
import { MapsSDK } from '$lib/sdk/map';
import { MapLoading, type ResourceLoading } from '$lib/utils/types/loading';
import { toast } from 'svelte-sonner';

export class MapDetailsManager {
	#sdk: MapsSDK;
	map = $state<Map>();
	loading = $state<ResourceLoading>({});

	loadSDK(token: string) {
		this.#sdk.newAuthToken(token);
	}

	constructor() {
		this.#sdk = new MapsSDK();
	}

	async loadMapDetails(id: string) {
		try {
			this.loading[MapLoading.FetchingMapById] = true;
			this.map = await this.#sdk.getMapById(id);
		} catch (error) {
			console.error(error);

			toast.error('Something went wrong while creating maps', {
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
