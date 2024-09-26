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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataTesting is the corresponding interface of SALDataTesting
type SALDataTesting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// IsSALDataTesting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataTesting()
	// CreateBuilder creates a SALDataTestingBuilder
	CreateSALDataTestingBuilder() SALDataTestingBuilder
}

// _SALDataTesting is the data-structure of this message
type _SALDataTesting struct {
	SALDataContract
}

var _ SALDataTesting = (*_SALDataTesting)(nil)
var _ SALDataRequirements = (*_SALDataTesting)(nil)

// NewSALDataTesting factory function for _SALDataTesting
func NewSALDataTesting(salData SALData) *_SALDataTesting {
	_result := &_SALDataTesting{
		SALDataContract: NewSALData(salData),
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataTestingBuilder is a builder for SALDataTesting
type SALDataTestingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SALDataTestingBuilder
	// Build builds the SALDataTesting or returns an error if something is wrong
	Build() (SALDataTesting, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataTesting
}

// NewSALDataTestingBuilder() creates a SALDataTestingBuilder
func NewSALDataTestingBuilder() SALDataTestingBuilder {
	return &_SALDataTestingBuilder{_SALDataTesting: new(_SALDataTesting)}
}

type _SALDataTestingBuilder struct {
	*_SALDataTesting

	err *utils.MultiError
}

var _ (SALDataTestingBuilder) = (*_SALDataTestingBuilder)(nil)

func (m *_SALDataTestingBuilder) WithMandatoryFields() SALDataTestingBuilder {
	return m
}

func (m *_SALDataTestingBuilder) Build() (SALDataTesting, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SALDataTesting.deepCopy(), nil
}

func (m *_SALDataTestingBuilder) MustBuild() SALDataTesting {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SALDataTestingBuilder) DeepCopy() any {
	return m.CreateSALDataTestingBuilder()
}

// CreateSALDataTestingBuilder creates a SALDataTestingBuilder
func (m *_SALDataTesting) CreateSALDataTestingBuilder() SALDataTestingBuilder {
	if m == nil {
		return NewSALDataTestingBuilder()
	}
	return &_SALDataTestingBuilder{_SALDataTesting: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTesting) GetApplicationId() ApplicationId {
	return ApplicationId_TESTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTesting) GetParent() SALDataContract {
	return m.SALDataContract
}

// Deprecated: use the interface for direct cast
func CastSALDataTesting(structType any) SALDataTesting {
	if casted, ok := structType.(SALDataTesting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTesting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTesting) GetTypeName() string {
	return "SALDataTesting"
}

func (m *_SALDataTesting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SALDataTesting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTesting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTesting SALDataTesting, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTesting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTesting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "TESTING Not yet implemented"})
	}

	if closeErr := readBuffer.CloseContext("SALDataTesting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTesting")
	}

	return m, nil
}

func (m *_SALDataTesting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTesting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTesting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTesting")
		}

		if popErr := writeBuffer.PopContext("SALDataTesting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTesting")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTesting) IsSALDataTesting() {}

func (m *_SALDataTesting) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataTesting) deepCopy() *_SALDataTesting {
	if m == nil {
		return nil
	}
	_SALDataTestingCopy := &_SALDataTesting{
		m.SALDataContract.(*_SALData).deepCopy(),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataTestingCopy
}

func (m *_SALDataTesting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
