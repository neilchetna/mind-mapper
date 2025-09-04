<script lang="ts">
	import type { Snippet } from 'svelte';
	import Actions from './actions.svelte';
	import { flowState } from './flow-manager.svelte';

	type Props = { children: Snippet; selected: boolean };
	const { selected, children }: Props = $props();
	const handleNewNodeClicked = () => {
		flowState.openCreateNewNode = true;
	};
</script>

<div class={[selected && 'outline-primary outline outline-offset-4']}>
	{#if selected}
		<div class="absolute inset-x-0 -top-12">
			<Actions onNewNodeClick={handleNewNodeClicked} />
		</div>
	{/if}
	{@render children?.()}
</div>
