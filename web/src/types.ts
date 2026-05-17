/**
 * TypeScript types for Capability Stack documents.
 * These mirror the Go types in the capstack package.
 */

export interface CapabilityStack {
  $schema?: string;
  metadata: Metadata;
  layers: Layer[];
  categories?: Category[];
  capabilities: Capability[];
  foundational?: Capability[];
  prismIntegration?: PRISMIntegration;
}

export interface Metadata {
  name: string;
  version: string;
  title?: string;
  description?: string;
  domain?: string;
  createdAt?: string;
  updatedAt?: string;
  authors?: string[];
}

export interface Layer {
  id: string;
  name: string;
  description?: string;
  order?: number;
  phase?: string;
  nistCsfFunction?: string;
}

export interface Category {
  id: string;
  name: string;
  description?: string;
  color?: string;
}

export interface Capability {
  id: string;
  name: string;
  fullName?: string;
  description?: string;
  layerId: string;
  categoryId?: string;
  status?: CapabilityStatus;
  priority?: Priority;
  targetDate?: string;
  implementedAt?: string;
  owner?: string;
  tooling?: Tool[];
  dependencies?: string[];
  enables?: string[];
  tags?: string[];
  frameworkMappings?: FrameworkMapping[];
  prismRef?: PRISMRef;
}

export type CapabilityStatus =
  | 'planned'
  | 'in-progress'
  | 'implemented'
  | 'operational'
  | 'deprecated';

export type Priority = 'critical' | 'high' | 'medium' | 'low';

export interface Tool {
  name: string;
  vendor?: string;
  type?: 'commercial' | 'open-source' | 'internal' | 'managed-service';
  url?: string;
  status?: 'evaluating' | 'piloting' | 'deployed' | 'deprecated';
}

export interface FrameworkMapping {
  framework: string;
  controls: string[];
}

export interface PRISMRef {
  domainId?: string;
  sliIds?: string[];
  levelCriteria?: LevelCriteria;
}

export interface LevelCriteria {
  M1?: string;
  M2?: string;
  M3?: string;
  M4?: string;
  M5?: string;
}

export interface PRISMIntegration {
  modelRef?: string;
  stateRef?: string;
  planRef?: string;
  defaultDomain?: string;
}

// Status colors matching the Go render package
export const STATUS_COLORS: Record<CapabilityStatus, { bg: string; text: string }> = {
  'operational': { bg: '#10b981', text: '#ffffff' },
  'implemented': { bg: '#3b82f6', text: '#ffffff' },
  'in-progress': { bg: '#f59e0b', text: '#000000' },
  'planned': { bg: '#9ca3af', text: '#000000' },
  'deprecated': { bg: '#ef4444', text: '#ffffff' },
};

export const STATUS_LABELS: Record<CapabilityStatus, string> = {
  'operational': 'Operational',
  'implemented': 'Implemented',
  'in-progress': 'In Progress',
  'planned': 'Planned',
  'deprecated': 'Deprecated',
};
