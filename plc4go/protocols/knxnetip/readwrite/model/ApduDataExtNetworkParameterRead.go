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

// ApduDataExtNetworkParameterRead is the corresponding interface of ApduDataExtNetworkParameterRead
type ApduDataExtNetworkParameterRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtNetworkParameterRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtNetworkParameterRead()
	// CreateBuilder creates a ApduDataExtNetworkParameterReadBuilder
	CreateApduDataExtNetworkParameterReadBuilder() ApduDataExtNetworkParameterReadBuilder
}

// _ApduDataExtNetworkParameterRead is the data-structure of this message
type _ApduDataExtNetworkParameterRead struct {
	ApduDataExtContract
}

var _ ApduDataExtNetworkParameterRead = (*_ApduDataExtNetworkParameterRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtNetworkParameterRead)(nil)

// NewApduDataExtNetworkParameterRead factory function for _ApduDataExtNetworkParameterRead
func NewApduDataExtNetworkParameterRead(length uint8) *_ApduDataExtNetworkParameterRead {
	_result := &_ApduDataExtNetworkParameterRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtNetworkParameterReadBuilder is a builder for ApduDataExtNetworkParameterRead
type ApduDataExtNetworkParameterReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtNetworkParameterReadBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtNetworkParameterRead or returns an error if something is wrong
	Build() (ApduDataExtNetworkParameterRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtNetworkParameterRead
}

// NewApduDataExtNetworkParameterReadBuilder() creates a ApduDataExtNetworkParameterReadBuilder
func NewApduDataExtNetworkParameterReadBuilder() ApduDataExtNetworkParameterReadBuilder {
	return &_ApduDataExtNetworkParameterReadBuilder{_ApduDataExtNetworkParameterRead: new(_ApduDataExtNetworkParameterRead)}
}

type _ApduDataExtNetworkParameterReadBuilder struct {
	*_ApduDataExtNetworkParameterRead

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtNetworkParameterReadBuilder) = (*_ApduDataExtNetworkParameterReadBuilder)(nil)

func (b *_ApduDataExtNetworkParameterReadBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtNetworkParameterRead
}

func (b *_ApduDataExtNetworkParameterReadBuilder) WithMandatoryFields() ApduDataExtNetworkParameterReadBuilder {
	return b
}

func (b *_ApduDataExtNetworkParameterReadBuilder) Build() (ApduDataExtNetworkParameterRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtNetworkParameterRead.deepCopy(), nil
}

func (b *_ApduDataExtNetworkParameterReadBuilder) MustBuild() ApduDataExtNetworkParameterRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtNetworkParameterReadBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtNetworkParameterReadBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtNetworkParameterReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtNetworkParameterReadBuilder().(*_ApduDataExtNetworkParameterReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtNetworkParameterReadBuilder creates a ApduDataExtNetworkParameterReadBuilder
func (b *_ApduDataExtNetworkParameterRead) CreateApduDataExtNetworkParameterReadBuilder() ApduDataExtNetworkParameterReadBuilder {
	if b == nil {
		return NewApduDataExtNetworkParameterReadBuilder()
	}
	return &_ApduDataExtNetworkParameterReadBuilder{_ApduDataExtNetworkParameterRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtNetworkParameterRead) GetExtApciType() uint8 {
	return 0x1A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtNetworkParameterRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtNetworkParameterRead(structType any) ApduDataExtNetworkParameterRead {
	if casted, ok := structType.(ApduDataExtNetworkParameterRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtNetworkParameterRead) GetTypeName() string {
	return "ApduDataExtNetworkParameterRead"
}

func (m *_ApduDataExtNetworkParameterRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtNetworkParameterRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtNetworkParameterRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtNetworkParameterRead ApduDataExtNetworkParameterRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterRead")
	}

	return m, nil
}

func (m *_ApduDataExtNetworkParameterRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtNetworkParameterRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtNetworkParameterRead) IsApduDataExtNetworkParameterRead() {}

func (m *_ApduDataExtNetworkParameterRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtNetworkParameterRead) deepCopy() *_ApduDataExtNetworkParameterRead {
	if m == nil {
		return nil
	}
	_ApduDataExtNetworkParameterReadCopy := &_ApduDataExtNetworkParameterRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtNetworkParameterReadCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtNetworkParameterReadCopy
}

func (m *_ApduDataExtNetworkParameterRead) String() string {
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
