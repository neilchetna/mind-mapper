import type { Map } from '$lib/modles';
import { getMapById } from '$lib/sdk/maps.mock';

class MapDetailsManager {
	map = $state<Map>();

	async init(id: string) {
		this.map = await getMapById(id);
	}
}

export const mapDetailsManager = new MapDetailsManager();
