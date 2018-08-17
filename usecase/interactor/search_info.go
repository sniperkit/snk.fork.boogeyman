/*
Sniperkit-Bot
- Status: analyzed
*/

package interactor

import (
	"github.com/pkg/errors"

	"github.com/sniperkit/snk.fork.boogeyman/config"
	"github.com/sniperkit/snk.fork.boogeyman/domain"
	"github.com/sniperkit/snk.fork.boogeyman/usecase/presenter"
	"github.com/sniperkit/snk.fork.boogeyman/usecase/repository"
)

type InfoSearch struct {
	ranker    *domain.Ranker
	poolRepo  repository.QueryResultPool
	presenter presenter.TextPresenter
}

func NewInfoSearch(
	presenter presenter.TextPresenter,
	poolRepo repository.QueryResultPool,
) *InfoSearch {
	return &InfoSearch{
		ranker:    domain.NewRanker(),
		poolRepo:  poolRepo,
		presenter: presenter,
	}
}

func (i *InfoSearch) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) (*domain.QueryResult, error) {
	resultPool, err := i.fetchData(domain.NewKeyword(queryString))
	if err != nil {
		return nil, errors.Wrap(err, "Error on fetch data from pool!\n")
	}
	switch strategy {
	case domain.TOP:
		return i.ranker.Top(resultPool)
	case domain.CROSS:
		return i.ranker.CrossMatch(resultPool)
	case domain.ALL:
		return i.ranker.All(resultPool,
			config.GetConfig().RankerConf.MaxReturnItems)
	default:
		return i.ranker.CrossMatch(resultPool)
	}
}

func (i *InfoSearch) PrintResults(results *domain.QueryResult) {
	i.presenter.PrintList(results)
}

func (i *InfoSearch) fetchData(
	keyword *domain.Keyword,
) (*domain.QueryResultPool, error) {
	pool, err := i.poolRepo.FetchData(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching data from search engines!\n")
	}
	return pool, nil
}
