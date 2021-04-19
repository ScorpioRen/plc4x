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
type BACnetErrorConfirmedPrivateTransfer struct {
	Parent *BACnetError
}

// The corresponding interface
type IBACnetErrorConfirmedPrivateTransfer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorConfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetErrorConfirmedPrivateTransfer) InitializeParent(parent *BACnetError) {
}

func NewBACnetErrorConfirmedPrivateTransfer() *BACnetError {
	child := &BACnetErrorConfirmedPrivateTransfer{
		Parent: NewBACnetError(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetErrorConfirmedPrivateTransfer(structType interface{}) *BACnetErrorConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetErrorConfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetErrorConfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorConfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorConfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorConfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetErrorConfirmedPrivateTransfer"
}

func (m *BACnetErrorConfirmedPrivateTransfer) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorConfirmedPrivateTransfer) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorConfirmedPrivateTransfer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorConfirmedPrivateTransferParse(io utils.ReadBuffer) (*BACnetError, error) {
	io.PullContext("BACnetErrorConfirmedPrivateTransfer")

	io.CloseContext("BACnetErrorConfirmedPrivateTransfer")

	// Create a partially initialized instance
	_child := &BACnetErrorConfirmedPrivateTransfer{
		Parent: &BACnetError{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetErrorConfirmedPrivateTransfer) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("BACnetErrorConfirmedPrivateTransfer")

		io.PopContext("BACnetErrorConfirmedPrivateTransfer")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetErrorConfirmedPrivateTransfer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetErrorConfirmedPrivateTransfer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetErrorConfirmedPrivateTransfer) String() string {
	return string(m.Box("", 120))
}

func (m BACnetErrorConfirmedPrivateTransfer) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetErrorConfirmedPrivateTransfer"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
