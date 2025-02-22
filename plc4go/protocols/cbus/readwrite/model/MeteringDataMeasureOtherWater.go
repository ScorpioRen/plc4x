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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MeteringDataMeasureOtherWater is the corresponding interface of MeteringDataMeasureOtherWater
type MeteringDataMeasureOtherWater interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// IsMeteringDataMeasureOtherWater is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataMeasureOtherWater()
	// CreateBuilder creates a MeteringDataMeasureOtherWaterBuilder
	CreateMeteringDataMeasureOtherWaterBuilder() MeteringDataMeasureOtherWaterBuilder
}

// _MeteringDataMeasureOtherWater is the data-structure of this message
type _MeteringDataMeasureOtherWater struct {
	MeteringDataContract
}

var _ MeteringDataMeasureOtherWater = (*_MeteringDataMeasureOtherWater)(nil)
var _ MeteringDataRequirements = (*_MeteringDataMeasureOtherWater)(nil)

// NewMeteringDataMeasureOtherWater factory function for _MeteringDataMeasureOtherWater
func NewMeteringDataMeasureOtherWater(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureOtherWater {
	_result := &_MeteringDataMeasureOtherWater{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MeteringDataMeasureOtherWaterBuilder is a builder for MeteringDataMeasureOtherWater
type MeteringDataMeasureOtherWaterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MeteringDataMeasureOtherWaterBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MeteringDataBuilder
	// Build builds the MeteringDataMeasureOtherWater or returns an error if something is wrong
	Build() (MeteringDataMeasureOtherWater, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataMeasureOtherWater
}

// NewMeteringDataMeasureOtherWaterBuilder() creates a MeteringDataMeasureOtherWaterBuilder
func NewMeteringDataMeasureOtherWaterBuilder() MeteringDataMeasureOtherWaterBuilder {
	return &_MeteringDataMeasureOtherWaterBuilder{_MeteringDataMeasureOtherWater: new(_MeteringDataMeasureOtherWater)}
}

type _MeteringDataMeasureOtherWaterBuilder struct {
	*_MeteringDataMeasureOtherWater

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataMeasureOtherWaterBuilder) = (*_MeteringDataMeasureOtherWaterBuilder)(nil)

func (b *_MeteringDataMeasureOtherWaterBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
	contract.(*_MeteringData)._SubType = b._MeteringDataMeasureOtherWater
}

func (b *_MeteringDataMeasureOtherWaterBuilder) WithMandatoryFields() MeteringDataMeasureOtherWaterBuilder {
	return b
}

func (b *_MeteringDataMeasureOtherWaterBuilder) Build() (MeteringDataMeasureOtherWater, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataMeasureOtherWater.deepCopy(), nil
}

func (b *_MeteringDataMeasureOtherWaterBuilder) MustBuild() MeteringDataMeasureOtherWater {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MeteringDataMeasureOtherWaterBuilder) Done() MeteringDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMeteringDataBuilder().(*_MeteringDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MeteringDataMeasureOtherWaterBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataMeasureOtherWaterBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataMeasureOtherWaterBuilder().(*_MeteringDataMeasureOtherWaterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataMeasureOtherWaterBuilder creates a MeteringDataMeasureOtherWaterBuilder
func (b *_MeteringDataMeasureOtherWater) CreateMeteringDataMeasureOtherWaterBuilder() MeteringDataMeasureOtherWaterBuilder {
	if b == nil {
		return NewMeteringDataMeasureOtherWaterBuilder()
	}
	return &_MeteringDataMeasureOtherWaterBuilder{_MeteringDataMeasureOtherWater: b.deepCopy()}
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

func (m *_MeteringDataMeasureOtherWater) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureOtherWater(structType any) MeteringDataMeasureOtherWater {
	if casted, ok := structType.(MeteringDataMeasureOtherWater); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureOtherWater); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureOtherWater) GetTypeName() string {
	return "MeteringDataMeasureOtherWater"
}

func (m *_MeteringDataMeasureOtherWater) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureOtherWater) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataMeasureOtherWater) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataMeasureOtherWater MeteringDataMeasureOtherWater, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureOtherWater"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureOtherWater")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureOtherWater"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureOtherWater")
	}

	return m, nil
}

func (m *_MeteringDataMeasureOtherWater) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureOtherWater) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureOtherWater"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureOtherWater")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureOtherWater"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureOtherWater")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureOtherWater) IsMeteringDataMeasureOtherWater() {}

func (m *_MeteringDataMeasureOtherWater) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataMeasureOtherWater) deepCopy() *_MeteringDataMeasureOtherWater {
	if m == nil {
		return nil
	}
	_MeteringDataMeasureOtherWaterCopy := &_MeteringDataMeasureOtherWater{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
	}
	_MeteringDataMeasureOtherWaterCopy.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataMeasureOtherWaterCopy
}

func (m *_MeteringDataMeasureOtherWater) String() string {
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
