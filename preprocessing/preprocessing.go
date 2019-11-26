package preprocessing

import (
	"strings"
	"unicode/utf8"

	"github.com/shogo82148/go-mecab"
)

type TaggerNode struct {
	Surface       string
	PartsOfSpeech string
	Subcategory1  string
	Subcategory2  string
	Subcategory3  string
	Utilization   string
	Pratical      string
	Origin        string
	Reading       string
	Pronunciation string
}

func GetKeys(fileName string) ([]string, error) {
	tagger, err := mecab.New(map[string]string{})
	if err != nil {
		return nil, err
	}
	defer tagger.Destroy()

	result, err := tagger.ParseToNode(fileName)
	if err != nil {
		return nil, err
	}

	noSymbolFileName := ""
	keys := []string{}
	result = result.Next()
	for ; !result.Next().IsZero(); result = result.Next() {
		resultStruct := nodeToStruct(result)
		if !resultStruct.isSymbol() { //記号を取り除いた元ファイル名の生成
			noSymbolFileName += resultStruct.getSurface()
		}
		if resultStruct.allowAdd() {
			keys = append(keys, resultStruct.getOrigin())
		}
	}

	noSokuonFileName := trimSokuon(noSymbolFileName)
	noLongVowelFileName := trimLongVowel(noSokuonFileName)
	keys = append(keys, noSymbolFileName)
	keys = append(keys, noSokuonFileName)
	keys = append(keys, noLongVowelFileName)

	return keys, nil
}

func (t TaggerNode) allowAdd() bool {
	reject := t.Surface == "っ" ||
		t.PartsOfSpeech == "記号" ||
		t.PartsOfSpeech == "助詞" ||
		t.PartsOfSpeech == "助動詞" ||
		t.PartsOfSpeech == "フィラー" ||
		t.PartsOfSpeech == "連体詞" ||
		t.PartsOfSpeech == "名詞" && t.Subcategory1 == "非自立" ||
		t.PartsOfSpeech == "接頭詞" && t.Subcategory1 == "名詞接続" ||
		t.PartsOfSpeech == "感動詞" && utf8.RuneCountInString(t.Origin) < 2 ||
		t.Subcategory1 == "代名詞" ||
		t.Subcategory1 == "非自立" && t.Pratical == "連用形" ||
		t.Subcategory1 == "接尾" && t.Subcategory2 == "人名" ||
		t.Utilization == "五段・カ行促音便" ||
		t.Utilization == "サ変・スル" ||
		t.Origin == "*" ||
		t.Pratical == "文語基本形" && t.PartsOfSpeech != "形容詞"

	return !reject
}

func (t TaggerNode) getOrigin() string {
	return t.Origin
}

func (t TaggerNode) getSurface() string {
	return t.Surface
}

func nodeToStruct(node mecab.Node) TaggerNode {
	ret := TaggerNode{}
	ret.Surface = node.Surface()

	feature := node.Feature()
	splitedFeature := strings.Split(feature, ",")

	ret.PartsOfSpeech = splitedFeature[0]
	ret.Subcategory1 = splitedFeature[1]
	ret.Subcategory2 = splitedFeature[2]
	ret.Subcategory3 = splitedFeature[3]
	ret.Utilization = splitedFeature[4]
	ret.Pratical = splitedFeature[5]
	ret.Origin = splitedFeature[6]
	// 発音と読みがないことがあるため
	if !(len(splitedFeature) < 8) {
		ret.Reading = splitedFeature[7]
		ret.Pronunciation = splitedFeature[8]
	}

	return ret
}

func (t TaggerNode) isSymbol() bool {
	if t.PartsOfSpeech == "記号" {
		return true
	}
	return false
}

func trimLongVowel(fileName string) string {
	return strings.ReplaceAll(fileName, "ー", "")
}

func trimSokuon(fileName string) string {
	return strings.ReplaceAll(fileName, "っ", "")
}
