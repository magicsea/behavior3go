package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//原生工程json类型
type RawProjectCfg struct {
	Name string       `json:"name"`
	Data BTProjectCfg `json:"data"`
	Path string       `json:"path"`
}

//加载原生工程
func LoadRawProjectCfg(path string) (*RawProjectCfg, bool) {

	var project RawProjectCfg
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("LoadRawProjectCfg fail:", err)
		return nil, false
	}
	err = json.Unmarshal(file, &project)
	if err != nil {
		fmt.Println("LoadRawProjectCfg fail, ummarshal:", err, len(file))
		return nil, false
	}

	//fmt.Println("load tree:", tree.Title, " nodes:", len(tree.Nodes))
	return &project, true
}
