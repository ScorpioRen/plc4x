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

// BACnetConstructedDataMinPresValue is the data-structure of this message
type BACnetConstructedDataMinPresValue struct {
	*BACnetConstructedData
	MinPresValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinPresValue is the corresponding interface of BACnetConstructedDataMinPresValue
type IBACnetConstructedDataMinPresValue interface {
	IBACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinPresValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinPresValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinPresValue) GetMinPresValue() *BACnetApplicationTagReal {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinPresValue factory function for BACnetConstructedDataMinPresValue
func NewBACnetConstructedDataMinPresValue(minPresValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinPresValue {
	_result := &BACnetConstructedDataMinPresValue{
		MinPresValue:          minPresValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinPresValue(structType interface{}) *BACnetConstructedDataMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataMinPresValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinPresValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinPresValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataMinPresValue"
}

func (m *BACnetConstructedDataMinPresValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinPresValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinPresValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinPresValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinPresValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minPresValue)
	if pullErr := readBuffer.PullContext("minPresValue"); pullErr != nil {
		return nil, pullErr
	}
	_minPresValue, _minPresValueErr := BACnetApplicationTagParse(readBuffer)
	if _minPresValueErr != nil {
		return nil, errors.Wrap(_minPresValueErr, "Error parsing 'minPresValue' field")
	}
	minPresValue := CastBACnetApplicationTagReal(_minPresValue)
	if closeErr := readBuffer.CloseContext("minPresValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinPresValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinPresValue{
		MinPresValue:          CastBACnetApplicationTagReal(minPresValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinPresValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinPresValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (minPresValue)
		if pushErr := writeBuffer.PushContext("minPresValue"); pushErr != nil {
			return pushErr
		}
		_minPresValueErr := m.MinPresValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("minPresValue"); popErr != nil {
			return popErr
		}
		if _minPresValueErr != nil {
			return errors.Wrap(_minPresValueErr, "Error serializing 'minPresValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinPresValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
