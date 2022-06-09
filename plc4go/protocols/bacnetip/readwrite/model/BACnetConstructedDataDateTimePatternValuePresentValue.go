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

// BACnetConstructedDataDateTimePatternValuePresentValue is the data-structure of this message
type BACnetConstructedDataDateTimePatternValuePresentValue struct {
	*BACnetConstructedData
	PresentValue *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDateTimePatternValuePresentValue is the corresponding interface of BACnetConstructedDataDateTimePatternValuePresentValue
type IBACnetConstructedDataDateTimePatternValuePresentValue interface {
	IBACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() *BACnetDateTime
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

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIMEPATTERN_VALUE
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetPresentValue() *BACnetDateTime {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDateTimePatternValuePresentValue factory function for BACnetConstructedDataDateTimePatternValuePresentValue
func NewBACnetConstructedDataDateTimePatternValuePresentValue(presentValue *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDateTimePatternValuePresentValue {
	_result := &BACnetConstructedDataDateTimePatternValuePresentValue{
		PresentValue:          presentValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDateTimePatternValuePresentValue(structType interface{}) *BACnetConstructedDataDateTimePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDateTimePatternValuePresentValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDateTimePatternValuePresentValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataDateTimePatternValuePresentValue"
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDateTimePatternValuePresentValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDateTimePatternValuePresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, pullErr
	}
	_presentValue, _presentValueErr := BACnetDateTimeParse(readBuffer)
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field")
	}
	presentValue := CastBACnetDateTime(_presentValue)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateTimePatternValuePresentValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDateTimePatternValuePresentValue{
		PresentValue:          CastBACnetDateTime(presentValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (presentValue)
		if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
			return pushErr
		}
		_presentValueErr := m.PresentValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
			return popErr
		}
		if _presentValueErr != nil {
			return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateTimePatternValuePresentValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDateTimePatternValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
