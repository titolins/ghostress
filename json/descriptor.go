package json

import (
	"encoding/json"
	"io/ioutil"
)

// DescriptorField -> represents a json field
type DescriptorField struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Generate bool              `json:"generate"`
	Value    interface{}       `json:"value"`
	Fields   []DescriptorField `json:"fields"`
	Format   DescriptorFormat  `json:"format"`
}

// DescriptorFormat -> Describes the format of the JSON payload to be sent
type DescriptorFormat struct {
	Shape string `json:"shape"`
	Size  int    `json:"size"`
}

// Descriptor -> an array of DescriptorField
type Descriptor struct {
	Fields []DescriptorField `json:"fields"`
	Format DescriptorFormat  `json:"options"`
}

// NewDescriptor -> constructor for a Descriptor
func NewDescriptor(descriptorPath string) *Descriptor {
	var descriptor Descriptor
	descriptorData, err := ioutil.ReadFile(descriptorPath)
	if err != nil {
		panic("Couldn't read descriptor file")
	}
	err = json.Unmarshal(descriptorData, &descriptor)
	if err != nil {
		panic("Error trying to unmarshal json")
	}
	return &descriptor
}
