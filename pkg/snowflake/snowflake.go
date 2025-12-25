package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

type Snowflake struct {
	nodeNo int64
	node   *snowflake.Node
}

func NewSnowflake(nodeNo int64) (*Snowflake, error) {
	node, err := snowflake.NewNode(nodeNo)
	if err != nil {
		return nil, err
	}
	var s = &Snowflake{
		nodeNo: nodeNo,
		node:   node,
	}
	return s, nil
}

func (s *Snowflake) GenerateID() int64 {
	return s.node.Generate().Int64()
}

func (s *Snowflake) GenerateStringID() string {
	return s.node.Generate().String()
}

func (s *Snowflake) GetNodeNo() int64 {
	return s.nodeNo
}

func (s *Snowflake) GetNode() *snowflake.Node {
	return s.node
}
