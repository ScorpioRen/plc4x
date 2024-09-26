/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadFileRecordRequestItem is the corresponding interface of ModbusPDUReadFileRecordRequestItem
type ModbusPDUReadFileRecordRequestItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordLength returns RecordLength (property field)
	GetRecordLength() uint16
	// IsModbusPDUReadFileRecordRequestItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadFileRecordRequestItem()
	// CreateBuilder creates a ModbusPDUReadFileRecordRequestItemBuilder
	CreateModbusPDUReadFileRecordRequestItemBuilder() ModbusPDUReadFileRecordRequestItemBuilder
}

// _ModbusPDUReadFileRecordRequestItem is the data-structure of this message
type _ModbusPDUReadFileRecordRequestItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordLength  uint16
}

var _ ModbusPDUReadFileRecordRequestItem = (*_ModbusPDUReadFileRecordRequestItem)(nil)

// NewModbusPDUReadFileRecordRequestItem factory function for _ModbusPDUReadFileRecordRequestItem
func NewModbusPDUReadFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) *_ModbusPDUReadFileRecordRequestItem {
	return &_ModbusPDUReadFileRecordRequestItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordLength: recordLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadFileRecordRequestItemBuilder is a builder for ModbusPDUReadFileRecordRequestItem
type ModbusPDUReadFileRecordRequestItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) ModbusPDUReadFileRecordRequestItemBuilder
	// WithReferenceType adds ReferenceType (property field)
	WithReferenceType(uint8) ModbusPDUReadFileRecordRequestItemBuilder
	// WithFileNumber adds FileNumber (property field)
	WithFileNumber(uint16) ModbusPDUReadFileRecordRequestItemBuilder
	// WithRecordNumber adds RecordNumber (property field)
	WithRecordNumber(uint16) ModbusPDUReadFileRecordRequestItemBuilder
	// WithRecordLength adds RecordLength (property field)
	WithRecordLength(uint16) ModbusPDUReadFileRecordRequestItemBuilder
	// Build builds the ModbusPDUReadFileRecordRequestItem or returns an error if something is wrong
	Build() (ModbusPDUReadFileRecordRequestItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadFileRecordRequestItem
}

// NewModbusPDUReadFileRecordRequestItemBuilder() creates a ModbusPDUReadFileRecordRequestItemBuilder
func NewModbusPDUReadFileRecordRequestItemBuilder() ModbusPDUReadFileRecordRequestItemBuilder {
	return &_ModbusPDUReadFileRecordRequestItemBuilder{_ModbusPDUReadFileRecordRequestItem: new(_ModbusPDUReadFileRecordRequestItem)}
}

type _ModbusPDUReadFileRecordRequestItemBuilder struct {
	*_ModbusPDUReadFileRecordRequestItem

	err *utils.MultiError
}

var _ (ModbusPDUReadFileRecordRequestItemBuilder) = (*_ModbusPDUReadFileRecordRequestItemBuilder)(nil)

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) WithMandatoryFields(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) ModbusPDUReadFileRecordRequestItemBuilder {
	return m.WithReferenceType(referenceType).WithFileNumber(fileNumber).WithRecordNumber(recordNumber).WithRecordLength(recordLength)
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) WithReferenceType(referenceType uint8) ModbusPDUReadFileRecordRequestItemBuilder {
	m.ReferenceType = referenceType
	return m
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) WithFileNumber(fileNumber uint16) ModbusPDUReadFileRecordRequestItemBuilder {
	m.FileNumber = fileNumber
	return m
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) WithRecordNumber(recordNumber uint16) ModbusPDUReadFileRecordRequestItemBuilder {
	m.RecordNumber = recordNumber
	return m
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) WithRecordLength(recordLength uint16) ModbusPDUReadFileRecordRequestItemBuilder {
	m.RecordLength = recordLength
	return m
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) Build() (ModbusPDUReadFileRecordRequestItem, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ModbusPDUReadFileRecordRequestItem.deepCopy(), nil
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) MustBuild() ModbusPDUReadFileRecordRequestItem {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ModbusPDUReadFileRecordRequestItemBuilder) DeepCopy() any {
	return m.CreateModbusPDUReadFileRecordRequestItemBuilder()
}

// CreateModbusPDUReadFileRecordRequestItemBuilder creates a ModbusPDUReadFileRecordRequestItemBuilder
func (m *_ModbusPDUReadFileRecordRequestItem) CreateModbusPDUReadFileRecordRequestItemBuilder() ModbusPDUReadFileRecordRequestItemBuilder {
	if m == nil {
		return NewModbusPDUReadFileRecordRequestItemBuilder()
	}
	return &_ModbusPDUReadFileRecordRequestItemBuilder{_ModbusPDUReadFileRecordRequestItem: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordRequestItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetRecordLength() uint16 {
	return m.RecordLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordRequestItem(structType any) ModbusPDUReadFileRecordRequestItem {
	if casted, ok := structType.(ModbusPDUReadFileRecordRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordRequestItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetTypeName() string {
	return "ModbusPDUReadFileRecordRequestItem"
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Simple field (recordLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordRequestItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadFileRecordRequestItemParse(ctx context.Context, theBytes []byte) (ModbusPDUReadFileRecordRequestItem, error) {
	return ModbusPDUReadFileRecordRequestItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUReadFileRecordRequestItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordRequestItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordRequestItem, error) {
		return ModbusPDUReadFileRecordRequestItemParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusPDUReadFileRecordRequestItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordRequestItem, error) {
	v, err := (&_ModbusPDUReadFileRecordRequestItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ModbusPDUReadFileRecordRequestItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__modbusPDUReadFileRecordRequestItem ModbusPDUReadFileRecordRequestItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordRequestItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordRequestItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceType, err := ReadSimpleField(ctx, "referenceType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceType' field"))
	}
	m.ReferenceType = referenceType

	fileNumber, err := ReadSimpleField(ctx, "fileNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileNumber' field"))
	}
	m.FileNumber = fileNumber

	recordNumber, err := ReadSimpleField(ctx, "recordNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordNumber' field"))
	}
	m.RecordNumber = recordNumber

	recordLength, err := ReadSimpleField(ctx, "recordLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordLength' field"))
	}
	m.RecordLength = recordLength

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordRequestItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordRequestItem")
	}

	return m, nil
}

func (m *_ModbusPDUReadFileRecordRequestItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFileRecordRequestItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordRequestItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordRequestItem")
	}

	if err := WriteSimpleField[uint8](ctx, "referenceType", m.GetReferenceType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'referenceType' field")
	}

	if err := WriteSimpleField[uint16](ctx, "fileNumber", m.GetFileNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'fileNumber' field")
	}

	if err := WriteSimpleField[uint16](ctx, "recordNumber", m.GetRecordNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordNumber' field")
	}

	if err := WriteSimpleField[uint16](ctx, "recordLength", m.GetRecordLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordLength' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordRequestItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordRequestItem")
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordRequestItem) IsModbusPDUReadFileRecordRequestItem() {}

func (m *_ModbusPDUReadFileRecordRequestItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadFileRecordRequestItem) deepCopy() *_ModbusPDUReadFileRecordRequestItem {
	if m == nil {
		return nil
	}
	_ModbusPDUReadFileRecordRequestItemCopy := &_ModbusPDUReadFileRecordRequestItem{
		m.ReferenceType,
		m.FileNumber,
		m.RecordNumber,
		m.RecordLength,
	}
	return _ModbusPDUReadFileRecordRequestItemCopy
}

func (m *_ModbusPDUReadFileRecordRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
