package main

import (
	"flag"
	"fmt"
	_"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Path string `yaml:"path"`
	Url string  `yaml:"url"`
}

func MapHandler(pathsToUrls []YamlConfig) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, v := range pathsToUrls {
			fmt.Println("This is path ", v.Url)
			if v.Path == request.URL.Path {
				http.Redirect(writer, request, v.Url, http.StatusFound)
			}
			_, _ = fmt.Fprint(writer, "fallbackResponse")
		}
	}
}

func YAMLHandler(file string) http.HandlerFunc {
	var yamlConfig []YamlConfig

	yamlFile, err := ioutil.ReadFile(file)

	err = yaml.Unmarshal(yamlFile, &yamlConfig)

	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Result: %v\n", yamlConfig)

	return MapHandler(yamlConfig)
}

func main() {
	filename := flag.String("filename", "redirect.yaml", "blah blah")

	flag.Parse()

	r := http.NewServeMux()

	r.HandleFunc("/", YAMLHandler(*filename))

	err := http.ListenAndServe(":9090", r)

	if err != nil {
		log.Fatal(err)
	}
}
