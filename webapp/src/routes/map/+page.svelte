<script lang="ts">
	import { goto } from '$app/navigation';
	import MapCard from '$lib/components/maps-list-page/map-card.svelte';
	import Topbar from '$lib/components/maps-list-page/topbar.svelte';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { MapsManager } from '$lib/manager/maps-manager.svelte';
	import type { Map } from '$lib/modles';
	import { MapLoading } from '$lib/utils/types/loading';
	import { useClerkContext } from 'svelte-clerk';

	const m = new MapsManager();
	const ctx = useClerkContext();
	const loading = $derived(!!m.loading[MapLoading.FetchingMaps]);
	const creatingMap = $derived(!!m.loading[MapLoading.CreatingMap]);

	const init = async () => {
		const token = await ctx.session?.getToken();
		if (typeof token === 'string') {
			m.loadSDK(token);
		}

		await m.loadMaps();
	};

	function redirectPath(map: Map): string {
		return `/map/${map.id}`;
	}

	async function handleCreateNewMap() {
		const map = await m.createEmptyMap();
		if (map) {
			goto(`/map/${map.id}`);
		}
	}

	$effect(() => {
		init();
	});
</script>

<Topbar {handleCreateNewMap} loading={creatingMap} />
{#if loading}
	<div class="flex gap-4 p-8">
		<Skeleton class="h-12 w-56 rounded" />
		<Skeleton class="h-12 w-56 rounded" />
		<Skeleton class="h-12 w-56 rounded" />
	</div>
{:else if m.maps}
	<div class="flex gap-4 p-8">
		{#each m.maps as map (map.id)}
			<MapCard {map} redirectPath={redirectPath(map)} />
		{/each}
	</div>
{/if}
