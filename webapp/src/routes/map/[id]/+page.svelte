<script lang="ts">
	import FlowCanvas from '$lib/components/map-page/flow-canvas.svelte';
	import Topbar from '$lib/components/map-page/topbar.svelte';
	import { mapDetailsManager } from '$lib/manager/map-details-manager.svelte';
	let loadingPreview = $state<boolean>(false);
	let showError = $state<boolean>(false);

	const getMap = async () => {
		try {
			loadingPreview = true;
			await mapDetailsManager.init('');
		} catch (error) {
			showError = true;
		} finally {
			loadingPreview = false;
		}
	};

	$effect(() => {
		getMap();
	});
</script>

{#if !loadingPreview && mapDetailsManager.map}
	<Topbar title={mapDetailsManager.map?.title} />
	<FlowCanvas />
{/if}
