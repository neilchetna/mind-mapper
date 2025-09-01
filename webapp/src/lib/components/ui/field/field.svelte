<script lang="ts">
	import type { Snippet } from 'svelte';
	import Label from '../label/label.svelte';

	interface Props {
		label: string;
		children: Snippet;
		error?: string;
		hint?: string;
	}

	const { label, children, hint, error }: Props = $props();
</script>

<div class="mb-8">
	<Label class="mb-2">{label}</Label>

	{@render children()}

	{#if hint && !error}
		<p class="mt-2 text-xs text-muted-foreground">{hint}</p>
	{/if}

	{#if error}
		<p class="mt-2 text-xs text-destructive">{error}</p>
	{/if}
</div>
