import{C as w}from"./ConfigPage-MDRitquN.js";import{d as _,u as h,r as u,O as b,f as I,c,t as l,h as D,e as p,o as k,_ as F}from"./index-0c4Jk4Xf.js";import{u as y,b as S,c as L,d as P}from"./instance-a9aKQOdb.js";import"./request-fr6f-AQz.js";const C={class:"__container_app_config"},N=_({__name:"configuration",setup(v){const s=h();let i=u({list:[{title:"instanceDomain.operatorLog",key:"log",form:{logFlag:!1},submit:a=>new Promise(e=>{e(d(a==null?void 0:a.logFlag))}),reset(a){a.logFlag=!1}},{title:"instanceDomain.flowDisabled",form:{flowDisabledFlag:!1},key:"flowDisabled",submit:a=>new Promise(e=>{e(m(a==null?void 0:a.flowDisabledFlag))}),reset(a){a.logFlag=!1}}],current:[0]});const f=async()=>{var e,o;const a=await L((e=s.params)==null?void 0:e.pathId,(o=s.params)==null?void 0:o.appName);(a==null?void 0:a.code)==200&&i.list.forEach(t=>{if(t.key==="log"){t.form.logFlag=a.data.operatorLog;return}})},d=async a=>{var o,t;const e=await y((o=s.params)==null?void 0:o.pathId,(t=s.params)==null?void 0:t.appName,a);(e==null?void 0:e.code)==200&&await f()},g=async()=>{var e,o;const a=await P((e=s.params)==null?void 0:e.pathId,(o=s.params)==null?void 0:o.appName);(a==null?void 0:a.code)==200&&i.list.forEach(t=>{t.key==="flowDisabled"&&(t.form.flowDisabledFlag=a.data.trafficDisable)})},m=async a=>{var o,t;const e=await S((o=s.params)==null?void 0:o.pathId,(t=s.params)==null?void 0:t.appName,a);console.log(e)};return b(()=>{console.log(333),f(),g()}),(a,e)=>{const o=p("a-switch"),t=p("a-form-item");return k(),I("div",C,[c(w,{options:D(i)},{form_log:l(({current:n})=>[c(t,{label:a.$t("instanceDomain.operatorLog"),name:"logFlag"},{default:l(()=>[c(o,{checked:n.form.logFlag,"onUpdate:checked":r=>n.form.logFlag=r},null,8,["checked","onUpdate:checked"])]),_:2},1032,["label"])]),form_flowDisabled:l(({current:n})=>[c(t,{label:a.$t("instanceDomain.flowDisabled"),name:"flowDisabledFlag"},{default:l(()=>[c(o,{checked:n.form.flowDisabledFlag,"onUpdate:checked":r=>n.form.flowDisabledFlag=r},null,8,["checked","onUpdate:checked"])]),_:2},1032,["label"])]),_:1},8,["options"])])}}}),x=F(N,[["__scopeId","data-v-e36478a9"]]);export{x as default};
