import type { BaseModel } from './base';

export interface Node extends BaseModel {
	isSeedNode: boolean;
	text: string;
	description: string;
	userId: string;
}

export type CreateNode = Pick<Node, 'text' | 'isSeedNode'>;
