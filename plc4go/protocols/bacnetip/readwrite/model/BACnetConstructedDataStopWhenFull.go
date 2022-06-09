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

// BACnetConstructedDataStopWhenFull is the data-structure of this message
type BACnetConstructedDataStopWhenFull struct {
	*BACnetConstructedData
	StopWhenFull *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataStopWhenFull is the corresponding interface of BACnetConstructedDataStopWhenFull
type IBACnetConstructedDataStopWhenFull interface {
	IBACnetConstructedData
	// GetStopWhenFull returns StopWhenFull (property field)
	GetStopWhenFull() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataStopWhenFull) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataStopWhenFull) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_WHEN_FULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataStopWhenFull) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataStopWhenFull) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataStopWhenFull) GetStopWhenFull() *BACnetApplicationTagBoolean {
	return m.StopWhenFull
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStopWhenFull factory function for BACnetConstructedDataStopWhenFull
func NewBACnetConstructedDataStopWhenFull(stopWhenFull *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataStopWhenFull {
	_result := &BACnetConstructedDataStopWhenFull{
		StopWhenFull:          stopWhenFull,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataStopWhenFull(structType interface{}) *BACnetConstructedDataStopWhenFull {
	if casted, ok := structType.(BACnetConstructedDataStopWhenFull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopWhenFull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataStopWhenFull(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataStopWhenFull(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataStopWhenFull) GetTypeName() string {
	return "BACnetConstructedDataStopWhenFull"
}

func (m *BACnetConstructedDataStopWhenFull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataStopWhenFull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (stopWhenFull)
	lengthInBits += m.StopWhenFull.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataStopWhenFull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStopWhenFullParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataStopWhenFull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopWhenFull"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stopWhenFull)
	if pullErr := readBuffer.PullContext("stopWhenFull"); pullErr != nil {
		return nil, pullErr
	}
	_stopWhenFull, _stopWhenFullErr := BACnetApplicationTagParse(readBuffer)
	if _stopWhenFullErr != nil {
		return nil, errors.Wrap(_stopWhenFullErr, "Error parsing 'stopWhenFull' field")
	}
	stopWhenFull := CastBACnetApplicationTagBoolean(_stopWhenFull)
	if closeErr := readBuffer.CloseContext("stopWhenFull"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopWhenFull"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataStopWhenFull{
		StopWhenFull:          CastBACnetApplicationTagBoolean(stopWhenFull),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataStopWhenFull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopWhenFull"); pushErr != nil {
			return pushErr
		}

		// Simple Field (stopWhenFull)
		if pushErr := writeBuffer.PushContext("stopWhenFull"); pushErr != nil {
			return pushErr
		}
		_stopWhenFullErr := m.StopWhenFull.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("stopWhenFull"); popErr != nil {
			return popErr
		}
		if _stopWhenFullErr != nil {
			return errors.Wrap(_stopWhenFullErr, "Error serializing 'stopWhenFull' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopWhenFull"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataStopWhenFull) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
