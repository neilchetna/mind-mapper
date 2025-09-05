import type { BaseModel } from './base';

export type Edge = { source: string; target: string; mapId: string } & BaseModel;
