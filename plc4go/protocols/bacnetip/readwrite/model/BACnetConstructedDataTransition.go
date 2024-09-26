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

// BACnetConstructedDataTransition is the corresponding interface of BACnetConstructedDataTransition
type BACnetConstructedDataTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTransition returns Transition (property field)
	GetTransition() BACnetLightingTransitionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingTransitionTagged
	// IsBACnetConstructedDataTransition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTransition()
	// CreateBuilder creates a BACnetConstructedDataTransitionBuilder
	CreateBACnetConstructedDataTransitionBuilder() BACnetConstructedDataTransitionBuilder
}

// _BACnetConstructedDataTransition is the data-structure of this message
type _BACnetConstructedDataTransition struct {
	BACnetConstructedDataContract
	Transition BACnetLightingTransitionTagged
}

var _ BACnetConstructedDataTransition = (*_BACnetConstructedDataTransition)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTransition)(nil)

// NewBACnetConstructedDataTransition factory function for _BACnetConstructedDataTransition
func NewBACnetConstructedDataTransition(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, transition BACnetLightingTransitionTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTransition {
	if transition == nil {
		panic("transition of type BACnetLightingTransitionTagged for BACnetConstructedDataTransition must not be nil")
	}
	_result := &_BACnetConstructedDataTransition{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Transition:                    transition,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTransitionBuilder is a builder for BACnetConstructedDataTransition
type BACnetConstructedDataTransitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(transition BACnetLightingTransitionTagged) BACnetConstructedDataTransitionBuilder
	// WithTransition adds Transition (property field)
	WithTransition(BACnetLightingTransitionTagged) BACnetConstructedDataTransitionBuilder
	// WithTransitionBuilder adds Transition (property field) which is build by the builder
	WithTransitionBuilder(func(BACnetLightingTransitionTaggedBuilder) BACnetLightingTransitionTaggedBuilder) BACnetConstructedDataTransitionBuilder
	// Build builds the BACnetConstructedDataTransition or returns an error if something is wrong
	Build() (BACnetConstructedDataTransition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTransition
}

// NewBACnetConstructedDataTransitionBuilder() creates a BACnetConstructedDataTransitionBuilder
func NewBACnetConstructedDataTransitionBuilder() BACnetConstructedDataTransitionBuilder {
	return &_BACnetConstructedDataTransitionBuilder{_BACnetConstructedDataTransition: new(_BACnetConstructedDataTransition)}
}

type _BACnetConstructedDataTransitionBuilder struct {
	*_BACnetConstructedDataTransition

	err *utils.MultiError
}

var _ (BACnetConstructedDataTransitionBuilder) = (*_BACnetConstructedDataTransitionBuilder)(nil)

func (m *_BACnetConstructedDataTransitionBuilder) WithMandatoryFields(transition BACnetLightingTransitionTagged) BACnetConstructedDataTransitionBuilder {
	return m.WithTransition(transition)
}

func (m *_BACnetConstructedDataTransitionBuilder) WithTransition(transition BACnetLightingTransitionTagged) BACnetConstructedDataTransitionBuilder {
	m.Transition = transition
	return m
}

func (m *_BACnetConstructedDataTransitionBuilder) WithTransitionBuilder(builderSupplier func(BACnetLightingTransitionTaggedBuilder) BACnetLightingTransitionTaggedBuilder) BACnetConstructedDataTransitionBuilder {
	builder := builderSupplier(m.Transition.CreateBACnetLightingTransitionTaggedBuilder())
	var err error
	m.Transition, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetLightingTransitionTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataTransitionBuilder) Build() (BACnetConstructedDataTransition, error) {
	if m.Transition == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'transition' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataTransition.deepCopy(), nil
}

func (m *_BACnetConstructedDataTransitionBuilder) MustBuild() BACnetConstructedDataTransition {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataTransitionBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataTransitionBuilder()
}

// CreateBACnetConstructedDataTransitionBuilder creates a BACnetConstructedDataTransitionBuilder
func (m *_BACnetConstructedDataTransition) CreateBACnetConstructedDataTransitionBuilder() BACnetConstructedDataTransitionBuilder {
	if m == nil {
		return NewBACnetConstructedDataTransitionBuilder()
	}
	return &_BACnetConstructedDataTransitionBuilder{_BACnetConstructedDataTransition: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTransition) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRANSITION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTransition) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetTransition() BACnetLightingTransitionTagged {
	return m.Transition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetActualValue() BACnetLightingTransitionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLightingTransitionTagged(m.GetTransition())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTransition(structType any) BACnetConstructedDataTransition {
	if casted, ok := structType.(BACnetConstructedDataTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTransition) GetTypeName() string {
	return "BACnetConstructedDataTransition"
}

func (m *_BACnetConstructedDataTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (transition)
	lengthInBits += m.Transition.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTransition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTransition BACnetConstructedDataTransition, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transition, err := ReadSimpleField[BACnetLightingTransitionTagged](ctx, "transition", ReadComplex[BACnetLightingTransitionTagged](BACnetLightingTransitionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transition' field"))
	}
	m.Transition = transition

	actualValue, err := ReadVirtualField[BACnetLightingTransitionTagged](ctx, "actualValue", (*BACnetLightingTransitionTagged)(nil), transition)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTransition")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTransition")
		}

		if err := WriteSimpleField[BACnetLightingTransitionTagged](ctx, "transition", m.GetTransition(), WriteComplex[BACnetLightingTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transition' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTransition")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTransition) IsBACnetConstructedDataTransition() {}

func (m *_BACnetConstructedDataTransition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTransition) deepCopy() *_BACnetConstructedDataTransition {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTransitionCopy := &_BACnetConstructedDataTransition{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Transition.DeepCopy().(BACnetLightingTransitionTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTransitionCopy
}

func (m *_BACnetConstructedDataTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
