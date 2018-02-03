package service

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
)

type Collector interface {
	Query(searchEngine search_engine.SearchEngineType, keyword *domain.Keyword) *search_engine.Base
}