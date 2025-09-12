<script lang="ts">
	import type { Edge, Node } from '$lib/models';
	import { computeLayout } from '$lib/utils';
	import {
		Background,
		SvelteFlow,
		useOnSelectionChange,
		type Edge as FlowEdge,
		type Node as FlowNode
	} from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import { type Snippet } from 'svelte';
	import EdgeComponent from './edge.svelte';
	import ExploredNode from './explored-node.svelte';
	import SeedNode from './seed-node.svelte';

	interface Props {
		nodeData: Node[];
		edgeData: Edge[];
		createNodeForm: Snippet;
		selectedNode: string;
	}
	let { nodeData, edgeData, createNodeForm, selectedNode = $bindable('') }: Props = $props();

	let { nodes, edges } = $derived<{ nodes: FlowNode[]; edges: FlowEdge[] }>(
		computeLayout(nodeData, edgeData)
	);

	const patternColor = 'var(--color-input)';
	const bgColor = 'var(--color-secondary)';

	useOnSelectionChange(({ nodes: selectedNodes }) => {
		selectedNode = selectedNodes[0]?.id || '';
	});
</script>

<SvelteFlow
	nodeTypes={{ seedNode: SeedNode, exploredNode: ExploredNode }}
	edgeTypes={{ custom: EdgeComponent }}
	bind:nodes
	bind:edges
	fitView
>
	{@render createNodeForm()}
	<Background {bgColor} {patternColor} />
</SvelteFlow>
