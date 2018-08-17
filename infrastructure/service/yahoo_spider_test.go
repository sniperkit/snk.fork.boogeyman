/*
Sniperkit-Bot
- Status: analyzed
*/

package service_test

import (
	"testing"

	"github.com/sniperkit/snk.fork.boogeyman/domain"
	"github.com/sniperkit/snk.fork.boogeyman/infrastructure/service"
)

func TestYahooSpider_Query(t *testing.T) {
	keyword := domain.NewKeyword("sample")
	bingSpider := service.NewYahooSpider()

	result, err := bingSpider.Query(keyword)
	if err != nil {
		t.Fatal("Fail test query data from search engine")
	}
	if len(*result.GetQueryResults()) < 1 {
		t.Fatal("Fail test query data from se, maybe error on internet connection")
	}
	if result.Type() != domain.YAHOO {
		t.Fatal("Fail test query data from se, error search engine type")
	}
}
