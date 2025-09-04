import type { Edge, Node } from '$lib/models';

export type CreateNode = Pick<Node, 'text' | 'isSeedNode'> & { parentId?: string };

export type CreateNodeResponse = {
	node: Node;
	edge: Edge;
};
