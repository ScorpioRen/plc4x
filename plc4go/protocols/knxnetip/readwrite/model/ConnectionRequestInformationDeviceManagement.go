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

// ConnectionRequestInformationDeviceManagement is the corresponding interface of ConnectionRequestInformationDeviceManagement
type ConnectionRequestInformationDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ConnectionRequestInformation
	// IsConnectionRequestInformationDeviceManagement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequestInformationDeviceManagement()
	// CreateBuilder creates a ConnectionRequestInformationDeviceManagementBuilder
	CreateConnectionRequestInformationDeviceManagementBuilder() ConnectionRequestInformationDeviceManagementBuilder
}

// _ConnectionRequestInformationDeviceManagement is the data-structure of this message
type _ConnectionRequestInformationDeviceManagement struct {
	ConnectionRequestInformationContract
}

var _ ConnectionRequestInformationDeviceManagement = (*_ConnectionRequestInformationDeviceManagement)(nil)
var _ ConnectionRequestInformationRequirements = (*_ConnectionRequestInformationDeviceManagement)(nil)

// NewConnectionRequestInformationDeviceManagement factory function for _ConnectionRequestInformationDeviceManagement
func NewConnectionRequestInformationDeviceManagement() *_ConnectionRequestInformationDeviceManagement {
	_result := &_ConnectionRequestInformationDeviceManagement{
		ConnectionRequestInformationContract: NewConnectionRequestInformation(),
	}
	_result.ConnectionRequestInformationContract.(*_ConnectionRequestInformation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionRequestInformationDeviceManagementBuilder is a builder for ConnectionRequestInformationDeviceManagement
type ConnectionRequestInformationDeviceManagementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ConnectionRequestInformationDeviceManagementBuilder
	// Build builds the ConnectionRequestInformationDeviceManagement or returns an error if something is wrong
	Build() (ConnectionRequestInformationDeviceManagement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionRequestInformationDeviceManagement
}

// NewConnectionRequestInformationDeviceManagementBuilder() creates a ConnectionRequestInformationDeviceManagementBuilder
func NewConnectionRequestInformationDeviceManagementBuilder() ConnectionRequestInformationDeviceManagementBuilder {
	return &_ConnectionRequestInformationDeviceManagementBuilder{_ConnectionRequestInformationDeviceManagement: new(_ConnectionRequestInformationDeviceManagement)}
}

type _ConnectionRequestInformationDeviceManagementBuilder struct {
	*_ConnectionRequestInformationDeviceManagement

	err *utils.MultiError
}

var _ (ConnectionRequestInformationDeviceManagementBuilder) = (*_ConnectionRequestInformationDeviceManagementBuilder)(nil)

func (m *_ConnectionRequestInformationDeviceManagementBuilder) WithMandatoryFields() ConnectionRequestInformationDeviceManagementBuilder {
	return m
}

func (m *_ConnectionRequestInformationDeviceManagementBuilder) Build() (ConnectionRequestInformationDeviceManagement, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ConnectionRequestInformationDeviceManagement.deepCopy(), nil
}

func (m *_ConnectionRequestInformationDeviceManagementBuilder) MustBuild() ConnectionRequestInformationDeviceManagement {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ConnectionRequestInformationDeviceManagementBuilder) DeepCopy() any {
	return m.CreateConnectionRequestInformationDeviceManagementBuilder()
}

// CreateConnectionRequestInformationDeviceManagementBuilder creates a ConnectionRequestInformationDeviceManagementBuilder
func (m *_ConnectionRequestInformationDeviceManagement) CreateConnectionRequestInformationDeviceManagementBuilder() ConnectionRequestInformationDeviceManagementBuilder {
	if m == nil {
		return NewConnectionRequestInformationDeviceManagementBuilder()
	}
	return &_ConnectionRequestInformationDeviceManagementBuilder{_ConnectionRequestInformationDeviceManagement: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequestInformationDeviceManagement) GetConnectionType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequestInformationDeviceManagement) GetParent() ConnectionRequestInformationContract {
	return m.ConnectionRequestInformationContract
}

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformationDeviceManagement(structType any) ConnectionRequestInformationDeviceManagement {
	if casted, ok := structType.(ConnectionRequestInformationDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformationDeviceManagement) GetTypeName() string {
	return "ConnectionRequestInformationDeviceManagement"
}

func (m *_ConnectionRequestInformationDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ConnectionRequestInformationDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionRequestInformationDeviceManagement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ConnectionRequestInformation) (__connectionRequestInformationDeviceManagement ConnectionRequestInformationDeviceManagement, err error) {
	m.ConnectionRequestInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationDeviceManagement")
	}

	return m, nil
}

func (m *_ConnectionRequestInformationDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequestInformationDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationDeviceManagement")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationDeviceManagement")
		}
		return nil
	}
	return m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequestInformationDeviceManagement) IsConnectionRequestInformationDeviceManagement() {
}

func (m *_ConnectionRequestInformationDeviceManagement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionRequestInformationDeviceManagement) deepCopy() *_ConnectionRequestInformationDeviceManagement {
	if m == nil {
		return nil
	}
	_ConnectionRequestInformationDeviceManagementCopy := &_ConnectionRequestInformationDeviceManagement{
		m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).deepCopy(),
	}
	m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation)._SubType = m
	return _ConnectionRequestInformationDeviceManagementCopy
}

func (m *_ConnectionRequestInformationDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
