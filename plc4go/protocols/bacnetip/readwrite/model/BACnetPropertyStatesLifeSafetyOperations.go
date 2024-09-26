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

// BACnetPropertyStatesLifeSafetyOperations is the corresponding interface of BACnetPropertyStatesLifeSafetyOperations
type BACnetPropertyStatesLifeSafetyOperations interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLifeSafetyOperations returns LifeSafetyOperations (property field)
	GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged
	// IsBACnetPropertyStatesLifeSafetyOperations is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLifeSafetyOperations()
	// CreateBuilder creates a BACnetPropertyStatesLifeSafetyOperationsBuilder
	CreateBACnetPropertyStatesLifeSafetyOperationsBuilder() BACnetPropertyStatesLifeSafetyOperationsBuilder
}

// _BACnetPropertyStatesLifeSafetyOperations is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyOperations struct {
	BACnetPropertyStatesContract
	LifeSafetyOperations BACnetLifeSafetyOperationTagged
}

var _ BACnetPropertyStatesLifeSafetyOperations = (*_BACnetPropertyStatesLifeSafetyOperations)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLifeSafetyOperations)(nil)

// NewBACnetPropertyStatesLifeSafetyOperations factory function for _BACnetPropertyStatesLifeSafetyOperations
func NewBACnetPropertyStatesLifeSafetyOperations(peekedTagHeader BACnetTagHeader, lifeSafetyOperations BACnetLifeSafetyOperationTagged) *_BACnetPropertyStatesLifeSafetyOperations {
	if lifeSafetyOperations == nil {
		panic("lifeSafetyOperations of type BACnetLifeSafetyOperationTagged for BACnetPropertyStatesLifeSafetyOperations must not be nil")
	}
	_result := &_BACnetPropertyStatesLifeSafetyOperations{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LifeSafetyOperations:         lifeSafetyOperations,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLifeSafetyOperationsBuilder is a builder for BACnetPropertyStatesLifeSafetyOperations
type BACnetPropertyStatesLifeSafetyOperationsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetPropertyStatesLifeSafetyOperationsBuilder
	// WithLifeSafetyOperations adds LifeSafetyOperations (property field)
	WithLifeSafetyOperations(BACnetLifeSafetyOperationTagged) BACnetPropertyStatesLifeSafetyOperationsBuilder
	// WithLifeSafetyOperationsBuilder adds LifeSafetyOperations (property field) which is build by the builder
	WithLifeSafetyOperationsBuilder(func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetPropertyStatesLifeSafetyOperationsBuilder
	// Build builds the BACnetPropertyStatesLifeSafetyOperations or returns an error if something is wrong
	Build() (BACnetPropertyStatesLifeSafetyOperations, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLifeSafetyOperations
}

// NewBACnetPropertyStatesLifeSafetyOperationsBuilder() creates a BACnetPropertyStatesLifeSafetyOperationsBuilder
func NewBACnetPropertyStatesLifeSafetyOperationsBuilder() BACnetPropertyStatesLifeSafetyOperationsBuilder {
	return &_BACnetPropertyStatesLifeSafetyOperationsBuilder{_BACnetPropertyStatesLifeSafetyOperations: new(_BACnetPropertyStatesLifeSafetyOperations)}
}

type _BACnetPropertyStatesLifeSafetyOperationsBuilder struct {
	*_BACnetPropertyStatesLifeSafetyOperations

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLifeSafetyOperationsBuilder) = (*_BACnetPropertyStatesLifeSafetyOperationsBuilder)(nil)

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) WithMandatoryFields(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetPropertyStatesLifeSafetyOperationsBuilder {
	return m.WithLifeSafetyOperations(lifeSafetyOperations)
}

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) WithLifeSafetyOperations(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetPropertyStatesLifeSafetyOperationsBuilder {
	m.LifeSafetyOperations = lifeSafetyOperations
	return m
}

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) WithLifeSafetyOperationsBuilder(builderSupplier func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetPropertyStatesLifeSafetyOperationsBuilder {
	builder := builderSupplier(m.LifeSafetyOperations.CreateBACnetLifeSafetyOperationTaggedBuilder())
	var err error
	m.LifeSafetyOperations, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetLifeSafetyOperationTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) Build() (BACnetPropertyStatesLifeSafetyOperations, error) {
	if m.LifeSafetyOperations == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lifeSafetyOperations' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesLifeSafetyOperations.deepCopy(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) MustBuild() BACnetPropertyStatesLifeSafetyOperations {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesLifeSafetyOperationsBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesLifeSafetyOperationsBuilder()
}

// CreateBACnetPropertyStatesLifeSafetyOperationsBuilder creates a BACnetPropertyStatesLifeSafetyOperationsBuilder
func (m *_BACnetPropertyStatesLifeSafetyOperations) CreateBACnetPropertyStatesLifeSafetyOperationsBuilder() BACnetPropertyStatesLifeSafetyOperationsBuilder {
	if m == nil {
		return NewBACnetPropertyStatesLifeSafetyOperationsBuilder()
	}
	return &_BACnetPropertyStatesLifeSafetyOperationsBuilder{_BACnetPropertyStatesLifeSafetyOperations: m.deepCopy()}
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

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged {
	return m.LifeSafetyOperations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyOperations(structType any) BACnetPropertyStatesLifeSafetyOperations {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyOperations); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyOperations); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyOperations"
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lifeSafetyOperations)
	lengthInBits += m.LifeSafetyOperations.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLifeSafetyOperations BACnetPropertyStatesLifeSafetyOperations, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyOperations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyOperations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyOperations, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyOperations' field"))
	}
	m.LifeSafetyOperations = lifeSafetyOperations

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyOperations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyOperations")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyOperations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyOperations")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", m.GetLifeSafetyOperations(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyOperations' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyOperations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyOperations")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) IsBACnetPropertyStatesLifeSafetyOperations() {}

func (m *_BACnetPropertyStatesLifeSafetyOperations) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) deepCopy() *_BACnetPropertyStatesLifeSafetyOperations {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLifeSafetyOperationsCopy := &_BACnetPropertyStatesLifeSafetyOperations{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.LifeSafetyOperations.DeepCopy().(BACnetLifeSafetyOperationTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLifeSafetyOperationsCopy
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
