package proj1util

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type ServerParams struct {
	ServerId, Protocol, IpAddress, Port, StorageContainer string
}

func ConfigureParameters(parameterFile string) (params ServerParams) {
	config, fErr := ioutil.ReadFile(parameterFile)
	if fErr != nil {
		panic(fErr)
	}
	json.Unmarshal(config, &params)
	fmt.Println(params)
	return params
}