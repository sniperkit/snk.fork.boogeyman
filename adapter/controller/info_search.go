/*
Sniperkit-Bot
- Status: analyzed
*/

package controller

import (
	"github.com/sniperkit/snk.fork.boogeyman/domain"
	"github.com/sniperkit/snk.fork.boogeyman/usecase/interactor"
	"github.com/sniperkit/snk.fork.boogeyman/usecase/presenter"
	"github.com/sniperkit/snk.fork.boogeyman/usecase/repository"
)

type InfoSearch struct {
	interactor *interactor.InfoSearch
}

func NewInfoSearch(
	presenter presenter.TextPresenter,
	resultPoolRepo repository.QueryResultPool,
) *InfoSearch {
	return &InfoSearch{
		interactor: interactor.NewInfoSearch(presenter, resultPoolRepo),
	}
}

func (b *InfoSearch) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) error {
	queryResults, err := b.interactor.Search(queryString, strategy)
	if err != nil {
		return err
	}
	b.interactor.PrintResults(queryResults)
	return nil
}
