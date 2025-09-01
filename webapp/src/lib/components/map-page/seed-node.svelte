<script lang="ts">
	import { type NodeProps } from '@xyflow/svelte';

	let { data }: NodeProps = $props();
</script>

<div class="bg-card max-w-56 rounded-sm border px-4 py-2 text-center shadow-md">
	<p>{data.text}</p>
</div>
