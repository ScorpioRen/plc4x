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

// BACnetConstructedDataStateChangeValues is the data-structure of this message
type BACnetConstructedDataStateChangeValues struct {
	*BACnetConstructedData
	StateChangeValues []*BACnetTimerStateChangeValue

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataStateChangeValues is the corresponding interface of BACnetConstructedDataStateChangeValues
type IBACnetConstructedDataStateChangeValues interface {
	IBACnetConstructedData
	// GetStateChangeValues returns StateChangeValues (property field)
	GetStateChangeValues() []*BACnetTimerStateChangeValue
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

func (m *BACnetConstructedDataStateChangeValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataStateChangeValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STATE_CHANGE_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataStateChangeValues) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataStateChangeValues) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataStateChangeValues) GetStateChangeValues() []*BACnetTimerStateChangeValue {
	return m.StateChangeValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStateChangeValues factory function for BACnetConstructedDataStateChangeValues
func NewBACnetConstructedDataStateChangeValues(stateChangeValues []*BACnetTimerStateChangeValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataStateChangeValues {
	_result := &BACnetConstructedDataStateChangeValues{
		StateChangeValues:     stateChangeValues,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataStateChangeValues(structType interface{}) *BACnetConstructedDataStateChangeValues {
	if casted, ok := structType.(BACnetConstructedDataStateChangeValues); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStateChangeValues); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataStateChangeValues(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataStateChangeValues(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataStateChangeValues) GetTypeName() string {
	return "BACnetConstructedDataStateChangeValues"
}

func (m *BACnetConstructedDataStateChangeValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataStateChangeValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.StateChangeValues) > 0 {
		for _, element := range m.StateChangeValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataStateChangeValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStateChangeValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataStateChangeValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStateChangeValues"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (stateChangeValues)
	if pullErr := readBuffer.PullContext("stateChangeValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	stateChangeValues := make([]*BACnetTimerStateChangeValue, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimerStateChangeValueParse(readBuffer, objectTypeArgument)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'stateChangeValues' field")
			}
			stateChangeValues = append(stateChangeValues, CastBACnetTimerStateChangeValue(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("stateChangeValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStateChangeValues"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataStateChangeValues{
		StateChangeValues:     stateChangeValues,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataStateChangeValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStateChangeValues"); pushErr != nil {
			return pushErr
		}

		// Array Field (stateChangeValues)
		if m.StateChangeValues != nil {
			if pushErr := writeBuffer.PushContext("stateChangeValues", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.StateChangeValues {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'stateChangeValues' field")
				}
			}
			if popErr := writeBuffer.PopContext("stateChangeValues", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStateChangeValues"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataStateChangeValues) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
