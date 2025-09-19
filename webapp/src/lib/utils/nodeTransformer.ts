import type { Node } from '$lib/models';
import { type Node as FlowNode } from '@xyflow/svelte';

export const node2FlowNode = (node: Node): FlowNode => {
	return {
		data: { text: node.text, description: node.description, isSuggested: node.isSuggested },
		id: node.id,
		position: { x: 0, y: 0 },
		type: node.isSeedNode ? 'seedNode' : 'exploredNode',
		draggable: false,
		selectable: true
	};
};
