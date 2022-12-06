package item

import (
	"es_load_test/internal/repositories/search"
	"github.com/olivere/elastic/v7"
)

func (qs *builderService) BuildSearchQuery(req search.Request) (elastic.Query, search.GetHits) {
	mustMatch := elastic.NewBoolQuery()

	if req.Query != "" {
		mustMatch = getQuery(req)
	}

	return mustMatch, search.DefaultGetHitsFunc
}

func getQuery(req search.Request) *elastic.BoolQuery {
	var constantScores []elastic.Query
	constantScores = append(constantScores,
		elastic.NewConstantScoreQuery(elastic.NewMatchPhraseQuery("name.analyzed", req.Query)).Boost(4),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhrasePrefixQuery("name.analyzed", req.Query)).Boost(2),
		elastic.NewConstantScoreQuery(elastic.NewMatchQuery("name.analyzed", req.Query).Operator("AND")).Boost(3),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhraseQuery("brand_name.analyzed", req.Query)).Boost(2),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhrasePrefixQuery("brand_name.analyzed", req.Query)).Boost(1),
		elastic.NewConstantScoreQuery(elastic.NewMatchQuery("brand_name.analyzed", req.Query).Operator("AND")).Boost(1.5),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhraseQuery("category_names.analyzed", req.Query)).Boost(2),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhrasePrefixQuery("category_names.analyzed", req.Query)).Boost(1),
		elastic.NewConstantScoreQuery(elastic.NewMatchQuery("category_names.analyzed", req.Query).Operator("AND")).Boost(1.5),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhraseQuery("tags.analyzed", req.Query)).Boost(4),
		elastic.NewConstantScoreQuery(elastic.NewMatchPhrasePrefixQuery("tags.analyzed", req.Query)).Boost(2),
		elastic.NewConstantScoreQuery(elastic.NewMatchQuery("tags.analyzed", req.Query).Operator("AND")).Boost(3))
	return elastic.NewBoolQuery().MinimumNumberShouldMatch(1).Should(
		constantScores...)
}
