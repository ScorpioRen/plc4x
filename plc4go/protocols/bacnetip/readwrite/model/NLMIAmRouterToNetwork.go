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

// NLMIAmRouterToNetwork is the corresponding interface of NLMIAmRouterToNetwork
type NLMIAmRouterToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetDestinationNetworkAddresses returns DestinationNetworkAddresses (property field)
	GetDestinationNetworkAddresses() []uint16
	// IsNLMIAmRouterToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMIAmRouterToNetwork()
	// CreateBuilder creates a NLMIAmRouterToNetworkBuilder
	CreateNLMIAmRouterToNetworkBuilder() NLMIAmRouterToNetworkBuilder
}

// _NLMIAmRouterToNetwork is the data-structure of this message
type _NLMIAmRouterToNetwork struct {
	NLMContract
	DestinationNetworkAddresses []uint16
}

var _ NLMIAmRouterToNetwork = (*_NLMIAmRouterToNetwork)(nil)
var _ NLMRequirements = (*_NLMIAmRouterToNetwork)(nil)

// NewNLMIAmRouterToNetwork factory function for _NLMIAmRouterToNetwork
func NewNLMIAmRouterToNetwork(destinationNetworkAddresses []uint16, apduLength uint16) *_NLMIAmRouterToNetwork {
	_result := &_NLMIAmRouterToNetwork{
		NLMContract:                 NewNLM(apduLength),
		DestinationNetworkAddresses: destinationNetworkAddresses,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMIAmRouterToNetworkBuilder is a builder for NLMIAmRouterToNetwork
type NLMIAmRouterToNetworkBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationNetworkAddresses []uint16) NLMIAmRouterToNetworkBuilder
	// WithDestinationNetworkAddresses adds DestinationNetworkAddresses (property field)
	WithDestinationNetworkAddresses(...uint16) NLMIAmRouterToNetworkBuilder
	// Build builds the NLMIAmRouterToNetwork or returns an error if something is wrong
	Build() (NLMIAmRouterToNetwork, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMIAmRouterToNetwork
}

// NewNLMIAmRouterToNetworkBuilder() creates a NLMIAmRouterToNetworkBuilder
func NewNLMIAmRouterToNetworkBuilder() NLMIAmRouterToNetworkBuilder {
	return &_NLMIAmRouterToNetworkBuilder{_NLMIAmRouterToNetwork: new(_NLMIAmRouterToNetwork)}
}

type _NLMIAmRouterToNetworkBuilder struct {
	*_NLMIAmRouterToNetwork

	err *utils.MultiError
}

var _ (NLMIAmRouterToNetworkBuilder) = (*_NLMIAmRouterToNetworkBuilder)(nil)

func (m *_NLMIAmRouterToNetworkBuilder) WithMandatoryFields(destinationNetworkAddresses []uint16) NLMIAmRouterToNetworkBuilder {
	return m.WithDestinationNetworkAddresses(destinationNetworkAddresses...)
}

func (m *_NLMIAmRouterToNetworkBuilder) WithDestinationNetworkAddresses(destinationNetworkAddresses ...uint16) NLMIAmRouterToNetworkBuilder {
	m.DestinationNetworkAddresses = destinationNetworkAddresses
	return m
}

func (m *_NLMIAmRouterToNetworkBuilder) Build() (NLMIAmRouterToNetwork, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._NLMIAmRouterToNetwork.deepCopy(), nil
}

func (m *_NLMIAmRouterToNetworkBuilder) MustBuild() NLMIAmRouterToNetwork {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_NLMIAmRouterToNetworkBuilder) DeepCopy() any {
	return m.CreateNLMIAmRouterToNetworkBuilder()
}

// CreateNLMIAmRouterToNetworkBuilder creates a NLMIAmRouterToNetworkBuilder
func (m *_NLMIAmRouterToNetwork) CreateNLMIAmRouterToNetworkBuilder() NLMIAmRouterToNetworkBuilder {
	if m == nil {
		return NewNLMIAmRouterToNetworkBuilder()
	}
	return &_NLMIAmRouterToNetworkBuilder{_NLMIAmRouterToNetwork: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMIAmRouterToNetwork) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMIAmRouterToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMIAmRouterToNetwork) GetDestinationNetworkAddresses() []uint16 {
	return m.DestinationNetworkAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMIAmRouterToNetwork(structType any) NLMIAmRouterToNetwork {
	if casted, ok := structType.(NLMIAmRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMIAmRouterToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMIAmRouterToNetwork) GetTypeName() string {
	return "NLMIAmRouterToNetwork"
}

func (m *_NLMIAmRouterToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Array field
	if len(m.DestinationNetworkAddresses) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddresses))
	}

	return lengthInBits
}

func (m *_NLMIAmRouterToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMIAmRouterToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMIAmRouterToNetwork NLMIAmRouterToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMIAmRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMIAmRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationNetworkAddresses, err := ReadLengthArrayField[uint16](ctx, "destinationNetworkAddresses", ReadUnsignedShort(readBuffer, uint8(16)), int(int32(apduLength)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddresses' field"))
	}
	m.DestinationNetworkAddresses = destinationNetworkAddresses

	if closeErr := readBuffer.CloseContext("NLMIAmRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMIAmRouterToNetwork")
	}

	return m, nil
}

func (m *_NLMIAmRouterToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMIAmRouterToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMIAmRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMIAmRouterToNetwork")
		}

		if err := WriteSimpleTypeArrayField(ctx, "destinationNetworkAddresses", m.GetDestinationNetworkAddresses(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddresses' field")
		}

		if popErr := writeBuffer.PopContext("NLMIAmRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMIAmRouterToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMIAmRouterToNetwork) IsNLMIAmRouterToNetwork() {}

func (m *_NLMIAmRouterToNetwork) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMIAmRouterToNetwork) deepCopy() *_NLMIAmRouterToNetwork {
	if m == nil {
		return nil
	}
	_NLMIAmRouterToNetworkCopy := &_NLMIAmRouterToNetwork{
		m.NLMContract.(*_NLM).deepCopy(),
		utils.DeepCopySlice[uint16, uint16](m.DestinationNetworkAddresses),
	}
	m.NLMContract.(*_NLM)._SubType = m
	return _NLMIAmRouterToNetworkCopy
}

func (m *_NLMIAmRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
