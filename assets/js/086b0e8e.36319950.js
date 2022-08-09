"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[3383],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),u=c(n),m=a,f=u["".concat(s,".").concat(m)]||u[m]||d[m]||i;return n?r.createElement(f,o(o({ref:t},p),{},{components:n})):r.createElement(f,o({ref:t},p))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=u;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var c=2;c<i;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},2122:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>o,default:()=>d,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var r=n(7462),a=(n(7294),n(3905));const i={},o="Introduction",l={unversionedId:"dataframe/introduction",id:"dataframe/introduction",title:"Introduction",description:"A DataFrame is one of the fundamental types used in gambas. It represents a tabular 2D data. DataFrame is a column-based data, meaning that it is made of several Series's.",source:"@site/docs/dataframe/introduction.md",sourceDirName:"dataframe",slug:"/dataframe/introduction",permalink:"/gambas/docs/dataframe/introduction",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Sorting",permalink:"/gambas/docs/series/sorting"},next:{title:"Printing",permalink:"/gambas/docs/dataframe/printing"}},s={},c=[],p={toc:c};function d(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"introduction"},"Introduction"),(0,a.kt)("p",null,"A ",(0,a.kt)("inlineCode",{parentName:"p"},"DataFrame")," is one of the fundamental types used in ",(0,a.kt)("inlineCode",{parentName:"p"},"gambas"),". It represents a tabular 2D data. ",(0,a.kt)("inlineCode",{parentName:"p"},"DataFrame")," is a column-based data, meaning that it is made of several ",(0,a.kt)("inlineCode",{parentName:"p"},"Series"),"'s."),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"DataFrame")," type is defined as per below:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type DataFrame struct {\n    series  []Series\n    index   IndexData\n    columns []string\n}\n\n")),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"series")," holds the ",(0,a.kt)("inlineCode",{parentName:"li"},"Series")," objects that make up the ",(0,a.kt)("inlineCode",{parentName:"li"},"DataFrame")," object."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"index")," holds the index of the ",(0,a.kt)("inlineCode",{parentName:"li"},"DataFrame")," object."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"columns")," is the names of columns in the ",(0,a.kt)("inlineCode",{parentName:"li"},"DataFrame")," object.")),(0,a.kt)("p",null,"The fields are private, but ",(0,a.kt)("inlineCode",{parentName:"p"},"gambas")," provides accessors to get these fields."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (df DataFrame) Series() []Series {\n    return df.series\n}\n\nfunc (df DataFrame) Index() IndexData {\n    return df.index\n}\nfunc (df DataFrame) Columns() []string {\n    return df.columns\n}\n")))}d.isMDXComponent=!0}}]);