<script lang="ts">
	import type { Edge, Node } from '$lib/models';
	import { node2FlowNode } from '$lib/utils/nodeTransformer';
	import {
		Background,
		SvelteFlow,
		useOnSelectionChange,
		type Edge as FlowEdge,
		type Node as FlowNode
	} from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import { type Snippet } from 'svelte';
	import ExploredNode from './explored-node.svelte';
	import SeedNode from './seed-node.svelte';

	interface Props {
		nodeData: Node[];
		edgeData: Edge[];
		createNodeForm: Snippet;
		selectedNode: string;
	}
	let { nodeData, edgeData, createNodeForm, selectedNode = $bindable('') }: Props = $props();

	let nodes = $derived<FlowNode[]>(nodeData.map(node2FlowNode));
	let edges = $derived<FlowEdge[]>(edgeData);

	const patternColor = 'var(--color-input)';
	const bgColor = 'var(--color-secondary)';

	useOnSelectionChange(({ nodes: selectedNodes }) => {
		selectedNode = selectedNodes[0]?.id || '';
	});
</script>

<SvelteFlow
	nodeTypes={{ seedNode: SeedNode, exploredNode: ExploredNode }}
	bind:nodes
	bind:edges
	fitView
>
	{@render createNodeForm()}
	<Background {bgColor} {patternColor} />
</SvelteFlow>
