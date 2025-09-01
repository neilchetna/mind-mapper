import { z } from 'zod/v4';

export const createMapSchema = z.object({
	seedNode: z
		.string()
		.min(4, { error: 'The fist node should be at least 4 characters long' })
		.max(22, { error: 'The node cannot be longer than 22 characters' }),
	explorationDetails: z.string().max(1000)
});

export type CreateMapSchema = z.infer<typeof createMapSchema>;
