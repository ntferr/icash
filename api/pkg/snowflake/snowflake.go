package snowflake

import (
	"errors"
	"log"

	"github.com/bwmarrin/snowflake"
)

var errOnValidateSnowflake = errors.New("error on validate snowflake")

func GenerateNew() (*string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Printf("failed to generate a new snowflake node")
		return nil, err
	}

	id := node.Generate().String()

	return &id, err
}

func Validate(id string) error {
	idLength := len([]rune(id))
	if idLength != 19 {
		log.Println("provide a valid id")
		return errOnValidateSnowflake
	}

	return nil
}
