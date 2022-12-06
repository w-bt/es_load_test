package item

const (
	mappings = `{
		  "settings": {
			"index": {
			  "number_of_shards": 3,
			  "number_of_replicas": 1,
			  "analysis": {
				"filter": {
				  "snowball_stemmer": {
					"type": "stemmer",
					"language": "english"
				  },
				  "shingles": {
					"max_shingle_size": "2",
					"type": "shingle",
					"token_separator": ""
				  },
				  "search_shingles": {
					"token_separator": "",
					"output_unigrams_if_no_shingles": "true",
					"output_unigrams": "false",
					"type": "shingle"
				  }
				},
				"analyzer": {
				  "index_analyzer": {
					"filter": [
					  "lowercase",
					  "asciifolding",
					  "shingles",
					  "snowball_stemmer"
					],
					"char_filter": [
					  "ampersand",
					  "single_quote",
					  "double_quote"
					],
					"type": "custom",
					"tokenizer": "standard"
				  }
				},
				"char_filter": {
				  "ampersand": {
					"type": "mapping",
					"mappings": [
					  "&=> and "
					]
				  },
				  "single_quote": {
					"type": "mapping",
					"mappings": [
					  "' => "
					]
				  },
				  "double_quote": {
					"type": "mapping",
					"mappings": [
					  "\" => "
					]
				  },
				  "comma": {
					"type": "mapping",
					"mappings": [",=> "]
				  }
				}
			  }
			}
		  },
		  "mappings": {
			"properties": {
			  "id": {
				"type": "long"
			  },
			  "brand_name": {
				"ignore_above": 300,
				"type": "keyword",
				"fields": {
				  "analyzed": {
					"analyzer": "index_analyzer",
					"type": "text"
				  }
				}
			  },
			  "price": {
				"type": "float"
			  },
			  "tags": {
				"ignore_above": 3000,
				"type": "keyword",
				"fields": {
				  "analyzed": {
					"analyzer": "index_analyzer",
					"type": "text"
				  }
				}
			  },
			  "name": {
				"ignore_above": 300,
				"type": "keyword",
				"fields": {
				  "analyzed": {
					"analyzer": "index_analyzer",
					"type": "text"
				  }
				}
			  },
			  "category_names": {
				"ignore_above": 300,
				"type": "keyword",
				"fields": {
				  "analyzed": {
					"analyzer": "index_analyzer",
					"type": "text"
				  }
				}
			  },
			  "code": {
				"type": "keyword"
			  }
			}
		  }
		}`
)
