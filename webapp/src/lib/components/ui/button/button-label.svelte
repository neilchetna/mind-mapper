<script lang="ts">
	import * as Tooltip from '../tooltip';
	import type { ButtonProps } from './button.svelte';
	import Button from './button.svelte';

	type Props = ButtonProps & { label: string };
	const { label, children, ...props }: Props = $props();
</script>

<Tooltip.Provider>
	<Tooltip.Root>
		<Tooltip.Trigger>
			<Button {...props}>
				{@render children?.()}
			</Button>
		</Tooltip.Trigger>
		<Tooltip.Content>
			<p class="text-xs">{label}</p>
		</Tooltip.Content>
	</Tooltip.Root>
</Tooltip.Provider>
