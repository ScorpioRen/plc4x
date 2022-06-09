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

// BACnetConstructedDataLinkSpeed is the data-structure of this message
type BACnetConstructedDataLinkSpeed struct {
	*BACnetConstructedData
	LinkSpeed *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLinkSpeed is the corresponding interface of BACnetConstructedDataLinkSpeed
type IBACnetConstructedDataLinkSpeed interface {
	IBACnetConstructedData
	// GetLinkSpeed returns LinkSpeed (property field)
	GetLinkSpeed() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataLinkSpeed) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLinkSpeed) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLinkSpeed) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLinkSpeed) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLinkSpeed) GetLinkSpeed() *BACnetApplicationTagReal {
	return m.LinkSpeed
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLinkSpeed factory function for BACnetConstructedDataLinkSpeed
func NewBACnetConstructedDataLinkSpeed(linkSpeed *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLinkSpeed {
	_result := &BACnetConstructedDataLinkSpeed{
		LinkSpeed:             linkSpeed,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLinkSpeed(structType interface{}) *BACnetConstructedDataLinkSpeed {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeed); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLinkSpeed(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLinkSpeed(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLinkSpeed) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeed"
}

func (m *BACnetConstructedDataLinkSpeed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLinkSpeed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (linkSpeed)
	lengthInBits += m.LinkSpeed.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLinkSpeed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLinkSpeedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLinkSpeed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (linkSpeed)
	if pullErr := readBuffer.PullContext("linkSpeed"); pullErr != nil {
		return nil, pullErr
	}
	_linkSpeed, _linkSpeedErr := BACnetApplicationTagParse(readBuffer)
	if _linkSpeedErr != nil {
		return nil, errors.Wrap(_linkSpeedErr, "Error parsing 'linkSpeed' field")
	}
	linkSpeed := CastBACnetApplicationTagReal(_linkSpeed)
	if closeErr := readBuffer.CloseContext("linkSpeed"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeed"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLinkSpeed{
		LinkSpeed:             CastBACnetApplicationTagReal(linkSpeed),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLinkSpeed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeed"); pushErr != nil {
			return pushErr
		}

		// Simple Field (linkSpeed)
		if pushErr := writeBuffer.PushContext("linkSpeed"); pushErr != nil {
			return pushErr
		}
		_linkSpeedErr := m.LinkSpeed.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("linkSpeed"); popErr != nil {
			return popErr
		}
		if _linkSpeedErr != nil {
			return errors.Wrap(_linkSpeedErr, "Error serializing 'linkSpeed' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeed"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLinkSpeed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
