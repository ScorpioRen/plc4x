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

// BACnetConstructedDataInstanceOf is the data-structure of this message
type BACnetConstructedDataInstanceOf struct {
	*BACnetConstructedData
	InstanceOf *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInstanceOf is the corresponding interface of BACnetConstructedDataInstanceOf
type IBACnetConstructedDataInstanceOf interface {
	IBACnetConstructedData
	// GetInstanceOf returns InstanceOf (property field)
	GetInstanceOf() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataInstanceOf) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInstanceOf) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANCE_OF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInstanceOf) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInstanceOf) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInstanceOf) GetInstanceOf() *BACnetApplicationTagCharacterString {
	return m.InstanceOf
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstanceOf factory function for BACnetConstructedDataInstanceOf
func NewBACnetConstructedDataInstanceOf(instanceOf *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInstanceOf {
	_result := &BACnetConstructedDataInstanceOf{
		InstanceOf:            instanceOf,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInstanceOf(structType interface{}) *BACnetConstructedDataInstanceOf {
	if casted, ok := structType.(BACnetConstructedDataInstanceOf); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstanceOf); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstanceOf(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstanceOf(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInstanceOf) GetTypeName() string {
	return "BACnetConstructedDataInstanceOf"
}

func (m *BACnetConstructedDataInstanceOf) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInstanceOf) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (instanceOf)
	lengthInBits += m.InstanceOf.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataInstanceOf) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstanceOfParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInstanceOf, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstanceOf"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (instanceOf)
	if pullErr := readBuffer.PullContext("instanceOf"); pullErr != nil {
		return nil, pullErr
	}
	_instanceOf, _instanceOfErr := BACnetApplicationTagParse(readBuffer)
	if _instanceOfErr != nil {
		return nil, errors.Wrap(_instanceOfErr, "Error parsing 'instanceOf' field")
	}
	instanceOf := CastBACnetApplicationTagCharacterString(_instanceOf)
	if closeErr := readBuffer.CloseContext("instanceOf"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstanceOf"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInstanceOf{
		InstanceOf:            CastBACnetApplicationTagCharacterString(instanceOf),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInstanceOf) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstanceOf"); pushErr != nil {
			return pushErr
		}

		// Simple Field (instanceOf)
		if pushErr := writeBuffer.PushContext("instanceOf"); pushErr != nil {
			return pushErr
		}
		_instanceOfErr := m.InstanceOf.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("instanceOf"); popErr != nil {
			return popErr
		}
		if _instanceOfErr != nil {
			return errors.Wrap(_instanceOfErr, "Error serializing 'instanceOf' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstanceOf"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInstanceOf) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
