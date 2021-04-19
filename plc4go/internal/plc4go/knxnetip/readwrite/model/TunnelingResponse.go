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
type TunnelingResponse struct {
	TunnelingResponseDataBlock *TunnelingResponseDataBlock
	Parent                     *KnxNetIpMessage
}

// The corresponding interface
type ITunnelingResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TunnelingResponse) MsgType() uint16 {
	return 0x0421
}

func (m *TunnelingResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewTunnelingResponse(tunnelingResponseDataBlock *TunnelingResponseDataBlock) *KnxNetIpMessage {
	child := &TunnelingResponse{
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
		Parent:                     NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastTunnelingResponse(structType interface{}) *TunnelingResponse {
	castFunc := func(typ interface{}) *TunnelingResponse {
		if casted, ok := typ.(TunnelingResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*TunnelingResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastTunnelingResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastTunnelingResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *TunnelingResponse) GetTypeName() string {
	return "TunnelingResponse"
}

func (m *TunnelingResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *TunnelingResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (tunnelingResponseDataBlock)
	lengthInBits += m.TunnelingResponseDataBlock.LengthInBits()

	return lengthInBits
}

func (m *TunnelingResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TunnelingResponseParse(io utils.ReadBuffer) (*KnxNetIpMessage, error) {
	io.PullContext("TunnelingResponse")

	// Simple Field (tunnelingResponseDataBlock)
	tunnelingResponseDataBlock, _tunnelingResponseDataBlockErr := TunnelingResponseDataBlockParse(io)
	if _tunnelingResponseDataBlockErr != nil {
		return nil, errors.Wrap(_tunnelingResponseDataBlockErr, "Error parsing 'tunnelingResponseDataBlock' field")
	}

	io.CloseContext("TunnelingResponse")

	// Create a partially initialized instance
	_child := &TunnelingResponse{
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
		Parent:                     &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *TunnelingResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("TunnelingResponse")

		// Simple Field (tunnelingResponseDataBlock)
		_tunnelingResponseDataBlockErr := m.TunnelingResponseDataBlock.Serialize(io)
		if _tunnelingResponseDataBlockErr != nil {
			return errors.Wrap(_tunnelingResponseDataBlockErr, "Error serializing 'tunnelingResponseDataBlock' field")
		}

		io.PopContext("TunnelingResponse")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *TunnelingResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "tunnelingResponseDataBlock":
				var data TunnelingResponseDataBlock
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TunnelingResponseDataBlock = &data
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

func (m *TunnelingResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.TunnelingResponseDataBlock, xml.StartElement{Name: xml.Name{Local: "tunnelingResponseDataBlock"}}); err != nil {
		return err
	}
	return nil
}

func (m TunnelingResponse) String() string {
	return string(m.Box("", 120))
}

func (m TunnelingResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "TunnelingResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.TunnelingResponseDataBlock.Box("tunnelingResponseDataBlock", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
