/**
 * @license
 * Copyright 2019 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const F = globalThis, q = F.ShadowRoot && (F.ShadyCSS === void 0 || F.ShadyCSS.nativeShadow) && "adoptedStyleSheets" in Document.prototype && "replace" in CSSStyleSheet.prototype, W = Symbol(), K = /* @__PURE__ */ new WeakMap();
let ct = class {
  constructor(t, e, i) {
    if (this._$cssResult$ = !0, i !== W) throw Error("CSSResult is not constructable. Use `unsafeCSS` or `css` instead.");
    this.cssText = t, this.t = e;
  }
  get styleSheet() {
    let t = this.o;
    const e = this.t;
    if (q && t === void 0) {
      const i = e !== void 0 && e.length === 1;
      i && (t = K.get(e)), t === void 0 && ((this.o = t = new CSSStyleSheet()).replaceSync(this.cssText), i && K.set(e, t));
    }
    return t;
  }
  toString() {
    return this.cssText;
  }
};
const gt = (r) => new ct(typeof r == "string" ? r : r + "", void 0, W), $t = (r, ...t) => {
  const e = r.length === 1 ? r[0] : t.reduce((i, s, o) => i + ((n) => {
    if (n._$cssResult$ === !0) return n.cssText;
    if (typeof n == "number") return n;
    throw Error("Value passed to 'css' function must be a 'css' function result: " + n + ". Use 'unsafeCSS' to pass non-literal values, but take care to ensure page security.");
  })(s) + r[o + 1], r[0]);
  return new ct(e, r, W);
}, mt = (r, t) => {
  if (q) r.adoptedStyleSheets = t.map((e) => e instanceof CSSStyleSheet ? e : e.styleSheet);
  else for (const e of t) {
    const i = document.createElement("style"), s = F.litNonce;
    s !== void 0 && i.setAttribute("nonce", s), i.textContent = e.cssText, r.appendChild(i);
  }
}, X = q ? (r) => r : (r) => r instanceof CSSStyleSheet ? ((t) => {
  let e = "";
  for (const i of t.cssRules) e += i.cssText;
  return gt(e);
})(r) : r;
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const { is: yt, defineProperty: bt, getOwnPropertyDescriptor: vt, getOwnPropertyNames: _t, getOwnPropertySymbols: At, getPrototypeOf: xt } = Object, b = globalThis, Y = b.trustedTypes, wt = Y ? Y.emptyScript : "", j = b.reactiveElementPolyfillSupport, M = (r, t) => r, L = { toAttribute(r, t) {
  switch (t) {
    case Boolean:
      r = r ? wt : null;
      break;
    case Object:
    case Array:
      r = r == null ? r : JSON.stringify(r);
  }
  return r;
}, fromAttribute(r, t) {
  let e = r;
  switch (t) {
    case Boolean:
      e = r !== null;
      break;
    case Number:
      e = r === null ? null : Number(r);
      break;
    case Object:
    case Array:
      try {
        e = JSON.parse(r);
      } catch {
        e = null;
      }
  }
  return e;
} }, J = (r, t) => !yt(r, t), Q = { attribute: !0, type: String, converter: L, reflect: !1, useDefault: !1, hasChanged: J };
Symbol.metadata ?? (Symbol.metadata = Symbol("metadata")), b.litPropertyMetadata ?? (b.litPropertyMetadata = /* @__PURE__ */ new WeakMap());
let w = class extends HTMLElement {
  static addInitializer(t) {
    this._$Ei(), (this.l ?? (this.l = [])).push(t);
  }
  static get observedAttributes() {
    return this.finalize(), this._$Eh && [...this._$Eh.keys()];
  }
  static createProperty(t, e = Q) {
    if (e.state && (e.attribute = !1), this._$Ei(), this.prototype.hasOwnProperty(t) && ((e = Object.create(e)).wrapped = !0), this.elementProperties.set(t, e), !e.noAccessor) {
      const i = Symbol(), s = this.getPropertyDescriptor(t, i, e);
      s !== void 0 && bt(this.prototype, t, s);
    }
  }
  static getPropertyDescriptor(t, e, i) {
    const { get: s, set: o } = vt(this.prototype, t) ?? { get() {
      return this[e];
    }, set(n) {
      this[e] = n;
    } };
    return { get: s, set(n) {
      const l = s == null ? void 0 : s.call(this);
      o == null || o.call(this, n), this.requestUpdate(t, l, i);
    }, configurable: !0, enumerable: !0 };
  }
  static getPropertyOptions(t) {
    return this.elementProperties.get(t) ?? Q;
  }
  static _$Ei() {
    if (this.hasOwnProperty(M("elementProperties"))) return;
    const t = xt(this);
    t.finalize(), t.l !== void 0 && (this.l = [...t.l]), this.elementProperties = new Map(t.elementProperties);
  }
  static finalize() {
    if (this.hasOwnProperty(M("finalized"))) return;
    if (this.finalized = !0, this._$Ei(), this.hasOwnProperty(M("properties"))) {
      const e = this.properties, i = [..._t(e), ...At(e)];
      for (const s of i) this.createProperty(s, e[s]);
    }
    const t = this[Symbol.metadata];
    if (t !== null) {
      const e = litPropertyMetadata.get(t);
      if (e !== void 0) for (const [i, s] of e) this.elementProperties.set(i, s);
    }
    this._$Eh = /* @__PURE__ */ new Map();
    for (const [e, i] of this.elementProperties) {
      const s = this._$Eu(e, i);
      s !== void 0 && this._$Eh.set(s, e);
    }
    this.elementStyles = this.finalizeStyles(this.styles);
  }
  static finalizeStyles(t) {
    const e = [];
    if (Array.isArray(t)) {
      const i = new Set(t.flat(1 / 0).reverse());
      for (const s of i) e.unshift(X(s));
    } else t !== void 0 && e.push(X(t));
    return e;
  }
  static _$Eu(t, e) {
    const i = e.attribute;
    return i === !1 ? void 0 : typeof i == "string" ? i : typeof t == "string" ? t.toLowerCase() : void 0;
  }
  constructor() {
    super(), this._$Ep = void 0, this.isUpdatePending = !1, this.hasUpdated = !1, this._$Em = null, this._$Ev();
  }
  _$Ev() {
    var t;
    this._$ES = new Promise((e) => this.enableUpdating = e), this._$AL = /* @__PURE__ */ new Map(), this._$E_(), this.requestUpdate(), (t = this.constructor.l) == null || t.forEach((e) => e(this));
  }
  addController(t) {
    var e;
    (this._$EO ?? (this._$EO = /* @__PURE__ */ new Set())).add(t), this.renderRoot !== void 0 && this.isConnected && ((e = t.hostConnected) == null || e.call(t));
  }
  removeController(t) {
    var e;
    (e = this._$EO) == null || e.delete(t);
  }
  _$E_() {
    const t = /* @__PURE__ */ new Map(), e = this.constructor.elementProperties;
    for (const i of e.keys()) this.hasOwnProperty(i) && (t.set(i, this[i]), delete this[i]);
    t.size > 0 && (this._$Ep = t);
  }
  createRenderRoot() {
    const t = this.shadowRoot ?? this.attachShadow(this.constructor.shadowRootOptions);
    return mt(t, this.constructor.elementStyles), t;
  }
  connectedCallback() {
    var t;
    this.renderRoot ?? (this.renderRoot = this.createRenderRoot()), this.enableUpdating(!0), (t = this._$EO) == null || t.forEach((e) => {
      var i;
      return (i = e.hostConnected) == null ? void 0 : i.call(e);
    });
  }
  enableUpdating(t) {
  }
  disconnectedCallback() {
    var t;
    (t = this._$EO) == null || t.forEach((e) => {
      var i;
      return (i = e.hostDisconnected) == null ? void 0 : i.call(e);
    });
  }
  attributeChangedCallback(t, e, i) {
    this._$AK(t, i);
  }
  _$ET(t, e) {
    var o;
    const i = this.constructor.elementProperties.get(t), s = this.constructor._$Eu(t, i);
    if (s !== void 0 && i.reflect === !0) {
      const n = (((o = i.converter) == null ? void 0 : o.toAttribute) !== void 0 ? i.converter : L).toAttribute(e, i.type);
      this._$Em = t, n == null ? this.removeAttribute(s) : this.setAttribute(s, n), this._$Em = null;
    }
  }
  _$AK(t, e) {
    var o, n;
    const i = this.constructor, s = i._$Eh.get(t);
    if (s !== void 0 && this._$Em !== s) {
      const l = i.getPropertyOptions(s), a = typeof l.converter == "function" ? { fromAttribute: l.converter } : ((o = l.converter) == null ? void 0 : o.fromAttribute) !== void 0 ? l.converter : L;
      this._$Em = s;
      const c = a.fromAttribute(e, l.type);
      this[s] = c ?? ((n = this._$Ej) == null ? void 0 : n.get(s)) ?? c, this._$Em = null;
    }
  }
  requestUpdate(t, e, i, s = !1, o) {
    var n;
    if (t !== void 0) {
      const l = this.constructor;
      if (s === !1 && (o = this[t]), i ?? (i = l.getPropertyOptions(t)), !((i.hasChanged ?? J)(o, e) || i.useDefault && i.reflect && o === ((n = this._$Ej) == null ? void 0 : n.get(t)) && !this.hasAttribute(l._$Eu(t, i)))) return;
      this.C(t, e, i);
    }
    this.isUpdatePending === !1 && (this._$ES = this._$EP());
  }
  C(t, e, { useDefault: i, reflect: s, wrapped: o }, n) {
    i && !(this._$Ej ?? (this._$Ej = /* @__PURE__ */ new Map())).has(t) && (this._$Ej.set(t, n ?? e ?? this[t]), o !== !0 || n !== void 0) || (this._$AL.has(t) || (this.hasUpdated || i || (e = void 0), this._$AL.set(t, e)), s === !0 && this._$Em !== t && (this._$Eq ?? (this._$Eq = /* @__PURE__ */ new Set())).add(t));
  }
  async _$EP() {
    this.isUpdatePending = !0;
    try {
      await this._$ES;
    } catch (e) {
      Promise.reject(e);
    }
    const t = this.scheduleUpdate();
    return t != null && await t, !this.isUpdatePending;
  }
  scheduleUpdate() {
    return this.performUpdate();
  }
  performUpdate() {
    var i;
    if (!this.isUpdatePending) return;
    if (!this.hasUpdated) {
      if (this.renderRoot ?? (this.renderRoot = this.createRenderRoot()), this._$Ep) {
        for (const [o, n] of this._$Ep) this[o] = n;
        this._$Ep = void 0;
      }
      const s = this.constructor.elementProperties;
      if (s.size > 0) for (const [o, n] of s) {
        const { wrapped: l } = n, a = this[o];
        l !== !0 || this._$AL.has(o) || a === void 0 || this.C(o, void 0, n, a);
      }
    }
    let t = !1;
    const e = this._$AL;
    try {
      t = this.shouldUpdate(e), t ? (this.willUpdate(e), (i = this._$EO) == null || i.forEach((s) => {
        var o;
        return (o = s.hostUpdate) == null ? void 0 : o.call(s);
      }), this.update(e)) : this._$EM();
    } catch (s) {
      throw t = !1, this._$EM(), s;
    }
    t && this._$AE(e);
  }
  willUpdate(t) {
  }
  _$AE(t) {
    var e;
    (e = this._$EO) == null || e.forEach((i) => {
      var s;
      return (s = i.hostUpdated) == null ? void 0 : s.call(i);
    }), this.hasUpdated || (this.hasUpdated = !0, this.firstUpdated(t)), this.updated(t);
  }
  _$EM() {
    this._$AL = /* @__PURE__ */ new Map(), this.isUpdatePending = !1;
  }
  get updateComplete() {
    return this.getUpdateComplete();
  }
  getUpdateComplete() {
    return this._$ES;
  }
  shouldUpdate(t) {
    return !0;
  }
  update(t) {
    this._$Eq && (this._$Eq = this._$Eq.forEach((e) => this._$ET(e, this[e]))), this._$EM();
  }
  updated(t) {
  }
  firstUpdated(t) {
  }
};
w.elementStyles = [], w.shadowRootOptions = { mode: "open" }, w[M("elementProperties")] = /* @__PURE__ */ new Map(), w[M("finalized")] = /* @__PURE__ */ new Map(), j == null || j({ ReactiveElement: w }), (b.reactiveElementVersions ?? (b.reactiveElementVersions = [])).push("2.1.2");
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const O = globalThis, tt = (r) => r, z = O.trustedTypes, et = z ? z.createPolicy("lit-html", { createHTML: (r) => r }) : void 0, dt = "$lit$", y = `lit$${Math.random().toFixed(9).slice(2)}$`, pt = "?" + y, St = `<${pt}>`, x = document, k = () => x.createComment(""), T = (r) => r === null || typeof r != "object" && typeof r != "function", G = Array.isArray, Et = (r) => G(r) || typeof (r == null ? void 0 : r[Symbol.iterator]) == "function", I = `[ 	
\f\r]`, P = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, st = /-->/g, it = />/g, v = RegExp(`>|${I}(?:([^\\s"'>=/]+)(${I}*=${I}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), rt = /'/g, ot = /"/g, ut = /^(?:script|style|textarea|title)$/i, Ct = (r) => (t, ...e) => ({ _$litType$: r, strings: t, values: e }), u = Ct(1), S = Symbol.for("lit-noChange"), d = Symbol.for("lit-nothing"), nt = /* @__PURE__ */ new WeakMap(), _ = x.createTreeWalker(x, 129);
function ft(r, t) {
  if (!G(r) || !r.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return et !== void 0 ? et.createHTML(t) : t;
}
const Pt = (r, t) => {
  const e = r.length - 1, i = [];
  let s, o = t === 2 ? "<svg>" : t === 3 ? "<math>" : "", n = P;
  for (let l = 0; l < e; l++) {
    const a = r[l];
    let c, p, h = -1, g = 0;
    for (; g < a.length && (n.lastIndex = g, p = n.exec(a), p !== null); ) g = n.lastIndex, n === P ? p[1] === "!--" ? n = st : p[1] !== void 0 ? n = it : p[2] !== void 0 ? (ut.test(p[2]) && (s = RegExp("</" + p[2], "g")), n = v) : p[3] !== void 0 && (n = v) : n === v ? p[0] === ">" ? (n = s ?? P, h = -1) : p[1] === void 0 ? h = -2 : (h = n.lastIndex - p[2].length, c = p[1], n = p[3] === void 0 ? v : p[3] === '"' ? ot : rt) : n === ot || n === rt ? n = v : n === st || n === it ? n = P : (n = v, s = void 0);
    const m = n === v && r[l + 1].startsWith("/>") ? " " : "";
    o += n === P ? a + St : h >= 0 ? (i.push(c), a.slice(0, h) + dt + a.slice(h) + y + m) : a + y + (h === -2 ? l : m);
  }
  return [ft(r, o + (r[e] || "<?>") + (t === 2 ? "</svg>" : t === 3 ? "</math>" : "")), i];
};
class H {
  constructor({ strings: t, _$litType$: e }, i) {
    let s;
    this.parts = [];
    let o = 0, n = 0;
    const l = t.length - 1, a = this.parts, [c, p] = Pt(t, e);
    if (this.el = H.createElement(c, i), _.currentNode = this.el.content, e === 2 || e === 3) {
      const h = this.el.content.firstChild;
      h.replaceWith(...h.childNodes);
    }
    for (; (s = _.nextNode()) !== null && a.length < l; ) {
      if (s.nodeType === 1) {
        if (s.hasAttributes()) for (const h of s.getAttributeNames()) if (h.endsWith(dt)) {
          const g = p[n++], m = s.getAttribute(h).split(y), R = /([.?@])?(.*)/.exec(g);
          a.push({ type: 1, index: o, name: R[2], strings: m, ctor: R[1] === "." ? Ot : R[1] === "?" ? Ut : R[1] === "@" ? kt : D }), s.removeAttribute(h);
        } else h.startsWith(y) && (a.push({ type: 6, index: o }), s.removeAttribute(h));
        if (ut.test(s.tagName)) {
          const h = s.textContent.split(y), g = h.length - 1;
          if (g > 0) {
            s.textContent = z ? z.emptyScript : "";
            for (let m = 0; m < g; m++) s.append(h[m], k()), _.nextNode(), a.push({ type: 2, index: ++o });
            s.append(h[g], k());
          }
        }
      } else if (s.nodeType === 8) if (s.data === pt) a.push({ type: 2, index: o });
      else {
        let h = -1;
        for (; (h = s.data.indexOf(y, h + 1)) !== -1; ) a.push({ type: 7, index: o }), h += y.length - 1;
      }
      o++;
    }
  }
  static createElement(t, e) {
    const i = x.createElement("template");
    return i.innerHTML = t, i;
  }
}
function E(r, t, e = r, i) {
  var n, l;
  if (t === S) return t;
  let s = i !== void 0 ? (n = e._$Co) == null ? void 0 : n[i] : e._$Cl;
  const o = T(t) ? void 0 : t._$litDirective$;
  return (s == null ? void 0 : s.constructor) !== o && ((l = s == null ? void 0 : s._$AO) == null || l.call(s, !1), o === void 0 ? s = void 0 : (s = new o(r), s._$AT(r, e, i)), i !== void 0 ? (e._$Co ?? (e._$Co = []))[i] = s : e._$Cl = s), s !== void 0 && (t = E(r, s._$AS(r, t.values), s, i)), t;
}
class Mt {
  constructor(t, e) {
    this._$AV = [], this._$AN = void 0, this._$AD = t, this._$AM = e;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t) {
    const { el: { content: e }, parts: i } = this._$AD, s = ((t == null ? void 0 : t.creationScope) ?? x).importNode(e, !0);
    _.currentNode = s;
    let o = _.nextNode(), n = 0, l = 0, a = i[0];
    for (; a !== void 0; ) {
      if (n === a.index) {
        let c;
        a.type === 2 ? c = new N(o, o.nextSibling, this, t) : a.type === 1 ? c = new a.ctor(o, a.name, a.strings, this, t) : a.type === 6 && (c = new Tt(o, this, t)), this._$AV.push(c), a = i[++l];
      }
      n !== (a == null ? void 0 : a.index) && (o = _.nextNode(), n++);
    }
    return _.currentNode = x, s;
  }
  p(t) {
    let e = 0;
    for (const i of this._$AV) i !== void 0 && (i.strings !== void 0 ? (i._$AI(t, i, e), e += i.strings.length - 2) : i._$AI(t[e])), e++;
  }
}
class N {
  get _$AU() {
    var t;
    return ((t = this._$AM) == null ? void 0 : t._$AU) ?? this._$Cv;
  }
  constructor(t, e, i, s) {
    this.type = 2, this._$AH = d, this._$AN = void 0, this._$AA = t, this._$AB = e, this._$AM = i, this.options = s, this._$Cv = (s == null ? void 0 : s.isConnected) ?? !0;
  }
  get parentNode() {
    let t = this._$AA.parentNode;
    const e = this._$AM;
    return e !== void 0 && (t == null ? void 0 : t.nodeType) === 11 && (t = e.parentNode), t;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t, e = this) {
    t = E(this, t, e), T(t) ? t === d || t == null || t === "" ? (this._$AH !== d && this._$AR(), this._$AH = d) : t !== this._$AH && t !== S && this._(t) : t._$litType$ !== void 0 ? this.$(t) : t.nodeType !== void 0 ? this.T(t) : Et(t) ? this.k(t) : this._(t);
  }
  O(t) {
    return this._$AA.parentNode.insertBefore(t, this._$AB);
  }
  T(t) {
    this._$AH !== t && (this._$AR(), this._$AH = this.O(t));
  }
  _(t) {
    this._$AH !== d && T(this._$AH) ? this._$AA.nextSibling.data = t : this.T(x.createTextNode(t)), this._$AH = t;
  }
  $(t) {
    var o;
    const { values: e, _$litType$: i } = t, s = typeof i == "number" ? this._$AC(t) : (i.el === void 0 && (i.el = H.createElement(ft(i.h, i.h[0]), this.options)), i);
    if (((o = this._$AH) == null ? void 0 : o._$AD) === s) this._$AH.p(e);
    else {
      const n = new Mt(s, this), l = n.u(this.options);
      n.p(e), this.T(l), this._$AH = n;
    }
  }
  _$AC(t) {
    let e = nt.get(t.strings);
    return e === void 0 && nt.set(t.strings, e = new H(t)), e;
  }
  k(t) {
    G(this._$AH) || (this._$AH = [], this._$AR());
    const e = this._$AH;
    let i, s = 0;
    for (const o of t) s === e.length ? e.push(i = new N(this.O(k()), this.O(k()), this, this.options)) : i = e[s], i._$AI(o), s++;
    s < e.length && (this._$AR(i && i._$AB.nextSibling, s), e.length = s);
  }
  _$AR(t = this._$AA.nextSibling, e) {
    var i;
    for ((i = this._$AP) == null ? void 0 : i.call(this, !1, !0, e); t !== this._$AB; ) {
      const s = tt(t).nextSibling;
      tt(t).remove(), t = s;
    }
  }
  setConnected(t) {
    var e;
    this._$AM === void 0 && (this._$Cv = t, (e = this._$AP) == null || e.call(this, t));
  }
}
class D {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t, e, i, s, o) {
    this.type = 1, this._$AH = d, this._$AN = void 0, this.element = t, this.name = e, this._$AM = s, this.options = o, i.length > 2 || i[0] !== "" || i[1] !== "" ? (this._$AH = Array(i.length - 1).fill(new String()), this.strings = i) : this._$AH = d;
  }
  _$AI(t, e = this, i, s) {
    const o = this.strings;
    let n = !1;
    if (o === void 0) t = E(this, t, e, 0), n = !T(t) || t !== this._$AH && t !== S, n && (this._$AH = t);
    else {
      const l = t;
      let a, c;
      for (t = o[0], a = 0; a < o.length - 1; a++) c = E(this, l[i + a], e, a), c === S && (c = this._$AH[a]), n || (n = !T(c) || c !== this._$AH[a]), c === d ? t = d : t !== d && (t += (c ?? "") + o[a + 1]), this._$AH[a] = c;
    }
    n && !s && this.j(t);
  }
  j(t) {
    t === d ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t ?? "");
  }
}
class Ot extends D {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t) {
    this.element[this.name] = t === d ? void 0 : t;
  }
}
class Ut extends D {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t) {
    this.element.toggleAttribute(this.name, !!t && t !== d);
  }
}
class kt extends D {
  constructor(t, e, i, s, o) {
    super(t, e, i, s, o), this.type = 5;
  }
  _$AI(t, e = this) {
    if ((t = E(this, t, e, 0) ?? d) === S) return;
    const i = this._$AH, s = t === d && i !== d || t.capture !== i.capture || t.once !== i.once || t.passive !== i.passive, o = t !== d && (i === d || s);
    s && this.element.removeEventListener(this.name, this, i), o && this.element.addEventListener(this.name, this, t), this._$AH = t;
  }
  handleEvent(t) {
    var e;
    typeof this._$AH == "function" ? this._$AH.call(((e = this.options) == null ? void 0 : e.host) ?? this.element, t) : this._$AH.handleEvent(t);
  }
}
class Tt {
  constructor(t, e, i) {
    this.element = t, this.type = 6, this._$AN = void 0, this._$AM = e, this.options = i;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t) {
    E(this, t);
  }
}
const B = O.litHtmlPolyfillSupport;
B == null || B(H, N), (O.litHtmlVersions ?? (O.litHtmlVersions = [])).push("3.3.3");
const Ht = (r, t, e) => {
  const i = (e == null ? void 0 : e.renderBefore) ?? t;
  let s = i._$litPart$;
  if (s === void 0) {
    const o = (e == null ? void 0 : e.renderBefore) ?? null;
    i._$litPart$ = s = new N(t.insertBefore(k(), o), o, void 0, e ?? {});
  }
  return s._$AI(r), s;
};
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const A = globalThis;
class U extends w {
  constructor() {
    super(...arguments), this.renderOptions = { host: this }, this._$Do = void 0;
  }
  createRenderRoot() {
    var e;
    const t = super.createRenderRoot();
    return (e = this.renderOptions).renderBefore ?? (e.renderBefore = t.firstChild), t;
  }
  update(t) {
    const e = this.render();
    this.hasUpdated || (this.renderOptions.isConnected = this.isConnected), super.update(t), this._$Do = Ht(e, this.renderRoot, this.renderOptions);
  }
  connectedCallback() {
    var t;
    super.connectedCallback(), (t = this._$Do) == null || t.setConnected(!0);
  }
  disconnectedCallback() {
    var t;
    super.disconnectedCallback(), (t = this._$Do) == null || t.setConnected(!1);
  }
  render() {
    return S;
  }
}
var ht;
U._$litElement$ = !0, U.finalized = !0, (ht = A.litElementHydrateSupport) == null || ht.call(A, { LitElement: U });
const V = A.litElementPolyfillSupport;
V == null || V({ LitElement: U });
(A.litElementVersions ?? (A.litElementVersions = [])).push("4.2.2");
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Nt = (r) => (t, e) => {
  e !== void 0 ? e.addInitializer(() => {
    customElements.define(r, t);
  }) : customElements.define(r, t);
};
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Rt = { attribute: !0, type: String, converter: L, reflect: !1, hasChanged: J }, Ft = (r = Rt, t, e) => {
  const { kind: i, metadata: s } = e;
  let o = globalThis.litPropertyMetadata.get(s);
  if (o === void 0 && globalThis.litPropertyMetadata.set(s, o = /* @__PURE__ */ new Map()), i === "setter" && ((r = Object.create(r)).wrapped = !0), o.set(e.name, r), i === "accessor") {
    const { name: n } = e;
    return { set(l) {
      const a = t.get.call(this);
      t.set.call(this, l), this.requestUpdate(n, a, r, !0, l);
    }, init(l) {
      return l !== void 0 && this.C(n, void 0, r, l), l;
    } };
  }
  if (i === "setter") {
    const { name: n } = e;
    return function(l) {
      const a = this[n];
      t.call(this, l), this.requestUpdate(n, a, r, !0, l);
    };
  }
  throw Error("Unsupported decorator location: " + i);
};
function C(r) {
  return (t, e) => typeof e == "object" ? Ft(r, t, e) : ((i, s, o) => {
    const n = s.hasOwnProperty(o);
    return s.constructor.createProperty(o, i), n ? Object.getOwnPropertyDescriptor(s, o) : void 0;
  })(r, t, e);
}
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
function Z(r) {
  return C({ ...r, state: !0, attribute: !1 });
}
const at = {
  operational: "#10b981",
  implemented: "#3b82f6",
  "in-progress": "#f59e0b",
  planned: "#9ca3af",
  deprecated: "#ef4444"
}, Lt = {
  operational: "#ffffff",
  implemented: "#ffffff",
  "in-progress": "#000000",
  planned: "#000000",
  deprecated: "#ffffff"
}, lt = {
  1: "#ef4444",
  // red
  2: "#f59e0b",
  // amber
  3: "#eab308",
  // yellow
  4: "#22c55e",
  // green
  5: "#3b82f6"
  // blue
};
var zt = Object.defineProperty, Dt = Object.getOwnPropertyDescriptor, $ = (r, t, e, i) => {
  for (var s = i > 1 ? void 0 : i ? Dt(t, e) : t, o = r.length - 1, n; o >= 0; o--)
    (n = r[o]) && (s = (i ? n(t, e, s) : n(s)) || s);
  return i && s && zt(t, e, s), s;
};
let f = class extends U {
  constructor() {
    super(...arguments), this.view = "by-layer", this.theme = "light", this.showLegend = !0, this.showViewToggle = !0, this.data = null, this.statusFilters = /* @__PURE__ */ new Set([
      "operational",
      "implemented",
      "in-progress",
      "planned",
      "deprecated"
    ]), this.maturityFilters = /* @__PURE__ */ new Set([1, 2, 3, 4, 5]);
  }
  async connectedCallback() {
    super.connectedCallback(), requestAnimationFrame(() => this.loadData());
  }
  updated(r) {
    r.has("src") && this.loadData();
  }
  async loadData() {
    if (this.src)
      try {
        const t = await fetch(this.src);
        this.data = await t.json();
        return;
      } catch (t) {
        console.error("Failed to load data from src:", t);
      }
    const r = this.querySelector('script[type="application/json"]');
    if (r != null && r.textContent)
      try {
        this.data = JSON.parse(r.textContent);
      } catch (t) {
        console.error("Failed to parse inline JSON:", t);
      }
  }
  getGroups() {
    if (!this.data) return [];
    const r = [], t = /* @__PURE__ */ new Map(), e = /* @__PURE__ */ new Map(), i = /* @__PURE__ */ new Map();
    if (this.view === "by-layer") {
      for (let s = 0; s < this.data.layers.length; s++) {
        const o = this.data.layers[s];
        e.set(o.id, o.order ?? s), i.set(o.id, o.name);
      }
      for (const s of this.data.capabilities) {
        const o = s.layerId || "other";
        t.has(o) || t.set(o, []), t.get(o).push(s), e.has(o) || (e.set(o, 100), i.set(o, o));
      }
    } else {
      for (let s = 0; s < this.data.categories.length; s++) {
        const o = this.data.categories[s];
        e.set(o.id, o.order ?? s), i.set(o.id, o.name);
      }
      for (const s of this.data.capabilities) {
        const o = s.categoryId || "uncategorized";
        t.has(o) || t.set(o, []), t.get(o).push(s), e.has(o) || (e.set(o, 100), i.set(o, o.charAt(0).toUpperCase() + o.slice(1)));
      }
    }
    for (const [s, o] of t)
      r.push({
        id: s,
        name: i.get(s) || s,
        order: e.get(s) || 100,
        capabilities: o
      });
    return r.sort((s, o) => s.order - o.order), r;
  }
  hasMaturityData() {
    var r;
    return !!((r = this.data) != null && r.maturity) && Object.keys(this.data.maturity).length > 0;
  }
  getMaturityLevel(r) {
    var t, e;
    return (e = (t = this.data) == null ? void 0 : t.maturity) != null && e[r] ? this.data.maturity[r].level : null;
  }
  isFiltered(r) {
    if (!this.statusFilters.has(r.status))
      return !0;
    if (this.hasMaturityData()) {
      const t = this.getMaturityLevel(r.id);
      if (t !== null && !this.maturityFilters.has(t))
        return !0;
    }
    return !1;
  }
  toggleStatusFilter(r) {
    const t = new Set(this.statusFilters);
    t.has(r) ? t.delete(r) : t.add(r), this.statusFilters = t;
  }
  toggleMaturityFilter(r) {
    const t = new Set(this.maturityFilters);
    t.has(r) ? t.delete(r) : t.add(r), this.maturityFilters = t;
  }
  selectAll() {
    this.statusFilters = /* @__PURE__ */ new Set([
      "operational",
      "implemented",
      "in-progress",
      "planned",
      "deprecated"
    ]), this.maturityFilters = /* @__PURE__ */ new Set([1, 2, 3, 4, 5]);
  }
  clearAll() {
    this.statusFilters = /* @__PURE__ */ new Set(), this.maturityFilters = /* @__PURE__ */ new Set();
  }
  setView(r) {
    this.view = r, this.dispatchEvent(new CustomEvent("view-change", { detail: { view: r } }));
  }
  renderFilters() {
    return this.showLegend ? u`
      <div class="filters">
        <div class="filter-group">
          <div class="filter-label">Status</div>
          <div class="filter-options">
            ${[
      "operational",
      "implemented",
      "in-progress",
      "planned",
      "deprecated"
    ].map(
      (t) => u`
                <button
                  class="filter-btn ${this.statusFilters.has(t) ? "" : "inactive"}"
                  @click=${() => this.toggleStatusFilter(t)}
                >
                  <span
                    class="filter-color"
                    style="background-color: ${at[t]}"
                  ></span>
                  <span>${this.formatStatus(t)}</span>
                </button>
              `
    )}
          </div>
        </div>

        ${this.hasMaturityData() ? u`
              <div class="filter-group">
                <div class="filter-label">Maturity Level</div>
                <div class="filter-options">
                  ${[1, 2, 3, 4, 5].map(
      (t) => u`
                      <button
                        class="filter-btn ${this.maturityFilters.has(t) ? "" : "inactive"}"
                        @click=${() => this.toggleMaturityFilter(t)}
                      >
                        <span
                          class="filter-color"
                          style="background-color: ${lt[t]}"
                        ></span>
                        <span>M${t}</span>
                      </button>
                    `
    )}
                </div>
              </div>
            ` : null}

        <div class="filter-actions">
          <button class="btn" @click=${this.selectAll}>Select All</button>
          <button class="btn" @click=${this.clearAll}>Clear All</button>
        </div>
      </div>
    ` : null;
  }
  formatStatus(r) {
    return r.split("-").map((t) => t.charAt(0).toUpperCase() + t.slice(1)).join(" ");
  }
  renderCapability(r) {
    const t = this.isFiltered(r), e = at[r.status] || "#e5e7eb", i = Lt[r.status] || "#000000", s = this.getMaturityLevel(r.id), o = [
      r.fullName,
      r.description,
      r.owner ? `Owner: ${r.owner}` : null,
      `Status: ${r.status}`,
      s !== null ? `Maturity: M${s}` : null
    ].filter(Boolean).join(" | ");
    return u`
      <div
        class="capability ${t ? "filtered-out" : ""}"
        style="background-color: ${e}; color: ${i}"
        title=${o}
      >
        <span class="cap-name">${r.name}</span>
        ${s !== null ? u`
              <span
                class="badge"
                style="background-color: ${lt[s]}; color: #ffffff"
              >
                M${s}
              </span>
            ` : null}
      </div>
    `;
  }
  render() {
    if (!this.data)
      return u`<div class="container">Loading...</div>`;
    const r = this.getGroups();
    return u`
      <div class="container">
        ${this.data.title ? u`<h1 class="title">${this.data.title}</h1>` : null}

        ${this.showViewToggle ? u`
              <div class="view-toggle">
                <button
                  class="view-btn ${this.view === "by-layer" ? "active" : ""}"
                  @click=${() => this.setView("by-layer")}
                >
                  By Layer
                </button>
                <button
                  class="view-btn ${this.view === "by-category" ? "active" : ""}"
                  @click=${() => this.setView("by-category")}
                >
                  By Category
                </button>
              </div>
            ` : null}

        ${this.renderFilters()}

        <div class="stack">
          ${r.map(
      (t) => u`
              <div class="layer">
                <div class="layer-header">${t.name}</div>
                <div class="capabilities">
                  ${t.capabilities.map((e) => this.renderCapability(e))}
                </div>
              </div>
            `
    )}
        </div>
      </div>
    `;
  }
};
f.styles = $t`
    :host {
      --mg-bg: #ffffff;
      --mg-text: #1f2937;
      --mg-border: #e5e7eb;
      --mg-layer-bg: #f8fafc;
      --mg-inactive-bg: #f3f4f6;
      --mg-inactive-text: #9ca3af;
      display: block;
      font-family: system-ui, -apple-system, sans-serif;
    }

    :host([theme='dark']) {
      --mg-bg: #0f172a;
      --mg-text: #f1f5f9;
      --mg-border: #334155;
      --mg-layer-bg: #1e293b;
      --mg-inactive-bg: #334155;
      --mg-inactive-text: #94a3b8;
    }

    .container {
      background: var(--mg-bg);
      color: var(--mg-text);
      padding: 24px;
      min-height: 100%;
      box-sizing: border-box;
    }

    .title {
      font-size: 1.75rem;
      font-weight: 700;
      margin: 0 0 24px 0;
      text-align: center;
      letter-spacing: -0.025em;
    }

    .filters {
      background: var(--mg-layer-bg);
      border: 1px solid var(--mg-border);
      border-radius: 12px;
      padding: 20px;
      margin-bottom: 24px;
    }

    .filter-group {
      margin-bottom: 16px;
    }

    .filter-group:last-of-type {
      margin-bottom: 0;
    }

    .filter-label {
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.05em;
      opacity: 0.7;
      margin-bottom: 10px;
    }

    .filter-options {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }

    .filter-btn {
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 8px 14px;
      border-radius: 8px;
      font-size: 0.8125rem;
      font-weight: 500;
      cursor: pointer;
      transition: all 0.15s ease;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      user-select: none;
    }

    .filter-btn:hover {
      background: rgba(255, 255, 255, 0.05);
    }

    .filter-btn.inactive {
      opacity: 0.4;
    }

    .filter-color {
      width: 14px;
      height: 14px;
      border-radius: 4px;
      flex-shrink: 0;
    }

    .filter-actions {
      display: flex;
      gap: 8px;
      margin-top: 16px;
      padding-top: 16px;
      border-top: 1px solid var(--mg-border);
    }

    .btn {
      padding: 6px 12px;
      border-radius: 6px;
      font-size: 0.75rem;
      font-weight: 500;
      cursor: pointer;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      transition: all 0.15s ease;
    }

    .btn:hover {
      background: rgba(255, 255, 255, 0.1);
    }

    .stack {
      display: flex;
      flex-direction: column;
      gap: 20px;
    }

    .layer {
      background: var(--mg-layer-bg);
      border: 1px solid var(--mg-border);
      border-radius: 12px;
      padding: 20px;
    }

    .layer-header {
      font-size: 0.875rem;
      font-weight: 600;
      margin-bottom: 12px;
      opacity: 0.8;
    }

    .capabilities {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
      gap: 10px;
    }

    .capability {
      padding: 14px 16px;
      border-radius: 8px;
      font-size: 0.875rem;
      font-weight: 500;
      text-align: center;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 6px;
      transition: all 0.2s ease;
      cursor: default;
    }

    .capability.filtered-out {
      background-color: var(--mg-inactive-bg) !important;
      color: var(--mg-inactive-text) !important;
      box-shadow: none;
      border: 1px solid var(--mg-border);
    }

    .capability.filtered-out .badge {
      background-color: transparent !important;
      color: var(--mg-inactive-text) !important;
      border: 1px solid var(--mg-border);
    }

    .cap-name {
      display: block;
      line-height: 1.3;
    }

    .badge {
      font-size: 0.625rem;
      font-weight: 700;
      padding: 3px 8px;
      border-radius: 4px;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }

    .view-toggle {
      display: flex;
      gap: 8px;
      margin-bottom: 16px;
      justify-content: center;
    }

    .view-btn {
      padding: 8px 16px;
      border-radius: 8px;
      font-size: 0.875rem;
      font-weight: 500;
      cursor: pointer;
      border: 1px solid var(--mg-border);
      background: transparent;
      color: var(--mg-text);
      transition: all 0.15s ease;
    }

    .view-btn.active {
      background: var(--mg-text);
      color: var(--mg-bg);
    }

    .view-btn:hover:not(.active) {
      background: rgba(255, 255, 255, 0.1);
    }
  `;
$([
  C({ type: String })
], f.prototype, "view", 2);
$([
  C({ type: String, reflect: !0 })
], f.prototype, "theme", 2);
$([
  C({ type: String })
], f.prototype, "src", 2);
$([
  C({ type: Boolean, attribute: "show-legend" })
], f.prototype, "showLegend", 2);
$([
  C({ type: Boolean, attribute: "show-view-toggle" })
], f.prototype, "showViewToggle", 2);
$([
  Z()
], f.prototype, "data", 2);
$([
  Z()
], f.prototype, "statusFilters", 2);
$([
  Z()
], f.prototype, "maturityFilters", 2);
f = $([
  Nt("maturity-grid")
], f);
export {
  lt as MATURITY_COLORS,
  f as MaturityGrid,
  at as STATUS_COLORS,
  Lt as STATUS_TEXT_COLORS
};
//# sourceMappingURL=prism-ui.js.map
