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

// BACnetConstructedDataTimeOfDeviceRestart is the data-structure of this message
type BACnetConstructedDataTimeOfDeviceRestart struct {
	*BACnetConstructedData
	TimeOfDeviceRestart *BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTimeOfDeviceRestart is the corresponding interface of BACnetConstructedDataTimeOfDeviceRestart
type IBACnetConstructedDataTimeOfDeviceRestart interface {
	IBACnetConstructedData
	// GetTimeOfDeviceRestart returns TimeOfDeviceRestart (property field)
	GetTimeOfDeviceRestart() *BACnetTimeStamp
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

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_DEVICE_RESTART
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimeOfDeviceRestart) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetTimeOfDeviceRestart() *BACnetTimeStamp {
	return m.TimeOfDeviceRestart
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfDeviceRestart factory function for BACnetConstructedDataTimeOfDeviceRestart
func NewBACnetConstructedDataTimeOfDeviceRestart(timeOfDeviceRestart *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTimeOfDeviceRestart {
	_result := &BACnetConstructedDataTimeOfDeviceRestart{
		TimeOfDeviceRestart:   timeOfDeviceRestart,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimeOfDeviceRestart(structType interface{}) *BACnetConstructedDataTimeOfDeviceRestart {
	if casted, ok := structType.(BACnetConstructedDataTimeOfDeviceRestart); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfDeviceRestart); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeOfDeviceRestart(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeOfDeviceRestart(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetTypeName() string {
	return "BACnetConstructedDataTimeOfDeviceRestart"
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeOfDeviceRestart)
	lengthInBits += m.TimeOfDeviceRestart.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeOfDeviceRestartParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTimeOfDeviceRestart, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfDeviceRestart"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfDeviceRestart)
	if pullErr := readBuffer.PullContext("timeOfDeviceRestart"); pullErr != nil {
		return nil, pullErr
	}
	_timeOfDeviceRestart, _timeOfDeviceRestartErr := BACnetTimeStampParse(readBuffer)
	if _timeOfDeviceRestartErr != nil {
		return nil, errors.Wrap(_timeOfDeviceRestartErr, "Error parsing 'timeOfDeviceRestart' field")
	}
	timeOfDeviceRestart := CastBACnetTimeStamp(_timeOfDeviceRestart)
	if closeErr := readBuffer.CloseContext("timeOfDeviceRestart"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfDeviceRestart"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimeOfDeviceRestart{
		TimeOfDeviceRestart:   CastBACnetTimeStamp(timeOfDeviceRestart),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfDeviceRestart"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeOfDeviceRestart)
		if pushErr := writeBuffer.PushContext("timeOfDeviceRestart"); pushErr != nil {
			return pushErr
		}
		_timeOfDeviceRestartErr := m.TimeOfDeviceRestart.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeOfDeviceRestart"); popErr != nil {
			return popErr
		}
		if _timeOfDeviceRestartErr != nil {
			return errors.Wrap(_timeOfDeviceRestartErr, "Error serializing 'timeOfDeviceRestart' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfDeviceRestart"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimeOfDeviceRestart) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
