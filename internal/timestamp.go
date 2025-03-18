package suntime

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return t.formatUnix()
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	stringData := ""
	err := json.Unmarshal(data, &stringData)
	if err != nil {
		return fmt.Errorf("unmarshalling field: %v", err)
	}


	return t.parseUnix(stringData)
}

func (t Timestamp) formatUnix() ([]byte, error) {
	return strconv.AppendInt(nil, t.Unix(), 10), nil
}

func (t *Timestamp) parseUnix(data string) error {
	i, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("parsing unix time: %v", err)
	}

	t.Time = time.Unix(int64(i), 0)
	return nil
}
