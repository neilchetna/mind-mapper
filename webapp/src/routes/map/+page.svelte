<script lang="ts">
	import MapCard from '$lib/components/maps-list-page/map-card.svelte';
	import Topbar from '$lib/components/maps-list-page/topbar.svelte';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { mapsManager } from '$lib/manager/maps-manager.svelte';
	import type { Map } from '$lib/modles';
	let loadingPreview = $state(false);
	let mockError = $state(false);

	const handleError = () => {
		mockError = true;
	};

	const getMapsData = async () => {
		try {
			loadingPreview = true;
			await mapsManager.init();
		} catch (error) {
			handleError();
		} finally {
			loadingPreview = false;
		}
	};

	$effect(() => {
		void getMapsData();
	});

	function redirectPath(map: Map): string {
		return `/map/${map.id}`;
	}
</script>

<Topbar />
{#if loadingPreview}
	<div class="flex gap-4 p-8">
		<Skeleton class="h-12 w-56 rounded" />
		<Skeleton class="h-12 w-56 rounded" />
		<Skeleton class="h-12 w-56 rounded" />
	</div>
{:else}
	<div class="flex gap-4 p-8">
		{#each mapsManager.maps as map (map.id)}
			<MapCard {map} redirectPath={redirectPath(map)} />
		{/each}
	</div>
{/if}
