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

// BACnetServiceAckAtomicReadFileStream is the corresponding interface of BACnetServiceAckAtomicReadFileStream
type BACnetServiceAckAtomicReadFileStream interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAckAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetFileData returns FileData (property field)
	GetFileData() BACnetApplicationTagOctetString
	// IsBACnetServiceAckAtomicReadFileStream is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicReadFileStream()
	// CreateBuilder creates a BACnetServiceAckAtomicReadFileStreamBuilder
	CreateBACnetServiceAckAtomicReadFileStreamBuilder() BACnetServiceAckAtomicReadFileStreamBuilder
}

// _BACnetServiceAckAtomicReadFileStream is the data-structure of this message
type _BACnetServiceAckAtomicReadFileStream struct {
	BACnetServiceAckAtomicReadFileStreamOrRecordContract
	FileStartPosition BACnetApplicationTagSignedInteger
	FileData          BACnetApplicationTagOctetString
}

var _ BACnetServiceAckAtomicReadFileStream = (*_BACnetServiceAckAtomicReadFileStream)(nil)
var _ BACnetServiceAckAtomicReadFileStreamOrRecordRequirements = (*_BACnetServiceAckAtomicReadFileStream)(nil)

// NewBACnetServiceAckAtomicReadFileStream factory function for _BACnetServiceAckAtomicReadFileStream
func NewBACnetServiceAckAtomicReadFileStream(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag, fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString) *_BACnetServiceAckAtomicReadFileStream {
	if fileStartPosition == nil {
		panic("fileStartPosition of type BACnetApplicationTagSignedInteger for BACnetServiceAckAtomicReadFileStream must not be nil")
	}
	if fileData == nil {
		panic("fileData of type BACnetApplicationTagOctetString for BACnetServiceAckAtomicReadFileStream must not be nil")
	}
	_result := &_BACnetServiceAckAtomicReadFileStream{
		BACnetServiceAckAtomicReadFileStreamOrRecordContract: NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartPosition: fileStartPosition,
		FileData:          fileData,
	}
	_result.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckAtomicReadFileStreamBuilder is a builder for BACnetServiceAckAtomicReadFileStream
type BACnetServiceAckAtomicReadFileStreamBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString) BACnetServiceAckAtomicReadFileStreamBuilder
	// WithFileStartPosition adds FileStartPosition (property field)
	WithFileStartPosition(BACnetApplicationTagSignedInteger) BACnetServiceAckAtomicReadFileStreamBuilder
	// WithFileStartPositionBuilder adds FileStartPosition (property field) which is build by the builder
	WithFileStartPositionBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetServiceAckAtomicReadFileStreamBuilder
	// WithFileData adds FileData (property field)
	WithFileData(BACnetApplicationTagOctetString) BACnetServiceAckAtomicReadFileStreamBuilder
	// WithFileDataBuilder adds FileData (property field) which is build by the builder
	WithFileDataBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetServiceAckAtomicReadFileStreamBuilder
	// Build builds the BACnetServiceAckAtomicReadFileStream or returns an error if something is wrong
	Build() (BACnetServiceAckAtomicReadFileStream, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckAtomicReadFileStream
}

// NewBACnetServiceAckAtomicReadFileStreamBuilder() creates a BACnetServiceAckAtomicReadFileStreamBuilder
func NewBACnetServiceAckAtomicReadFileStreamBuilder() BACnetServiceAckAtomicReadFileStreamBuilder {
	return &_BACnetServiceAckAtomicReadFileStreamBuilder{_BACnetServiceAckAtomicReadFileStream: new(_BACnetServiceAckAtomicReadFileStream)}
}

type _BACnetServiceAckAtomicReadFileStreamBuilder struct {
	*_BACnetServiceAckAtomicReadFileStream

	err *utils.MultiError
}

var _ (BACnetServiceAckAtomicReadFileStreamBuilder) = (*_BACnetServiceAckAtomicReadFileStreamBuilder)(nil)

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) WithMandatoryFields(fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString) BACnetServiceAckAtomicReadFileStreamBuilder {
	return m.WithFileStartPosition(fileStartPosition).WithFileData(fileData)
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) WithFileStartPosition(fileStartPosition BACnetApplicationTagSignedInteger) BACnetServiceAckAtomicReadFileStreamBuilder {
	m.FileStartPosition = fileStartPosition
	return m
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) WithFileStartPositionBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetServiceAckAtomicReadFileStreamBuilder {
	builder := builderSupplier(m.FileStartPosition.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	m.FileStartPosition, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) WithFileData(fileData BACnetApplicationTagOctetString) BACnetServiceAckAtomicReadFileStreamBuilder {
	m.FileData = fileData
	return m
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) WithFileDataBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetServiceAckAtomicReadFileStreamBuilder {
	builder := builderSupplier(m.FileData.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	m.FileData, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return m
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) Build() (BACnetServiceAckAtomicReadFileStream, error) {
	if m.FileStartPosition == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'fileStartPosition' not set"))
	}
	if m.FileData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'fileData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetServiceAckAtomicReadFileStream.deepCopy(), nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) MustBuild() BACnetServiceAckAtomicReadFileStream {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetServiceAckAtomicReadFileStreamBuilder) DeepCopy() any {
	return m.CreateBACnetServiceAckAtomicReadFileStreamBuilder()
}

// CreateBACnetServiceAckAtomicReadFileStreamBuilder creates a BACnetServiceAckAtomicReadFileStreamBuilder
func (m *_BACnetServiceAckAtomicReadFileStream) CreateBACnetServiceAckAtomicReadFileStreamBuilder() BACnetServiceAckAtomicReadFileStreamBuilder {
	if m == nil {
		return NewBACnetServiceAckAtomicReadFileStreamBuilder()
	}
	return &_BACnetServiceAckAtomicReadFileStreamBuilder{_BACnetServiceAckAtomicReadFileStream: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicReadFileStream) GetParent() BACnetServiceAckAtomicReadFileStreamOrRecordContract {
	return m.BACnetServiceAckAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStream) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetFileData() BACnetApplicationTagOctetString {
	return m.FileData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileStream(structType any) BACnetServiceAckAtomicReadFileStream {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStream); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStream"
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckAtomicReadFileStream) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAckAtomicReadFileStreamOrRecord) (__bACnetServiceAckAtomicReadFileStream BACnetServiceAckAtomicReadFileStream, err error) {
	m.BACnetServiceAckAtomicReadFileStreamOrRecordContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStream"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileStream")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	fileData, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileData' field"))
	}
	m.FileData = fileData

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStream"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileStream")
	}

	return m, nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStream"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileStream")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", m.GetFileData(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileData' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStream"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileStream")
		}
		return nil
	}
	return m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicReadFileStream) IsBACnetServiceAckAtomicReadFileStream() {}

func (m *_BACnetServiceAckAtomicReadFileStream) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckAtomicReadFileStream) deepCopy() *_BACnetServiceAckAtomicReadFileStream {
	if m == nil {
		return nil
	}
	_BACnetServiceAckAtomicReadFileStreamCopy := &_BACnetServiceAckAtomicReadFileStream{
		m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord).deepCopy(),
		m.FileStartPosition.DeepCopy().(BACnetApplicationTagSignedInteger),
		m.FileData.DeepCopy().(BACnetApplicationTagOctetString),
	}
	m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord)._SubType = m
	return _BACnetServiceAckAtomicReadFileStreamCopy
}

func (m *_BACnetServiceAckAtomicReadFileStream) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
