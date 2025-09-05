<script lang="ts">
	import type { Edge, Node } from '$lib/models';
	import { node2FlowNode } from '$lib/utils/nodeTransformer';
	import {
		Background,
		SvelteFlow,
		useOnSelectionChange,
		type Node as FlowNode,
		type Edge as FlowEdge
	} from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import SeedNode from './seed-node.svelte';
	import { onDestroy, type Snippet } from 'svelte';
	import { flowState } from './flow-manager.svelte';
	import ExploredNode from './explored-node.svelte';

	interface Props {
		nodeData: Node[];
		edgeData: Edge[];
		createNodeForm: Snippet<[{ open: boolean; onClose(value: boolean): void }]>;
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

	onDestroy(() => {
		flowState.destroy();
	});
</script>

<SvelteFlow
	nodeTypes={{ seedNode: SeedNode, exploredNode: ExploredNode }}
	bind:nodes
	bind:edges
	fitView
>
	{@render createNodeForm({
		open: flowState.openCreateNewNode,
		onClose: (value) => (flowState.openCreateNewNode = value)
	})}
	<Background {bgColor} {patternColor} />
</SvelteFlow>
