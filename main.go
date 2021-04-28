package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os/exec"
)

var (
	services []string
)

type Attrs struct {
	Path      string   `yaml:"path,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
}

func executeService(path string) (out []byte, err error) {
	out, err = exec.Command("/bin/sh", path).Output()
	if err != nil {
		return
	}

	return
}

func readYamlFile(path string) ([]byte, error) {
	var data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func unmarshalYaml(contents []byte) (m map[string]Attrs, err error) {
	m = make(map[string]Attrs)
	err = yaml.Unmarshal(contents, &m)
	if err != nil {
		return
	}
	return
}

func buildDepGraph(m map[string]Attrs) {
	for k, _ := range m {
		services = append(services, k)
	}

	for k, v := range m {
		for _, dep := range v.DependsOn {
			if findService(dep) {
				addDependency(vertex{name: k}, vertex{name: dep})
			} else {
				fmt.Printf("Unknown service %v", dep)
			}
		}
	}
}

func main() {
	contents, err := readYamlFile("test.yml")
	if err != nil {
		panic(err)
	}

	m, err := unmarshalYaml(contents)
	if err != nil {
		panic(err)
	}

	buildDepGraph(m)

	var resolved = make([]vertex, 0)
	determineOrder(vertices[2], &resolved)
	for _, value := range resolved {
		fmt.Println(value)
	}
}
