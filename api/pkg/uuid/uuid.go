package uuid

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func GenUUID5(data any) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(
			fmt.Sprintf("failed to serialize data: %s", err.Error()),
		)
	}
	return uuid.NewSHA1(
		uuid.New(),
		b,
	).String()
}

func Validate(id string) error {
	return uuid.Validate(id)
}
