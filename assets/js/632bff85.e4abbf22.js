"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[4167],{3905:(e,n,r)=>{r.d(n,{Zo:()=>s,kt:()=>u});var a=r(7294);function t(e,n,r){return n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r,e}function o(e,n){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),r.push.apply(r,a)}return r}function l(e){for(var n=1;n<arguments.length;n++){var r=null!=arguments[n]?arguments[n]:{};n%2?o(Object(r),!0).forEach((function(n){t(e,n,r[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(r,n))}))}return e}function i(e,n){if(null==e)return{};var r,a,t=function(e,n){if(null==e)return{};var r,a,t={},o=Object.keys(e);for(a=0;a<o.length;a++)r=o[a],n.indexOf(r)>=0||(t[r]=e[r]);return t}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)r=o[a],n.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(t[r]=e[r])}return t}var c=a.createContext({}),p=function(e){var n=a.useContext(c),r=n;return e&&(r="function"==typeof e?e(n):l(l({},n),e)),r},s=function(e){var n=p(e.components);return a.createElement(c.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},m=a.forwardRef((function(e,n){var r=e.components,t=e.mdxType,o=e.originalType,c=e.parentName,s=i(e,["components","mdxType","originalType","parentName"]),m=p(r),u=t,f=m["".concat(c,".").concat(u)]||m[u]||d[u]||o;return r?a.createElement(f,l(l({ref:n},s),{},{components:r})):a.createElement(f,l({ref:n},s))}));function u(e,n){var r=arguments,t=n&&n.mdxType;if("string"==typeof e||t){var o=r.length,l=new Array(o);l[0]=m;var i={};for(var c in n)hasOwnProperty.call(n,c)&&(i[c]=n[c]);i.originalType=e,i.mdxType="string"==typeof e?e:t,l[1]=i;for(var p=2;p<o;p++)l[p]=r[p];return a.createElement.apply(null,l)}return a.createElement.apply(null,r)}m.displayName="MDXCreateElement"},5065:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>i,toc:()=>p});var a=r(7462),t=(r(7294),r(3905));const o={},l="Arithmetic Operations",i={unversionedId:"dataframe/arithmetic-operations",id:"dataframe/arithmetic-operations",title:"Arithmetic Operations",description:"You can apply column-wide arithmetic operations to each of your DataFrame object's columns.",source:"@site/docs/dataframe/arithmetic-operations.md",sourceDirName:"dataframe",slug:"/dataframe/arithmetic-operations",permalink:"/gambas/docs/dataframe/arithmetic-operations",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Indexing",permalink:"/gambas/docs/dataframe/indexing"},next:{title:"Editing Properties",permalink:"/gambas/docs/dataframe/editing-properties"}},c={},p=[{value:"ColAdd",id:"coladd",level:2},{value:"ColSub",id:"colsub",level:2},{value:"ColMul",id:"colmul",level:2},{value:"ColDiv",id:"coldiv",level:2},{value:"ColMod",id:"colmod",level:2},{value:"ColGt",id:"colgt",level:2},{value:"ColLt",id:"collt",level:2},{value:"ColEq",id:"coleq",level:2}],s={toc:p};function d(e){let{components:n,...r}=e;return(0,t.kt)("wrapper",(0,a.Z)({},s,r,{components:n,mdxType:"MDXLayout"}),(0,t.kt)("h1",{id:"arithmetic-operations"},"Arithmetic Operations"),(0,t.kt)("p",null,"You can apply column-wide arithmetic operations to each of your ",(0,t.kt)("inlineCode",{parentName:"p"},"DataFrame")," object's columns."),(0,t.kt)("p",null,"The data used in the example ",(0,t.kt)("inlineCode",{parentName:"p"},"2019.csv")," is UN's 2019 World Happiness Report, sourced from ",(0,t.kt)("a",{parentName:"p",href:"https://www.kaggle.com/datasets/unsdsn/world-happiness"},"Kaggle"),"."),(0,t.kt)("h2",{id:"coladd"},"ColAdd"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColAdd(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColAdd")," adds the given value to each element in the specified column."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColAdd("GDP per capita", 2.5)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    3.84              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      3.883             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    3.988             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    3.88              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    3.896             1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"colsub"},"ColSub"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColSub(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColSub")," subtracts the given value from each element in the specified column."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColSub("GDP per capita", 0.5)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    0.8400000000000001    1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      0.883                 1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    0.988                 1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    0.8799999999999999    1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    0.8959999999999999    1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"colmul"},"ColMul"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColMul(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColMul")," multiplies each element in the specified column by the given value."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColMul("GDP per capita", 1.5)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    2.0100000000000002    1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      2.0745                1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    2.232                 1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    2.07                  1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    2.094                 1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"coldiv"},"ColDiv"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColDiv(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColDiv")," divides each element in the specified column by the given value."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColDiv("GDP per capita", 1.5)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    0.8933333333333334    1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      0.922                 1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    0.992                 1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    0.9199999999999999    1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    0.9306666666666666    1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"colmod"},"ColMod"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColMod(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColMod")," applies modulus calculations on each element in the specified column, returning the remainder."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColMod("GDP per capita", 1.0)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    0.3400000000000001    1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      0.383                 1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    0.488                 1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    0.3799999999999999    1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    0.3959999999999999    1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"colgt"},"ColGt"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColGt(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColGt")," checks if each element in the specified column is greater than the given value."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColGt("GDP per capita", 1.4)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    false             1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      false             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    true              1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    false             1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    false             1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"collt"},"ColLt"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColLt(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColLt")," checks if each element in the specified column is less than the given value."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColLt("GDP per capita", 1.4)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    true              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      true              1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    false             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    true              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    true              1.522             0.999                      0.557                           0.322         0.298                        \n")),(0,t.kt)("h2",{id:"coleq"},"ColEq"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) ColEq(colname string, value float64) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"ColEq")," checks if each element in the specified column is equal to the given value."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.ColEq("GDP per capita", 1.38)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    false             1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              7.6      false             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    false             1.582             1.028                      0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    true              1.624             1.026                      0.591                           0.354         0.118                        \n4    |    5               Netherlands          7.488    false             1.522             0.999                      0.557                           0.322         0.298                        \n")))}d.isMDXComponent=!0}}]);