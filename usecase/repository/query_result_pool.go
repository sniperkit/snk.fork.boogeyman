/*
Sniperkit-Bot
- Status: analyzed
*/

package repository

import (
	"github.com/sniperkit/snk.fork.boogeyman/domain"
)

type QueryResultPool interface {
	FetchData(keyword *domain.Keyword) (*domain.QueryResultPool, error)
}
