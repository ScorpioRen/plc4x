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

// SALDataAirConditioning is the corresponding interface of SALDataAirConditioning
type SALDataAirConditioning interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetAirConditioningData returns AirConditioningData (property field)
	GetAirConditioningData() AirConditioningData
	// IsSALDataAirConditioning is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataAirConditioning()
	// CreateBuilder creates a SALDataAirConditioningBuilder
	CreateSALDataAirConditioningBuilder() SALDataAirConditioningBuilder
}

// _SALDataAirConditioning is the data-structure of this message
type _SALDataAirConditioning struct {
	SALDataContract
	AirConditioningData AirConditioningData
}

var _ SALDataAirConditioning = (*_SALDataAirConditioning)(nil)
var _ SALDataRequirements = (*_SALDataAirConditioning)(nil)

// NewSALDataAirConditioning factory function for _SALDataAirConditioning
func NewSALDataAirConditioning(salData SALData, airConditioningData AirConditioningData) *_SALDataAirConditioning {
	if airConditioningData == nil {
		panic("airConditioningData of type AirConditioningData for SALDataAirConditioning must not be nil")
	}
	_result := &_SALDataAirConditioning{
		SALDataContract:     NewSALData(salData),
		AirConditioningData: airConditioningData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataAirConditioningBuilder is a builder for SALDataAirConditioning
type SALDataAirConditioningBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(airConditioningData AirConditioningData) SALDataAirConditioningBuilder
	// WithAirConditioningData adds AirConditioningData (property field)
	WithAirConditioningData(AirConditioningData) SALDataAirConditioningBuilder
	// WithAirConditioningDataBuilder adds AirConditioningData (property field) which is build by the builder
	WithAirConditioningDataBuilder(func(AirConditioningDataBuilder) AirConditioningDataBuilder) SALDataAirConditioningBuilder
	// Build builds the SALDataAirConditioning or returns an error if something is wrong
	Build() (SALDataAirConditioning, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataAirConditioning
}

// NewSALDataAirConditioningBuilder() creates a SALDataAirConditioningBuilder
func NewSALDataAirConditioningBuilder() SALDataAirConditioningBuilder {
	return &_SALDataAirConditioningBuilder{_SALDataAirConditioning: new(_SALDataAirConditioning)}
}

type _SALDataAirConditioningBuilder struct {
	*_SALDataAirConditioning

	err *utils.MultiError
}

var _ (SALDataAirConditioningBuilder) = (*_SALDataAirConditioningBuilder)(nil)

func (m *_SALDataAirConditioningBuilder) WithMandatoryFields(airConditioningData AirConditioningData) SALDataAirConditioningBuilder {
	return m.WithAirConditioningData(airConditioningData)
}

func (m *_SALDataAirConditioningBuilder) WithAirConditioningData(airConditioningData AirConditioningData) SALDataAirConditioningBuilder {
	m.AirConditioningData = airConditioningData
	return m
}

func (m *_SALDataAirConditioningBuilder) WithAirConditioningDataBuilder(builderSupplier func(AirConditioningDataBuilder) AirConditioningDataBuilder) SALDataAirConditioningBuilder {
	builder := builderSupplier(m.AirConditioningData.CreateAirConditioningDataBuilder())
	var err error
	m.AirConditioningData, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "AirConditioningDataBuilder failed"))
	}
	return m
}

func (m *_SALDataAirConditioningBuilder) Build() (SALDataAirConditioning, error) {
	if m.AirConditioningData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'airConditioningData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SALDataAirConditioning.deepCopy(), nil
}

func (m *_SALDataAirConditioningBuilder) MustBuild() SALDataAirConditioning {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SALDataAirConditioningBuilder) DeepCopy() any {
	return m.CreateSALDataAirConditioningBuilder()
}

// CreateSALDataAirConditioningBuilder creates a SALDataAirConditioningBuilder
func (m *_SALDataAirConditioning) CreateSALDataAirConditioningBuilder() SALDataAirConditioningBuilder {
	if m == nil {
		return NewSALDataAirConditioningBuilder()
	}
	return &_SALDataAirConditioningBuilder{_SALDataAirConditioning: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAirConditioning) GetApplicationId() ApplicationId {
	return ApplicationId_AIR_CONDITIONING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAirConditioning) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAirConditioning) GetAirConditioningData() AirConditioningData {
	return m.AirConditioningData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataAirConditioning(structType any) SALDataAirConditioning {
	if casted, ok := structType.(SALDataAirConditioning); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAirConditioning); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAirConditioning) GetTypeName() string {
	return "SALDataAirConditioning"
}

func (m *_SALDataAirConditioning) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (airConditioningData)
	lengthInBits += m.AirConditioningData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataAirConditioning) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataAirConditioning) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataAirConditioning SALDataAirConditioning, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAirConditioning"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAirConditioning")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	airConditioningData, err := ReadSimpleField[AirConditioningData](ctx, "airConditioningData", ReadComplex[AirConditioningData](AirConditioningDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'airConditioningData' field"))
	}
	m.AirConditioningData = airConditioningData

	if closeErr := readBuffer.CloseContext("SALDataAirConditioning"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAirConditioning")
	}

	return m, nil
}

func (m *_SALDataAirConditioning) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataAirConditioning) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAirConditioning"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAirConditioning")
		}

		if err := WriteSimpleField[AirConditioningData](ctx, "airConditioningData", m.GetAirConditioningData(), WriteComplex[AirConditioningData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'airConditioningData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataAirConditioning"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAirConditioning")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataAirConditioning) IsSALDataAirConditioning() {}

func (m *_SALDataAirConditioning) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataAirConditioning) deepCopy() *_SALDataAirConditioning {
	if m == nil {
		return nil
	}
	_SALDataAirConditioningCopy := &_SALDataAirConditioning{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.AirConditioningData.DeepCopy().(AirConditioningData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataAirConditioningCopy
}

func (m *_SALDataAirConditioning) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
