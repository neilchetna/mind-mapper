<script lang="ts">
	import type { DialogRootProps } from 'bits-ui';
	import { Button, type ButtonProps } from '../ui/button';
	import * as Dialog from '../ui/dialog';

	type SubmitButton = {
		title: string;
		config?: ButtonProps;
	};

	interface Props extends DialogRootProps {
		open: boolean;
		submitButton: SubmitButton;
		onSumbit: (e: Event) => void | Promise<void>;
		title?: string;
		description?: string;
	}

	let {
		open = $bindable(false),
		title,
		description,
		submitButton,
		children,
		onSumbit,
		...props
	}: Props = $props();
	const { title: buttonTitle, config: buttonProps } = submitButton;
</script>

<Dialog.Root bind:open {...props}>
	<Dialog.Trigger />
	<Dialog.Content>
		<Dialog.Header>
			{#if title}
				<Dialog.Title>{title}</Dialog.Title>
			{/if}
			{#if description}
				<Dialog.Description>{description}</Dialog.Description>
			{/if}
		</Dialog.Header>
		<form onsubmit={onSumbit}>
			{@render children?.()}
			<Dialog.Footer><Button type="submit" {...buttonProps}>{buttonTitle}</Button></Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
