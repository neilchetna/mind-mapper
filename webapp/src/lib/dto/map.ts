import type { Map } from '$lib/models';
import type { CreateNode } from './node';

export type CreateMap = Partial<Pick<Map, 'explorationDetails'>> & { nodes: [CreateNode] };
