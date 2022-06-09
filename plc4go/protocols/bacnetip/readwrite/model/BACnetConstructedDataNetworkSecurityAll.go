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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataNetworkSecurityAll is the data-structure of this message
type BACnetConstructedDataNetworkSecurityAll struct {
	*BACnetConstructedData

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataNetworkSecurityAll is the corresponding interface of BACnetConstructedDataNetworkSecurityAll
type IBACnetConstructedDataNetworkSecurityAll interface {
	IBACnetConstructedData
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

func (m *BACnetConstructedDataNetworkSecurityAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_NETWORK_SECURITY
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNetworkSecurityAll) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

// NewBACnetConstructedDataNetworkSecurityAll factory function for BACnetConstructedDataNetworkSecurityAll
func NewBACnetConstructedDataNetworkSecurityAll(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataNetworkSecurityAll {
	_result := &BACnetConstructedDataNetworkSecurityAll{
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNetworkSecurityAll(structType interface{}) *BACnetConstructedDataNetworkSecurityAll {
	if casted, ok := structType.(BACnetConstructedDataNetworkSecurityAll); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkSecurityAll); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkSecurityAll(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkSecurityAll(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetTypeName() string {
	return "BACnetConstructedDataNetworkSecurityAll"
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConstructedDataNetworkSecurityAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkSecurityAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataNetworkSecurityAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkSecurityAll"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkSecurityAll"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNetworkSecurityAll{
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNetworkSecurityAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkSecurityAll"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkSecurityAll"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNetworkSecurityAll) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
