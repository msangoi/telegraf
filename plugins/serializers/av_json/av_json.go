package av_json

import (
	"encoding/json"
	"strconv"

	"github.com/influxdata/telegraf"
)

type AvJsonSerializer struct {
}

func (s *AvJsonSerializer) Serialize(metric telegraf.Metric) ([]string, error) {

	// Get the metric name
	name := metric.Name()

	// Ignore tags for now ("host"?)

	// Convert UnixNano to Unix timestamps
	timestamp := strconv.FormatInt(metric.UnixNano()/1000000, 10)

	// json message: { <timestamp>: { "<dataname>": <value>, "<dataname>": <value>}, ... }
	data := make(map[string]interface{})
	for field_name, value := range metric.Fields() {
		data[name+"."+field_name] = value
	}
	b, _ := json.Marshal(map[string]interface{}{timestamp: data})

	return []string{string(b[:])}, nil
}
