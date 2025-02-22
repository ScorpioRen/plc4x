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

// SALDataPoolsSpasPondsFountainsControl is the corresponding interface of SALDataPoolsSpasPondsFountainsControl
type SALDataPoolsSpasPondsFountainsControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetPoolsSpaPondsFountainsData returns PoolsSpaPondsFountainsData (property field)
	GetPoolsSpaPondsFountainsData() LightingData
	// IsSALDataPoolsSpasPondsFountainsControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataPoolsSpasPondsFountainsControl()
	// CreateBuilder creates a SALDataPoolsSpasPondsFountainsControlBuilder
	CreateSALDataPoolsSpasPondsFountainsControlBuilder() SALDataPoolsSpasPondsFountainsControlBuilder
}

// _SALDataPoolsSpasPondsFountainsControl is the data-structure of this message
type _SALDataPoolsSpasPondsFountainsControl struct {
	SALDataContract
	PoolsSpaPondsFountainsData LightingData
}

var _ SALDataPoolsSpasPondsFountainsControl = (*_SALDataPoolsSpasPondsFountainsControl)(nil)
var _ SALDataRequirements = (*_SALDataPoolsSpasPondsFountainsControl)(nil)

// NewSALDataPoolsSpasPondsFountainsControl factory function for _SALDataPoolsSpasPondsFountainsControl
func NewSALDataPoolsSpasPondsFountainsControl(salData SALData, poolsSpaPondsFountainsData LightingData) *_SALDataPoolsSpasPondsFountainsControl {
	if poolsSpaPondsFountainsData == nil {
		panic("poolsSpaPondsFountainsData of type LightingData for SALDataPoolsSpasPondsFountainsControl must not be nil")
	}
	_result := &_SALDataPoolsSpasPondsFountainsControl{
		SALDataContract:            NewSALData(salData),
		PoolsSpaPondsFountainsData: poolsSpaPondsFountainsData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataPoolsSpasPondsFountainsControlBuilder is a builder for SALDataPoolsSpasPondsFountainsControl
type SALDataPoolsSpasPondsFountainsControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(poolsSpaPondsFountainsData LightingData) SALDataPoolsSpasPondsFountainsControlBuilder
	// WithPoolsSpaPondsFountainsData adds PoolsSpaPondsFountainsData (property field)
	WithPoolsSpaPondsFountainsData(LightingData) SALDataPoolsSpasPondsFountainsControlBuilder
	// WithPoolsSpaPondsFountainsDataBuilder adds PoolsSpaPondsFountainsData (property field) which is build by the builder
	WithPoolsSpaPondsFountainsDataBuilder(func(LightingDataBuilder) LightingDataBuilder) SALDataPoolsSpasPondsFountainsControlBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SALDataBuilder
	// Build builds the SALDataPoolsSpasPondsFountainsControl or returns an error if something is wrong
	Build() (SALDataPoolsSpasPondsFountainsControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataPoolsSpasPondsFountainsControl
}

// NewSALDataPoolsSpasPondsFountainsControlBuilder() creates a SALDataPoolsSpasPondsFountainsControlBuilder
func NewSALDataPoolsSpasPondsFountainsControlBuilder() SALDataPoolsSpasPondsFountainsControlBuilder {
	return &_SALDataPoolsSpasPondsFountainsControlBuilder{_SALDataPoolsSpasPondsFountainsControl: new(_SALDataPoolsSpasPondsFountainsControl)}
}

type _SALDataPoolsSpasPondsFountainsControlBuilder struct {
	*_SALDataPoolsSpasPondsFountainsControl

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataPoolsSpasPondsFountainsControlBuilder) = (*_SALDataPoolsSpasPondsFountainsControlBuilder)(nil)

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
	contract.(*_SALData)._SubType = b._SALDataPoolsSpasPondsFountainsControl
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) WithMandatoryFields(poolsSpaPondsFountainsData LightingData) SALDataPoolsSpasPondsFountainsControlBuilder {
	return b.WithPoolsSpaPondsFountainsData(poolsSpaPondsFountainsData)
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) WithPoolsSpaPondsFountainsData(poolsSpaPondsFountainsData LightingData) SALDataPoolsSpasPondsFountainsControlBuilder {
	b.PoolsSpaPondsFountainsData = poolsSpaPondsFountainsData
	return b
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) WithPoolsSpaPondsFountainsDataBuilder(builderSupplier func(LightingDataBuilder) LightingDataBuilder) SALDataPoolsSpasPondsFountainsControlBuilder {
	builder := builderSupplier(b.PoolsSpaPondsFountainsData.CreateLightingDataBuilder())
	var err error
	b.PoolsSpaPondsFountainsData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LightingDataBuilder failed"))
	}
	return b
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) Build() (SALDataPoolsSpasPondsFountainsControl, error) {
	if b.PoolsSpaPondsFountainsData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'poolsSpaPondsFountainsData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataPoolsSpasPondsFountainsControl.deepCopy(), nil
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) MustBuild() SALDataPoolsSpasPondsFountainsControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) Done() SALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSALDataBuilder().(*_SALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataPoolsSpasPondsFountainsControlBuilder) DeepCopy() any {
	_copy := b.CreateSALDataPoolsSpasPondsFountainsControlBuilder().(*_SALDataPoolsSpasPondsFountainsControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataPoolsSpasPondsFountainsControlBuilder creates a SALDataPoolsSpasPondsFountainsControlBuilder
func (b *_SALDataPoolsSpasPondsFountainsControl) CreateSALDataPoolsSpasPondsFountainsControlBuilder() SALDataPoolsSpasPondsFountainsControlBuilder {
	if b == nil {
		return NewSALDataPoolsSpasPondsFountainsControlBuilder()
	}
	return &_SALDataPoolsSpasPondsFountainsControlBuilder{_SALDataPoolsSpasPondsFountainsControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) GetApplicationId() ApplicationId {
	return ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) GetPoolsSpaPondsFountainsData() LightingData {
	return m.PoolsSpaPondsFountainsData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataPoolsSpasPondsFountainsControl(structType any) SALDataPoolsSpasPondsFountainsControl {
	if casted, ok := structType.(SALDataPoolsSpasPondsFountainsControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataPoolsSpasPondsFountainsControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetTypeName() string {
	return "SALDataPoolsSpasPondsFountainsControl"
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (poolsSpaPondsFountainsData)
	lengthInBits += m.PoolsSpaPondsFountainsData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataPoolsSpasPondsFountainsControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataPoolsSpasPondsFountainsControl SALDataPoolsSpasPondsFountainsControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataPoolsSpasPondsFountainsControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataPoolsSpasPondsFountainsControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	poolsSpaPondsFountainsData, err := ReadSimpleField[LightingData](ctx, "poolsSpaPondsFountainsData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'poolsSpaPondsFountainsData' field"))
	}
	m.PoolsSpaPondsFountainsData = poolsSpaPondsFountainsData

	if closeErr := readBuffer.CloseContext("SALDataPoolsSpasPondsFountainsControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataPoolsSpasPondsFountainsControl")
	}

	return m, nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataPoolsSpasPondsFountainsControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataPoolsSpasPondsFountainsControl")
		}

		if err := WriteSimpleField[LightingData](ctx, "poolsSpaPondsFountainsData", m.GetPoolsSpaPondsFountainsData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'poolsSpaPondsFountainsData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataPoolsSpasPondsFountainsControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataPoolsSpasPondsFountainsControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataPoolsSpasPondsFountainsControl) IsSALDataPoolsSpasPondsFountainsControl() {}

func (m *_SALDataPoolsSpasPondsFountainsControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataPoolsSpasPondsFountainsControl) deepCopy() *_SALDataPoolsSpasPondsFountainsControl {
	if m == nil {
		return nil
	}
	_SALDataPoolsSpasPondsFountainsControlCopy := &_SALDataPoolsSpasPondsFountainsControl{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[LightingData](m.PoolsSpaPondsFountainsData),
	}
	_SALDataPoolsSpasPondsFountainsControlCopy.SALDataContract.(*_SALData)._SubType = m
	return _SALDataPoolsSpasPondsFountainsControlCopy
}

func (m *_SALDataPoolsSpasPondsFountainsControl) String() string {
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
