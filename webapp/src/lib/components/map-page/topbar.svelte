<script lang="ts">
	import { IconArrowLeft } from '@tabler/icons-svelte';
	import { ButtonLabel } from '../ui/button';

	const { title }: { title: string } = $props();
</script>

<div class="flex items-center gap-4 border bg-accent px-4 py-2 dark:border-input/30">
	<ButtonLabel label="Back" size="icon" variant="outline" href="/map"><IconArrowLeft /></ButtonLabel
	>
	<p>{title}</p>
</div>
