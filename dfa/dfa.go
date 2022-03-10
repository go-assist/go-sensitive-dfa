package dfa

import (
	"regexp"
	"strings"
)

type Sensitive struct {
	Words 		string
	Filter 		string 	// " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,`,"
	Compile 	bool 	// 是否需要替换,true,匹配到敏感词则按规则替换
	Rule  		string 	// 替换规则,正则表达式, `(\.|!|/| |~|@|#|$|%|^|&|\$|\*|\(|\)|_|\+|=|\?|<|>|-|,|—|，|。|/|\\|\||《|》|？|;|:|:|\\|‘|；|“|¥|·|\{|}|/|\|"|~|""|"|)`
	Repl 		string  // 替换的目标字符串
}

// MakeInvalidSensitiveWordsDFA 生成跳过验证的map.
func (sw *Sensitive) MakeInvalidSensitiveWordsDFA () (dfa map[string]interface{}) {
	words := strings.Split(sw.Filter,",")
	invalid := make(map[string]interface{})
	for _, v := range words {
		invalid[v] = nil
	}
	dfa = invalid
	return
}

// MakeInitSensitiveWordsDFA 敏感词集合, map[不:map[isEnd:false 合:map[isEnd:false 法:map[isEnd:true]]]
func (sw *Sensitive) MakeInitSensitiveWordsDFA (set map[string]interface{}) (dfa map[string]interface{}) {
	sensitiveWord := make(map[string]interface{})
	for key := range set {
		str := []rune(key)
		nowMap := sensitiveWord
		for i := 0; i < len(str); i++ {
			if _, ok := nowMap[string(str[i])]; !ok {
				thisMap := make(map[string]interface{})
				thisMap["isEnd"] = false
				nowMap[string(str[i])] = thisMap
				nowMap = thisMap
			} else {
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}
			if i == len(str) - 1 {
				nowMap["isEnd"] = true
			}
		}
	}
	dfa = sensitiveWord
	return
}

// FilterSensitiveWordsDFA 过滤字符串
func (sw *Sensitive) FilterSensitiveWordsDFA(sensitive, invalidWordDFA map[string]interface{}) (sensitiveWordsArr []string) {
	str := []rune(sw.Words)
	nowMap := sensitive
	start := -1
	tag := -1
	for i := 0; i < len(str); i++ {
		// 如果是无效词汇直接跳过
		if _, ok:= invalidWordDFA[(string(str[i]))]; ok {
			continue
		}
		if thisMap, ok := nowMap[string(str[i])].(map[string]interface{}); ok {
			// 记录敏感词第一个文字的位置
			tag++
			if  tag == 0 {
				start = i
			}
			// 判断是否为敏感词的最后一个文字
			if isEnd, _ := thisMap["isEnd"].(bool); isEnd {
				// 42 为 *,可选返回类型
				// 记录匹配到的值
				fw := ``
				for y := start; y < i+1; y++ {
					fw = fw + string(str[y])
					str[y] = 42
				}
				if len(fw) > 0 {
					if sw.Compile {
						fw = regexp.MustCompile(sw.Rule).ReplaceAllString(fw, sw.Repl)
					}
					sensitiveWordsArr = append(sensitiveWordsArr, fw)
				}
				// 重置标识位
				nowMap = sensitive
				start = -1
				tag = -1
			} else {// 不是最后一个,则将其包含的map赋值给nowMap
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}
		} else {  // 如果敏感词不是全匹配,则终止此敏感词查找.从开始位置的第二个文字继续判断
			if start != -1 {
				i = start + 1
			}
			// 重置标识位
			nowMap = sensitive
			start = -1
			tag = -1
		}
	}
	return
}