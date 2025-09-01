import { ApiClient } from '$lib/api';
import { type Node } from '$lib/models';

export class NodeSDK extends ApiClient {
	public async getAllNodes(mapId: string) {
		const url = `/charts/${mapId}/nodes`;
		return this.get<Node[]>(url).then((res) => res.data);
	}
}
