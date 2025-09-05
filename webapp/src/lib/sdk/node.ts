import { ApiClient } from '$lib/api';
import type { CreateNode, CreateNodeResponse } from '$lib/dto';
import { type Node } from '$lib/models';

export class NodeSDK extends ApiClient {
	public async getAllNodes(mapId: string) {
		const url = `/charts/${mapId}/nodes`;
		return this.get<Node[]>(url).then((res) => res.data);
	}

	public async createNode(mapId: string, node: CreateNode) {
		const url = `/charts/${mapId}/nodes`;
		return this.post<CreateNode, CreateNodeResponse>(url, node).then((res) => res.data);
	}
}
