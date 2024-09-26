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

// BACnetPropertyStatesReliability is the corresponding interface of BACnetPropertyStatesReliability
type BACnetPropertyStatesReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// IsBACnetPropertyStatesReliability is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesReliability()
	// CreateBuilder creates a BACnetPropertyStatesReliabilityBuilder
	CreateBACnetPropertyStatesReliabilityBuilder() BACnetPropertyStatesReliabilityBuilder
}

// _BACnetPropertyStatesReliability is the data-structure of this message
type _BACnetPropertyStatesReliability struct {
	BACnetPropertyStatesContract
	Reliability BACnetReliabilityTagged
}

var _ BACnetPropertyStatesReliability = (*_BACnetPropertyStatesReliability)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesReliability)(nil)

// NewBACnetPropertyStatesReliability factory function for _BACnetPropertyStatesReliability
func NewBACnetPropertyStatesReliability(peekedTagHeader BACnetTagHeader, reliability BACnetReliabilityTagged) *_BACnetPropertyStatesReliability {
	if reliability == nil {
		panic("reliability of type BACnetReliabilityTagged for BACnetPropertyStatesReliability must not be nil")
	}
	_result := &_BACnetPropertyStatesReliability{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Reliability:                  reliability,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesReliabilityBuilder is a builder for BACnetPropertyStatesReliability
type BACnetPropertyStatesReliabilityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reliability BACnetReliabilityTagged) BACnetPropertyStatesReliabilityBuilder
	// WithReliability adds Reliability (property field)
	WithReliability(BACnetReliabilityTagged) BACnetPropertyStatesReliabilityBuilder
	// WithReliabilityBuilder adds Reliability (property field) which is build by the builder
	WithReliabilityBuilder(func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetPropertyStatesReliabilityBuilder
	// Build builds the BACnetPropertyStatesReliability or returns an error if something is wrong
	Build() (BACnetPropertyStatesReliability, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesReliability
}

// NewBACnetPropertyStatesReliabilityBuilder() creates a BACnetPropertyStatesReliabilityBuilder
func NewBACnetPropertyStatesReliabilityBuilder() BACnetPropertyStatesReliabilityBuilder {
	return &_BACnetPropertyStatesReliabilityBuilder{_BACnetPropertyStatesReliability: new(_BACnetPropertyStatesReliability)}
}

type _BACnetPropertyStatesReliabilityBuilder struct {
	*_BACnetPropertyStatesReliability

	err *utils.MultiError
}

var _ (BACnetPropertyStatesReliabilityBuilder) = (*_BACnetPropertyStatesReliabilityBuilder)(nil)

func (m *_BACnetPropertyStatesReliabilityBuilder) WithMandatoryFields(reliability BACnetReliabilityTagged) BACnetPropertyStatesReliabilityBuilder {
	return m.WithReliability(reliability)
}

func (m *_BACnetPropertyStatesReliabilityBuilder) WithReliability(reliability BACnetReliabilityTagged) BACnetPropertyStatesReliabilityBuilder {
	m.Reliability = reliability
	return m
}

func (m *_BACnetPropertyStatesReliabilityBuilder) WithReliabilityBuilder(builderSupplier func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetPropertyStatesReliabilityBuilder {
	builder := builderSupplier(m.Reliability.CreateBACnetReliabilityTaggedBuilder())
	var err error
	m.Reliability, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetReliabilityTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesReliabilityBuilder) Build() (BACnetPropertyStatesReliability, error) {
	if m.Reliability == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'reliability' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesReliability.deepCopy(), nil
}

func (m *_BACnetPropertyStatesReliabilityBuilder) MustBuild() BACnetPropertyStatesReliability {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesReliabilityBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesReliabilityBuilder()
}

// CreateBACnetPropertyStatesReliabilityBuilder creates a BACnetPropertyStatesReliabilityBuilder
func (m *_BACnetPropertyStatesReliability) CreateBACnetPropertyStatesReliabilityBuilder() BACnetPropertyStatesReliabilityBuilder {
	if m == nil {
		return NewBACnetPropertyStatesReliabilityBuilder()
	}
	return &_BACnetPropertyStatesReliabilityBuilder{_BACnetPropertyStatesReliability: m.deepCopy()}
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

func (m *_BACnetPropertyStatesReliability) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesReliability(structType any) BACnetPropertyStatesReliability {
	if casted, ok := structType.(BACnetPropertyStatesReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesReliability) GetTypeName() string {
	return "BACnetPropertyStatesReliability"
}

func (m *_BACnetPropertyStatesReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesReliability) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesReliability BACnetPropertyStatesReliability, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reliability, err := ReadSimpleField[BACnetReliabilityTagged](ctx, "reliability", ReadComplex[BACnetReliabilityTagged](BACnetReliabilityTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reliability' field"))
	}
	m.Reliability = reliability

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesReliability")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesReliability")
		}

		if err := WriteSimpleField[BACnetReliabilityTagged](ctx, "reliability", m.GetReliability(), WriteComplex[BACnetReliabilityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reliability' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesReliability")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesReliability) IsBACnetPropertyStatesReliability() {}

func (m *_BACnetPropertyStatesReliability) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesReliability) deepCopy() *_BACnetPropertyStatesReliability {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesReliabilityCopy := &_BACnetPropertyStatesReliability{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.Reliability.DeepCopy().(BACnetReliabilityTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesReliabilityCopy
}

func (m *_BACnetPropertyStatesReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
