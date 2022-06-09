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

// BACnetConstructedDataCarDriveStatus is the data-structure of this message
type BACnetConstructedDataCarDriveStatus struct {
	*BACnetConstructedData
	CarDriveStatus *BACnetLiftCarDriveStatusTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCarDriveStatus is the corresponding interface of BACnetConstructedDataCarDriveStatus
type IBACnetConstructedDataCarDriveStatus interface {
	IBACnetConstructedData
	// GetCarDriveStatus returns CarDriveStatus (property field)
	GetCarDriveStatus() *BACnetLiftCarDriveStatusTagged
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

func (m *BACnetConstructedDataCarDriveStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCarDriveStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_DRIVE_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCarDriveStatus) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCarDriveStatus) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCarDriveStatus) GetCarDriveStatus() *BACnetLiftCarDriveStatusTagged {
	return m.CarDriveStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarDriveStatus factory function for BACnetConstructedDataCarDriveStatus
func NewBACnetConstructedDataCarDriveStatus(carDriveStatus *BACnetLiftCarDriveStatusTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCarDriveStatus {
	_result := &BACnetConstructedDataCarDriveStatus{
		CarDriveStatus:        carDriveStatus,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCarDriveStatus(structType interface{}) *BACnetConstructedDataCarDriveStatus {
	if casted, ok := structType.(BACnetConstructedDataCarDriveStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDriveStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarDriveStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarDriveStatus(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCarDriveStatus) GetTypeName() string {
	return "BACnetConstructedDataCarDriveStatus"
}

func (m *BACnetConstructedDataCarDriveStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCarDriveStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (carDriveStatus)
	lengthInBits += m.CarDriveStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCarDriveStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarDriveStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCarDriveStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDriveStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carDriveStatus)
	if pullErr := readBuffer.PullContext("carDriveStatus"); pullErr != nil {
		return nil, pullErr
	}
	_carDriveStatus, _carDriveStatusErr := BACnetLiftCarDriveStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _carDriveStatusErr != nil {
		return nil, errors.Wrap(_carDriveStatusErr, "Error parsing 'carDriveStatus' field")
	}
	carDriveStatus := CastBACnetLiftCarDriveStatusTagged(_carDriveStatus)
	if closeErr := readBuffer.CloseContext("carDriveStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDriveStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCarDriveStatus{
		CarDriveStatus:        CastBACnetLiftCarDriveStatusTagged(carDriveStatus),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCarDriveStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDriveStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (carDriveStatus)
		if pushErr := writeBuffer.PushContext("carDriveStatus"); pushErr != nil {
			return pushErr
		}
		_carDriveStatusErr := m.CarDriveStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("carDriveStatus"); popErr != nil {
			return popErr
		}
		if _carDriveStatusErr != nil {
			return errors.Wrap(_carDriveStatusErr, "Error serializing 'carDriveStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDriveStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCarDriveStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
