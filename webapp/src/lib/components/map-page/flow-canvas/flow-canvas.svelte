<script lang="ts">
	import type { Node } from '$lib/models';
	import { node2FlowNode } from '$lib/utils/nodeTransformer';
	import { Background, SvelteFlow, type Node as FlowNode } from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import SeedNode from './seed-node.svelte';
	import { onDestroy, type Snippet } from 'svelte';
	import { flowState } from './flow-manager.svelte';

	interface Props {
		nodeData: Node[];
		createNodeForm: Snippet<[{ open: boolean; onClose(value: boolean): void }]>;
	}
	const { nodeData, createNodeForm }: Props = $props();

	let nodes = $state.raw<FlowNode[]>(nodeData.map(node2FlowNode));
	const patternColor = 'var(--color-input)';
	const bgColor = 'var(--color-secondary)';

	onDestroy(() => {
		flowState.destroy();
	});
</script>

<SvelteFlow nodeTypes={{ seedNode: SeedNode }} bind:nodes fitView>
	{@render createNodeForm({
		open: flowState.openCreateNewNode,
		onClose: (value) => (flowState.openCreateNewNode = value)
	})}
	<Background {bgColor} {patternColor} />
</SvelteFlow>
