(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const r of document.querySelectorAll('link[rel="modulepreload"]'))l(r);new MutationObserver(r=>{for(const i of r)if(i.type==="childList")for(const s of i.addedNodes)s.tagName==="LINK"&&s.rel==="modulepreload"&&l(s)}).observe(document,{childList:!0,subtree:!0});function n(r){const i={};return r.integrity&&(i.integrity=r.integrity),r.referrerpolicy&&(i.referrerPolicy=r.referrerpolicy),r.crossorigin==="use-credentials"?i.credentials="include":r.crossorigin==="anonymous"?i.credentials="omit":i.credentials="same-origin",i}function l(r){if(r.ep)return;r.ep=!0;const i=n(r);fetch(r.href,i)}})();function x(){}function B(e){return e()}function C(){return Object.create(null)}function k(e){e.forEach(B)}function I(e){return typeof e=="function"}function K(e,t){return e!=e?t==t:e!==t||e&&typeof e=="object"||typeof e=="function"}let E;function W(e,t){return E||(E=document.createElement("a")),E.href=t,e===E.href}function z(e){return Object.keys(e).length===0}function d(e,t){e.appendChild(t)}function D(e,t,n){e.insertBefore(t,n||null)}function M(e){e.parentNode&&e.parentNode.removeChild(e)}function _(e){return document.createElement(e)}function T(e){return document.createTextNode(e)}function A(){return T(" ")}function P(e,t,n,l){return e.addEventListener(t,n,l),()=>e.removeEventListener(t,n,l)}function c(e,t,n){n==null?e.removeAttribute(t):e.getAttribute(t)!==n&&e.setAttribute(t,n)}function H(e){return Array.from(e.childNodes)}function J(e,t){t=""+t,e.data!==t&&(e.data=t)}function S(e,t){e.value=t==null?"":t}let j;function v(e){j=e}const g=[],q=[];let y=[];const G=[],Q=Promise.resolve();let L=!1;function R(){L||(L=!0,Q.then(F))}function N(e){y.push(e)}const O=new Set;let m=0;function F(){if(m!==0)return;const e=j;do{try{for(;m<g.length;){const t=g[m];m++,v(t),U(t.$$)}}catch(t){throw g.length=0,m=0,t}for(v(null),g.length=0,m=0;q.length;)q.pop()();for(let t=0;t<y.length;t+=1){const n=y[t];O.has(n)||(O.add(n),n())}y.length=0}while(g.length);for(;G.length;)G.pop()();L=!1,O.clear(),v(e)}function U(e){if(e.fragment!==null){e.update(),k(e.before_update);const t=e.dirty;e.dirty=[-1],e.fragment&&e.fragment.p(e.ctx,t),e.after_update.forEach(N)}}function V(e){const t=[],n=[];y.forEach(l=>e.indexOf(l)===-1?t.push(l):n.push(l)),n.forEach(l=>l()),y=t}const X=new Set;function Y(e,t){e&&e.i&&(X.delete(e),e.i(t))}function Z(e,t,n,l){const{fragment:r,after_update:i}=e.$$;r&&r.m(t,n),l||N(()=>{const s=e.$$.on_mount.map(B).filter(I);e.$$.on_destroy?e.$$.on_destroy.push(...s):k(s),e.$$.on_mount=[]}),i.forEach(N)}function ee(e,t){const n=e.$$;n.fragment!==null&&(V(n.after_update),k(n.on_destroy),n.fragment&&n.fragment.d(t),n.on_destroy=n.fragment=null,n.ctx=[])}function te(e,t){e.$$.dirty[0]===-1&&(g.push(e),R(),e.$$.dirty.fill(0)),e.$$.dirty[t/31|0]|=1<<t%31}function ne(e,t,n,l,r,i,s,p=[-1]){const f=j;v(e);const o=e.$$={fragment:null,ctx:[],props:i,update:x,not_equal:r,bound:C(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(t.context||(f?f.$$.context:[])),callbacks:C(),dirty:p,skip_bound:!1,root:t.target||f.$$.root};s&&s(o.root);let $=!1;if(o.ctx=n?n(e,t.props||{},(u,h,...b)=>{const a=b.length?b[0]:h;return o.ctx&&r(o.ctx[u],o.ctx[u]=a)&&(!o.skip_bound&&o.bound[u]&&o.bound[u](a),$&&te(e,u)),h}):[],o.update(),$=!0,k(o.before_update),o.fragment=l?l(o.ctx):!1,t.target){if(t.hydrate){const u=H(t.target);o.fragment&&o.fragment.l(u),u.forEach(M)}else o.fragment&&o.fragment.c();t.intro&&Y(e.$$.fragment),Z(e,t.target,t.anchor,t.customElement),F()}v(f)}class re{$destroy(){ee(this,1),this.$destroy=x}$on(t,n){if(!I(n))return x;const l=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return l.push(n),()=>{const r=l.indexOf(n);r!==-1&&l.splice(r,1)}}$set(t){this.$$set&&!z(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}const oe="/assets/logo-universal.f9dae61b.png";function le(e){return window.go.main.App.Greet(e)}function ie(e){let t,n,l,r,i,s,p,f,o,$,u,h,b;return{c(){t=_("main"),n=_("img"),r=A(),i=_("div"),s=T(e[0]),p=A(),f=_("div"),o=_("input"),$=A(),u=_("button"),u.textContent="Greet",c(n,"alt","Wails logo"),c(n,"id","logo"),W(n.src,l=oe)||c(n,"src",l),c(n,"class","svelte-1yppp05"),c(i,"class","result svelte-1yppp05"),c(i,"id","result"),c(o,"autocomplete","off"),c(o,"class","input svelte-1yppp05"),c(o,"id","name"),c(o,"type","text"),c(u,"class","btn svelte-1yppp05"),c(f,"class","input-box svelte-1yppp05"),c(f,"id","input")},m(a,w){D(a,t,w),d(t,n),d(t,r),d(t,i),d(i,s),d(t,p),d(t,f),d(f,o),S(o,e[1]),d(f,$),d(f,u),h||(b=[P(o,"input",e[3]),P(u,"click",e[2])],h=!0)},p(a,[w]){w&1&&J(s,a[0]),w&2&&o.value!==a[1]&&S(o,a[1])},i:x,o:x,d(a){a&&M(t),h=!1,k(b)}}}function se(e,t,n){let l="Please enter your name below \u{1F447}",r;function i(){le(r).then(p=>n(0,l=p))}function s(){r=this.value,n(1,r)}return[l,r,i,s]}class ue extends re{constructor(t){super(),ne(this,t,se,ie,K,{})}}new ue({target:document.getElementById("app")});
