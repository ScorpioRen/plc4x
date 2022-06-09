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

// BACnetConstructedDataDerivativeConstant is the data-structure of this message
type BACnetConstructedDataDerivativeConstant struct {
	*BACnetConstructedData
	DerivativeConstant *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDerivativeConstant is the corresponding interface of BACnetConstructedDataDerivativeConstant
type IBACnetConstructedDataDerivativeConstant interface {
	IBACnetConstructedData
	// GetDerivativeConstant returns DerivativeConstant (property field)
	GetDerivativeConstant() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataDerivativeConstant) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDerivativeConstant) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DERIVATIVE_CONSTANT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDerivativeConstant) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDerivativeConstant) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDerivativeConstant) GetDerivativeConstant() *BACnetApplicationTagReal {
	return m.DerivativeConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDerivativeConstant factory function for BACnetConstructedDataDerivativeConstant
func NewBACnetConstructedDataDerivativeConstant(derivativeConstant *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDerivativeConstant {
	_result := &BACnetConstructedDataDerivativeConstant{
		DerivativeConstant:    derivativeConstant,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDerivativeConstant(structType interface{}) *BACnetConstructedDataDerivativeConstant {
	if casted, ok := structType.(BACnetConstructedDataDerivativeConstant); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDerivativeConstant); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDerivativeConstant(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDerivativeConstant(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDerivativeConstant) GetTypeName() string {
	return "BACnetConstructedDataDerivativeConstant"
}

func (m *BACnetConstructedDataDerivativeConstant) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDerivativeConstant) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (derivativeConstant)
	lengthInBits += m.DerivativeConstant.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDerivativeConstant) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDerivativeConstantParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDerivativeConstant, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDerivativeConstant"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (derivativeConstant)
	if pullErr := readBuffer.PullContext("derivativeConstant"); pullErr != nil {
		return nil, pullErr
	}
	_derivativeConstant, _derivativeConstantErr := BACnetApplicationTagParse(readBuffer)
	if _derivativeConstantErr != nil {
		return nil, errors.Wrap(_derivativeConstantErr, "Error parsing 'derivativeConstant' field")
	}
	derivativeConstant := CastBACnetApplicationTagReal(_derivativeConstant)
	if closeErr := readBuffer.CloseContext("derivativeConstant"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDerivativeConstant"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDerivativeConstant{
		DerivativeConstant:    CastBACnetApplicationTagReal(derivativeConstant),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDerivativeConstant) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDerivativeConstant"); pushErr != nil {
			return pushErr
		}

		// Simple Field (derivativeConstant)
		if pushErr := writeBuffer.PushContext("derivativeConstant"); pushErr != nil {
			return pushErr
		}
		_derivativeConstantErr := m.DerivativeConstant.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("derivativeConstant"); popErr != nil {
			return popErr
		}
		if _derivativeConstantErr != nil {
			return errors.Wrap(_derivativeConstantErr, "Error serializing 'derivativeConstant' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDerivativeConstant"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDerivativeConstant) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
