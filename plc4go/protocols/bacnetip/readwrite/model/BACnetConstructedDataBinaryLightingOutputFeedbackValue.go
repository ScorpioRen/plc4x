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

// BACnetConstructedDataBinaryLightingOutputFeedbackValue is the corresponding interface of BACnetConstructedDataBinaryLightingOutputFeedbackValue
type BACnetConstructedDataBinaryLightingOutputFeedbackValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetBinaryLightingPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryLightingPVTagged
	// IsBACnetConstructedDataBinaryLightingOutputFeedbackValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryLightingOutputFeedbackValue()
	// CreateBuilder creates a BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
	CreateBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder() BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
}

// _BACnetConstructedDataBinaryLightingOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputFeedbackValue struct {
	BACnetConstructedDataContract
	FeedbackValue BACnetBinaryLightingPVTagged
}

var _ BACnetConstructedDataBinaryLightingOutputFeedbackValue = (*_BACnetConstructedDataBinaryLightingOutputFeedbackValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryLightingOutputFeedbackValue)(nil)

// NewBACnetConstructedDataBinaryLightingOutputFeedbackValue factory function for _BACnetConstructedDataBinaryLightingOutputFeedbackValue
func NewBACnetConstructedDataBinaryLightingOutputFeedbackValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, feedbackValue BACnetBinaryLightingPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	if feedbackValue == nil {
		panic("feedbackValue of type BACnetBinaryLightingPVTagged for BACnetConstructedDataBinaryLightingOutputFeedbackValue must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryLightingOutputFeedbackValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FeedbackValue:                 feedbackValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder is a builder for BACnetConstructedDataBinaryLightingOutputFeedbackValue
type BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(feedbackValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
	// WithFeedbackValue adds FeedbackValue (property field)
	WithFeedbackValue(BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
	// WithFeedbackValueBuilder adds FeedbackValue (property field) which is build by the builder
	WithFeedbackValueBuilder(func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryLightingOutputFeedbackValue or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryLightingOutputFeedbackValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryLightingOutputFeedbackValue
}

// NewBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder() creates a BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
func NewBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder() BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder {
	return &_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder{_BACnetConstructedDataBinaryLightingOutputFeedbackValue: new(_BACnetConstructedDataBinaryLightingOutputFeedbackValue)}
}

type _BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder struct {
	*_BACnetConstructedDataBinaryLightingOutputFeedbackValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) = (*_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder)(nil)

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBinaryLightingOutputFeedbackValue
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) WithMandatoryFields(feedbackValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder {
	return b.WithFeedbackValue(feedbackValue)
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) WithFeedbackValue(feedbackValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder {
	b.FeedbackValue = feedbackValue
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) WithFeedbackValueBuilder(builderSupplier func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder {
	builder := builderSupplier(b.FeedbackValue.CreateBACnetBinaryLightingPVTaggedBuilder())
	var err error
	b.FeedbackValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBinaryLightingPVTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) Build() (BACnetConstructedDataBinaryLightingOutputFeedbackValue, error) {
	if b.FeedbackValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'feedbackValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryLightingOutputFeedbackValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) MustBuild() BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder().(*_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder creates a BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder
func (b *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) CreateBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder() BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder()
	}
	return &_BACnetConstructedDataBinaryLightingOutputFeedbackValueBuilder{_BACnetConstructedDataBinaryLightingOutputFeedbackValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetFeedbackValue() BACnetBinaryLightingPVTagged {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetActualValue() BACnetBinaryLightingPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryLightingPVTagged(m.GetFeedbackValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputFeedbackValue(structType any) BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputFeedbackValue"
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryLightingOutputFeedbackValue BACnetConstructedDataBinaryLightingOutputFeedbackValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	feedbackValue, err := ReadSimpleField[BACnetBinaryLightingPVTagged](ctx, "feedbackValue", ReadComplex[BACnetBinaryLightingPVTagged](BACnetBinaryLightingPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'feedbackValue' field"))
	}
	m.FeedbackValue = feedbackValue

	actualValue, err := ReadVirtualField[BACnetBinaryLightingPVTagged](ctx, "actualValue", (*BACnetBinaryLightingPVTagged)(nil), feedbackValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}

		if err := WriteSimpleField[BACnetBinaryLightingPVTagged](ctx, "feedbackValue", m.GetFeedbackValue(), WriteComplex[BACnetBinaryLightingPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'feedbackValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) IsBACnetConstructedDataBinaryLightingOutputFeedbackValue() {
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) deepCopy() *_BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryLightingOutputFeedbackValueCopy := &_BACnetConstructedDataBinaryLightingOutputFeedbackValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetBinaryLightingPVTagged](m.FeedbackValue),
	}
	_BACnetConstructedDataBinaryLightingOutputFeedbackValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryLightingOutputFeedbackValueCopy
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) String() string {
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
