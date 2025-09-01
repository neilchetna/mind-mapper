<script lang="ts">
	import { goto } from '$app/navigation';
	import { createForm } from '$lib';
	import { FormDialog } from '$lib/components/common/form-dialog';
	import { MapCard, Topbar } from '$lib/components/maps-list-page';
	import { Field } from '$lib/components/ui/field';
	import { Input } from '$lib/components/ui/input';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Textarea } from '$lib/components/ui/textarea';
	import { MapsManager } from '$lib/manager';
	import type { Map } from '$lib/models';
	import { createMapSchema } from '$lib/schema/create-map';
	import { MapLoading } from '$lib/utils/types/loading';
	import { useClerkContext } from 'svelte-clerk';

	const m = new MapsManager();
	const ctx = useClerkContext();
	const loading = $derived(!!m.loading[MapLoading.FetchingMaps]);
	const creatingMap = $derived(!!m.loading[MapLoading.CreatingMap]);
	let openCreateMapDialog = $state<boolean>(false);

	const init = async () => {
		const token = await ctx.session?.getToken();
		if (typeof token === 'string') {
			m.loadSDK(token);
		}

		await m.loadMaps();
	};

	const { values, errors, handleSubmit, reset } = createForm(
		createMapSchema,
		{ seedNode: '', explorationDetails: '' },
		handleCreateNewMap
	);

	function redirectPath(map: Map): string {
		return `/map/${map.id}`;
	}

	async function handleCreateNewMap() {
		const map = await m.createMap(values);
		if (map) {
			goto(`/map/${map.id}`);
		}
	}

	$effect(() => {
		init();
	});
</script>

<Topbar handleNewMapClick={() => (openCreateMapDialog = true)} loading={creatingMap} />
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
	<FormDialog
		onSumbit={handleSubmit}
		submitButton={{ title: 'Create Map' }}
		onOpenChange={(value) => {
			openCreateMapDialog = value;
			reset();
		}}
		open={openCreateMapDialog}
		description="Enter the first builidng block of this mind map, it could be the agenda or central topic explained briefly"
		title="Add first node"
		><Field error={errors.seedNode} label="Enter topic">
			<Input
				aria-invalid={!!errors.seedNode}
				type="text"
				placeholder="e.g. History of America"
				bind:value={values.seedNode}
			/>
		</Field>
		<Field
			error={errors.explorationDetails}
			hint="How would you like to explore this topic"
			label="Exploration Details"
		>
			<Textarea
				aria-invalid={!!errors.explorationDetails}
				placeholder="The economical impact of America on the world"
				bind:value={values.explorationDetails}
			></Textarea>
		</Field>
	</FormDialog>
{/if}
