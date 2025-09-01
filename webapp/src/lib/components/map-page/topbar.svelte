<script lang="ts">
	import { IconArrowLeft } from '@tabler/icons-svelte';
	import Button from '../ui/button/button.svelte';

	const { title }: { title: string } = $props();
</script>

<div class="flex items-center gap-4 border bg-accent px-4 py-2 dark:border-input/30">
	<Button size="icon" variant="outline" href="/map"><IconArrowLeft /></Button>
	<p>{title}</p>
</div>
