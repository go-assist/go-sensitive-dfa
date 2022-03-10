# go-sensitive-dfa
# Helper

##### 1.golang dfa检测敏感词

##### 2.获取 ❤❤❤

go get github.com/golangtoolkit/go-sensitive-dfa

##### 3. 示例 for example

```Golang
package main

import (
	"fmt"
	"github.com/golangtoolkit/go-sensitive-dfa/dfa"
)

func example() {
	var sw dfa.Sensitive
	sw.Filter = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,`,"
	sw.Rule = `(\.|!|/| |~|@|#|$|%|^|&|\$|\*|\(|\)|_|\+|=|\?|<|>|-|,|—|，|。|/|\\|\||《|》|？|;|:|:|\\|‘|；|“|¥|·|\{|}|/|\|"|~|""|"|)`
	sw.Compile = true
	sw.Repl = ``
	sw.Words = `一只草##泥马`
	invalid := sw.MakeInvalidSensitiveWordsDFA()
	m := map[string]interface{}{
		"草拟吗": nil,
		"曹尼玛": nil,
		"草泥马": nil,
	}
	dfa1 := sw.MakeInitSensitiveWordsDFA(m)
	filter := sw.FilterSensitiveWordsDFA(dfa1,invalid)
	fmt.Println(filter)
}
```

##### 4. 使用过程如有疑问欢迎issue ｡◕‿◕｡
