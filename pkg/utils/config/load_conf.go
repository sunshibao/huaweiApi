package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var (
	NL  = []byte{'\n'}
	ANT = []byte{'#'}
)

func MustLoadConf(conf interface{}, confName string) {
	data, err := loadFileData(confName)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	err = decoder.Decode(conf)
	if err != nil {
		logrus.Fatal("Parse conf failed:", err)
	}
}

func loadFileData(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Fatal("Load conf failed:", err)
	}
	data = trimComments(data)
	logrus.Info("Load conf succeeded with path ", filePath)
	logrus.Info(string(data))
	return data, err
}

func trimComments(data []byte) (data1 []byte) {

	configLines := bytes.Split(data, NL)
	for k, line := range configLines {
		configLines[k] = trimCommentsLine(line)
	}
	return bytes.Join(configLines, NL)
}

func trimCommentsLine(line []byte) []byte {

	var newLine []byte
	var i, quoteCount int
	lastIdx := len(line) - 1
	for i = 0; i <= lastIdx; i++ {
		if line[i] == '\\' {
			if i != lastIdx && (line[i+1] == '\\' || line[i+1] == '"') {
				newLine = append(newLine, line[i], line[i+1])
				i++
				continue
			}
		}
		if line[i] == '"' {
			quoteCount++
		}
		if line[i] == '#' {
			if quoteCount%2 == 0 {
				break
			}
		}
		newLine = append(newLine, line[i])
	}
	return newLine
}
