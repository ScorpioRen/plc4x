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

// BACnetPropertyStatesPolarity is the corresponding interface of BACnetPropertyStatesPolarity
type BACnetPropertyStatesPolarity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetPolarity returns Polarity (property field)
	GetPolarity() BACnetPolarityTagged
	// IsBACnetPropertyStatesPolarity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesPolarity()
	// CreateBuilder creates a BACnetPropertyStatesPolarityBuilder
	CreateBACnetPropertyStatesPolarityBuilder() BACnetPropertyStatesPolarityBuilder
}

// _BACnetPropertyStatesPolarity is the data-structure of this message
type _BACnetPropertyStatesPolarity struct {
	BACnetPropertyStatesContract
	Polarity BACnetPolarityTagged
}

var _ BACnetPropertyStatesPolarity = (*_BACnetPropertyStatesPolarity)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesPolarity)(nil)

// NewBACnetPropertyStatesPolarity factory function for _BACnetPropertyStatesPolarity
func NewBACnetPropertyStatesPolarity(peekedTagHeader BACnetTagHeader, polarity BACnetPolarityTagged) *_BACnetPropertyStatesPolarity {
	if polarity == nil {
		panic("polarity of type BACnetPolarityTagged for BACnetPropertyStatesPolarity must not be nil")
	}
	_result := &_BACnetPropertyStatesPolarity{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Polarity:                     polarity,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesPolarityBuilder is a builder for BACnetPropertyStatesPolarity
type BACnetPropertyStatesPolarityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(polarity BACnetPolarityTagged) BACnetPropertyStatesPolarityBuilder
	// WithPolarity adds Polarity (property field)
	WithPolarity(BACnetPolarityTagged) BACnetPropertyStatesPolarityBuilder
	// WithPolarityBuilder adds Polarity (property field) which is build by the builder
	WithPolarityBuilder(func(BACnetPolarityTaggedBuilder) BACnetPolarityTaggedBuilder) BACnetPropertyStatesPolarityBuilder
	// Build builds the BACnetPropertyStatesPolarity or returns an error if something is wrong
	Build() (BACnetPropertyStatesPolarity, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesPolarity
}

// NewBACnetPropertyStatesPolarityBuilder() creates a BACnetPropertyStatesPolarityBuilder
func NewBACnetPropertyStatesPolarityBuilder() BACnetPropertyStatesPolarityBuilder {
	return &_BACnetPropertyStatesPolarityBuilder{_BACnetPropertyStatesPolarity: new(_BACnetPropertyStatesPolarity)}
}

type _BACnetPropertyStatesPolarityBuilder struct {
	*_BACnetPropertyStatesPolarity

	err *utils.MultiError
}

var _ (BACnetPropertyStatesPolarityBuilder) = (*_BACnetPropertyStatesPolarityBuilder)(nil)

func (m *_BACnetPropertyStatesPolarityBuilder) WithMandatoryFields(polarity BACnetPolarityTagged) BACnetPropertyStatesPolarityBuilder {
	return m.WithPolarity(polarity)
}

func (m *_BACnetPropertyStatesPolarityBuilder) WithPolarity(polarity BACnetPolarityTagged) BACnetPropertyStatesPolarityBuilder {
	m.Polarity = polarity
	return m
}

func (m *_BACnetPropertyStatesPolarityBuilder) WithPolarityBuilder(builderSupplier func(BACnetPolarityTaggedBuilder) BACnetPolarityTaggedBuilder) BACnetPropertyStatesPolarityBuilder {
	builder := builderSupplier(m.Polarity.CreateBACnetPolarityTaggedBuilder())
	var err error
	m.Polarity, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetPolarityTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesPolarityBuilder) Build() (BACnetPropertyStatesPolarity, error) {
	if m.Polarity == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'polarity' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesPolarity.deepCopy(), nil
}

func (m *_BACnetPropertyStatesPolarityBuilder) MustBuild() BACnetPropertyStatesPolarity {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesPolarityBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesPolarityBuilder()
}

// CreateBACnetPropertyStatesPolarityBuilder creates a BACnetPropertyStatesPolarityBuilder
func (m *_BACnetPropertyStatesPolarity) CreateBACnetPropertyStatesPolarityBuilder() BACnetPropertyStatesPolarityBuilder {
	if m == nil {
		return NewBACnetPropertyStatesPolarityBuilder()
	}
	return &_BACnetPropertyStatesPolarityBuilder{_BACnetPropertyStatesPolarity: m.deepCopy()}
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

func (m *_BACnetPropertyStatesPolarity) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesPolarity) GetPolarity() BACnetPolarityTagged {
	return m.Polarity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesPolarity(structType any) BACnetPropertyStatesPolarity {
	if casted, ok := structType.(BACnetPropertyStatesPolarity); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesPolarity); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesPolarity) GetTypeName() string {
	return "BACnetPropertyStatesPolarity"
}

func (m *_BACnetPropertyStatesPolarity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (polarity)
	lengthInBits += m.Polarity.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesPolarity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesPolarity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesPolarity BACnetPropertyStatesPolarity, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesPolarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesPolarity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	polarity, err := ReadSimpleField[BACnetPolarityTagged](ctx, "polarity", ReadComplex[BACnetPolarityTagged](BACnetPolarityTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'polarity' field"))
	}
	m.Polarity = polarity

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesPolarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesPolarity")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesPolarity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesPolarity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesPolarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesPolarity")
		}

		if err := WriteSimpleField[BACnetPolarityTagged](ctx, "polarity", m.GetPolarity(), WriteComplex[BACnetPolarityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'polarity' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesPolarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesPolarity")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesPolarity) IsBACnetPropertyStatesPolarity() {}

func (m *_BACnetPropertyStatesPolarity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesPolarity) deepCopy() *_BACnetPropertyStatesPolarity {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesPolarityCopy := &_BACnetPropertyStatesPolarity{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.Polarity.DeepCopy().(BACnetPolarityTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesPolarityCopy
}

func (m *_BACnetPropertyStatesPolarity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
