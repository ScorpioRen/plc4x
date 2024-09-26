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

// BACnetConstructedDataTrendLogMultipleLogBuffer is the corresponding interface of BACnetConstructedDataTrendLogMultipleLogBuffer
type BACnetConstructedDataTrendLogMultipleLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetLogMultipleRecord
	// IsBACnetConstructedDataTrendLogMultipleLogBuffer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTrendLogMultipleLogBuffer()
	// CreateBuilder creates a BACnetConstructedDataTrendLogMultipleLogBufferBuilder
	CreateBACnetConstructedDataTrendLogMultipleLogBufferBuilder() BACnetConstructedDataTrendLogMultipleLogBufferBuilder
}

// _BACnetConstructedDataTrendLogMultipleLogBuffer is the data-structure of this message
type _BACnetConstructedDataTrendLogMultipleLogBuffer struct {
	BACnetConstructedDataContract
	FloorText []BACnetLogMultipleRecord
}

var _ BACnetConstructedDataTrendLogMultipleLogBuffer = (*_BACnetConstructedDataTrendLogMultipleLogBuffer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTrendLogMultipleLogBuffer)(nil)

// NewBACnetConstructedDataTrendLogMultipleLogBuffer factory function for _BACnetConstructedDataTrendLogMultipleLogBuffer
func NewBACnetConstructedDataTrendLogMultipleLogBuffer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, floorText []BACnetLogMultipleRecord, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogMultipleLogBuffer {
	_result := &_BACnetConstructedDataTrendLogMultipleLogBuffer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FloorText:                     floorText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTrendLogMultipleLogBufferBuilder is a builder for BACnetConstructedDataTrendLogMultipleLogBuffer
type BACnetConstructedDataTrendLogMultipleLogBufferBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floorText []BACnetLogMultipleRecord) BACnetConstructedDataTrendLogMultipleLogBufferBuilder
	// WithFloorText adds FloorText (property field)
	WithFloorText(...BACnetLogMultipleRecord) BACnetConstructedDataTrendLogMultipleLogBufferBuilder
	// Build builds the BACnetConstructedDataTrendLogMultipleLogBuffer or returns an error if something is wrong
	Build() (BACnetConstructedDataTrendLogMultipleLogBuffer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTrendLogMultipleLogBuffer
}

// NewBACnetConstructedDataTrendLogMultipleLogBufferBuilder() creates a BACnetConstructedDataTrendLogMultipleLogBufferBuilder
func NewBACnetConstructedDataTrendLogMultipleLogBufferBuilder() BACnetConstructedDataTrendLogMultipleLogBufferBuilder {
	return &_BACnetConstructedDataTrendLogMultipleLogBufferBuilder{_BACnetConstructedDataTrendLogMultipleLogBuffer: new(_BACnetConstructedDataTrendLogMultipleLogBuffer)}
}

type _BACnetConstructedDataTrendLogMultipleLogBufferBuilder struct {
	*_BACnetConstructedDataTrendLogMultipleLogBuffer

	err *utils.MultiError
}

var _ (BACnetConstructedDataTrendLogMultipleLogBufferBuilder) = (*_BACnetConstructedDataTrendLogMultipleLogBufferBuilder)(nil)

func (m *_BACnetConstructedDataTrendLogMultipleLogBufferBuilder) WithMandatoryFields(floorText []BACnetLogMultipleRecord) BACnetConstructedDataTrendLogMultipleLogBufferBuilder {
	return m.WithFloorText(floorText...)
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBufferBuilder) WithFloorText(floorText ...BACnetLogMultipleRecord) BACnetConstructedDataTrendLogMultipleLogBufferBuilder {
	m.FloorText = floorText
	return m
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBufferBuilder) Build() (BACnetConstructedDataTrendLogMultipleLogBuffer, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataTrendLogMultipleLogBuffer.deepCopy(), nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBufferBuilder) MustBuild() BACnetConstructedDataTrendLogMultipleLogBuffer {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBufferBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataTrendLogMultipleLogBufferBuilder()
}

// CreateBACnetConstructedDataTrendLogMultipleLogBufferBuilder creates a BACnetConstructedDataTrendLogMultipleLogBufferBuilder
func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) CreateBACnetConstructedDataTrendLogMultipleLogBufferBuilder() BACnetConstructedDataTrendLogMultipleLogBufferBuilder {
	if m == nil {
		return NewBACnetConstructedDataTrendLogMultipleLogBufferBuilder()
	}
	return &_BACnetConstructedDataTrendLogMultipleLogBufferBuilder{_BACnetConstructedDataTrendLogMultipleLogBuffer: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG_MULTIPLE
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetFloorText() []BACnetLogMultipleRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogMultipleLogBuffer(structType any) BACnetConstructedDataTrendLogMultipleLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataTrendLogMultipleLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogMultipleLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataTrendLogMultipleLogBuffer"
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrendLogMultipleLogBuffer BACnetConstructedDataTrendLogMultipleLogBuffer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogMultipleLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorText, err := ReadTerminatedArrayField[BACnetLogMultipleRecord](ctx, "floorText", ReadComplex[BACnetLogMultipleRecord](BACnetLogMultipleRecordParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	m.FloorText = floorText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogMultipleLogBuffer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogMultipleLogBuffer")
		}

		if err := WriteComplexTypeArrayField(ctx, "floorText", m.GetFloorText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'floorText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogMultipleLogBuffer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) IsBACnetConstructedDataTrendLogMultipleLogBuffer() {
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) deepCopy() *_BACnetConstructedDataTrendLogMultipleLogBuffer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTrendLogMultipleLogBufferCopy := &_BACnetConstructedDataTrendLogMultipleLogBuffer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLogMultipleRecord, BACnetLogMultipleRecord](m.FloorText),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTrendLogMultipleLogBufferCopy
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
