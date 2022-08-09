"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[5586],{3905:(e,n,t)=>{t.d(n,{Zo:()=>m,kt:()=>u});var r=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var o=r.createContext({}),c=function(e){var n=r.useContext(o),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},m=function(e){var n=c(e.components);return r.createElement(o.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},p=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,i=e.originalType,o=e.parentName,m=s(e,["components","mdxType","originalType","parentName"]),p=c(t),u=a,f=p["".concat(o,".").concat(u)]||p[u]||d[u]||i;return t?r.createElement(f,l(l({ref:n},m),{},{components:t})):r.createElement(f,l({ref:n},m))}));function u(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var i=t.length,l=new Array(i);l[0]=p;var s={};for(var o in n)hasOwnProperty.call(n,o)&&(s[o]=n[o]);s.originalType=e,s.mdxType="string"==typeof e?e:a,l[1]=s;for(var c=2;c<i;c++)l[c]=t[c];return r.createElement.apply(null,l)}return r.createElement.apply(null,t)}p.displayName="MDXCreateElement"},1382:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>o,contentTitle:()=>l,default:()=>d,frontMatter:()=>i,metadata:()=>s,toc:()=>c});var r=t(7462),a=(t(7294),t(3905));const i={},l="Statistics Functions",s={unversionedId:"statistics/statistics-functions",id:"statistics/statistics-functions",title:"Statistics Functions",description:"The data used in the example neo_v2.csv is NASA's list of Nearest Earth Objects, sourced from Kaggle.",source:"@site/docs/statistics/statistics-functions.md",sourceDirName:"statistics",slug:"/statistics/statistics-functions",permalink:"/gambas/docs/statistics/statistics-functions",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Introduction",permalink:"/gambas/docs/statistics/introduction"}},o={},c=[{value:"Count",id:"count",level:2},{value:"Mean",id:"mean",level:2},{value:"Median",id:"median",level:2},{value:"Std",id:"std",level:2},{value:"Min",id:"min",level:2},{value:"Max",id:"max",level:2},{value:"Q1",id:"q1",level:2},{value:"Q2",id:"q2",level:2},{value:"Q3",id:"q3",level:2}],m={toc:c};function d(e){let{components:n,...t}=e;return(0,a.kt)("wrapper",(0,r.Z)({},m,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"statistics-functions"},"Statistics Functions"),(0,a.kt)("p",null,"The data used in the example ",(0,a.kt)("inlineCode",{parentName:"p"},"neo_v2.csv")," is NASA's list of Nearest Earth Objects, sourced from ",(0,a.kt)("a",{parentName:"p",href:"https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects"},"Kaggle"),"."),(0,a.kt)("h2",{id:"count"},"Count"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Count(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Count")," counts the number of non-NaN elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Count(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Count\n90836\n<nil>\n")),(0,a.kt)("h2",{id:"mean"},"Mean"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Mean(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Mean")," returns the mean of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Mean(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Mean\n0.127\n<nil>\n")),(0,a.kt)("h2",{id:"median"},"Median"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Median(dataset []interface{}) StatsResult \n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Median")," returns the median of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Median(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Median\n0.048\n<nil>\n")),(0,a.kt)("h2",{id:"std"},"Std"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Std(dataset []interface{}) StatsResult \n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Std")," returns the sample standard deviation of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Std(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Std\n0.299\n<nil>\n")),(0,a.kt)("h2",{id:"min"},"Min"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Min(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Min")," returns the smallest element in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Min(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Min\n0.0006089126\n<nil>\n")),(0,a.kt)("h2",{id:"max"},"Max"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Max(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Max")," returns the largest element is a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Max(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Max\n37.8926498379\n<nil>\n")),(0,a.kt)("h2",{id:"q1"},"Q1"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Q1(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Q1")," returns the lower quartile (25%) of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),". This does not include the median during calculation."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Q1(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Q1\n0.0192555078\n<nil>\n")),(0,a.kt)("h2",{id:"q2"},"Q2"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Q2(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Q2")," returns the middle quartile (50%) of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),". This accomplishes the same thing as ",(0,a.kt)("inlineCode",{parentName:"p"},"Median"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Q2(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Q2\n0.048\n<nil>\n")),(0,a.kt)("h2",{id:"q3"},"Q3"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Q3(dataset []interface{}) StatsResult\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Q3")," returns the upper quartile (75%) of the elements in a ",(0,a.kt)("inlineCode",{parentName:"p"},"dataset"),". This does not include the median during calculation."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})\nif err != nil {\n    fmt.Println(err)\n}\n\ncol1, err := df.LocColsItems("est_diameter_min")\nif err != nil {\n    fmt.Println(err)\n}\n\nres := gambas.Q3(col1[0])\nfmt.Println(res.UsedFunc)\nfmt.Println(res.Result)\nfmt.Println(res.Err)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"Q3\n0.1434019235\n<nil>\n")))}d.isMDXComponent=!0}}]);