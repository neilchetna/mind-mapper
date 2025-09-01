<script lang="ts">
	import { type NodeProps } from '@xyflow/svelte';

	let { data }: NodeProps = $props();
</script>

<div class="max-w-56 rounded-sm border bg-card px-4 py-2 text-center shadow-md">
	<p>{data.text}</p>
</div>
