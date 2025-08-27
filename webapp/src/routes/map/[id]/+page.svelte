<script lang="ts">
	import VanillaDialog from '$lib/components/common/vanilla-dialog.svelte';
	import FlowCanvas from '$lib/components/map-page/flow-canvas.svelte';
	import Topbar from '$lib/components/map-page/topbar.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Input } from '$lib/components/ui/input/index.js';
	import { MapDetailsManager } from '$lib/manager/map-details-manager.svelte.js';
	import { MapLoading } from '$lib/utils/types/loading.js';
	import { onMount } from 'svelte';
	import { useClerkContext } from 'svelte-clerk';

	const { params } = $props();
	const m = new MapDetailsManager();
	const ctx = useClerkContext();
	const loading = $derived(m.loading[MapLoading.FetchingMapById]);
	let seedDialogOpen = $state(false);

	const init = async () => {
		const token = await ctx.session?.getToken();
		if (typeof token === 'string') {
			m.loadSDK(token);
		}

		m.loadMapDetails(params.id);
	};

	onMount(() => {
		seedDialogOpen = true;
	});

	$effect(() => {
		init();
	});
</script>

{#if !loading && m.map}
	<Topbar title={m.map.title} />
	<FlowCanvas {portal} />
{/if}

{#snippet portal()}
	<VanillaDialog
		{body}
		{footer}
		open={seedDialogOpen}
		description="Enter the first builidng block of this mind map, it could be the agenda or central topic explained briefly"
		title="Add first node"
	/>
	{#snippet body()}
		<p>Enter topic</p>
		<Input type="text" placeholder="e.g. History of America" />
	{/snippet}

	{#snippet footer()}
		<Button>Start map</Button>
	{/snippet}
{/snippet}
