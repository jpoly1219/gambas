"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[7309],{3905:(e,n,r)=>{r.d(n,{Zo:()=>p,kt:()=>u});var t=r(7294);function a(e,n,r){return n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r,e}function i(e,n){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),r.push.apply(r,t)}return r}function l(e){for(var n=1;n<arguments.length;n++){var r=null!=arguments[n]?arguments[n]:{};n%2?i(Object(r),!0).forEach((function(n){a(e,n,r[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(r,n))}))}return e}function m(e,n){if(null==e)return{};var r,t,a=function(e,n){if(null==e)return{};var r,t,a={},i=Object.keys(e);for(t=0;t<i.length;t++)r=i[t],n.indexOf(r)>=0||(a[r]=e[r]);return a}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(t=0;t<i.length;t++)r=i[t],n.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=t.createContext({}),o=function(e){var n=t.useContext(s),r=n;return e&&(r="function"==typeof e?e(n):l(l({},n),e)),r},p=function(e){var n=o(e.components);return t.createElement(s.Provider,{value:n},e.children)},c={inlineCode:"code",wrapper:function(e){var n=e.children;return t.createElement(t.Fragment,{},n)}},d=t.forwardRef((function(e,n){var r=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,p=m(e,["components","mdxType","originalType","parentName"]),d=o(r),u=a,y=d["".concat(s,".").concat(u)]||d[u]||c[u]||i;return r?t.createElement(y,l(l({ref:n},p),{},{components:r})):t.createElement(y,l({ref:n},p))}));function u(e,n){var r=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var i=r.length,l=new Array(i);l[0]=d;var m={};for(var s in n)hasOwnProperty.call(n,s)&&(m[s]=n[s]);m.originalType=e,m.mdxType="string"==typeof e?e:a,l[1]=m;for(var o=2;o<i;o++)l[o]=r[o];return t.createElement.apply(null,l)}return t.createElement.apply(null,r)}d.displayName="MDXCreateElement"},2602:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>s,contentTitle:()=>l,default:()=>c,frontMatter:()=>i,metadata:()=>m,toc:()=>o});var t=r(7462),a=(r(7294),r(3905));const i={},l="Indexing",m={unversionedId:"series/indexing",id:"series/indexing",title:"Indexing",description:"You can index your Series with gambas's built-in indexing tools.",source:"@site/docs/series/indexing.md",sourceDirName:"series",slug:"/series/indexing",permalink:"/gambas/docs/series/indexing",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Printing",permalink:"/gambas/docs/series/printing"},next:{title:"Summary Statistics",permalink:"/gambas/docs/series/summary-statistics"}},s={},o=[{value:"At",id:"at",level:2},{value:"Example 1: Single index",id:"example-1-single-index",level:3},{value:"Example 2: Multiindex",id:"example-2-multiindex",level:3},{value:"IAt",id:"iat",level:2},{value:"Loc",id:"loc",level:2},{value:"Example 1: Indexing a single item",id:"example-1-indexing-a-single-item",level:3},{value:"Example 2: Indexing multiple items",id:"example-2-indexing-multiple-items",level:3},{value:"Example 3: Multiindex (all index)",id:"example-3-multiindex-all-index",level:3},{value:"Example 4: Multiindex (first index)",id:"example-4-multiindex-first-index",level:3},{value:"LocItems",id:"locitems",level:2},{value:"ILoc",id:"iloc",level:2}],p={toc:o};function c(e){let{components:n,...r}=e;return(0,a.kt)("wrapper",(0,t.Z)({},p,r,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"indexing"},"Indexing"),(0,a.kt)("p",null,"You can index your ",(0,a.kt)("inlineCode",{parentName:"p"},"Series")," with ",(0,a.kt)("inlineCode",{parentName:"p"},"gambas"),"'s built-in indexing tools."),(0,a.kt)("h2",{id:"at"},"At"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *Series) At(ind ...interface{}) (interface{}, error)\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"At")," returns an element at a given index."),(0,a.kt)("p",null,"For multiindex, you need to pass in the whole index tuple."),(0,a.kt)("h3",{id:"example-1-single-index"},"Example 1: Single index"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry"}\nmyName := "Fruit"\n\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.At(1)\nif err != nil {\n    fmt.Println(err)\n}\nfmt.Println(res)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"banana\n")),(0,a.kt)("h3",{id:"example-2-multiindex"},"Example 2: Multiindex"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry"}\nmyName := "Fruit"\nmyIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"c", "red"}}\n\nmyIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})\nif err != nil {\n    fmt.Println(err)\n}\n\nmySeries, err := gambas.NewSeries(myData, myName, &myIndexData)\nif err != nil {\n    fmt.Println(err)\n}\nmySeries.Print()\n\nres, err := mySeries.At("b", "yellow")\nif err != nil {\n    fmt.Println(err)\n}\nfmt.Println(res)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"key    color     |    Fruit     \na      red       |    apple     \nb      yellow    |    banana    \nc      red       |    cherry    \nbanana\n")),(0,a.kt)("h2",{id:"iat"},"IAt"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *Series) IAt(ind int) (interface{}, error)\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"IAt")," returns an element at a given integer index."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry"}\nmyName := "Fruit"\n\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.IAt(2)\nif err != nil {\n    fmt.Println(err)\n}\nfmt.Println(res)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"cherry\n")),(0,a.kt)("h2",{id:"loc"},"Loc"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *Series) Loc(idx ...[]interface{}) (Series, error)\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Loc")," accepts index tuples and returns a ",(0,a.kt)("inlineCode",{parentName:"p"},"Series")," object containing data at the given rows."),(0,a.kt)("p",null,"Each ",(0,a.kt)("inlineCode",{parentName:"p"},"idx")," item should contain the index of the data you would like to query. For multiindex ",(0,a.kt)("inlineCode",{parentName:"p"},"Series"),", you can either pass in the whole index tuple, or the first index."),(0,a.kt)("h3",{id:"example-1-indexing-a-single-item"},"Example 1: Indexing a single item"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.Loc([]interface{}{2})\nif err != nil {\n    fmt.Println(err)\n}\nres.Print()\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"     |    Fruit     \n2    |    cherry\n")),(0,a.kt)("h3",{id:"example-2-indexing-multiple-items"},"Example 2: Indexing multiple items"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.Loc([]interface{}{2}, []interface{}{4})\nif err != nil {\n    fmt.Println(err)\n}\nres.Print()\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"     |    Fruit     \n2    |    cherry    \n4    |    kiwi\n")),(0,a.kt)("h3",{id:"example-3-multiindex-all-index"},"Example 3: Multiindex (all index)"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\nmyIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"c", "red"}, {"d", "brown"}, {"e", "green"}}\n\nmyIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})\nif err != nil {\n    fmt.Println(err)\n}\n\nmySeries, err := gambas.NewSeries(myData, myName, &myIndexData)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.Loc([]interface{}{"a", "red"}, []interface{}{"d", "brown"})\nif err != nil {\n    fmt.Println(err)\n}\nres.Print()\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"key    color    |    Fruit     \na      red      |    apple     \nd      brown    |    durian\n")),(0,a.kt)("h3",{id:"example-4-multiindex-first-index"},"Example 4: Multiindex (first index)"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\nmyIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"a", "red"}, {"b", "brown"}, {"a", "green"}}\n\nmyIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})\nif err != nil {\n    fmt.Println(err)\n}\n\nmySeries, err := gambas.NewSeries(myData, myName, &myIndexData)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.Loc([]interface{}{"a"})\nif err != nil {\n    fmt.Println(err)\n}\nres.Print()\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"key    color    |    Fruit     \na      red      |    apple     \na      red      |    cherry    \na      green    |    kiwi\n")),(0,a.kt)("h2",{id:"locitems"},"LocItems"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *Series) LocItems(idx ...[]interface{}) ([]interface{}, error)\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"LocItems")," acts the exact same as Loc, but returns data as ",(0,a.kt)("inlineCode",{parentName:"p"},"[]interface{}")," instead of ",(0,a.kt)("inlineCode",{parentName:"p"},"Series"),". For usage, refer to ",(0,a.kt)("inlineCode",{parentName:"p"},"Loc"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\n\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.LocItems([]interface{}{1}, []interface{}{3}, []interface{}{4})\nif err != nil {\n    fmt.Println(err)\n}\nfmt.Println(res)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"[banana durian kiwi]\n")),(0,a.kt)("h2",{id:"iloc"},"ILoc"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *Series) ILoc(min, max int) ([]interface{}, error)\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"ILoc")," returns an array of elements at a given integer index range."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}\nmyName := "Fruit"\n\nmySeries, err := gambas.NewSeries(myData, myName, nil)\nif err != nil {\n    fmt.Println(err)\n}\n\nres, err := mySeries.ILoc(1, 3)\nif err != nil {\n    fmt.Println(err)\n}\nfmt.Println(res)\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"[banana cherry]\n")))}c.isMDXComponent=!0}}]);