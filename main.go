/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"os"
	"strings"

	"github.com/sniperkit/snk.fork.boogeyman/adapter/controller"
	"github.com/sniperkit/snk.fork.boogeyman/adapter/persistent/repository"
	"github.com/sniperkit/snk.fork.boogeyman/adapter/persistent/service"
	"github.com/sniperkit/snk.fork.boogeyman/adapter/presenter/console"
	"github.com/sniperkit/snk.fork.boogeyman/cross_cutting/io"
	"github.com/sniperkit/snk.fork.boogeyman/domain"
	"github.com/sniperkit/snk.fork.boogeyman/infrastructure/cmd"
	"github.com/sniperkit/snk.fork.boogeyman/infrastructure/meta_info"
	spiderPool "github.com/sniperkit/snk.fork.boogeyman/infrastructure/service"
)

var (
	version   string
	revision  string
	buildDate string
	goVersion string
	mode      string
)

var metaInfo = meta_info.NewMetaInfo(
	version,
	revision,
	buildDate,
	goVersion,
	mode,
)

func main() {
	commandParser := cmd.NewCommandParser()

	// parse command params
	cmdParams := commandParser.ParseParams()

	// check meta_info
	if cmdParams.ShowVersion {
		ShowMetaInfo(metaInfo)
	}

	resultPoolRepo := MaterialPoolFactory(cmdParams.Engine)
	textPresenter := console.NewColorfulTextPresenter()
	queryStrategy := SetQueryStrategy(cmdParams.Strategy)

	infoSearchCtl := controller.NewInfoSearch(textPresenter, resultPoolRepo)

	err := infoSearchCtl.Search(cmdParams.QueryString, queryStrategy)
	if err != nil {
		io.Errorln(err)
		os.Exit(1)
	}
}

func ShowMetaInfo(metaInfo *meta_info.MetaInfo) {
	io.Infof(metaInfo.GetMetaInfo())
	os.Exit(0)
}

func MaterialPoolFactory(selectedEngine string) *repository.QueryResultPool {
	collectors := service.EmptyCollectorList()
	switch strings.ToUpper(selectedEngine) {
	case domain.GOOGLE.String():
		collectors.Add(spiderPool.NewGoogleSpider())
		break
	case domain.BING.String():
		collectors.Add(spiderPool.NewBingSpider())
		break
	case domain.ASK.String():
		collectors.Add(spiderPool.NewAskSpider())
		break
	case domain.YAHOO.String():
		collectors.Add(spiderPool.NewYahooSpider())
		break
	default:
		collectors.Add(spiderPool.NewAskSpider())
		collectors.Add(spiderPool.NewBingSpider())
		collectors.Add(spiderPool.NewGoogleSpider())
		collectors.Add(spiderPool.NewYahooSpider())
	}
	return repository.NewResultPool(*collectors)
}

func SetQueryStrategy(selectedStrategy string) domain.RankerStrategyType {
	switch strings.ToUpper(selectedStrategy) {
	case domain.TOP.String():
		return domain.TOP
	case domain.CROSS.String():
		return domain.CROSS
	default:
		return domain.ALL
	}
}
