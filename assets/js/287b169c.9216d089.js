"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[27],{7149:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>s,contentTitle:()=>l,default:()=>u,frontMatter:()=>i,metadata:()=>a,toc:()=>c});var o=r(4848),t=r(8453);const i={sidebar_position:2,title:"Workflow"},l="What is a workflow?",a={id:"concepts/workflow",title:"Workflow",description:"A specification and container for a topology of connected builders that generate a final data. It has the following meta:",source:"@site/docs/concepts/workflow.md",sourceDirName:"concepts",slug:"/concepts/workflow",permalink:"/polaris/concepts/workflow",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:2,frontMatter:{sidebar_position:2,title:"Workflow"},sidebar:"tutorialSidebar",previous:{title:"Polaris",permalink:"/polaris/concepts/polaris"},next:{title:"Builder",permalink:"/polaris/concepts/builder"}},s={},c=[{value:"Definition",id:"definition",level:3},{value:"Registering a workflow",id:"registering-a-workflow",level:3},{value:"Executing a workflow",id:"executing-a-workflow",level:3}];function d(e){const n={code:"code",h1:"h1",h3:"h3",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,t.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"what-is-a-workflow",children:"What is a workflow?"}),"\n",(0,o.jsx)(n.p,{children:"A specification and container for a topology of connected builders that generate a final data. It has the following meta:"}),"\n",(0,o.jsxs)(n.ul,{children:["\n",(0,o.jsxs)(n.li,{children:[(0,o.jsx)(n.strong,{children:"Builders"})," - List of ",(0,o.jsx)("a",{href:"/polaris/concepts/builder",children:"Builders"})]}),"\n",(0,o.jsxs)(n.li,{children:[(0,o.jsx)(n.strong,{children:"Target Data"})," - The name of the ",(0,o.jsx)("a",{href:"/polaris/concepts/data",children:"Data"})," being generated by this data flow. Once this is produced, the workflow is complete. It can, however, be re-opened by feeding new data."]}),"\n"]}),"\n",(0,o.jsx)(n.h3,{id:"definition",children:"Definition"}),"\n",(0,o.jsxs)(n.p,{children:["Workflows must implement the ",(0,o.jsx)(n.code,{children:"IWorkflow"})," interface."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-go",children:"type IWorkflow interface {\n\tGetWorkflowMeta() WorkflowMeta\n}\n"})}),"\n",(0,o.jsx)(n.p,{children:"Let's take the example of a cab ride workflow. Essentially, for a cab ride workflow, builders (units of work) could be:"}),"\n",(0,o.jsxs)(n.ul,{children:["\n",(0,o.jsx)(n.li,{children:"User initiating a request"}),"\n",(0,o.jsx)(n.li,{children:"Cabbie match"}),"\n",(0,o.jsx)(n.li,{children:"Cabbie reaches source"}),"\n",(0,o.jsx)(n.li,{children:"Ride starts"}),"\n",(0,o.jsx)(n.li,{children:"Cabbie reaches destination"}),"\n",(0,o.jsx)(n.li,{children:"User makes payment"}),"\n",(0,o.jsx)(n.li,{children:"Ride ends"}),"\n"]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-go",children:"type CabRideWorkflow struct {\n}\n\nfunc (cr CabRideWorkflow) GetWorkflowMeta() WorkflowMeta {\n\treturn WorkflowMeta{\n\t\tBuilders: []IBuilder{\n                    UserInitiation{},\n                    CabbieMatching{},\n                    CabbieArrivalAtSource{},\n                    CabDepartureFromSource{},\n                    CabArrivalAtDest{},\n                    UserPayment{},\n                    RideEnds{}\n\t\t},\n\t\tTargetData: WorkflowTerminated{},\n\t}\n}\n"})}),"\n",(0,o.jsxs)(n.p,{children:["You don't have to sequentially define the builders in order of execution. Polaris will figure it out. However, ",(0,o.jsx)(n.strong,{children:"you should if you can. It helps readability."})]}),"\n",(0,o.jsx)(n.h3,{id:"registering-a-workflow",children:"Registering a workflow"}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-go",children:"polaris.RegisterWorkflow(workflowKey, workflow)\n"})}),"\n",(0,o.jsx)(n.h3,{id:"executing-a-workflow",children:"Executing a workflow"}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-go",children:'executor := polaris.Executor{\n  Before: func(builder reflect.Type, delta []IData) {\n    fmt.Printf("Builder %s is about to be run with new data %v\\n", builder, delta)\n  }\n  After: func(builder reflect.Type, produced IData) {\n    fmt.Printf("Builder %s produced %s\\n", builder, produced)\n  }\n}\n// sequentially execute builders\nresponse, err := executor.Sequential(workflowKey, workflowId, dataDelta)\n\n// concurrently execute builders (does not guarantee parallelism!)\nresponse, err := executor.Parallel(workflowKey, workflowId, dataDelta)\n'})})]})}function u(e={}){const{wrapper:n}={...(0,t.R)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(d,{...e})}):d(e)}},8453:(e,n,r)=>{r.d(n,{R:()=>l,x:()=>a});var o=r(6540);const t={},i=o.createContext(t);function l(e){const n=o.useContext(i);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:l(e.components),o.createElement(i.Provider,{value:n},e.children)}}}]);