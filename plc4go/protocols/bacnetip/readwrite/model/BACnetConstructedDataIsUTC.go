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

// BACnetConstructedDataIsUTC is the data-structure of this message
type BACnetConstructedDataIsUTC struct {
	*BACnetConstructedData
	IsUtc *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIsUTC is the corresponding interface of BACnetConstructedDataIsUTC
type IBACnetConstructedDataIsUTC interface {
	IBACnetConstructedData
	// GetIsUtc returns IsUtc (property field)
	GetIsUtc() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataIsUTC) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIsUTC) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IS_UTC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIsUTC) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIsUTC) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIsUTC) GetIsUtc() *BACnetApplicationTagBoolean {
	return m.IsUtc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIsUTC factory function for BACnetConstructedDataIsUTC
func NewBACnetConstructedDataIsUTC(isUtc *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIsUTC {
	_result := &BACnetConstructedDataIsUTC{
		IsUtc:                 isUtc,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIsUTC(structType interface{}) *BACnetConstructedDataIsUTC {
	if casted, ok := structType.(BACnetConstructedDataIsUTC); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIsUTC); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIsUTC(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIsUTC(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIsUTC) GetTypeName() string {
	return "BACnetConstructedDataIsUTC"
}

func (m *BACnetConstructedDataIsUTC) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIsUTC) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isUtc)
	lengthInBits += m.IsUtc.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIsUTC) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIsUTCParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIsUTC, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIsUTC"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isUtc)
	if pullErr := readBuffer.PullContext("isUtc"); pullErr != nil {
		return nil, pullErr
	}
	_isUtc, _isUtcErr := BACnetApplicationTagParse(readBuffer)
	if _isUtcErr != nil {
		return nil, errors.Wrap(_isUtcErr, "Error parsing 'isUtc' field")
	}
	isUtc := CastBACnetApplicationTagBoolean(_isUtc)
	if closeErr := readBuffer.CloseContext("isUtc"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIsUTC"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIsUTC{
		IsUtc:                 CastBACnetApplicationTagBoolean(isUtc),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIsUTC) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIsUTC"); pushErr != nil {
			return pushErr
		}

		// Simple Field (isUtc)
		if pushErr := writeBuffer.PushContext("isUtc"); pushErr != nil {
			return pushErr
		}
		_isUtcErr := m.IsUtc.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("isUtc"); popErr != nil {
			return popErr
		}
		if _isUtcErr != nil {
			return errors.Wrap(_isUtcErr, "Error serializing 'isUtc' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIsUTC"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIsUTC) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
