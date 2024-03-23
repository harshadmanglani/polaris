"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[194],{7197:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>o,default:()=>u,frontMatter:()=>i,metadata:()=>a,toc:()=>l});var s=t(4848),r=t(8453);const i={sidebar_position:3},o="Builder",a={id:"concepts/builder",title:"Builder",description:"An actor that consumes a bunch of data and produces another data. It has the following meta associated with it:",source:"@site/docs/concepts/builder.md",sourceDirName:"concepts",slug:"/concepts/builder",permalink:"/polaris/concepts/builder",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:3,frontMatter:{sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Workflow",permalink:"/polaris/concepts/workflow"},next:{title:"Data",permalink:"/polaris/concepts/data"}},c={},l=[];function d(e){const n={code:"code",h1:"h1",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,r.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.h1,{id:"builder",children:"Builder"}),"\n",(0,s.jsx)(n.p,{children:"An actor that consumes a bunch of data and produces another data. It has the following meta associated with it:"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"Name"})," - Name of the builder."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"Consumes"})," - A set of data that the builder consumes."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"Produces"})," - Data that the builder produces."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"Optionals"})," - Data that the builder can optionally consume; one possible use case for this: if a builder wants to be re-run on demand with the same set of consumable data already present, add an optional data in the Builder and restart the workflow by passing an instance of the optional data in the data-delta"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"Access"})," - Data that the builder will just access and has no effect on the sequence of execution of builders"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.strong,{children:"BuilderContext"})," - A wrapper to access the data given to that builder to process."]}),"\n"]}),"\n",(0,s.jsxs)(n.p,{children:["A Builder is a unit of work in the workflow. Builders must implement the ",(0,s.jsx)(n.code,{children:"IBuilder"})," interface."]}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-go",children:"type IBuilder interface {\n  GetBuilderInfo() BuilderInfo\n  Process(BuilderContext) IData\n}\n"})}),"\n",(0,s.jsx)(n.p,{children:"Following the same example, for the first unit of work in a cab ride workflow:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-go",children:"var database Database\nvar cabbieHttpClient CabbieHttpClient \n\ntype UserInitiation struct {\n}\n\nfunc (uI UserInitiation) GetBuilderInfo() BuilderInfo {\n  return BuilderInfo{\n    Consumes: []IData{\n        UserInitiationRequest{},\n    },\n    Produces:  UserInitiationResponse{},\n    Optionals: nil,\n    Accesses:  nil,\n  }\n}\n\nfunc (uI UserInitiation) Process(context BuilderContext) IData {\n  userInitReq := context.Get(UserInitiationRequest{})\n  // save the request in a database (different from Polaris storing workflows in `IDataStore`)\n  database.save(userInitReq)\n\n  // call another service to place a request, and wait for the response\n  cabbieResponse := cabbieHttpClient.request(RideRequest{\n    userId: userInitReq.userId,\n    source: userInitReq.source,\n    dest: userInitReq.dest\n  })\n\n  // once done, return the `Produces` of the data\n  return UserInitiationResponse{\n    success: true,\n    etaForCabbie: cabbieResponse.eta\n  }\n}\n"})})]})}function u(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,s.jsx)(n,{...e,children:(0,s.jsx)(d,{...e})}):d(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>o,x:()=>a});var s=t(6540);const r={},i=s.createContext(r);function o(e){const n=s.useContext(i);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:o(e.components),s.createElement(i.Provider,{value:n},e.children)}}}]);