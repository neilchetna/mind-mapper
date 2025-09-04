<script lang="ts">
	import { Handle, Position, type NodeProps } from '@xyflow/svelte';
	import NodeRoot from './node-root.svelte';

	let { data, selected }: NodeProps = $props();
</script>

<NodeRoot {selected}>
	<div class="bg-card max-w-56 rounded-sm border px-4 py-2 text-center shadow-md">
		<Handle class="opacity-0" isConnectable={false} type="source" position={Position.Right} />
		<p>{data.text}</p>
		<Handle class="opacity-0" isConnectable={false} type="source" position={Position.Left} />
	</div>
</NodeRoot>
