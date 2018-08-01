package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"strings"
	"testing"
)

func TestAnalyzer(t *testing.T) {
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://10.10.11.200:9200")) //http://10.10.11.200:9200,http://47.96.172.174:9200，http://10.10.7.127:9200
	if err != nil {
		t.Error(err)
	}
	res, err := client.IndexAnalyze().Analyzer("ik_smart").Text("全国性交流会议在武汉举办").Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	fmt.Println("")
	fmt.Printf("%+v", res)
	//tokens :=[len(res.Tokens)]string
	tokens := make([]string, len(res.Tokens))
	for i, v := range res.Tokens {
		tokens[i] = v.Token
	}
	content := strings.Join(tokens, "|")
	fmt.Print(content)
}
