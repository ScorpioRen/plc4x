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

// BACnetConstructedDataThreatAuthority is the data-structure of this message
type BACnetConstructedDataThreatAuthority struct {
	*BACnetConstructedData
	ThreatAuthority *BACnetAccessThreatLevel

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataThreatAuthority is the corresponding interface of BACnetConstructedDataThreatAuthority
type IBACnetConstructedDataThreatAuthority interface {
	IBACnetConstructedData
	// GetThreatAuthority returns ThreatAuthority (property field)
	GetThreatAuthority() *BACnetAccessThreatLevel
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

func (m *BACnetConstructedDataThreatAuthority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataThreatAuthority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_THREAT_AUTHORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataThreatAuthority) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataThreatAuthority) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataThreatAuthority) GetThreatAuthority() *BACnetAccessThreatLevel {
	return m.ThreatAuthority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataThreatAuthority factory function for BACnetConstructedDataThreatAuthority
func NewBACnetConstructedDataThreatAuthority(threatAuthority *BACnetAccessThreatLevel, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataThreatAuthority {
	_result := &BACnetConstructedDataThreatAuthority{
		ThreatAuthority:       threatAuthority,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataThreatAuthority(structType interface{}) *BACnetConstructedDataThreatAuthority {
	if casted, ok := structType.(BACnetConstructedDataThreatAuthority); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataThreatAuthority); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataThreatAuthority(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataThreatAuthority(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataThreatAuthority) GetTypeName() string {
	return "BACnetConstructedDataThreatAuthority"
}

func (m *BACnetConstructedDataThreatAuthority) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataThreatAuthority) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (threatAuthority)
	lengthInBits += m.ThreatAuthority.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataThreatAuthority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataThreatAuthorityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataThreatAuthority, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataThreatAuthority"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (threatAuthority)
	if pullErr := readBuffer.PullContext("threatAuthority"); pullErr != nil {
		return nil, pullErr
	}
	_threatAuthority, _threatAuthorityErr := BACnetAccessThreatLevelParse(readBuffer)
	if _threatAuthorityErr != nil {
		return nil, errors.Wrap(_threatAuthorityErr, "Error parsing 'threatAuthority' field")
	}
	threatAuthority := CastBACnetAccessThreatLevel(_threatAuthority)
	if closeErr := readBuffer.CloseContext("threatAuthority"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataThreatAuthority"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataThreatAuthority{
		ThreatAuthority:       CastBACnetAccessThreatLevel(threatAuthority),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataThreatAuthority) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataThreatAuthority"); pushErr != nil {
			return pushErr
		}

		// Simple Field (threatAuthority)
		if pushErr := writeBuffer.PushContext("threatAuthority"); pushErr != nil {
			return pushErr
		}
		_threatAuthorityErr := m.ThreatAuthority.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("threatAuthority"); popErr != nil {
			return popErr
		}
		if _threatAuthorityErr != nil {
			return errors.Wrap(_threatAuthorityErr, "Error serializing 'threatAuthority' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataThreatAuthority"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataThreatAuthority) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
