import{d as C,k as T,u as R,b as V,r as d,O as D,a as M,f as m,c as t,t as a,e as n,h as r,P as p,K as O,o as l,F as P,x as B,v as _,I as v,y as f,z as h,H as E,J as A,m as F,_ as Q}from"./index-ydEeV_dL.js";import{g as Y,a as j}from"./serverInfo-SBK1Jza-.js";import{S as q,a as z}from"./SearchUtil-pmlYZKt1.js";import{c as H}from"./app-xsf0myXG.js";import"./request-TC7KOJ9l.js";const J={class:"__container_app_service"},K={class:"statistic-icon-big"},L=C({__name:"service",setup($){T(e=>({"383e04d6":r(p)+"22","03b8a93b":r(p)}));const g=R(),b=V();let s=d({info:{},report:{}}),x=d({info:{}});D(async()=>{o.tableStyle={scrollX:"100",scrollY:"calc(100vh - 600px)"};let e=(await Y({})).data;x.info=(await j({})).data,s.info=e,s.report={providers:{icon:"carbon:branch",value:s.info.providers},consumers:{icon:"mdi:merge",value:s.info.consumers}}});const y=[{title:"idx",key:"idx"},{title:"服务",dataIndex:"serviceName",key:"serviceName",sorter:!0,width:"30%"},{title:"接口数",dataIndex:"interfaceNum",key:"interfaceNum",sorter:!0,width:"10%"},{title:"近 1min QPS",dataIndex:"avgQPS",key:"avgQPS",sorter:!0,width:"15%"},{title:"近 1min RT",dataIndex:"avgRT",key:"avgRT",sorter:!0,width:"15%"},{title:"近 1min 请求总量",dataIndex:"requestTotal",key:"requestTotal",sorter:!0,width:"15%"}],I=M(()=>{var e;return(e=g.params)==null?void 0:e.pathId}),o=d(new q([{label:"",param:"side",defaultValue:"provider",dict:[{label:"providers",value:"provider"},{label:"consumers",value:"consumer"}],dictType:"BUTTON"},{label:"serviceName",param:"serviceName"},{label:"",param:"appName",defaultValue:I}],H,y,{pageSize:4},!0));o.onSearch();const N=e=>{b.push("/resources/services/detail/"+e)};return O(F.SEARCH_DOMAIN,o),(e,U)=>{const k=n("a-statistic"),u=n("a-flex"),w=n("a-card"),S=n("a-button");return l(),m("div",J,[t(u,{wrap:"wrap",gap:"small",vertical:!1,justify:"start",align:"left"},{default:a(()=>[(l(!0),m(P,null,B(r(s).report,(i,c)=>(l(),_(w,{class:"statistic-card"},{default:a(()=>[t(u,{gap:"middle",vertical:!1,justify:"space-between",align:"center"},{default:a(()=>[t(k,{value:i.value,class:"statistic"},{prefix:a(()=>[t(r(v),{class:"statistic-icon",icon:"solar:target-line-duotone"})]),title:a(()=>[f(h(e.$t(c.toString())),1)]),_:2},1032,["value"]),E("div",K,[t(r(v),{icon:i.icon},null,8,["icon"])])]),_:2},1024)]),_:2},1024))),256))]),_:1}),t(z,{"search-domain":o},{bodyCell:a(({column:i,text:c})=>[i.dataIndex==="serviceName"?(l(),_(S,{key:0,type:"link",onClick:X=>N(c)},{default:a(()=>[f(h(c),1)]),_:2},1032,["onClick"])):A("",!0)]),_:1},8,["search-domain"])])}}}),te=Q(L,[["__scopeId","data-v-95988178"]]);export{te as default};
