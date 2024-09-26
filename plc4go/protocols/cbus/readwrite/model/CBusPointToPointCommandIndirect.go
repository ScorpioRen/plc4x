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

// CBusPointToPointCommandIndirect is the corresponding interface of CBusPointToPointCommandIndirect
type CBusPointToPointCommandIndirect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusPointToPointCommand
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// IsCBusPointToPointCommandIndirect is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointCommandIndirect()
	// CreateBuilder creates a CBusPointToPointCommandIndirectBuilder
	CreateCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder
}

// _CBusPointToPointCommandIndirect is the data-structure of this message
type _CBusPointToPointCommandIndirect struct {
	CBusPointToPointCommandContract
	BridgeAddress BridgeAddress
	NetworkRoute  NetworkRoute
	UnitAddress   UnitAddress
}

var _ CBusPointToPointCommandIndirect = (*_CBusPointToPointCommandIndirect)(nil)
var _ CBusPointToPointCommandRequirements = (*_CBusPointToPointCommandIndirect)(nil)

// NewCBusPointToPointCommandIndirect factory function for _CBusPointToPointCommandIndirect
func NewCBusPointToPointCommandIndirect(bridgeAddressCountPeek uint16, calData CALData, bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress, cBusOptions CBusOptions) *_CBusPointToPointCommandIndirect {
	if bridgeAddress == nil {
		panic("bridgeAddress of type BridgeAddress for CBusPointToPointCommandIndirect must not be nil")
	}
	if networkRoute == nil {
		panic("networkRoute of type NetworkRoute for CBusPointToPointCommandIndirect must not be nil")
	}
	if unitAddress == nil {
		panic("unitAddress of type UnitAddress for CBusPointToPointCommandIndirect must not be nil")
	}
	_result := &_CBusPointToPointCommandIndirect{
		CBusPointToPointCommandContract: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, cBusOptions),
		BridgeAddress:                   bridgeAddress,
		NetworkRoute:                    networkRoute,
		UnitAddress:                     unitAddress,
	}
	_result.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusPointToPointCommandIndirectBuilder is a builder for CBusPointToPointCommandIndirect
type CBusPointToPointCommandIndirectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder
	// WithBridgeAddress adds BridgeAddress (property field)
	WithBridgeAddress(BridgeAddress) CBusPointToPointCommandIndirectBuilder
	// WithBridgeAddressBuilder adds BridgeAddress (property field) which is build by the builder
	WithBridgeAddressBuilder(func(BridgeAddressBuilder) BridgeAddressBuilder) CBusPointToPointCommandIndirectBuilder
	// WithNetworkRoute adds NetworkRoute (property field)
	WithNetworkRoute(NetworkRoute) CBusPointToPointCommandIndirectBuilder
	// WithNetworkRouteBuilder adds NetworkRoute (property field) which is build by the builder
	WithNetworkRouteBuilder(func(NetworkRouteBuilder) NetworkRouteBuilder) CBusPointToPointCommandIndirectBuilder
	// WithUnitAddress adds UnitAddress (property field)
	WithUnitAddress(UnitAddress) CBusPointToPointCommandIndirectBuilder
	// WithUnitAddressBuilder adds UnitAddress (property field) which is build by the builder
	WithUnitAddressBuilder(func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandIndirectBuilder
	// Build builds the CBusPointToPointCommandIndirect or returns an error if something is wrong
	Build() (CBusPointToPointCommandIndirect, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusPointToPointCommandIndirect
}

// NewCBusPointToPointCommandIndirectBuilder() creates a CBusPointToPointCommandIndirectBuilder
func NewCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder {
	return &_CBusPointToPointCommandIndirectBuilder{_CBusPointToPointCommandIndirect: new(_CBusPointToPointCommandIndirect)}
}

type _CBusPointToPointCommandIndirectBuilder struct {
	*_CBusPointToPointCommandIndirect

	err *utils.MultiError
}

var _ (CBusPointToPointCommandIndirectBuilder) = (*_CBusPointToPointCommandIndirectBuilder)(nil)

func (m *_CBusPointToPointCommandIndirectBuilder) WithMandatoryFields(bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder {
	return m.WithBridgeAddress(bridgeAddress).WithNetworkRoute(networkRoute).WithUnitAddress(unitAddress)
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithBridgeAddress(bridgeAddress BridgeAddress) CBusPointToPointCommandIndirectBuilder {
	m.BridgeAddress = bridgeAddress
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithBridgeAddressBuilder(builderSupplier func(BridgeAddressBuilder) BridgeAddressBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(m.BridgeAddress.CreateBridgeAddressBuilder())
	var err error
	m.BridgeAddress, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BridgeAddressBuilder failed"))
	}
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithNetworkRoute(networkRoute NetworkRoute) CBusPointToPointCommandIndirectBuilder {
	m.NetworkRoute = networkRoute
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithNetworkRouteBuilder(builderSupplier func(NetworkRouteBuilder) NetworkRouteBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(m.NetworkRoute.CreateNetworkRouteBuilder())
	var err error
	m.NetworkRoute, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "NetworkRouteBuilder failed"))
	}
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithUnitAddress(unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder {
	m.UnitAddress = unitAddress
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) WithUnitAddressBuilder(builderSupplier func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(m.UnitAddress.CreateUnitAddressBuilder())
	var err error
	m.UnitAddress, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "UnitAddressBuilder failed"))
	}
	return m
}

func (m *_CBusPointToPointCommandIndirectBuilder) Build() (CBusPointToPointCommandIndirect, error) {
	if m.BridgeAddress == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'bridgeAddress' not set"))
	}
	if m.NetworkRoute == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'networkRoute' not set"))
	}
	if m.UnitAddress == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'unitAddress' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CBusPointToPointCommandIndirect.deepCopy(), nil
}

func (m *_CBusPointToPointCommandIndirectBuilder) MustBuild() CBusPointToPointCommandIndirect {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CBusPointToPointCommandIndirectBuilder) DeepCopy() any {
	return m.CreateCBusPointToPointCommandIndirectBuilder()
}

// CreateCBusPointToPointCommandIndirectBuilder creates a CBusPointToPointCommandIndirectBuilder
func (m *_CBusPointToPointCommandIndirect) CreateCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder {
	if m == nil {
		return NewCBusPointToPointCommandIndirectBuilder()
	}
	return &_CBusPointToPointCommandIndirectBuilder{_CBusPointToPointCommandIndirect: m.deepCopy()}
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

func (m *_CBusPointToPointCommandIndirect) GetParent() CBusPointToPointCommandContract {
	return m.CBusPointToPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommandIndirect) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointCommandIndirect) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointCommandIndirect) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommandIndirect(structType any) CBusPointToPointCommandIndirect {
	if casted, ok := structType.(CBusPointToPointCommandIndirect); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandIndirect); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommandIndirect) GetTypeName() string {
	return "CBusPointToPointCommandIndirect"
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).getLengthInBits(ctx))

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToPointCommandIndirect) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToPointCommand, cBusOptions CBusOptions) (__cBusPointToPointCommandIndirect CBusPointToPointCommandIndirect, err error) {
	m.CBusPointToPointCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandIndirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandIndirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bridgeAddress, err := ReadSimpleField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	m.BridgeAddress = bridgeAddress

	networkRoute, err := ReadSimpleField[NetworkRoute](ctx, "networkRoute", ReadComplex[NetworkRoute](NetworkRouteParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkRoute' field"))
	}
	m.NetworkRoute = networkRoute

	unitAddress, err := ReadSimpleField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	m.UnitAddress = unitAddress

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandIndirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandIndirect")
	}

	return m, nil
}

func (m *_CBusPointToPointCommandIndirect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointCommandIndirect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandIndirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandIndirect")
		}

		if err := WriteSimpleField[BridgeAddress](ctx, "bridgeAddress", m.GetBridgeAddress(), WriteComplex[BridgeAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
		}

		if err := WriteSimpleField[NetworkRoute](ctx, "networkRoute", m.GetNetworkRoute(), WriteComplex[NetworkRoute](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkRoute' field")
		}

		if err := WriteSimpleField[UnitAddress](ctx, "unitAddress", m.GetUnitAddress(), WriteComplex[UnitAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandIndirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandIndirect")
		}
		return nil
	}
	return m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointCommandIndirect) IsCBusPointToPointCommandIndirect() {}

func (m *_CBusPointToPointCommandIndirect) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToPointCommandIndirect) deepCopy() *_CBusPointToPointCommandIndirect {
	if m == nil {
		return nil
	}
	_CBusPointToPointCommandIndirectCopy := &_CBusPointToPointCommandIndirect{
		m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).deepCopy(),
		m.BridgeAddress.DeepCopy().(BridgeAddress),
		m.NetworkRoute.DeepCopy().(NetworkRoute),
		m.UnitAddress.DeepCopy().(UnitAddress),
	}
	m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = m
	return _CBusPointToPointCommandIndirectCopy
}

func (m *_CBusPointToPointCommandIndirect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
