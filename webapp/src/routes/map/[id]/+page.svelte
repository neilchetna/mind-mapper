<script lang="ts">
	import { createForm } from '$lib';
	import { FormDialog } from '$lib/components/common/form-dialog/index.js';
	import { FlowCanvas, TopBar } from '$lib/components/map-page';
	import Field from '$lib/components/ui/field/field.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import { MapDetailsManager } from '$lib/manager';
	import { createNodeSchema, type CreateNodeSchema } from '$lib/schema/create-node.js';
	import { MapLoading, NodeLoading } from '$lib/utils/types/loading.js';
	import { useClerkContext } from 'svelte-clerk';

	type CreateNodeFormProps = {
		open: boolean;
		onClose(value: boolean): void;
	};

	const { params } = $props();
	const m = new MapDetailsManager(params.id);
	const ctx = useClerkContext();
	const loading = $derived(
		!!m.loading[MapLoading.FetchingMapById] || !!m.loading[NodeLoading.FetchingNodes]
	);

	const handleCreateNode = async (formData: CreateNodeSchema) => {
		return m.createNewNode(formData);
	};

	const { values, errors, handleSubmit, reset } = createForm(
		createNodeSchema,
		{ nodeText: '' },
		handleCreateNode
	);

	const init = async () => {
		const token = await ctx.session?.getToken();
		if (typeof token === 'string') {
			m.loadSDK(token);
		}

		m.loadMapDetails();
		m.loadingNodes();
	};

	$effect(() => {
		init();
	});
</script>

{#if !loading && m.map && m.nodes}
	<TopBar title={m.map.title} />
	<FlowCanvas {createNodeForm} nodeData={m.nodes} />
{/if}

{#snippet createNodeForm({ open, onClose }: CreateNodeFormProps)}
	<FormDialog
		onOpenChange={(value: boolean) => {
			onClose(value);
			reset();
		}}
		title="Create new node"
		description="Enter the text that is relevent to the previous node"
		onSumbit={handleSubmit}
		submitButton={{ title: 'Create Node' }}
		{open}
	>
		<Field label="Enter new node's text" error={errors.nodeText}>
			<Input type="text" aria-invalid={!!errors.nodeText} bind:value={values.nodeText} />
		</Field>
	</FormDialog>
{/snippet}
