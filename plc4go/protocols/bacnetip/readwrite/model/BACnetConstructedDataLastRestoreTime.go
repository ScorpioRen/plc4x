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

// BACnetConstructedDataLastRestoreTime is the data-structure of this message
type BACnetConstructedDataLastRestoreTime struct {
	*BACnetConstructedData
	LastRestoreTime *BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLastRestoreTime is the corresponding interface of BACnetConstructedDataLastRestoreTime
type IBACnetConstructedDataLastRestoreTime interface {
	IBACnetConstructedData
	// GetLastRestoreTime returns LastRestoreTime (property field)
	GetLastRestoreTime() *BACnetTimeStamp
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

func (m *BACnetConstructedDataLastRestoreTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLastRestoreTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_RESTORE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLastRestoreTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLastRestoreTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLastRestoreTime) GetLastRestoreTime() *BACnetTimeStamp {
	return m.LastRestoreTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastRestoreTime factory function for BACnetConstructedDataLastRestoreTime
func NewBACnetConstructedDataLastRestoreTime(lastRestoreTime *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLastRestoreTime {
	_result := &BACnetConstructedDataLastRestoreTime{
		LastRestoreTime:       lastRestoreTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLastRestoreTime(structType interface{}) *BACnetConstructedDataLastRestoreTime {
	if casted, ok := structType.(BACnetConstructedDataLastRestoreTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestoreTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastRestoreTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastRestoreTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLastRestoreTime) GetTypeName() string {
	return "BACnetConstructedDataLastRestoreTime"
}

func (m *BACnetConstructedDataLastRestoreTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLastRestoreTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastRestoreTime)
	lengthInBits += m.LastRestoreTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLastRestoreTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastRestoreTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLastRestoreTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestoreTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastRestoreTime)
	if pullErr := readBuffer.PullContext("lastRestoreTime"); pullErr != nil {
		return nil, pullErr
	}
	_lastRestoreTime, _lastRestoreTimeErr := BACnetTimeStampParse(readBuffer)
	if _lastRestoreTimeErr != nil {
		return nil, errors.Wrap(_lastRestoreTimeErr, "Error parsing 'lastRestoreTime' field")
	}
	lastRestoreTime := CastBACnetTimeStamp(_lastRestoreTime)
	if closeErr := readBuffer.CloseContext("lastRestoreTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestoreTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLastRestoreTime{
		LastRestoreTime:       CastBACnetTimeStamp(lastRestoreTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLastRestoreTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestoreTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lastRestoreTime)
		if pushErr := writeBuffer.PushContext("lastRestoreTime"); pushErr != nil {
			return pushErr
		}
		_lastRestoreTimeErr := m.LastRestoreTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lastRestoreTime"); popErr != nil {
			return popErr
		}
		if _lastRestoreTimeErr != nil {
			return errors.Wrap(_lastRestoreTimeErr, "Error serializing 'lastRestoreTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestoreTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLastRestoreTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
