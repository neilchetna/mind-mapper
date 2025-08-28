<script lang="ts">
	import { IconPlus } from '@tabler/icons-svelte';
	import Button from '../ui/button/button.svelte';

	interface Props {
		handleNewMapClick: () => void;
		loading: boolean;
	}

	const { handleNewMapClick, loading }: Props = $props();
</script>

<div class="dark:border-input/30 bg-accent flex w-full justify-end border px-4 py-2">
	<Button disabled={loading} onclick={handleNewMapClick}>New Map <IconPlus /></Button>
</div>
