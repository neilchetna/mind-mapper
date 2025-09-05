import { ApiClient } from '$lib/api';
import type { Edge } from '$lib/models';

export class EdgeSDK extends ApiClient {
	async getEdges(mapId: string) {
		const url = `/charts/${mapId}/edges`;
		return this.get<Edge[]>(url).then((res) => res.data);
	}
}
