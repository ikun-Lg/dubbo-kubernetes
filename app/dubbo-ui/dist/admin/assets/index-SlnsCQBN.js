import{d as D,k as T,r as N,O as E,f as u,c as C,t as m,s as f,h as _,P as O,K as R,e as v,o as s,Q as A,H as P,I as V,y as p,z as d,J as c,v as y,Y as L,Z as Y,F as g,x,m as $,_ as B}from"./index-QTFC1BX7.js";import{s as M}from"./instance-l-z-gAsN.js";import{S as F,a as H,s as o}from"./SearchUtil-gzPkvecP.js";import{f as U}from"./DateUtil-9OgzTl0-.js";import{q as b}from"./promQuery-PGswEixs.js";import{b as q}from"./ByteUtil-YdHlSEeW.js";import"./request-bLCExP-g.js";const J={class:"__container_resources_application_index"},K=["onClick"],z=D({__name:"index",setup(G){T(t=>({"66e1b67c":_(O)}));let S=[{title:"instanceDomain.instanceIP",key:"ip",dataIndex:"ip",sorter:(t,e)=>o(t.ip,e.ip),width:200},{title:"instanceDomain.instanceName",key:"name",dataIndex:"name",sorter:(t,e)=>o(t.name,e.name),width:140},{title:"instanceDomain.deployState",key:"deployState",dataIndex:"deployState",width:120,sorter:(t,e)=>o(t.deployState,e.deployState)},{title:"instanceDomain.deployCluster",key:"deployCluster",dataIndex:"deployCluster",sorter:(t,e)=>o(t.deployCluster,e.deployCluster),width:120},{title:"instanceDomain.registerState",key:"registerState",dataIndex:"registerState",sorter:(t,e)=>o(t.registerState,e.registerState),width:120},{title:"instanceDomain.registerCluster",key:"registerClusters",dataIndex:"registerClusters",sorter:(t,e)=>o(t.registerClusters,e.registerClusters),width:140},{title:"instanceDomain.CPU",key:"cpu",dataIndex:"cpu",sorter:(t,e)=>o(t.cpu,e.cpu),width:140},{title:"instanceDomain.memory",key:"memory",dataIndex:"memory",sorter:(t,e)=>o(t.memory,e.memory),width:100},{title:"instanceDomain.startTime_k8s",key:"startTime_k8s",dataIndex:"startTime",sorter:(t,e)=>o(t.startTime,e.startTime),width:200},{title:"instanceDomain.registerTime",key:"registerTime",dataIndex:"registerTime",sorter:(t,e)=>o(t.registerTime,e.registerTime),width:200},{title:"instanceDomain.labels",key:"labels",dataIndex:"labels",width:800}];function w(t){return M(t).then(async e=>{var a;let l=(a=e==null?void 0:e.data)==null?void 0:a.list;try{for(let i of l){let I=i.ip.split(":")[0],r=await b(`sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{container!=""}) by (pod) * on (pod) group_left(pod_ip)
        kube_pod_info{pod_ip="${I}"}`),n=await b(`sum(container_memory_working_set_bytes{container!=""}) by (pod)
* on (pod) group_left(pod_ip)
kube_pod_info{pod_ip="${I}"}`);i.cpu=f.isNumber(r)?r.toFixed(3)+"u":r,i.memory=f.isNumber(n)?q(n):n}}catch(i){console.error(i)}return e})}const k=N(new F([{label:"appName",param:"keywords",placeholder:"typeAppName",style:{width:"200px"}}],w,S));return E(()=>{k.tableStyle={scrollY:"calc(100vh - 200px)"},k.onSearch()}),R($.SEARCH_DOMAIN,k),(t,e)=>{const l=v("a-tag");return s(),u("div",J,[C(H,{"search-domain":k},{bodyCell:m(({text:a,record:i,index:I,column:r})=>[r.dataIndex==="ip"?(s(),u("span",{key:0,class:"app-link",onClick:n=>_(A).replace(`detail/${i[r.key]}/${i.name}`)},[P("b",null,[C(_(V),{style:{"margin-bottom":"-2px"},icon:"material-symbols:attach-file-rounded"}),p(" "+d(a),1)])],8,K)):c("",!0),r.dataIndex==="deployState"?(s(),y(l,{key:1,color:_(L)[a.toUpperCase()]},{default:m(()=>[p(d(a),1)]),_:2},1032,["color"])):c("",!0),r.dataIndex==="deployCluster"?(s(),y(l,{key:2,color:"grey"},{default:m(()=>[p(d(a),1)]),_:2},1024)):c("",!0),r.dataIndex==="registerState"?(s(),y(l,{key:3,color:_(Y)[a.toUpperCase()]},{default:m(()=>[p(d(a),1)]),_:2},1032,["color"])):c("",!0),r.dataIndex==="registerClusters"?(s(!0),u(g,{key:4},x(a,n=>(s(),y(l,{color:"grey"},{default:m(()=>[p(d(n),1)]),_:2},1024))),256)):c("",!0),r.dataIndex==="registerTime"?(s(),u(g,{key:5},[p(d(_(U)(a)),1)],64)):c("",!0),r.dataIndex==="labels"?(s(!0),u(g,{key:6},x(a,(n,h)=>(s(),y(l,{key:h},{default:m(()=>[p(d(h)+":"+d(n),1)]),_:2},1024))),128)):c("",!0)]),_:1},8,["search-domain"])])}}}),ae=B(z,[["__scopeId","data-v-b69b68d9"]]);export{ae as default};
