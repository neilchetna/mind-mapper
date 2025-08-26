import { ApiClient } from '$lib/api';
import type { Map } from '$lib/modles';

export class MapsSDK extends ApiClient {
	public async getMaps() {
		return this.get<Map[]>('/charts').then((res) => res.data);
	}

	public async getMapById(id: string) {
		return this.get<Map>(`/charts/${id}`).then((res) => res.data);
	}

	public async createNewMap(map: Partial<Map>) {
		return this.post<Partial<Map>, Map>('/charts', map).then((res) => res.data);
	}
}
