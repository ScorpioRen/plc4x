/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCarDoorCommand is the data-structure of this message
type BACnetConstructedDataCarDoorCommand struct {
	*BACnetConstructedData
	CarDoorCommand []*BACnetLiftCarDoorCommandTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCarDoorCommand is the corresponding interface of BACnetConstructedDataCarDoorCommand
type IBACnetConstructedDataCarDoorCommand interface {
	IBACnetConstructedData
	// GetCarDoorCommand returns CarDoorCommand (property field)
	GetCarDoorCommand() []*BACnetLiftCarDoorCommandTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataCarDoorCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCarDoorCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_DOOR_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCarDoorCommand) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCarDoorCommand) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCarDoorCommand) GetCarDoorCommand() []*BACnetLiftCarDoorCommandTagged {
	return m.CarDoorCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarDoorCommand factory function for BACnetConstructedDataCarDoorCommand
func NewBACnetConstructedDataCarDoorCommand(carDoorCommand []*BACnetLiftCarDoorCommandTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCarDoorCommand {
	_result := &BACnetConstructedDataCarDoorCommand{
		CarDoorCommand:        carDoorCommand,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCarDoorCommand(structType interface{}) *BACnetConstructedDataCarDoorCommand {
	if casted, ok := structType.(BACnetConstructedDataCarDoorCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDoorCommand); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarDoorCommand(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarDoorCommand(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCarDoorCommand) GetTypeName() string {
	return "BACnetConstructedDataCarDoorCommand"
}

func (m *BACnetConstructedDataCarDoorCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCarDoorCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.CarDoorCommand) > 0 {
		for _, element := range m.CarDoorCommand {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataCarDoorCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarDoorCommandParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCarDoorCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDoorCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (carDoorCommand)
	if pullErr := readBuffer.PullContext("carDoorCommand", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	carDoorCommand := make([]*BACnetLiftCarDoorCommandTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLiftCarDoorCommandTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'carDoorCommand' field")
			}
			carDoorCommand = append(carDoorCommand, CastBACnetLiftCarDoorCommandTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("carDoorCommand", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDoorCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCarDoorCommand{
		CarDoorCommand:        carDoorCommand,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCarDoorCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDoorCommand"); pushErr != nil {
			return pushErr
		}

		// Array Field (carDoorCommand)
		if m.CarDoorCommand != nil {
			if pushErr := writeBuffer.PushContext("carDoorCommand", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.CarDoorCommand {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'carDoorCommand' field")
				}
			}
			if popErr := writeBuffer.PopContext("carDoorCommand", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDoorCommand"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCarDoorCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
