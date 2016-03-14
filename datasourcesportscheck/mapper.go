package datasourcesportscheck

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzers/keyword_analyzer"
    
	//"github.com/blevesearch/bleve/analysis/language/en"
    
    "github.com/blevesearch/blevex/lang/de"
)

// BuildMapping ... maps the data struts into the db.
func BuildMapping() *bleve.IndexMapping {
    
	deFieldMapping := bleve.NewTextFieldMapping()
	deFieldMapping.Analyzer = de.AnalyzerName
   
    

	productMapping := bleve.NewDocumentMapping()
	productMapping.AddFieldMappingsAt("productname", deFieldMapping)
    productMapping.AddFieldMappingsAt("productshort", deFieldMapping)
	productMapping.AddFieldMappingsAt("productlong", deFieldMapping)

	kwFieldMapping := bleve.NewTextFieldMapping()
	kwFieldMapping.Analyzer = keyword_analyzer.Name

	productMapping.AddFieldMappingsAt("productid", kwFieldMapping)
	productMapping.AddFieldMappingsAt("productcategory", kwFieldMapping)
    productMapping.AddFieldMappingsAt("colour", kwFieldMapping)
    productMapping.AddFieldMappingsAt("size", kwFieldMapping)
    productMapping.AddFieldMappingsAt("brand", kwFieldMapping)
    productMapping.AddFieldMappingsAt("price", kwFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = productMapping
	mapping.DefaultAnalyzer = de.AnalyzerName
    

	return mapping
}
