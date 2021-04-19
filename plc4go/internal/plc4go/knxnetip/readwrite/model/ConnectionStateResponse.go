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
type ConnectionStateResponse struct {
	CommunicationChannelId uint8
	Status                 Status
	Parent                 *KnxNetIpMessage
}

// The corresponding interface
type IConnectionStateResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionStateResponse) MsgType() uint16 {
	return 0x0208
}

func (m *ConnectionStateResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewConnectionStateResponse(communicationChannelId uint8, status Status) *KnxNetIpMessage {
	child := &ConnectionStateResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastConnectionStateResponse(structType interface{}) *ConnectionStateResponse {
	castFunc := func(typ interface{}) *ConnectionStateResponse {
		if casted, ok := typ.(ConnectionStateResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionStateResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastConnectionStateResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastConnectionStateResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionStateResponse) GetTypeName() string {
	return "ConnectionStateResponse"
}

func (m *ConnectionStateResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ConnectionStateResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionStateResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionStateResponseParse(io utils.ReadBuffer) (*KnxNetIpMessage, error) {
	io.PullContext("ConnectionStateResponse")

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := io.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (status)
	status, _statusErr := StatusParse(io)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	io.CloseContext("ConnectionStateResponse")

	// Create a partially initialized instance
	_child := &ConnectionStateResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ConnectionStateResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("ConnectionStateResponse")

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.CommunicationChannelId)
		_communicationChannelIdErr := io.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		_statusErr := m.Status.Serialize(io)
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		io.PopContext("ConnectionStateResponse")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ConnectionStateResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "communicationChannelId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommunicationChannelId = data
			case "status":
				var data Status
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
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

func (m *ConnectionStateResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	return nil
}

func (m ConnectionStateResponse) String() string {
	return string(m.Box("", 120))
}

func (m ConnectionStateResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "ConnectionStateResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("CommunicationChannelId", m.CommunicationChannelId, -1))
		// Complex field (case complex)
		boxes = append(boxes, m.Status.Box("status", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
