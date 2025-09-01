import type z from 'zod';

type Errors<T> = Partial<Record<keyof T, string>>;

export function createForm<T extends z.ZodRawShape>(
	schmea: z.ZodObject<T>,
	initial: z.infer<typeof schmea>,
	onSumbit: (values: z.infer<typeof schmea>) => Promise<void> | void
) {
	let values = $state({ ...initial });
	let errors = $state<Errors<typeof initial>>({});

	async function handleSubmit(e: Event) {
		e.preventDefault();

		const result = schmea.safeParse(values);

		if (!result.success) {
			result.error.issues.forEach((issue) => {
				const field = issue.path[0] as keyof typeof initial;
				errors[field] = issue.message;
			});

			return;
		}

		errors = {};
		await onSumbit(result.data);
	}

	function setField<K extends keyof typeof initial>(key: K, value: unknown) {
		values = { ...values, [key]: value };
	}

	function reset(): void {
		(Object.keys(initial) as Array<keyof typeof initial>).forEach((key) => {
			errors[key] = '';
			values[key] = initial[key];
		});
	}

	return { values, errors, handleSubmit, setField, reset };
}
