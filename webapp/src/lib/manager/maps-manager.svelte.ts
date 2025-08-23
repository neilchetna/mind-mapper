import type { Map } from '$lib/modles';
import { getMaps } from '$lib/sdk/maps.mock';

class MapsManager {
	maps: Map[] = $state<Map[]>([]);

	async init() {
		this.maps = await getMaps();
	}
}

export const mapsManager = new MapsManager();
