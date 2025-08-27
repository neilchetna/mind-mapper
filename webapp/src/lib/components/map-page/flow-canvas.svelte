<script lang="ts">
	import { Background, SvelteFlow, ViewportPortal, type Node } from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import CustomNode from './custom-node.svelte';
	import type { Snippet } from 'svelte';
	const bgColor = 'var(--color-background)';
	const patternColor = 'var(--color-input)';

	let nodes = $state.raw<Node[]>([
		{
			id: '1',
			position: { x: 0, y: 0 },
			data: { label: 'Hello' },
			type: 'textNode'
		},
		{
			id: '2',
			position: { x: 100, y: 100 },
			data: { label: 'World' },
			type: 'textNode'
		}
	]);

	let { portal }: { portal: Snippet } = $props();
</script>

<SvelteFlow nodeTypes={{ textNode: CustomNode }} bind:nodes>
	<Background {bgColor} {patternColor} />
	<ViewportPortal target="front">
		{@render portal()}
	</ViewportPortal>
</SvelteFlow>
