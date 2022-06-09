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

// BACnetConstructedDataBackupPreparationTime is the data-structure of this message
type BACnetConstructedDataBackupPreparationTime struct {
	*BACnetConstructedData
	BackupPreparationTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBackupPreparationTime is the corresponding interface of BACnetConstructedDataBackupPreparationTime
type IBACnetConstructedDataBackupPreparationTime interface {
	IBACnetConstructedData
	// GetBackupPreparationTime returns BackupPreparationTime (property field)
	GetBackupPreparationTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataBackupPreparationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBackupPreparationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_PREPARATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBackupPreparationTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBackupPreparationTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBackupPreparationTime) GetBackupPreparationTime() *BACnetApplicationTagUnsignedInteger {
	return m.BackupPreparationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBackupPreparationTime factory function for BACnetConstructedDataBackupPreparationTime
func NewBACnetConstructedDataBackupPreparationTime(backupPreparationTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBackupPreparationTime {
	_result := &BACnetConstructedDataBackupPreparationTime{
		BackupPreparationTime: backupPreparationTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBackupPreparationTime(structType interface{}) *BACnetConstructedDataBackupPreparationTime {
	if casted, ok := structType.(BACnetConstructedDataBackupPreparationTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupPreparationTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBackupPreparationTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBackupPreparationTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBackupPreparationTime) GetTypeName() string {
	return "BACnetConstructedDataBackupPreparationTime"
}

func (m *BACnetConstructedDataBackupPreparationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBackupPreparationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (backupPreparationTime)
	lengthInBits += m.BackupPreparationTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBackupPreparationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBackupPreparationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBackupPreparationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupPreparationTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (backupPreparationTime)
	if pullErr := readBuffer.PullContext("backupPreparationTime"); pullErr != nil {
		return nil, pullErr
	}
	_backupPreparationTime, _backupPreparationTimeErr := BACnetApplicationTagParse(readBuffer)
	if _backupPreparationTimeErr != nil {
		return nil, errors.Wrap(_backupPreparationTimeErr, "Error parsing 'backupPreparationTime' field")
	}
	backupPreparationTime := CastBACnetApplicationTagUnsignedInteger(_backupPreparationTime)
	if closeErr := readBuffer.CloseContext("backupPreparationTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupPreparationTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBackupPreparationTime{
		BackupPreparationTime: CastBACnetApplicationTagUnsignedInteger(backupPreparationTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBackupPreparationTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupPreparationTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (backupPreparationTime)
		if pushErr := writeBuffer.PushContext("backupPreparationTime"); pushErr != nil {
			return pushErr
		}
		_backupPreparationTimeErr := m.BackupPreparationTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("backupPreparationTime"); popErr != nil {
			return popErr
		}
		if _backupPreparationTimeErr != nil {
			return errors.Wrap(_backupPreparationTimeErr, "Error serializing 'backupPreparationTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupPreparationTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBackupPreparationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
