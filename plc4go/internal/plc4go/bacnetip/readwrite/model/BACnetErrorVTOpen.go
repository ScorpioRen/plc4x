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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetErrorVTOpen struct {
	Parent *BACnetError
}

// The corresponding interface
type IBACnetErrorVTOpen interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorVTOpen) ServiceChoice() uint8 {
	return 0x15
}

func (m *BACnetErrorVTOpen) InitializeParent(parent *BACnetError) {
}

func NewBACnetErrorVTOpen() *BACnetError {
	child := &BACnetErrorVTOpen{
		Parent: NewBACnetError(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetErrorVTOpen(structType interface{}) *BACnetErrorVTOpen {
	castFunc := func(typ interface{}) *BACnetErrorVTOpen {
		if casted, ok := typ.(BACnetErrorVTOpen); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorVTOpen); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorVTOpen(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorVTOpen(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorVTOpen) GetTypeName() string {
	return "BACnetErrorVTOpen"
}

func (m *BACnetErrorVTOpen) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorVTOpen) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorVTOpen) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorVTOpenParse(io utils.ReadBuffer) (*BACnetError, error) {
	io.PullContext("BACnetErrorVTOpen")

	io.CloseContext("BACnetErrorVTOpen")

	// Create a partially initialized instance
	_child := &BACnetErrorVTOpen{
		Parent: &BACnetError{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetErrorVTOpen) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("BACnetErrorVTOpen")

		io.PopContext("BACnetErrorVTOpen")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetErrorVTOpen) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetErrorVTOpen) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetErrorVTOpen) String() string {
	return string(m.Box("", 120))
}

func (m BACnetErrorVTOpen) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetErrorVTOpen"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
