"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[6750],{3905:(e,n,r)=>{r.d(n,{Zo:()=>d,kt:()=>u});var a=r(7294);function t(e,n,r){return n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r,e}function o(e,n){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),r.push.apply(r,a)}return r}function l(e){for(var n=1;n<arguments.length;n++){var r=null!=arguments[n]?arguments[n]:{};n%2?o(Object(r),!0).forEach((function(n){t(e,n,r[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(r,n))}))}return e}function i(e,n){if(null==e)return{};var r,a,t=function(e,n){if(null==e)return{};var r,a,t={},o=Object.keys(e);for(a=0;a<o.length;a++)r=o[a],n.indexOf(r)>=0||(t[r]=e[r]);return t}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)r=o[a],n.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(t[r]=e[r])}return t}var p=a.createContext({}),c=function(e){var n=a.useContext(p),r=n;return e&&(r="function"==typeof e?e(n):l(l({},n),e)),r},d=function(e){var n=c(e.components);return a.createElement(p.Provider,{value:n},e.children)},s={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},m=a.forwardRef((function(e,n){var r=e.components,t=e.mdxType,o=e.originalType,p=e.parentName,d=i(e,["components","mdxType","originalType","parentName"]),m=c(r),u=t,f=m["".concat(p,".").concat(u)]||m[u]||s[u]||o;return r?a.createElement(f,l(l({ref:n},d),{},{components:r})):a.createElement(f,l({ref:n},d))}));function u(e,n){var r=arguments,t=n&&n.mdxType;if("string"==typeof e||t){var o=r.length,l=new Array(o);l[0]=m;var i={};for(var p in n)hasOwnProperty.call(n,p)&&(i[p]=n[p]);i.originalType=e,i.mdxType="string"==typeof e?e:t,l[1]=i;for(var c=2;c<o;c++)l[c]=r[c];return a.createElement.apply(null,l)}return a.createElement.apply(null,r)}m.displayName="MDXCreateElement"},8949:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>p,contentTitle:()=>l,default:()=>s,frontMatter:()=>o,metadata:()=>i,toc:()=>c});var a=r(7462),t=(r(7294),r(3905));const o={},l="Editing Properties",i={unversionedId:"dataframe/editing-properties",id:"dataframe/editing-properties",title:"Editing Properties",description:"You can edit your DataFrame object using these functions.",source:"@site/docs/dataframe/editing-properties.md",sourceDirName:"dataframe",slug:"/dataframe/editing-properties",permalink:"/docs/dataframe/editing-properties",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Arithmetic Operations",permalink:"/docs/dataframe/arithmetic-operations"},next:{title:"Merging",permalink:"/docs/dataframe/merging"}},p={},c=[{value:"NewCol",id:"newcol",level:2},{value:"NewDerivedCol",id:"newderivedcol",level:2},{value:"RenameCol",id:"renamecol",level:2},{value:"DropNaN",id:"dropnan",level:2},{value:"Example 1: Dropping rows",id:"example-1-dropping-rows",level:3},{value:"Example 1: Dropping columns",id:"example-1-dropping-columns",level:3}],d={toc:c};function s(e){let{components:n,...r}=e;return(0,t.kt)("wrapper",(0,a.Z)({},d,r,{components:n,mdxType:"MDXLayout"}),(0,t.kt)("h1",{id:"editing-properties"},"Editing Properties"),(0,t.kt)("p",null,"You can edit your ",(0,t.kt)("inlineCode",{parentName:"p"},"DataFrame")," object using these functions."),(0,t.kt)("p",null,"The data used in the example ",(0,t.kt)("inlineCode",{parentName:"p"},"2019.csv")," is UN's 2019 World Happiness Report, sourced from ",(0,t.kt)("a",{parentName:"p",href:"https://www.kaggle.com/datasets/unsdsn/world-happiness"},"Kaggle"),"."),(0,t.kt)("h2",{id:"newcol"},"NewCol"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) NewCol(colname string, data []interface{}) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"NewCol")," creates a new column with the given data and column name."),(0,t.kt)("p",null,"To create a blank column, pass in ",(0,t.kt)("inlineCode",{parentName:"p"},"nil"),"."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1.Head(5)\nfmt.Println("")\n\nres, err := df1.NewCol("NewCol", nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    \n0    |    1               Finland              7.769    1.34              1.587             \n1    |    2               Denmark              7.6      1.383             1.573             \n2    |    3               Norway               7.554    1.488             1.582             \n3    |    4               Iceland              7.494    1.38              1.624             \n4    |    5               Netherlands          7.488    1.396             1.522             \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    NewCol    \n0    |    1               Finland              7.769    1.34              1.587             NaN       \n1    |    2               Denmark              7.6      1.383             1.573             NaN       \n2    |    3               Norway               7.554    1.488             1.582             NaN       \n3    |    4               Iceland              7.494    1.38              1.624             NaN       \n4    |    5               Netherlands          7.488    1.396             1.522             NaN       \n")),(0,t.kt)("h2",{id:"newderivedcol"},"NewDerivedCol"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) NewDerivedCol(colname, srcCol string) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"NewDerivedCol")," creates a new column derived from an existing column."),(0,t.kt)("p",null,"It copies over the data from ",(0,t.kt)("inlineCode",{parentName:"p"},"srcCol")," into a new column."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1.Head(5)\nfmt.Println("")\n\nres, err := df1.NewDerivedCol("NewCol", "Social support")\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    \n0    |    1               Finland              7.769    1.34              1.587             \n1    |    2               Denmark              7.6      1.383             1.573             \n2    |    3               Norway               7.554    1.488             1.582             \n3    |    4               Iceland              7.494    1.38              1.624             \n4    |    5               Netherlands          7.488    1.396             1.522             \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    NewCol    \n0    |    1               Finland              7.769    1.34              1.587             1.587     \n1    |    2               Denmark              7.6      1.383             1.573             1.573     \n2    |    3               Norway               7.554    1.488             1.582             1.582     \n3    |    4               Iceland              7.494    1.38              1.624             1.624     \n4    |    5               Netherlands          7.488    1.396             1.522             1.522     \n")),(0,t.kt)("h2",{id:"renamecol"},"RenameCol"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) RenameCol(colnames map[string]string) error\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"RenameCol")," renames columns in a ",(0,t.kt)("inlineCode",{parentName:"p"},"DataFrame"),"."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1.Head(5)\nfmt.Println("")\n\nerr = df1.RenameCol(map[string]string{"Social support": "Some kind of support"})\nif err != nil {\n    fmt.Println(err)\n}\n\ndf1.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    \n0    |    1               Finland              7.769    1.34              1.587             \n1    |    2               Denmark              7.6      1.383             1.573             \n2    |    3               Norway               7.554    1.488             1.582             \n3    |    4               Iceland              7.494    1.38              1.624             \n4    |    5               Netherlands          7.488    1.396             1.522             \n\nindex does not exist: Social support\n     |    Overall rank    Country or region    Score    GDP per capita    Some kind of support    \n0    |    1               Finland              7.769    1.34              1.587                   \n1    |    2               Denmark              7.6      1.383             1.573                   \n2    |    3               Norway               7.554    1.488             1.582                   \n3    |    4               Iceland              7.494    1.38              1.624                   \n4    |    5               Netherlands          7.488    1.396             1.522                   \n")),(0,t.kt)("h2",{id:"dropnan"},"DropNaN"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},"func (df *DataFrame) DropNaN(axis int) (DataFrame, error)\n")),(0,t.kt)("p",null,(0,t.kt)("inlineCode",{parentName:"p"},"DropNaN")," drops rows or columns with ",(0,t.kt)("inlineCode",{parentName:"p"},"NaN")," values."),(0,t.kt)("p",null,"Specify axis to choose whether to remove rows with ",(0,t.kt)("inlineCode",{parentName:"p"},"NaN")," or columns with ",(0,t.kt)("inlineCode",{parentName:"p"},"NaN"),". ",(0,t.kt)("inlineCode",{parentName:"p"},"axis=0")," is row, ",(0,t.kt)("inlineCode",{parentName:"p"},"axis=1")," is column."),(0,t.kt)("h3",{id:"example-1-dropping-rows"},"Example 1: Dropping rows"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019-with-nan.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.DropNaN(0)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              NaN      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             NaN                        0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           NaN           0.118                        \n4    |    5               Netherlands          7.488    1.396             NaN               0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n5    |    6               Switzerland          7.48     1.452             1.526             1.052                      0.572                           0.263         0.343                        \n6    |    7               Sweden               7.343    1.387             1.487             1.009                      0.574                           0.267         0.373                        \n7    |    8               New Zealand          7.307    1.303             1.557             1.026                      0.585                           0.33          0.38                         \n8    |    9               Canada               7.278    1.365             1.505             1.039                      0.584                           0.285         0.308                        \n")),(0,t.kt)("h3",{id:"example-1-dropping-columns"},"Example 1: Dropping columns"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre",className:"language-go"},'df, err := gambas.ReadCsv(filepath.Join(".", "2019-with-nan.csv"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\ndf.Head(5)\nfmt.Println("")\n\nres, err := df.DropNaN(1)\nif err != nil {\n    fmt.Println(err)\n}\n\nres.Head(5)\n')),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    \n0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        \n1    |    2               Denmark              NaN      1.383             1.573             0.996                      0.592                           0.252         0.41                         \n2    |    3               Norway               7.554    1.488             1.582             NaN                        0.603                           0.271         0.341                        \n3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           NaN           0.118                        \n4    |    5               Netherlands          7.488    1.396             NaN               0.999                      0.557                           0.322         0.298                        \n\n     |    Overall rank    Country or region    GDP per capita    Healthy life expectancy    Freedom to make life choices    Perceptions of corruption    \n0    |    1               Finland              1.34              0.986                      0.596                           0.393                        \n1    |    2               Denmark              1.383             0.996                      0.592                           0.41                         \n2    |    3               Norway               1.488             NaN                        0.603                           0.341                        \n3    |    4               Iceland              1.38              1.026                      0.591                           0.118                        \n4    |    5               Netherlands          1.396             0.999                      0.557                           0.298                        \n")))}s.isMDXComponent=!0}}]);