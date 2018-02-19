package writer

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func WriteToFS(path string, entity interface{}) error {
	res, _ := json.Marshal(entity)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		ioutil.WriteFile(path, res, 0777)
	}
	return nil
}
