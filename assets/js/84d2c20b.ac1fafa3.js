"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[5633],{3905:(e,n,t)=>{t.d(n,{Zo:()=>m,kt:()=>d});var a=t(7294);function r(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){r(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function i(e,n){if(null==e)return{};var t,a,r=function(e,n){if(null==e)return{};var t,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||(r[t]=e[t]);return r}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(r[t]=e[t])}return r}var s=a.createContext({}),p=function(e){var n=a.useContext(s),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},m=function(e){var n=p(e.components);return a.createElement(s.Provider,{value:n},e.children)},c={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},u=a.forwardRef((function(e,n){var t=e.components,r=e.mdxType,o=e.originalType,s=e.parentName,m=i(e,["components","mdxType","originalType","parentName"]),u=p(t),d=r,f=u["".concat(s,".").concat(d)]||u[d]||c[d]||o;return t?a.createElement(f,l(l({ref:n},m),{},{components:t})):a.createElement(f,l({ref:n},m))}));function d(e,n){var t=arguments,r=n&&n.mdxType;if("string"==typeof e||r){var o=t.length,l=new Array(o);l[0]=u;var i={};for(var s in n)hasOwnProperty.call(n,s)&&(i[s]=n[s]);i.originalType=e,i.mdxType="string"==typeof e?e:r,l[1]=i;for(var p=2;p<o;p++)l[p]=t[p];return a.createElement.apply(null,l)}return a.createElement.apply(null,t)}u.displayName="MDXCreateElement"},154:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>s,contentTitle:()=>l,default:()=>c,frontMatter:()=>o,metadata:()=>i,toc:()=>p});var a=t(7462),r=(t(7294),t(3905));const o={},l="JSON",i={unversionedId:"io/json",id:"io/json",title:"JSON",description:"ReadJsonByColumns",source:"@site/docs/io/json.md",sourceDirName:"io",slug:"/io/json",permalink:"/docs/io/json",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"CSV",permalink:"/docs/io/csv"},next:{title:"Excel",permalink:"/docs/io/excel"}},s={},p=[{value:"ReadJsonByColumns",id:"readjsonbycolumns",level:2},{value:"ReadJsonStream",id:"readjsonstream",level:2},{value:"WriteJson",id:"writejson",level:2}],m={toc:p};function c(e){let{components:n,...t}=e;return(0,r.kt)("wrapper",(0,a.Z)({},m,t,{components:n,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"json"},"JSON"),(0,r.kt)("h2",{id:"readjsonbycolumns"},"ReadJsonByColumns"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"func ReadJsonByColumns(pathToFile string, indexCols []string) (DataFrame, error)\n")),(0,r.kt)("p",null,(0,r.kt)("inlineCode",{parentName:"p"},"ReadJsonByColumns")," reads a JSON file and returns a new DataFrame object. It is recommended to generate ",(0,r.kt)("inlineCode",{parentName:"p"},"pathToFile")," using the ",(0,r.kt)("inlineCode",{parentName:"p"},"path/filepath")," package for cross-platform compatibility."),(0,r.kt)("p",null,"The JSON file should be in this format:\n",(0,r.kt)("inlineCode",{parentName:"p"},'{"col1":[val1, val2, ...], "col2":[val1, val2, ...], ...}')),(0,r.kt)("p",null,"You can either set a column to be the index, or set it as ",(0,r.kt)("inlineCode",{parentName:"p"},"nil"),". If ",(0,r.kt)("inlineCode",{parentName:"p"},"nil"),", a new RangeIndex will be created.\nYour index column should not have any missing values. Columns will be alphabetically sorted, but the index column will always come first."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'/*\npeople.json\n\n{\n    "Name": ["Avery", "Bradley", "Candice"],\n    "Age": [19.0, 26.0, 23.0],\n    "Sex": ["Male", "Male", "Female"]\n}\n*/\n\nmyDf, err := gambas.ReadJsonByColumns(filepath.Join(".", "people.json"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nmyDf.Print()\n')),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"     |    Age    Name       Sex       \n0    |    19     Avery      Male      \n1    |    26     Bradley    Male      \n2    |    23     Candice    Female\n")),(0,r.kt)("h2",{id:"readjsonstream"},"ReadJsonStream"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"func ReadJsonStream(pathToFile string, indexCols []string) (DataFrame, error)\n")),(0,r.kt)("p",null,(0,r.kt)("inlineCode",{parentName:"p"},"ReadJsonStream")," reads a JSON stream and returns a new DataFrame object. It is recommended to generate ",(0,r.kt)("inlineCode",{parentName:"p"},"pathToFile")," using the ",(0,r.kt)("inlineCode",{parentName:"p"},"path/filepath")," package for cross-platform compatibility."),(0,r.kt)("p",null,"The JSON file should be in this format: ",(0,r.kt)("inlineCode",{parentName:"p"},'{"col1":val1, "col2":val2, ...}{"col1":val1, "col2":val2, ...}')),(0,r.kt)("p",null,"You can either set a column to be the index, or set it as ",(0,r.kt)("inlineCode",{parentName:"p"},"nil"),". If ",(0,r.kt)("inlineCode",{parentName:"p"},"nil"),", a new RangeIndex will be created.\nYour index column should not have any missing values. Columns will be alphabetically sorted, but the index column will always come first."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'/*\npeople.json\n\n[\n    {\n        "Name": "Avery",\n        "Age": 19.0,\n        "Sex": "Male"\n    },\n    {\n        "Name": "Bradley",\n        "Age": 26.0,\n        "Sex": "Male"\n    },\n    {\n        "Name": "Candice",\n        "Age": 23.0,\n        "Sex": "Female"\n    }\n]\n*/\n\nmyDf, err := gambas.ReadJsonStream(filepath.Join(".", "people.json"), nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nmyDf.Print()\n')),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"     |    Age    Name       Sex       \n0    |    19     Avery      Male      \n1    |    26     Bradley    Male      \n2    |    23     Candice    Female\n")),(0,r.kt)("h2",{id:"writejson"},"WriteJson"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"func WriteJson(df DataFrame, pathToFile string) (os.FileInfo, error)\n")),(0,r.kt)("p",null,(0,r.kt)("inlineCode",{parentName:"p"},"WriteJson")," writes a DataFrame object to a file. It is recommended to generate ",(0,r.kt)("inlineCode",{parentName:"p"},"pathToFile")," using the ",(0,r.kt)("inlineCode",{parentName:"p"},"path/filepath")," package for cross-platform compatibility."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}\nmyCols := []string{"group a", "group b", "group c"}\nmyIndexCols := []string{"group a"}\n\nmyDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)\nif err != nil {\n    fmt.Println(err)\n}\n\ngambas.WriteJson(myDf, filepath.Join(".", "output.json"))\n')),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-json"},'// output.json\n\n{\n    "series": [\n        {\n            "data": [\n                1,\n                2,\n                3\n            ],\n            "name": "group a",\n            "dtype": "int"\n        },\n        {\n            "data": [\n                4,\n                5,\n                6\n            ],\n            "name": "group b",\n            "dtype": "int"\n        },\n        {\n            "data": [\n                7,\n                8,\n                9\n            ],\n            "name": "group c",\n            "dtype": "int"\n        }\n    ],\n    "columns": [\n        "group a",\n        "group b",\n        "group c"\n    ]\n}\n')))}c.isMDXComponent=!0}}]);