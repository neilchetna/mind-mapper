import type { BaseModel } from './base';
import type { Node } from './node';

export interface Map extends BaseModel {
	id: string;
	title: string;
	createdAt: string;
	updatedAt: string;
	userId: string;
	nodes: Array<Node>;
	explorationDetails: string;
}
