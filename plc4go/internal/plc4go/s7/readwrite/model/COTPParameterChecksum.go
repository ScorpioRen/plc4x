//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type COTPParameterChecksum struct {
	Crc    uint8
	Parent *COTPParameter
}

// The corresponding interface
type ICOTPParameterChecksum interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPParameterChecksum) ParameterType() uint8 {
	return 0xC3
}

func (m *COTPParameterChecksum) InitializeParent(parent *COTPParameter) {
}

func NewCOTPParameterChecksum(crc uint8) *COTPParameter {
	child := &COTPParameterChecksum{
		Crc:    crc,
		Parent: NewCOTPParameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPParameterChecksum(structType interface{}) *COTPParameterChecksum {
	castFunc := func(typ interface{}) *COTPParameterChecksum {
		if casted, ok := typ.(COTPParameterChecksum); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPParameterChecksum); ok {
			return casted
		}
		if casted, ok := typ.(COTPParameter); ok {
			return CastCOTPParameterChecksum(casted.Child)
		}
		if casted, ok := typ.(*COTPParameter); ok {
			return CastCOTPParameterChecksum(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPParameterChecksum) GetTypeName() string {
	return "COTPParameterChecksum"
}

func (m *COTPParameterChecksum) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPParameterChecksum) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (crc)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPParameterChecksum) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPParameterChecksumParse(io utils.ReadBuffer) (*COTPParameter, error) {
	io.PullContext("COTPParameterChecksum")

	// Simple Field (crc)
	crc, _crcErr := io.ReadUint8("crc", 8)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}

	io.CloseContext("COTPParameterChecksum")

	// Create a partially initialized instance
	_child := &COTPParameterChecksum{
		Crc:    crc,
		Parent: &COTPParameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPParameterChecksum) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("COTPParameterChecksum")

		// Simple Field (crc)
		crc := uint8(m.Crc)
		_crcErr := io.WriteUint8("crc", 8, (crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		io.PopContext("COTPParameterChecksum")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPParameterChecksum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "crc":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Crc = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *COTPParameterChecksum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
		return err
	}
	return nil
}

func (m COTPParameterChecksum) String() string {
	return string(m.Box("", 120))
}

func (m COTPParameterChecksum) Box(name string, width int) utils.AsciiBox {
	boxName := "COTPParameterChecksum"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Crc", m.Crc, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
