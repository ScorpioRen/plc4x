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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = 2

// The data-structure of this message
type CEMIAdditionalInformationRelativeTimestamp struct {
	RelativeTimestamp *RelativeTimestamp
	Parent            *CEMIAdditionalInformation
}

// The corresponding interface
type ICEMIAdditionalInformationRelativeTimestamp interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CEMIAdditionalInformationRelativeTimestamp) AdditionalInformationType() uint8 {
	return 0x04
}

func (m *CEMIAdditionalInformationRelativeTimestamp) InitializeParent(parent *CEMIAdditionalInformation) {
}

func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp *RelativeTimestamp) *CEMIAdditionalInformation {
	child := &CEMIAdditionalInformationRelativeTimestamp{
		RelativeTimestamp: relativeTimestamp,
		Parent:            NewCEMIAdditionalInformation(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCEMIAdditionalInformationRelativeTimestamp(structType interface{}) *CEMIAdditionalInformationRelativeTimestamp {
	castFunc := func(typ interface{}) *CEMIAdditionalInformationRelativeTimestamp {
		if casted, ok := typ.(CEMIAdditionalInformationRelativeTimestamp); ok {
			return &casted
		}
		if casted, ok := typ.(*CEMIAdditionalInformationRelativeTimestamp); ok {
			return casted
		}
		if casted, ok := typ.(CEMIAdditionalInformation); ok {
			return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
		}
		if casted, ok := typ.(*CEMIAdditionalInformation); ok {
			return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
	return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *CEMIAdditionalInformationRelativeTimestamp) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (len)
	lengthInBits += 8

	// Simple field (relativeTimestamp)
	lengthInBits += m.RelativeTimestamp.LengthInBits()

	return lengthInBits
}

func (m *CEMIAdditionalInformationRelativeTimestamp) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CEMIAdditionalInformationRelativeTimestampParse(io utils.ReadBuffer) (*CEMIAdditionalInformation, error) {
	io.PullContext("CEMIAdditionalInformationRelativeTimestamp")

	// Const Field (len)
	len, _lenErr := io.ReadUint8("len", 8)
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field")
	}
	if len != CEMIAdditionalInformationRelativeTimestamp_LEN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CEMIAdditionalInformationRelativeTimestamp_LEN) + " but got " + fmt.Sprintf("%d", len))
	}

	// Simple Field (relativeTimestamp)
	relativeTimestamp, _relativeTimestampErr := RelativeTimestampParse(io)
	if _relativeTimestampErr != nil {
		return nil, errors.Wrap(_relativeTimestampErr, "Error parsing 'relativeTimestamp' field")
	}

	io.CloseContext("CEMIAdditionalInformationRelativeTimestamp")

	// Create a partially initialized instance
	_child := &CEMIAdditionalInformationRelativeTimestamp{
		RelativeTimestamp: relativeTimestamp,
		Parent:            &CEMIAdditionalInformation{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *CEMIAdditionalInformationRelativeTimestamp) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("CEMIAdditionalInformationRelativeTimestamp")

		// Const Field (len)
		_lenErr := io.WriteUint8("len", 8, 2)
		if _lenErr != nil {
			return errors.Wrap(_lenErr, "Error serializing 'len' field")
		}

		// Simple Field (relativeTimestamp)
		_relativeTimestampErr := m.RelativeTimestamp.Serialize(io)
		if _relativeTimestampErr != nil {
			return errors.Wrap(_relativeTimestampErr, "Error serializing 'relativeTimestamp' field")
		}

		io.PopContext("CEMIAdditionalInformationRelativeTimestamp")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relativeTimestamp":
				var data RelativeTimestamp
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.RelativeTimestamp = &data
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

func (m *CEMIAdditionalInformationRelativeTimestamp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.RelativeTimestamp, xml.StartElement{Name: xml.Name{Local: "relativeTimestamp"}}); err != nil {
		return err
	}
	return nil
}

func (m CEMIAdditionalInformationRelativeTimestamp) String() string {
	return string(m.Box("", 120))
}

func (m CEMIAdditionalInformationRelativeTimestamp) Box(name string, width int) utils.AsciiBox {
	boxName := "CEMIAdditionalInformationRelativeTimestamp"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Const Field (len)
		boxes = append(boxes, utils.BoxAnything("Len", uint8(2), -1))
		// Complex field (case complex)
		boxes = append(boxes, m.RelativeTimestamp.Box("relativeTimestamp", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
