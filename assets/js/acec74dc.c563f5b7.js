"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[1900],{3905:(e,n,t)=>{t.d(n,{Zo:()=>c,kt:()=>f});var r=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function l(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var p=r.createContext({}),s=function(e){var n=r.useContext(p),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},c=function(e){var n=s(e.components);return r.createElement(p.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},u=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,a=e.originalType,p=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),u=s(t),f=i,m=u["".concat(p,".").concat(f)]||u[f]||d[f]||a;return t?r.createElement(m,o(o({ref:n},c),{},{components:t})):r.createElement(m,o({ref:n},c))}));function f(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var a=t.length,o=new Array(a);o[0]=u;var l={};for(var p in n)hasOwnProperty.call(n,p)&&(l[p]=n[p]);l.originalType=e,l.mdxType="string"==typeof e?e:i,o[1]=l;for(var s=2;s<a;s++)o[s]=t[s];return r.createElement.apply(null,o)}return r.createElement.apply(null,t)}u.displayName="MDXCreateElement"},9403:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>p,contentTitle:()=>o,default:()=>d,frontMatter:()=>a,metadata:()=>l,toc:()=>s});var r=t(7462),i=(t(7294),t(3905));const a={},o="Fitting",l={unversionedId:"plotting/fitting",id:"plotting/fitting",title:"Fitting",description:"Fit",source:"@site/docs/plotting/fitting.md",sourceDirName:"plotting",slug:"/plotting/fitting",permalink:"/docs/plotting/fitting",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Plotting Options",permalink:"/docs/plotting/plotting-options"},next:{title:"Introduction",permalink:"/docs/groupby/introduction"}},p={},s=[{value:"Fit",id:"fit",level:2}],c={toc:s};function d(e){let{components:n,...t}=e;return(0,i.kt)("wrapper",(0,r.Z)({},c,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"fitting"},"Fitting"),(0,i.kt)("h2",{id:"fit"},"Fit"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func Fit(ff string, pd PlotData, viaOpts ...GnuplotOpt) error\n")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"Fit")," fits a user-defined function ",(0,i.kt)("inlineCode",{parentName:"p"},"ff")," to data given in ",(0,i.kt)("inlineCode",{parentName:"p"},"PlotData")," ",(0,i.kt)("inlineCode",{parentName:"p"},"pd"),", and prints out the results."),(0,i.kt)("p",null,"Pass options such as ",(0,i.kt)("inlineCode",{parentName:"p"},"using")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"pd"),", but ",(0,i.kt)("inlineCode",{parentName:"p"},"via")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"viaOpts"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "explike-data.csv"), []string{"index"})\nif err != nil {\n    fmt.Println(err)\n}\n\npd := gambas.PlotData{\n    Df:       &df,\n    Columns:  []string{"index", "data", "error"},\n    Function: "",\n    Opts:     []gambas.GnuplotOpt{gambas.Using("0:1:2 yerrors")},\n}\n\nerr = gambas.Fit("a*exp(-b*x)", pd, gambas.Via("a,b"))\nif err != nil {\n    fmt.Println(err)\n}\n')),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"// explike-data.csv\nindex,data,error\n0,1.129778763,0.1\n1,0.953320457,0.081873075\n2,0.732501277,0.067032005\n3,0.561141921,0.054881164\n4,0.434710291,0.044932896\n5,0.349139904,0.036787944\n6,0.308242881,0.030119421\n7,0.26402811,0.024659696\n8,0.18692481,0.020189652\n9,0.159507023,0.016529889\n10,0.15148364,0.013533528\n11,0.086788709,0.011080316\n12,0.094842665,0.009071795\n13,0.081597097,0.007427358\n14,0.052315262,0.006081006\n15,0.05010193,0.004978707\n16,0.043722877,0.00407622\n17,0.036272859,0.003337327\n18,0.025103659,0.002732372\n19,0.019053795,0.002237077\n20,0.018988513,0.001831564\n")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"iter      chisq       delta/lim  lambda   a             b            \n   0 3.2009094356e+06   0.00e+00  1.81e-01    1.000000e+00   1.000000e+00\n   * -nan              -nan       1.81e+00   -3.218104e+00  -2.204204e+01\n   * 7.4311004197e+27   1.00e+05  1.81e+01    1.213369e+00  -1.394988e+00\n   1 3.2009089058e+06  -1.66e-02  1.81e+00    1.004569e+00   9.730191e-01\niter      chisq       delta/lim  lambda   a             b            \n\nAfter 1 iterations the fit converged.\nfinal sum of squares of residuals : 3.20091e+06\nrel. change during last iteration : -1.65518e-07\n\ndegrees of freedom    (FIT_NDF)                        : 19\nrms of residuals      (FIT_STDFIT) = sqrt(WSSR/ndf)    : 410.45\nvariance of residuals (reduced chisquare) = WSSR/ndf   : 168469\np-value of the Chisq distribution (FIT_P)              : 0\n\nFinal set of parameters            Asymptotic Standard Error\n=======================            ==========================\na               = 1.00457          +/- 452.4        (4.504e+04%)\nb               = 0.973019         +/- 673.1        (6.917e+04%)\n\ncorrelation matrix of the fit parameters:\n                a      b      \na               1.000 \nb               0.411  1.000\n")))}d.isMDXComponent=!0}}]);