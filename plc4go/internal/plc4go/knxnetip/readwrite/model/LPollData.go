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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type LPollData struct {
	SourceAddress          *KnxAddress
	TargetAddress          []int8
	NumberExpectedPollData uint8
	Parent                 *LDataFrame
}

// The corresponding interface
type ILPollData interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LPollData) NotAckFrame() bool {
	return true
}

func (m *LPollData) Polling() bool {
	return true
}

func (m *LPollData) InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.Parent.FrameType = frameType
	m.Parent.NotRepeated = notRepeated
	m.Parent.Priority = priority
	m.Parent.AcknowledgeRequested = acknowledgeRequested
	m.Parent.ErrorFlag = errorFlag
}

func NewLPollData(sourceAddress *KnxAddress, targetAddress []int8, numberExpectedPollData uint8, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
	child := &LPollData{
		SourceAddress:          sourceAddress,
		TargetAddress:          targetAddress,
		NumberExpectedPollData: numberExpectedPollData,
		Parent:                 NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLPollData(structType interface{}) *LPollData {
	castFunc := func(typ interface{}) *LPollData {
		if casted, ok := typ.(LPollData); ok {
			return &casted
		}
		if casted, ok := typ.(*LPollData); ok {
			return casted
		}
		if casted, ok := typ.(LDataFrame); ok {
			return CastLPollData(casted.Child)
		}
		if casted, ok := typ.(*LDataFrame); ok {
			return CastLPollData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LPollData) GetTypeName() string {
	return "LPollData"
}

func (m *LPollData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LPollData) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.LengthInBits()

	// Array field
	if len(m.TargetAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.TargetAddress))
	}

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (numberExpectedPollData)
	lengthInBits += 6

	return lengthInBits
}

func (m *LPollData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LPollDataParse(io utils.ReadBuffer) (*LDataFrame, error) {
	io.PullContext("LPollData")

	// Simple Field (sourceAddress)
	sourceAddress, _sourceAddressErr := KnxAddressParse(io)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}

	// Array field (targetAddress)
	io.PullContext("targetAddress")
	// Count array
	targetAddress := make([]int8, uint16(2))
	for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'targetAddress' field")
		}
		targetAddress[curItem] = _item
	}
	io.CloseContext("targetAddress")

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (numberExpectedPollData)
	numberExpectedPollData, _numberExpectedPollDataErr := io.ReadUint8("numberExpectedPollData", 6)
	if _numberExpectedPollDataErr != nil {
		return nil, errors.Wrap(_numberExpectedPollDataErr, "Error parsing 'numberExpectedPollData' field")
	}

	io.CloseContext("LPollData")

	// Create a partially initialized instance
	_child := &LPollData{
		SourceAddress:          sourceAddress,
		TargetAddress:          targetAddress,
		NumberExpectedPollData: numberExpectedPollData,
		Parent:                 &LDataFrame{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LPollData) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("LPollData")

		// Simple Field (sourceAddress)
		_sourceAddressErr := m.SourceAddress.Serialize(io)
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Array Field (targetAddress)
		if m.TargetAddress != nil {
			io.PushContext("targetAddress")
			for _, _element := range m.TargetAddress {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'targetAddress' field")
				}
			}
			io.PopContext("targetAddress")
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8("reserved", 4, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (numberExpectedPollData)
		numberExpectedPollData := uint8(m.NumberExpectedPollData)
		_numberExpectedPollDataErr := io.WriteUint8("numberExpectedPollData", 6, (numberExpectedPollData))
		if _numberExpectedPollDataErr != nil {
			return errors.Wrap(_numberExpectedPollDataErr, "Error serializing 'numberExpectedPollData' field")
		}

		io.PopContext("LPollData")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LPollData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sourceAddress":
				var data KnxAddress
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAddress = &data
			case "targetAddress":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.TargetAddress = utils.ByteArrayToInt8Array(_decoded[0:_len])
			case "numberExpectedPollData":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NumberExpectedPollData = data
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

func (m *LPollData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.SourceAddress, xml.StartElement{Name: xml.Name{Local: "sourceAddress"}}); err != nil {
		return err
	}
	_encodedTargetAddress := hex.EncodeToString(utils.Int8ArrayToByteArray(m.TargetAddress))
	_encodedTargetAddress = strings.ToUpper(_encodedTargetAddress)
	if err := e.EncodeElement(_encodedTargetAddress, xml.StartElement{Name: xml.Name{Local: "targetAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.NumberExpectedPollData, xml.StartElement{Name: xml.Name{Local: "numberExpectedPollData"}}); err != nil {
		return err
	}
	return nil
}

func (m LPollData) String() string {
	return string(m.Box("", 120))
}

func (m LPollData) Box(name string, width int) utils.AsciiBox {
	boxName := "LPollData"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.SourceAddress.Box("sourceAddress", width-2))
		// Array Field (targetAddress)
		if m.TargetAddress != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.TargetAddress {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("TargetAddress", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint8(0x00), -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("NumberExpectedPollData", m.NumberExpectedPollData, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
