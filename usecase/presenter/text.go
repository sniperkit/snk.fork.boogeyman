/*
Sniperkit-Bot
- Status: analyzed
*/

package presenter

import (
	"github.com/sniperkit/snk.fork.boogeyman/domain"
)

type TextPresenter interface {
	PrintList(results *domain.QueryResult)
}
