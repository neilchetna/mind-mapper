import type { BaseModel } from './base';

export interface Node extends BaseModel {
	isSeedNode: boolean;
	text: string;
	description: string;
	userId: string;
	parentId: string;
	isSuggested: boolean;
}
