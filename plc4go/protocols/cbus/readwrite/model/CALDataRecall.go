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

// CALDataRecall is the corresponding interface of CALDataRecall
type CALDataRecall interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCount returns Count (property field)
	GetCount() uint8
	// IsCALDataRecall is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALDataRecall()
	// CreateBuilder creates a CALDataRecallBuilder
	CreateCALDataRecallBuilder() CALDataRecallBuilder
}

// _CALDataRecall is the data-structure of this message
type _CALDataRecall struct {
	CALDataContract
	ParamNo Parameter
	Count   uint8
}

var _ CALDataRecall = (*_CALDataRecall)(nil)
var _ CALDataRequirements = (*_CALDataRecall)(nil)

// NewCALDataRecall factory function for _CALDataRecall
func NewCALDataRecall(commandTypeContainer CALCommandTypeContainer, additionalData CALData, paramNo Parameter, count uint8, requestContext RequestContext) *_CALDataRecall {
	_result := &_CALDataRecall{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		Count:           count,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CALDataRecallBuilder is a builder for CALDataRecall
type CALDataRecallBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(paramNo Parameter, count uint8) CALDataRecallBuilder
	// WithParamNo adds ParamNo (property field)
	WithParamNo(Parameter) CALDataRecallBuilder
	// WithCount adds Count (property field)
	WithCount(uint8) CALDataRecallBuilder
	// Build builds the CALDataRecall or returns an error if something is wrong
	Build() (CALDataRecall, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CALDataRecall
}

// NewCALDataRecallBuilder() creates a CALDataRecallBuilder
func NewCALDataRecallBuilder() CALDataRecallBuilder {
	return &_CALDataRecallBuilder{_CALDataRecall: new(_CALDataRecall)}
}

type _CALDataRecallBuilder struct {
	*_CALDataRecall

	err *utils.MultiError
}

var _ (CALDataRecallBuilder) = (*_CALDataRecallBuilder)(nil)

func (m *_CALDataRecallBuilder) WithMandatoryFields(paramNo Parameter, count uint8) CALDataRecallBuilder {
	return m.WithParamNo(paramNo).WithCount(count)
}

func (m *_CALDataRecallBuilder) WithParamNo(paramNo Parameter) CALDataRecallBuilder {
	m.ParamNo = paramNo
	return m
}

func (m *_CALDataRecallBuilder) WithCount(count uint8) CALDataRecallBuilder {
	m.Count = count
	return m
}

func (m *_CALDataRecallBuilder) Build() (CALDataRecall, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CALDataRecall.deepCopy(), nil
}

func (m *_CALDataRecallBuilder) MustBuild() CALDataRecall {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CALDataRecallBuilder) DeepCopy() any {
	return m.CreateCALDataRecallBuilder()
}

// CreateCALDataRecallBuilder creates a CALDataRecallBuilder
func (m *_CALDataRecall) CreateCALDataRecallBuilder() CALDataRecallBuilder {
	if m == nil {
		return NewCALDataRecallBuilder()
	}
	return &_CALDataRecallBuilder{_CALDataRecall: m.deepCopy()}
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

func (m *_CALDataRecall) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataRecall) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataRecall) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCALDataRecall(structType any) CALDataRecall {
	if casted, ok := structType.(CALDataRecall); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataRecall); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataRecall) GetTypeName() string {
	return "CALDataRecall"
}

func (m *_CALDataRecall) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataRecall) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataRecall) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, requestContext RequestContext) (__cALDataRecall CALDataRecall, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRecall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataRecall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	count, err := ReadSimpleField(ctx, "count", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	if closeErr := readBuffer.CloseContext("CALDataRecall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataRecall")
	}

	return m, nil
}

func (m *_CALDataRecall) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataRecall) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRecall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataRecall")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[uint8](ctx, "count", m.GetCount(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRecall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataRecall")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataRecall) IsCALDataRecall() {}

func (m *_CALDataRecall) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALDataRecall) deepCopy() *_CALDataRecall {
	if m == nil {
		return nil
	}
	_CALDataRecallCopy := &_CALDataRecall{
		m.CALDataContract.(*_CALData).deepCopy(),
		m.ParamNo,
		m.Count,
	}
	m.CALDataContract.(*_CALData)._SubType = m
	return _CALDataRecallCopy
}

func (m *_CALDataRecall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
