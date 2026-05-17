import { LitElement, html, css, nothing } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import type { CapabilityStack, Capability, Layer, Category, CapabilityStatus } from './types.js';
import { STATUS_COLORS, STATUS_LABELS } from './types.js';

/**
 * A web component for rendering Capability Stack diagrams.
 *
 * @example
 * ```html
 * <capability-stack src="stack.json"></capability-stack>
 * ```
 *
 * @example
 * ```html
 * <capability-stack
 *   src="stack.json"
 *   theme="dark"
 *   show-legend
 *   interactive>
 * </capability-stack>
 * ```
 */
@customElement('capability-stack')
export class CapabilityStackElement extends LitElement {
  static override styles = css`
    :host {
      --cs-font-family: system-ui, -apple-system, sans-serif;
      --cs-bg: #ffffff;
      --cs-text: #1f2937;
      --cs-border: #e5e7eb;
      --cs-layer-bg: #f8fafc;
      --cs-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
      --cs-radius: 8px;
      --cs-radius-sm: 6px;
      --cs-gap: 16px;
      --cs-gap-sm: 8px;

      display: block;
      font-family: var(--cs-font-family);
      background: var(--cs-bg);
      color: var(--cs-text);
    }

    :host([theme="dark"]) {
      --cs-bg: #1f2937;
      --cs-text: #f9fafb;
      --cs-border: #374151;
      --cs-layer-bg: #111827;
    }

    .container {
      padding: var(--cs-gap);
    }

    .title {
      font-size: 1.5rem;
      font-weight: 600;
      margin: 0 0 var(--cs-gap) 0;
      text-align: center;
    }

    .stack {
      display: flex;
      flex-direction: column;
      gap: var(--cs-gap);
    }

    .layer {
      background: var(--cs-layer-bg);
      border: 1px solid var(--cs-border);
      border-radius: var(--cs-radius);
      padding: var(--cs-gap);
    }

    .layer-header {
      font-size: 0.875rem;
      font-weight: 600;
      color: var(--cs-text);
      margin-bottom: var(--cs-gap-sm);
      opacity: 0.8;
    }

    .capabilities {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
      gap: var(--cs-gap-sm);
    }

    .capability {
      padding: 12px 16px;
      border-radius: var(--cs-radius-sm);
      font-size: 0.875rem;
      font-weight: 500;
      text-align: center;
      box-shadow: var(--cs-shadow);
      cursor: default;
      transition: transform 0.15s ease, box-shadow 0.15s ease;
      position: relative;
    }

    :host([interactive]) .capability {
      cursor: pointer;
    }

    :host([interactive]) .capability:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    }

    .capability-tooltip {
      display: none;
      position: absolute;
      bottom: 100%;
      left: 50%;
      transform: translateX(-50%);
      background: var(--cs-bg);
      border: 1px solid var(--cs-border);
      border-radius: var(--cs-radius-sm);
      padding: 12px;
      min-width: 200px;
      max-width: 300px;
      box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
      z-index: 100;
      text-align: left;
      margin-bottom: 8px;
      color: var(--cs-text);
    }

    :host([interactive]) .capability:hover .capability-tooltip {
      display: block;
    }

    .capability-tooltip::after {
      content: '';
      position: absolute;
      top: 100%;
      left: 50%;
      transform: translateX(-50%);
      border: 6px solid transparent;
      border-top-color: var(--cs-border);
    }

    .tooltip-title {
      font-weight: 600;
      margin-bottom: 4px;
    }

    .tooltip-desc {
      font-size: 0.75rem;
      opacity: 0.8;
      margin-bottom: 8px;
    }

    .tooltip-meta {
      font-size: 0.75rem;
      display: flex;
      flex-direction: column;
      gap: 4px;
    }

    .tooltip-meta-item {
      display: flex;
      justify-content: space-between;
    }

    .tooltip-meta-label {
      opacity: 0.6;
    }

    .foundational {
      background: #fef3c7;
      border: 2px dashed #f59e0b;
    }

    :host([theme="dark"]) .foundational {
      background: #78350f;
    }

    .legend {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      gap: var(--cs-gap-sm);
      margin-top: var(--cs-gap);
      padding-top: var(--cs-gap);
      border-top: 1px solid var(--cs-border);
    }

    .legend-item {
      display: flex;
      align-items: center;
      gap: 6px;
      font-size: 0.75rem;
    }

    .legend-color {
      width: 16px;
      height: 16px;
      border-radius: 4px;
    }

    .loading, .error {
      padding: 2rem;
      text-align: center;
    }

    .error {
      color: #dc2626;
    }
  `;

  /** URL to the capability stack JSON file */
  @property({ type: String })
  src = '';

  /** Inline JSON data (alternative to src) */
  @property({ type: Object })
  data?: CapabilityStack;

  /** Color theme */
  @property({ type: String, reflect: true })
  theme: 'light' | 'dark' = 'light';

  /** Show the status legend */
  @property({ type: Boolean, attribute: 'show-legend' })
  showLegend = true;

  /** Enable interactive hover tooltips */
  @property({ type: Boolean, reflect: true })
  interactive = false;

  /** Show foundational capabilities */
  @property({ type: Boolean, attribute: 'show-foundational' })
  showFoundational = true;

  /** Filter by status */
  @property({ type: String, attribute: 'filter-status' })
  filterStatus?: CapabilityStatus;

  @state()
  private _stack?: CapabilityStack;

  @state()
  private _loading = false;

  @state()
  private _error?: string;

  override connectedCallback() {
    super.connectedCallback();
    this._loadData();
  }

  override updated(changedProperties: Map<string, unknown>) {
    if (changedProperties.has('src') || changedProperties.has('data')) {
      this._loadData();
    }
  }

  private async _loadData() {
    if (this.data) {
      this._stack = this.data;
      return;
    }

    if (!this.src) {
      return;
    }

    this._loading = true;
    this._error = undefined;

    try {
      const response = await fetch(this.src);
      if (!response.ok) {
        throw new Error(`Failed to load: ${response.statusText}`);
      }
      this._stack = await response.json();
    } catch (e) {
      this._error = e instanceof Error ? e.message : 'Failed to load data';
    } finally {
      this._loading = false;
    }
  }

  private _getCapabilitiesForLayer(layerId: string): Capability[] {
    if (!this._stack) return [];
    let caps = this._stack.capabilities.filter(c => c.layerId === layerId);
    if (this.filterStatus) {
      caps = caps.filter(c => c.status === this.filterStatus);
    }
    return caps;
  }

  private _getCategoryColor(categoryId?: string): string | undefined {
    if (!categoryId || !this._stack?.categories) return undefined;
    const cat = this._stack.categories.find(c => c.id === categoryId);
    return cat?.color;
  }

  private _getCapabilityStyle(cap: Capability): string {
    const status = cap.status || 'planned';
    const colors = STATUS_COLORS[status] || STATUS_COLORS.planned;
    return `background-color: ${colors.bg}; color: ${colors.text};`;
  }

  private _renderCapability(cap: Capability) {
    return html`
      <div class="capability" style=${this._getCapabilityStyle(cap)}>
        ${cap.name}
        ${this.interactive ? this._renderTooltip(cap) : nothing}
      </div>
    `;
  }

  private _renderTooltip(cap: Capability) {
    return html`
      <div class="capability-tooltip">
        <div class="tooltip-title">${cap.fullName || cap.name}</div>
        ${cap.description ? html`<div class="tooltip-desc">${cap.description}</div>` : nothing}
        <div class="tooltip-meta">
          ${cap.status ? html`
            <div class="tooltip-meta-item">
              <span class="tooltip-meta-label">Status</span>
              <span>${STATUS_LABELS[cap.status]}</span>
            </div>
          ` : nothing}
          ${cap.owner ? html`
            <div class="tooltip-meta-item">
              <span class="tooltip-meta-label">Owner</span>
              <span>${cap.owner}</span>
            </div>
          ` : nothing}
          ${cap.priority ? html`
            <div class="tooltip-meta-item">
              <span class="tooltip-meta-label">Priority</span>
              <span>${cap.priority}</span>
            </div>
          ` : nothing}
          ${cap.tooling?.length ? html`
            <div class="tooltip-meta-item">
              <span class="tooltip-meta-label">Tooling</span>
              <span>${cap.tooling.map(t => t.name).join(', ')}</span>
            </div>
          ` : nothing}
        </div>
      </div>
    `;
  }

  private _renderLayer(layer: Layer) {
    const caps = this._getCapabilitiesForLayer(layer.id);
    if (caps.length === 0) return nothing;

    return html`
      <div class="layer">
        <div class="layer-header">${layer.name}</div>
        <div class="capabilities">
          ${caps.map(cap => this._renderCapability(cap))}
        </div>
      </div>
    `;
  }

  private _renderFoundational() {
    if (!this.showFoundational || !this._stack?.foundational?.length) {
      return nothing;
    }

    let caps = this._stack.foundational;
    if (this.filterStatus) {
      caps = caps.filter(c => c.status === this.filterStatus);
    }
    if (caps.length === 0) return nothing;

    return html`
      <div class="layer foundational">
        <div class="layer-header">Foundational</div>
        <div class="capabilities">
          ${caps.map(cap => this._renderCapability(cap))}
        </div>
      </div>
    `;
  }

  private _renderLegend() {
    if (!this.showLegend) return nothing;

    const statuses: CapabilityStatus[] = ['operational', 'implemented', 'in-progress', 'planned', 'deprecated'];

    return html`
      <div class="legend">
        ${statuses.map(status => html`
          <div class="legend-item">
            <div class="legend-color" style="background-color: ${STATUS_COLORS[status].bg}"></div>
            <span>${STATUS_LABELS[status]}</span>
          </div>
        `)}
      </div>
    `;
  }

  override render() {
    if (this._loading) {
      return html`<div class="loading">Loading...</div>`;
    }

    if (this._error) {
      return html`<div class="error">Error: ${this._error}</div>`;
    }

    if (!this._stack) {
      return html`<div class="loading">No data</div>`;
    }

    return html`
      <div class="container">
        ${this._stack.metadata.title
          ? html`<h1 class="title">${this._stack.metadata.title}</h1>`
          : nothing}
        <div class="stack">
          ${this._stack.layers.map(layer => this._renderLayer(layer))}
          ${this._renderFoundational()}
        </div>
        ${this._renderLegend()}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'capability-stack': CapabilityStackElement;
  }
}
