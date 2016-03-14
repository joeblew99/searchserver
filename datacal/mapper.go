package datacal

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzers/keyword_analyzer"
	"github.com/blevesearch/bleve/analysis/language/en"
)

// BuildMapping ... maps the data struts into the db.
func BuildMapping() *bleve.IndexMapping {
	enFieldMapping := bleve.NewTextFieldMapping()
	enFieldMapping.Analyzer = en.AnalyzerName
    
    

	eventMapping := bleve.NewDocumentMapping()
	eventMapping.AddFieldMappingsAt("summary", enFieldMapping)
	eventMapping.AddFieldMappingsAt("description", enFieldMapping)

	kwFieldMapping := bleve.NewTextFieldMapping()
	kwFieldMapping.Analyzer = keyword_analyzer.Name

	eventMapping.AddFieldMappingsAt("url", kwFieldMapping)
	eventMapping.AddFieldMappingsAt("category", kwFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = eventMapping
	mapping.DefaultAnalyzer = en.AnalyzerName

	return mapping
}
