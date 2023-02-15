package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func GenerateNew() (*string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Printf("failed to generate a new snowflake node")
		return nil, err
	}

	id := node.Generate().String()

	return &id, err
}
