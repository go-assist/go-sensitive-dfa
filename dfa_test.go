package go_sensitive_dfa

import (
	"github.com/golangtoolkit/go-sensitive-dfa/dfa"
	"testing"
)

func TestDfa(t *testing.T) {
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
	if len(filter) == 0 {
		t.Errorf("errors .")
	}
}
