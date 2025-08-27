<script lang="ts">
	import { IconArrowLeft } from '@tabler/icons-svelte';
	import Button from '../ui/button/button.svelte';

	const { title }: { title: string } = $props();
</script>

<div class="bg-accent dark:border-input/30 flex items-center gap-4 border px-4 py-2">
	<Button size="icon" variant="outline" href="/map"><IconArrowLeft /></Button>
	<p>{title}</p>
</div>
