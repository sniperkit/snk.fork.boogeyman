package repository

import (
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
)

type MaterialPool interface {
	GetItemsBySearchEngine(searchEngineType search_engine.SearchEngineType) (search_engine.Base, error)
}