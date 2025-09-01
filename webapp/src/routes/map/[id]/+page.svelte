<script lang="ts">
	import FlowCanvas from '$lib/components/map-page/flow-canvas.svelte';
	import Topbar from '$lib/components/map-page/topbar.svelte';
	import { MapDetailsManager } from '$lib/manager/map-details-manager.svelte.js';
	import { MapLoading, NodeLoading } from '$lib/utils/types/loading.js';
	import { useClerkContext } from 'svelte-clerk';

	const { params } = $props();
	const m = new MapDetailsManager();
	const ctx = useClerkContext();
	const loading = $derived(
		!!m.loading[MapLoading.FetchingMapById] || !!m.loading[NodeLoading.FetchingNodes]
	);

	const init = async () => {
		const token = await ctx.session?.getToken();
		if (typeof token === 'string') {
			m.loadSDK(token);
		}

		m.loadMapDetails(params.id);
		m.loadingNodes(params.id);
	};

	$effect(() => {
		init();
	});
</script>

{#if !loading && m.map && m.nodes}
	<Topbar title={m.map.title} />
	<FlowCanvas nodeData={m.nodes} />
{/if}
