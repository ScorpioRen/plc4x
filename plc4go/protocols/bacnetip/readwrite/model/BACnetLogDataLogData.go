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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogData is the corresponding interface of BACnetLogDataLogData
type BACnetLogDataLogData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogData
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetLogData returns LogData (property field)
	GetLogData() []BACnetLogDataLogDataEntry
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetLogDataLogData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogData()
	// CreateBuilder creates a BACnetLogDataLogDataBuilder
	CreateBACnetLogDataLogDataBuilder() BACnetLogDataLogDataBuilder
}

// _BACnetLogDataLogData is the data-structure of this message
type _BACnetLogDataLogData struct {
	BACnetLogDataContract
	InnerOpeningTag BACnetOpeningTag
	LogData         []BACnetLogDataLogDataEntry
	InnerClosingTag BACnetClosingTag
}

var _ BACnetLogDataLogData = (*_BACnetLogDataLogData)(nil)
var _ BACnetLogDataRequirements = (*_BACnetLogDataLogData)(nil)

// NewBACnetLogDataLogData factory function for _BACnetLogDataLogData
func NewBACnetLogDataLogData(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, logData []BACnetLogDataLogDataEntry, innerClosingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogDataLogData {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetLogDataLogData must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetLogDataLogData must not be nil")
	}
	_result := &_BACnetLogDataLogData{
		BACnetLogDataContract: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
		InnerOpeningTag:       innerOpeningTag,
		LogData:               logData,
		InnerClosingTag:       innerClosingTag,
	}
	_result.BACnetLogDataContract.(*_BACnetLogData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogDataBuilder is a builder for BACnetLogDataLogData
type BACnetLogDataLogDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, logData []BACnetLogDataLogDataEntry, innerClosingTag BACnetClosingTag) BACnetLogDataLogDataBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetLogDataLogDataBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLogDataLogDataBuilder
	// WithLogData adds LogData (property field)
	WithLogData(...BACnetLogDataLogDataEntry) BACnetLogDataLogDataBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetLogDataLogDataBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLogDataLogDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogDataBuilder
	// Build builds the BACnetLogDataLogData or returns an error if something is wrong
	Build() (BACnetLogDataLogData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogData
}

// NewBACnetLogDataLogDataBuilder() creates a BACnetLogDataLogDataBuilder
func NewBACnetLogDataLogDataBuilder() BACnetLogDataLogDataBuilder {
	return &_BACnetLogDataLogDataBuilder{_BACnetLogDataLogData: new(_BACnetLogDataLogData)}
}

type _BACnetLogDataLogDataBuilder struct {
	*_BACnetLogDataLogData

	parentBuilder *_BACnetLogDataBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogDataBuilder) = (*_BACnetLogDataLogDataBuilder)(nil)

func (b *_BACnetLogDataLogDataBuilder) setParent(contract BACnetLogDataContract) {
	b.BACnetLogDataContract = contract
	contract.(*_BACnetLogData)._SubType = b._BACnetLogDataLogData
}

func (b *_BACnetLogDataLogDataBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, logData []BACnetLogDataLogDataEntry, innerClosingTag BACnetClosingTag) BACnetLogDataLogDataBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithLogData(logData...).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetLogDataLogDataBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetLogDataLogDataBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetLogDataLogDataBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLogDataLogDataBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataBuilder) WithLogData(logData ...BACnetLogDataLogDataEntry) BACnetLogDataLogDataBuilder {
	b.LogData = logData
	return b
}

func (b *_BACnetLogDataLogDataBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetLogDataLogDataBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetLogDataLogDataBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLogDataLogDataBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataBuilder) Build() (BACnetLogDataLogData, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogData.deepCopy(), nil
}

func (b *_BACnetLogDataLogDataBuilder) MustBuild() BACnetLogDataLogData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogDataLogDataBuilder) Done() BACnetLogDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogDataBuilder().(*_BACnetLogDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogDataLogDataBuilder) buildForBACnetLogData() (BACnetLogData, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogDataBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogDataBuilder().(*_BACnetLogDataLogDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogDataBuilder creates a BACnetLogDataLogDataBuilder
func (b *_BACnetLogDataLogData) CreateBACnetLogDataLogDataBuilder() BACnetLogDataLogDataBuilder {
	if b == nil {
		return NewBACnetLogDataLogDataBuilder()
	}
	return &_BACnetLogDataLogDataBuilder{_BACnetLogDataLogData: b.deepCopy()}
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

func (m *_BACnetLogDataLogData) GetParent() BACnetLogDataContract {
	return m.BACnetLogDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogData) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetLogDataLogData) GetLogData() []BACnetLogDataLogDataEntry {
	return m.LogData
}

func (m *_BACnetLogDataLogData) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogData(structType any) BACnetLogDataLogData {
	if casted, ok := structType.(BACnetLogDataLogData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogData) GetTypeName() string {
	return "BACnetLogDataLogData"
}

func (m *_BACnetLogDataLogData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataContract.(*_BACnetLogData).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.LogData) > 0 {
		for _, element := range m.LogData {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogData, tagNumber uint8) (__bACnetLogDataLogData BACnetLogDataLogData, err error) {
	m.BACnetLogDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	logData, err := ReadTerminatedArrayField[BACnetLogDataLogDataEntry](ctx, "logData", ReadComplex[BACnetLogDataLogDataEntry](BACnetLogDataLogDataEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logData' field"))
	}
	m.LogData = logData

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogData")
	}

	return m, nil
}

func (m *_BACnetLogDataLogData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogData")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "logData", m.GetLogData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'logData' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogData")
		}
		return nil
	}
	return m.BACnetLogDataContract.(*_BACnetLogData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogData) IsBACnetLogDataLogData() {}

func (m *_BACnetLogDataLogData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogData) deepCopy() *_BACnetLogDataLogData {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataCopy := &_BACnetLogDataLogData{
		m.BACnetLogDataContract.(*_BACnetLogData).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopySlice[BACnetLogDataLogDataEntry, BACnetLogDataLogDataEntry](m.LogData),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	_BACnetLogDataLogDataCopy.BACnetLogDataContract.(*_BACnetLogData)._SubType = m
	return _BACnetLogDataLogDataCopy
}

func (m *_BACnetLogDataLogData) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
