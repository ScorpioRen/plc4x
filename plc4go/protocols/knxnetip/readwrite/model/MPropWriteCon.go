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

// MPropWriteCon is the corresponding interface of MPropWriteCon
type MPropWriteCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMPropWriteCon is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMPropWriteCon()
	// CreateBuilder creates a MPropWriteConBuilder
	CreateMPropWriteConBuilder() MPropWriteConBuilder
}

// _MPropWriteCon is the data-structure of this message
type _MPropWriteCon struct {
	CEMIContract
}

var _ MPropWriteCon = (*_MPropWriteCon)(nil)
var _ CEMIRequirements = (*_MPropWriteCon)(nil)

// NewMPropWriteCon factory function for _MPropWriteCon
func NewMPropWriteCon(size uint16) *_MPropWriteCon {
	_result := &_MPropWriteCon{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MPropWriteConBuilder is a builder for MPropWriteCon
type MPropWriteConBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MPropWriteConBuilder
	// Build builds the MPropWriteCon or returns an error if something is wrong
	Build() (MPropWriteCon, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MPropWriteCon
}

// NewMPropWriteConBuilder() creates a MPropWriteConBuilder
func NewMPropWriteConBuilder() MPropWriteConBuilder {
	return &_MPropWriteConBuilder{_MPropWriteCon: new(_MPropWriteCon)}
}

type _MPropWriteConBuilder struct {
	*_MPropWriteCon

	err *utils.MultiError
}

var _ (MPropWriteConBuilder) = (*_MPropWriteConBuilder)(nil)

func (m *_MPropWriteConBuilder) WithMandatoryFields() MPropWriteConBuilder {
	return m
}

func (m *_MPropWriteConBuilder) Build() (MPropWriteCon, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._MPropWriteCon.deepCopy(), nil
}

func (m *_MPropWriteConBuilder) MustBuild() MPropWriteCon {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_MPropWriteConBuilder) DeepCopy() any {
	return m.CreateMPropWriteConBuilder()
}

// CreateMPropWriteConBuilder creates a MPropWriteConBuilder
func (m *_MPropWriteCon) CreateMPropWriteConBuilder() MPropWriteConBuilder {
	if m == nil {
		return NewMPropWriteConBuilder()
	}
	return &_MPropWriteConBuilder{_MPropWriteCon: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropWriteCon) GetMessageCode() uint8 {
	return 0xF5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropWriteCon) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMPropWriteCon(structType any) MPropWriteCon {
	if casted, ok := structType.(MPropWriteCon); ok {
		return casted
	}
	if casted, ok := structType.(*MPropWriteCon); ok {
		return *casted
	}
	return nil
}

func (m *_MPropWriteCon) GetTypeName() string {
	return "MPropWriteCon"
}

func (m *_MPropWriteCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MPropWriteCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropWriteCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropWriteCon MPropWriteCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropWriteCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropWriteCon")
	}

	return m, nil
}

func (m *_MPropWriteCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropWriteCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropWriteCon")
		}

		if popErr := writeBuffer.PopContext("MPropWriteCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropWriteCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropWriteCon) IsMPropWriteCon() {}

func (m *_MPropWriteCon) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MPropWriteCon) deepCopy() *_MPropWriteCon {
	if m == nil {
		return nil
	}
	_MPropWriteConCopy := &_MPropWriteCon{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MPropWriteConCopy
}

func (m *_MPropWriteCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
