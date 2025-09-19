import type { Edge, Node } from '$lib/models';
import D from '@dagrejs/dagre';
import type { Edge as FlowEdge, Node as FlowNode } from '@xyflow/svelte';
import { ERRORS, LAYOUT } from '.';
import { node2FlowNode } from './nodeTransformer';
import splitNodes from './split-nodes-arr';

export function computeLayout(
	nodes: Node[],
	edges: Edge[]
): {
	nodes: FlowNode[];
	edges: FlowEdge[];
} {
	const seed = nodes.find((node) => node.isSeedNode);
	if (!seed) throw new Error(ERRORS.NO_SEED_NODE);

	const flowSeedNode = node2FlowNode(seed);
	const { left, right } = splitNodes(nodes, edges);

	const leftArr = [flowSeedNode, ...left.nodes.map(node2FlowNode)];
	const rightArr = [flowSeedNode, ...right.nodes.map(node2FlowNode)];
	const leftNodes = setGraphLayout({
		seedId: seed.id,
		nodes: leftArr,
		side: 'left',
		edges: left.edges
	});
	const rightNodes = setGraphLayout({
		seedId: seed.id,
		nodes: rightArr,
		side: 'right',
		edges: right.edges
	});

	const suggestedNodes = new Set(nodes.filter((node) => node.isSuggested).map((node) => node.id));
	const flowEdges = buildEdges(edges, seed.id, left.edges, suggestedNodes);

	return {
		nodes: [...leftNodes, ...rightNodes, flowSeedNode],
		edges: flowEdges
	};
}

function buildEdges(
	edges: Edge[],
	seedId: string,
	leftEdges: Edge[],
	suggestedNodes: Set<string>
): FlowEdge[] {
	const leftSet = new Set(leftEdges.map((edge) => edge.id));
	return edges.map((edge) => ({
		...edge,
		data: {
			isSeedEdge: edge.source === seedId,
			side: leftSet.has(edge.id) ? 'left' : 'right',
			isDashed: suggestedNodes.has(edge.target)
		},
		selectable: false,
		type: 'custom'
	}));
}

function setGraphLayout({
	seedId,
	nodes,
	edges,
	side
}: {
	seedId: string;
	side: 'left' | 'right';
	nodes: FlowNode[];
	edges: FlowEdge[];
}): FlowNode[] {
	const g = new D.graphlib.Graph();
	g.setDefaultEdgeLabel(() => ({}));
	g.setGraph({ rankdir: 'LR' });

	nodes.forEach((node) => {
		g.setNode(node.id, { width: node.width, height: node.height });
	});
	edges.forEach((e) => {
		g.setEdge(e.source, e.target);
	});

	D.layout(g);

	const seed = g.node(seedId);
	return nodes
		.map((node) => {
			const { rank = 0, y: layoutY } = g.node(node.id);
			// Divide the graph along the Y-axis based on the side
			const MULT = side === 'left' ? -1 : 1;
			// Extra padding between seed node and first leaf nodes
			const RANK_MULT = rank === 1 ? LAYOUT.GRAPH_FIRST_RANK_SPACING : 1;
			const x = rank * LAYOUT.GRAPH_X_SPACING * MULT * RANK_MULT;
			// a. Dagre nodes are returned in the first quadrant. The algoritm also accounts for this
			// by shifting the first parent node up along the Y-axis. Since we will be placing this first
			// node on (0, 0). So we have to shift all the leaf nodes with the height of root node

			// b. For some reason there is a padding of 15 units in the Y-cordinates of Dagre layout
			// This is why we remove it with GRAPH_Y_NORMALIZATION
			const y = layoutY - (seed.y - LAYOUT.GRAPH_Y_NORMALIZATION);
			return {
				...node,
				data: { ...node.data, side },
				position: { x, y }
			};
		})
		.filter((node) => node.id != seedId);
}
