<script lang="ts">
	import type { Map } from '$lib/models';
	const { map, redirectPath }: { map: Map; redirectPath: string } = $props();
</script>

<a href={redirectPath} class="bg-card dark:border-input/30 w-56 rounded-sm border p-4">
	<p class="truncate text-xs">{map.title}</p>
</a>
