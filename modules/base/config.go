package base

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type configLoad struct {
	ShowPwd bool
}

//Config Load File and store in this
var Config configLoad

func init() {

	source, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &Config)
	if err != nil {
		panic(err)
	}
}
