package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const defaultTimestampFormat = time.RFC3339

var markers = [2]string{"sourcecode", "golang"}

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["date"] = entry.Time.Format(defaultTimestampFormat)
	data["message"] = entry.Message
	data["level"] = entry.Level.String()
	data["@marker"] = markers

	serialized, err := jsonMarshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
