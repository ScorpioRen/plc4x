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
type DescriptionRequest struct {
	HpaiControlEndpoint *HPAIControlEndpoint
	Parent              *KnxNetIpMessage
}

// The corresponding interface
type IDescriptionRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DescriptionRequest) MsgType() uint16 {
	return 0x0203
}

func (m *DescriptionRequest) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDescriptionRequest(hpaiControlEndpoint *HPAIControlEndpoint) *KnxNetIpMessage {
	child := &DescriptionRequest{
		HpaiControlEndpoint: hpaiControlEndpoint,
		Parent:              NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDescriptionRequest(structType interface{}) *DescriptionRequest {
	castFunc := func(typ interface{}) *DescriptionRequest {
		if casted, ok := typ.(DescriptionRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*DescriptionRequest); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastDescriptionRequest(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastDescriptionRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DescriptionRequest) GetTypeName() string {
	return "DescriptionRequest"
}

func (m *DescriptionRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DescriptionRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.LengthInBits()

	return lengthInBits
}

func (m *DescriptionRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DescriptionRequestParse(io utils.ReadBuffer) (*KnxNetIpMessage, error) {
	io.PullContext("DescriptionRequest")

	// Simple Field (hpaiControlEndpoint)
	hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(io)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field")
	}

	io.CloseContext("DescriptionRequest")

	// Create a partially initialized instance
	_child := &DescriptionRequest{
		HpaiControlEndpoint: hpaiControlEndpoint,
		Parent:              &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DescriptionRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("DescriptionRequest")

		// Simple Field (hpaiControlEndpoint)
		_hpaiControlEndpointErr := m.HpaiControlEndpoint.Serialize(io)
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		io.PopContext("DescriptionRequest")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *DescriptionRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "hpaiControlEndpoint":
				var data HPAIControlEndpoint
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.HpaiControlEndpoint = &data
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

func (m *DescriptionRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.HpaiControlEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiControlEndpoint"}}); err != nil {
		return err
	}
	return nil
}

func (m DescriptionRequest) String() string {
	return string(m.Box("", 120))
}

func (m DescriptionRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "DescriptionRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.HpaiControlEndpoint.Box("hpaiControlEndpoint", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
