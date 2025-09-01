<script lang="ts">
	import { Background, SvelteFlow, type Node as FlowNode } from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import SeedNode from './seed-node.svelte';
	import type { Node } from '$lib/models';
	import { node2FlowNode } from '$lib/utils/nodeTransformer';
	const bgColor = 'var(--color-secondary)';
	const patternColor = 'var(--color-input)';
	interface Props {
		nodeData: Node[];
	}

	const { nodeData }: Props = $props();

	let nodes = $state.raw<FlowNode[]>(nodeData.map(node2FlowNode));
</script>

<SvelteFlow nodeTypes={{ seedNode: SeedNode }} bind:nodes fitView>
	<Background {bgColor} {patternColor} />
</SvelteFlow>
