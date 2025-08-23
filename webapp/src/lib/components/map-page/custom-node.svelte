<script lang="ts">
	import { Handle, Position, type NodeProps } from '@xyflow/svelte';

	let { data }: NodeProps = $props();
</script>

<div class="bg-card rounded-sm border px-4 py-2 shadow-md">
	<p>{data.label}</p>
</div>
