<script lang="ts">
	import type { Snippet } from 'svelte';
	import * as Dialog from '../ui/dialog/index';
	import type { DialogRootProps } from 'bits-ui';

	interface Props extends DialogRootProps {
		open: boolean;
		title: string;
		description: string;
		body: Snippet;
		footer: Snippet;
	}

	let { open = $bindable(false), title, description, body, footer, ...props }: Props = $props();
</script>

<Dialog.Root bind:open {...props}>
	<Dialog.Trigger />
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title level={1}>
				{title}
			</Dialog.Title>
			<Dialog.Description>
				{description}
			</Dialog.Description>
		</Dialog.Header>
		{@render body()}
		<Dialog.Footer>
			{@render footer()}
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
