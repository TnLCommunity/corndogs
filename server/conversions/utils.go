package conversions

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func StructToProto(source interface{}, dest proto.Message) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = protojson.Unmarshal(bytes, dest)
	if err != nil {
		return err
	}
	return nil
}
