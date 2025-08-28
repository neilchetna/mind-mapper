<script lang="ts">
	import { type CreateMapSchema } from '$lib/schema/create-map';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';
	import Textarea from '../ui/textarea/textarea.svelte';

	interface Props {
		handleSubmit: (formData: CreateMapSchema) => void;
	}
	const { handleSubmit }: Props = $props();
</script>

<form
	onsubmit={(e) => {
		e.preventDefault();
	}}
>
	<div class="mb-6">
		<Label>Enter topic</Label>
		<Input class="my-2" type="text" placeholder="e.g. History of America" />
		<p class="text-muted-foreground text-xs"></p>
	</div>
	<div class="mb-6">
		<Label>Exploration Details</Label>
		<Textarea class="my-2" placeholder="The economical impact of America on the world"></Textarea>
		<p class="text-muted-foreground text-xs">How would you like to explore this topic</p>
	</div>
</form>
