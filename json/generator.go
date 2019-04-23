package json

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"math/rand"
	"time"
)

const (
	defaultStringSize    = 10
	defaultIntSize       = 5
	defaultJSONArraySize = 3
)

type jsonObject map[string]interface{}
type jsonArray []jsonObject

// Generator -> Receives a Descriptor and generates the appropriate json payload
type Generator struct {
	Descriptor *Descriptor
}

// randomInt -> generates an int between a min and max values
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// generateRandomIntLen -> Generates a random int with a given length l
func generateRandomIntLen(l int) int {
	min := int(math.Pow10(l - 1))
	max := int((math.Pow10(l) - 1))
	return randomInt(min, max)
}

// generateRandomBytes -> returns a random byte slice
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// generateRandomString -> generates a random base64 encoded string from random
// bytes
func generateRandomString(s int) string {
	b, err := generateRandomBytes(s)
	if err != nil {
		panic("Failed to generate random bytes")
	}

	return base64.URLEncoding.EncodeToString(b)
}

func buildField(field DescriptorField) interface{} {
	// we have to implement functionality for fields with nested fields
	// for now, let's just build a single field
	if field.Generate {
		size := field.Format.Size
		switch field.Type {
		case "string":
			if size <= 0 {
				size = defaultStringSize
			}
			return generateRandomString(size)
		case "int":
			if size <= 0 {
				size = defaultIntSize
			}
			return generateRandomIntLen(size)
		default:
			return nil
		}
	} else {
		return field.Value
	}
}

// horrible solution, but the interface didn't work for some reason..
// worth checking this out better later
func buildNestedField(field DescriptorField) jsonObject {
	obj := make(jsonObject)

	for _, f := range field.Fields {
		if f.Fields != nil {
			obj[f.Name] = buildNestedField(f)
		} else {
			obj[f.Name] = buildField(f)
		}
	}
	return obj
}

// buildObject -> Generates the fields and builds a payload object
func (gen *Generator) buildObject() jsonObject {
	rand.Seed(time.Now().UnixNano())
	obj := make(jsonObject)

	for _, f := range gen.Descriptor.Fields {
		if f.Fields != nil {
			obj[f.Name] = buildNestedField(f)
		} else {
			obj[f.Name] = buildField(f)
		}
	}
	return obj
}

// GetData -> Returns the result of json marshal on the object
func (gen *Generator) GetData() []byte {
	if gen.Descriptor.Format.Shape == "list" {
		arrSize := gen.Descriptor.Format.Size
		if arrSize <= 0 {
			arrSize = int(defaultJSONArraySize)
		}
		arr := make(jsonArray, arrSize)
		for i := 0; i < arrSize; i++ {
			arr[i] = gen.buildObject()
		}

		b, err := json.Marshal(arr)
		if err != nil {
			panic("Failed to marshal obj")
		}

		return b
	}

	obj := gen.buildObject()

	b, err := json.Marshal(obj)
	if err != nil {
		panic("Failed to marshal obj")
	}

	return b
}

// BuildPayload -> main function, responsible for generating the json payload
func (gen *Generator) BuildPayload() *interface{} {
	return nil
	/*
		if gen.Descriptor.Options.Format.Shape == "list" {
			var payload []interface{}
			for i, f := range gen.Descriptor.Fields {
				payload = append(payload, buildField(f))
			}
			return payload
		} else {
			var payload interface{}
			return payload
		}
	*/
}
