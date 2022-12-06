package config

type Elasticsearch struct {
	Hosts         []string
	HTTPTimeoutMS int
}

func newESConfig() Elasticsearch {
	esConfig := Elasticsearch{
		Hosts:         getStringArray("ES_CLUSTER_HOSTS"),
		HTTPTimeoutMS: getIntOrDefault("ES_HTTP_TIMEOUT_MS", 1000),
	}

	return esConfig
}
