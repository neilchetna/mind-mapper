<script lang="ts">
	import type { Map } from '$lib/models';
	const { map, redirectPath }: { map: Map; redirectPath: string } = $props();
</script>

<a href={redirectPath} class="w-56 rounded-sm border bg-card p-4 dark:border-input/30">
	<p class="truncate text-xs">{map.title}</p>
</a>
