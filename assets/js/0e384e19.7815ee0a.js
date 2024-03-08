"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[976],{1512:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>r,default:()=>h,frontMatter:()=>i,metadata:()=>a,toc:()=>d});var o=t(4848),s=t(8453);const i={sidebar_position:1},r="Introduction",a={id:"intro",title:"Introduction",description:"Light themed headerDark themed header",source:"@site/docs/intro.md",sourceDirName:".",slug:"/intro",permalink:"/polaris/docs/intro",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",next:{title:"Getting Started",permalink:"/polaris/docs/getting-started"}},l={},d=[{value:"What is a workflow?",id:"what-is-a-workflow",level:2},{value:"Use cases",id:"use-cases",level:2},{value:"Limitations",id:"limitations",level:2},{value:"How does the framework perform at scale?",id:"how-does-the-framework-perform-at-scale",level:2}];function c(e){const n={h1:"h1",h2:"h2",img:"img",li:"li",ol:"ol",p:"p",...(0,s.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"introduction",children:"Introduction"}),"\n",(0,o.jsxs)(n.p,{children:[(0,o.jsx)(n.img,{alt:"Light themed header",src:t(8780).A+"#gh-light-mode-only",width:"1280",height:"640"}),(0,o.jsx)(n.img,{alt:"Dark themed header",src:t(2484).A+"#gh-dark-mode-only",width:"1280",height:"640"})]}),"\n",(0,o.jsx)(n.h2,{id:"what-is-a-workflow",children:"What is a workflow?"}),"\n",(0,o.jsx)(n.p,{children:"A workflow is a series of multiple steps, and can often be long running."}),"\n",(0,o.jsx)(n.p,{children:"Workflow orchestrators (like Polaris) help break workflows down into chunks, so they can be processed asynchronously. If a workflow is waiting on an event, state is stored and the logical execution of the workflow is paused (meaning, the CPU is free to move on with other tasks). When the event is received, state is recovered and execution is resumed."}),"\n",(0,o.jsx)(n.h2,{id:"use-cases",children:"Use cases"}),"\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsx)(n.li,{children:"You have multi-step workflow executions where each step is dependent on data generated from previous steps."}),"\n",(0,o.jsx)(n.li,{children:"Executions can span one request scope or multiple scopes."}),"\n",(0,o.jsx)(n.li,{children:"Your systems works with reusable components that can be combined in different ways to generate different end-results."}),"\n",(0,o.jsx)(n.li,{children:"Your workflows can pause, resume or even restart from the beginning."}),"\n"]}),"\n",(0,o.jsx)(n.h2,{id:"limitations",children:"Limitations"}),"\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsxs)(n.li,{children:["Workflow versioning is tricky to implement:","\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsx)(n.li,{children:"Unless you can afford a 100% downtime ensuring all active workflows move into a terminal state, deploying new code requires ensuring backward compatibility."}),"\n",(0,o.jsx)(n.li,{children:"What this means is - you'll need to a deploy a version of code that is backward compatible for older non terminal workflows while newer ones will execute on the new code."}),"\n",(0,o.jsx)(n.li,{children:"Once the older workflows have completed, a deployment to clean up stale code will be required."}),"\n"]}),"\n"]}),"\n",(0,o.jsxs)(n.li,{children:["The level of abstraction is lower in this framework compared to Cadence, Conductor:","\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsx)(n.li,{children:"Workflows can be made fault oblivious if there is an external (reliable) service giving callbacks per workflow id."}),"\n",(0,o.jsx)(n.li,{children:"Instrumentation can be set up by adding your custom code to push events via listeners."}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,o.jsx)(n.h2,{id:"how-does-the-framework-perform-at-scale",children:"How does the framework perform at scale?"}),"\n",(0,o.jsx)(n.p,{children:"The framework itself has extremely low overhead. Since execution graphs are generated pre-runtime, all the orchestrator will do at runtime is use the graph and available data to run whichever builders can be run."})]})}function h(e={}){const{wrapper:n}={...(0,s.R)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(c,{...e})}):c(e)}},2484:(e,n,t)=>{t.d(n,{A:()=>o});const o=t.p+"assets/images/polaris-header-dark-fa2838a623ff57751307ab6c4bd8c637.svg"},8780:(e,n,t)=>{t.d(n,{A:()=>o});const o=t.p+"assets/images/polaris-header-light-6b9b55f756cc7fdeb57d4985c182b572.svg"},8453:(e,n,t)=>{t.d(n,{R:()=>r,x:()=>a});var o=t(6540);const s={},i=o.createContext(s);function r(e){const n=o.useContext(i);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:r(e.components),o.createElement(i.Provider,{value:n},e.children)}}}]);