import { z } from 'zod/v4';

export const createNodeSchema = z.object({
	nodeText: z.string().min(1).max(35, { error: 'Nodes text should not be longer than 35' })
});

export type CreateNodeSchema = z.infer<typeof createNodeSchema>;
