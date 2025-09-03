<script lang="ts">
	import { ButtonLabel } from '$lib/components/ui/button';
	import { IconPlus } from '@tabler/icons-svelte';

	type Props = {
		onNewNodeClick(): void;
	};

	const { onNewNodeClick }: Props = $props();
</script>

<div class="flex justify-center gap-0">
	<ButtonLabel onclick={onNewNodeClick} variant="outline" label="Add new node" size="icon"
		><IconPlus /></ButtonLabel
	>
</div>
