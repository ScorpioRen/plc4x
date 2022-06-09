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

// BACnetConstructedDataAccumulatorLowLimit is the data-structure of this message
type BACnetConstructedDataAccumulatorLowLimit struct {
	*BACnetConstructedData
	LowLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAccumulatorLowLimit is the corresponding interface of BACnetConstructedDataAccumulatorLowLimit
type IBACnetConstructedDataAccumulatorLowLimit interface {
	IBACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAccumulatorLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccumulatorLowLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccumulatorLowLimit) GetLowLimit() *BACnetApplicationTagUnsignedInteger {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorLowLimit factory function for BACnetConstructedDataAccumulatorLowLimit
func NewBACnetConstructedDataAccumulatorLowLimit(lowLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAccumulatorLowLimit {
	_result := &BACnetConstructedDataAccumulatorLowLimit{
		LowLimit:              lowLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccumulatorLowLimit(structType interface{}) *BACnetConstructedDataAccumulatorLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorLowLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorLowLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorLowLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorLowLimit"
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAccumulatorLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccumulatorLowLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAccumulatorLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorLowLimit"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, pullErr
	}
	_lowLimit, _lowLimitErr := BACnetApplicationTagParse(readBuffer)
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field")
	}
	lowLimit := CastBACnetApplicationTagUnsignedInteger(_lowLimit)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorLowLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccumulatorLowLimit{
		LowLimit:              CastBACnetApplicationTagUnsignedInteger(lowLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccumulatorLowLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorLowLimit"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return pushErr
		}
		_lowLimitErr := m.LowLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return popErr
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorLowLimit"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccumulatorLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
