import { ApiClient } from '$lib/api';
import type { CreateMap } from '$lib/dto';
import type { Map } from '$lib/models';

export class MapsSDK extends ApiClient {
	public async getMaps() {
		return this.get<Map[]>('/charts').then((res) => res.data);
	}

	public async getMapById(id: string) {
		return this.get<Map>(`/charts/${id}`).then((res) => res.data);
	}

	public async createNewMap(map: CreateMap) {
		return this.post<CreateMap, Map>('/charts', map).then((res) => res.data);
	}
}
