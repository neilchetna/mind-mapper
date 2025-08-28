import { z } from 'zod/v3';

export const createMapSchema = z.object({
	seedNode: z.string().max(22),
	explorationSeed: z.string().max(1000).optional()
});

export type CreateMapSchema = typeof createMapSchema;
