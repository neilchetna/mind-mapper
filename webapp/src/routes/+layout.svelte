<script lang="ts">
	import { PUBLIC_CLERK_PUBLISHABLE_KEY } from '$env/static/public';
	import favicon from '$lib/assets/favicon.svg';
	import { ModeWatcher } from 'mode-watcher';
	import { ClerkProvider } from 'svelte-clerk/client';
	import '../app.css';
	import { Toaster } from '$lib/components/ui/sonner';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<ModeWatcher />
<ClerkProvider publishableKey={PUBLIC_CLERK_PUBLISHABLE_KEY}>
	<div class="bg-background text-foreground h-dvh w-full">
		{@render children?.()}
		<Toaster />
	</div>
</ClerkProvider>
