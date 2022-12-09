import{bs as C,af as X,aA as A,Q as h,ac as g,bt as I,bu as B,bv as _,a9 as $,bw as G,bx as W,aY as D,as as L,ao as P,by as N,bz as O,k as p,bA as V,al as F,r as S,aw as J,bB as Z}from"./index.3ecae71b.js";var se=C({name:"QPageContainer",setup(e,{slots:s}){const{proxy:{$q:n}}=$(),t=X(B,()=>{console.error("QPageContainer needs to be child of QLayout")});A(_,!0);const r=h(()=>{const a={};return t.header.space===!0&&(a.paddingTop=`${t.header.size}px`),t.right.space===!0&&(a[`padding${n.lang.rtl===!0?"Left":"Right"}`]=`${t.right.size}px`),t.footer.space===!0&&(a.paddingBottom=`${t.footer.size}px`),t.left.space===!0&&(a[`padding${n.lang.rtl===!0?"Right":"Left"}`]=`${t.left.size}px`),a});return()=>g("div",{class:"q-page-container",style:r.value},I(s.default))}});const ee=[null,document,document.body,document.scrollingElement,document.documentElement];function te(e,s){let n=G(s);if(n===void 0){if(e==null)return window;n=e.closest(".scroll,.scroll-y,.overflow-auto")}return ee.includes(n)?window:n}function ne(e){return e===window?window.pageYOffset||window.scrollY||document.body.scrollTop||0:e.scrollTop}function oe(e){return e===window?window.pageXOffset||window.scrollX||document.body.scrollLeft||0:e.scrollLeft}let T;function E(){if(T!==void 0)return T;const e=document.createElement("p"),s=document.createElement("div");W(e,{width:"100%",height:"200px"}),W(s,{position:"absolute",top:"0px",left:"0px",visibility:"hidden",width:"200px",height:"150px",overflow:"hidden"}),s.appendChild(e),document.body.appendChild(s);const n=e.offsetWidth;s.style.overflow="scroll";let t=e.offsetWidth;return n===t&&(t=s.clientWidth),s.remove(),T=n-t,T}const{passive:M}=O,ie=["both","horizontal","vertical"];var le=C({name:"QScrollObserver",props:{axis:{type:String,validator:e=>ie.includes(e),default:"vertical"},debounce:[String,Number],scrollTarget:{default:void 0}},emits:["scroll"],setup(e,{emit:s}){const n={position:{top:0,left:0},direction:"down",directionChanged:!1,delta:{top:0,left:0},inflectionPoint:{top:0,left:0}};let t=null,r,a;D(()=>e.scrollTarget,()=>{l(),f()});function u(){t!==null&&t();const v=Math.max(0,ne(r)),w=oe(r),d={top:v-n.position.top,left:w-n.position.left};if(e.axis==="vertical"&&d.top===0||e.axis==="horizontal"&&d.left===0)return;const z=Math.abs(d.top)>=Math.abs(d.left)?d.top<0?"up":"down":d.left<0?"left":"right";n.position={top:v,left:w},n.directionChanged=n.direction!==z,n.delta=d,n.directionChanged===!0&&(n.direction=z,n.inflectionPoint=n.position),s("scroll",{...n})}function f(){r=te(a,e.scrollTarget),r.addEventListener("scroll",i,M),i(!0)}function l(){r!==void 0&&(r.removeEventListener("scroll",i,M),r=void 0)}function i(v){if(v===!0||e.debounce===0||e.debounce==="0")u();else if(t===null){const[w,d]=e.debounce?[setTimeout(u,e.debounce),clearTimeout]:[requestAnimationFrame(u),cancelAnimationFrame];t=()=>{d(w),t=null}}}const{proxy:m}=$();return L(()=>{a=m.$el.parentNode,f()}),P(()=>{t!==null&&t(),l()}),Object.assign(m,{trigger:i,getPosition:()=>n}),N}});function re(){const e=p(!V.value);return e.value===!1&&L(()=>{e.value=!0}),e}const Y=typeof ResizeObserver!="undefined",j=Y===!0?{}:{style:"display:block;position:absolute;top:0;left:0;right:0;bottom:0;height:100%;width:100%;overflow:hidden;pointer-events:none;z-index:-1;",url:"about:blank"};var k=C({name:"QResizeObserver",props:{debounce:{type:[String,Number],default:100}},emits:["resize"],setup(e,{emit:s}){let n=null,t,r={width:-1,height:-1};function a(l){l===!0||e.debounce===0||e.debounce==="0"?u():n===null&&(n=setTimeout(u,e.debounce))}function u(){if(clearTimeout(n),n=null,t){const{offsetWidth:l,offsetHeight:i}=t;(l!==r.width||i!==r.height)&&(r={width:l,height:i},s("resize",r))}}const{proxy:f}=$();if(Y===!0){let l;return L(()=>{F(()=>{t=f.$el.parentNode,t&&(l=new ResizeObserver(a),l.observe(t),u())})}),P(()=>{clearTimeout(n),l!==void 0&&(l.disconnect!==void 0?l.disconnect():t&&l.unobserve(t))}),N}else{let m=function(){clearTimeout(n),i!==void 0&&(i.removeEventListener!==void 0&&i.removeEventListener("resize",a,O.passive),i=void 0)},v=function(){m(),t&&t.contentDocument&&(i=t.contentDocument.defaultView,i.addEventListener("resize",a,O.passive),u())};const l=re();let i;return L(()=>{F(()=>{t=f.$el,t&&v()})}),P(m),f.trigger=a,()=>{if(l.value===!0)return g("object",{style:j.style,tabindex:-1,type:"text/html",data:j.url,"aria-hidden":"true",onLoad:v})}}}}),ce=C({name:"QLayout",props:{container:Boolean,view:{type:String,default:"hhh lpr fff",validator:e=>/^(h|l)h(h|r) lpr (f|l)f(f|r)$/.test(e.toLowerCase())},onScroll:Function,onScrollHeight:Function,onResize:Function},setup(e,{slots:s,emit:n}){const{proxy:{$q:t}}=$(),r=p(null),a=p(t.screen.height),u=p(e.container===!0?0:t.screen.width),f=p({position:0,direction:"down",inflectionPoint:0}),l=p(0),i=p(V.value===!0?0:E()),m=h(()=>"q-layout q-layout--"+(e.container===!0?"containerized":"standard")),v=h(()=>e.container===!1?{minHeight:t.screen.height+"px"}:null),w=h(()=>i.value!==0?{[t.lang.rtl===!0?"left":"right"]:`${i.value}px`}:null),d=h(()=>i.value!==0?{[t.lang.rtl===!0?"right":"left"]:0,[t.lang.rtl===!0?"left":"right"]:`-${i.value}px`,width:`calc(100% + ${i.value}px)`}:null);function z(o){if(e.container===!0||document.qScrollPrevented!==!0){const c={position:o.position.top,direction:o.direction,directionChanged:o.directionChanged,inflectionPoint:o.inflectionPoint.top,delta:o.delta.top};f.value=c,e.onScroll!==void 0&&n("scroll",c)}}function K(o){const{height:c,width:b}=o;let y=!1;a.value!==c&&(y=!0,a.value=c,e.onScrollHeight!==void 0&&n("scroll-height",c),Q()),u.value!==b&&(y=!0,u.value=b),y===!0&&e.onResize!==void 0&&n("resize",o)}function U({height:o}){l.value!==o&&(l.value=o,Q())}function Q(){if(e.container===!0){const o=a.value>l.value?E():0;i.value!==o&&(i.value=o)}}let x;const q={instances:{},view:h(()=>e.view),isContainer:h(()=>e.container),rootRef:r,height:a,containerHeight:l,scrollbarWidth:i,totalWidth:h(()=>u.value+i.value),rows:h(()=>{const o=e.view.toLowerCase().split(" ");return{top:o[0].split(""),middle:o[1].split(""),bottom:o[2].split("")}}),header:S({size:0,offset:0,space:!1}),right:S({size:300,offset:0,space:!1}),footer:S({size:0,offset:0,space:!1}),left:S({size:300,offset:0,space:!1}),scroll:f,animate(){x!==void 0?clearTimeout(x):document.body.classList.add("q-body--layout-animate"),x=setTimeout(()=>{document.body.classList.remove("q-body--layout-animate"),x=void 0},155)},update(o,c,b){q[o][c]=b}};if(A(B,q),E()>0){let b=function(){o=null,c.classList.remove("hide-scrollbar")},y=function(){if(o===null){if(c.scrollHeight>t.screen.height)return;c.classList.add("hide-scrollbar")}else clearTimeout(o);o=setTimeout(b,300)},R=function(H){o!==null&&H==="remove"&&(clearTimeout(o),b()),window[`${H}EventListener`]("resize",y)},o=null;const c=document.body;D(()=>e.container!==!0?"add":"remove",R),e.container!==!0&&R("add"),J(()=>{R("remove")})}return()=>{const o=Z(s.default,[g(le,{onScroll:z}),g(k,{onResize:K})]),c=g("div",{class:m.value,style:v.value,ref:e.container===!0?void 0:r,tabindex:-1},o);return e.container===!0?g("div",{class:"q-layout-container overflow-hidden",ref:r},[g(k,{onResize:U}),g("div",{class:"absolute-full",style:w.value},[g("div",{class:"scroll",style:d.value},[c])])]):c}}});export{ce as Q,se as a,te as b,k as c,E as g};
