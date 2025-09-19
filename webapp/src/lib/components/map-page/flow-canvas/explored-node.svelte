<script lang="ts">
	import { Handle, Position, type NodeProps } from '@xyflow/svelte';
	import NodeRoot from './node-root.svelte';

	const { data, selected }: NodeProps = $props();
	const sourcePosition = data.side === 'left' ? Position.Left : Position.Right;
	const targetPosition = data.side === 'left' ? Position.Right : Position.Left;
</script>

<NodeRoot isSuggested={!!data.isSuggested} {selected}>
	<Handle class="opacity-0" isConnectable={false} type="target" position={targetPosition} />
	<p class="text-xs">{data.text}</p>
	<Handle class="opacity-0" isConnectable={false} type="source" position={sourcePosition} />
</NodeRoot>
