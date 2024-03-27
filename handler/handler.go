package handler

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}

}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml
	var configs []YamlConfig
	if err := yaml.Unmarshal([]byte(yml), &configs); err != nil {
		panic(err)
	}

	// create a map
	pathsToUrls := map[string]string{}
	for _, config := range configs {
		pathsToUrls[config.Path] = config.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
