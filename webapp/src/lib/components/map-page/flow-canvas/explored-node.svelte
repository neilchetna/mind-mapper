<script lang="ts">
	import { Handle, Position, type NodeProps } from '@xyflow/svelte';
	import NodeRoot from './node-root.svelte';

	const { data, selected }: NodeProps = $props();
</script>

<NodeRoot {selected}>
	<Handle class="opacity-0" isConnectable={false} type="target" position={Position.Right} />
	<p class="text-xs">{data.text}</p>
	<Handle class="opacity-0" isConnectable={false} type="source" position={Position.Left} />
</NodeRoot>
