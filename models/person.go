package models

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Person struct {
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	HairColour string    `json:"hair_colour"`
	Friends    []*Friend `json:"friends"`
}

type Friend Person

type PersonWrapper struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	HairColour string `json:"hair_colour"`
	Friends    []byte `json:"friends"`
}

func (p *Person) UnmarshalJSON(data []byte) error {
	wrapper := &PersonWrapper{}
	if err := json.Unmarshal(data, wrapper); err != nil {
		return err
	}

	p.Name = wrapper.Name
	p.Age = wrapper.Age
	p.HairColour = wrapper.HairColour
	p.Friends = []*Friend{}

	if len(wrapper.Friends) > 0 {
		buf := bytes.NewBuffer(wrapper.Friends)
		reader, err := gzip.NewReader(buf)
		if err != nil {
			return err
		}

		if err := reader.Close(); err != nil {
			return err
		}

		raw, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		return json.Unmarshal(raw, &p.Friends)
	}

	return nil
}

func (p *Person) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(p.Friends)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	defer writer.Close()

	if _, err := writer.Write(raw); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper := &PersonWrapper{
		Name:       p.Name,
		Age:        p.Age,
		HairColour: p.HairColour,
		Friends:    buf.Bytes(),
	}

	return json.Marshal(wrapper)
}
