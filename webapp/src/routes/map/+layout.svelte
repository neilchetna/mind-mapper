<script lang="ts">
	import { useClerkContext } from 'svelte-clerk';

	const { children } = $props();

	const ctx = useClerkContext();
	const authLoaded = $derived(ctx.isLoaded);
</script>

{#if authLoaded}
	{@render children?.()}
{/if}
