import type { Edge, Node } from '$lib/models';
import { ERRORS } from './types';

type Result = {
	left: { nodes: Node[]; edges: Edge[] };
	right: { nodes: Node[]; edges: Edge[] };
};
export default function splitNodes(nodes: Node[], edges: Edge[]): Result {
	const seedNode = nodes.find((node) => node.isSeedNode);
	if (!seedNode) throw new Error(ERRORS.NO_SEED_NODE);

	const nodesMap = new Map(nodes.map((node) => [node.id, node]));
	const sideMap = new Map<string, 'left' | 'right'>();

	const children = nodes.filter((node) => node.parentId === seedNode.id);
	children.forEach((child, i) => {
		sideMap.set(child.id, i % 2 === 0 ? 'left' : 'right');
	});

	function assignSide(node: Node): 'left' | 'right' {
		if (sideMap.has(node.id)) return sideMap.get(node.id) as 'left' | 'right';

		const parent = nodesMap.get(node.parentId);
		if (!parent) throw new Error(ERRORS.NO_PARENT_NODE(node.id));

		const side = assignSide(parent);
		sideMap.set(node.id, side);

		return side;
	}

	const leftNodes: Node[] = [];
	const rightNodes: Node[] = [];

	nodes.forEach((node) => {
		if (node.id === seedNode.id) return;
		const side = assignSide(node);
		(side === 'left' ? leftNodes : rightNodes).push(node);
	});

	const leftEdges: Edge[] = [];
	const rightEdges: Edge[] = [];

	edges.forEach((edge) => {
		const side = sideMap.get(edge.target);
		(side === 'left' ? leftEdges : rightEdges).push(edge);
	});

	return {
		left: { nodes: leftNodes, edges: leftEdges },
		right: { nodes: rightNodes, edges: rightEdges }
	};
}
