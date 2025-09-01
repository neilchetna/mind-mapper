import type { BaseModel } from './base';
import type { CreateNode, Node } from './node';

export interface Map extends BaseModel {
	id: string;
	title: string;
	createdAt: string;
	updatedAt: string;
	userId: string;
	nodes: Array<Node>;
	explorationDetails: string;
}

export type CreateMap = Partial<Pick<Map, 'explorationDetails'>> & { nodes: [CreateNode] };
