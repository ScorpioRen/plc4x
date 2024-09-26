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

// SecurityDataZoneName is the corresponding interface of SecurityDataZoneName
type SecurityDataZoneName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// GetZoneName returns ZoneName (property field)
	GetZoneName() string
	// IsSecurityDataZoneName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataZoneName()
	// CreateBuilder creates a SecurityDataZoneNameBuilder
	CreateSecurityDataZoneNameBuilder() SecurityDataZoneNameBuilder
}

// _SecurityDataZoneName is the data-structure of this message
type _SecurityDataZoneName struct {
	SecurityDataContract
	ZoneNumber uint8
	ZoneName   string
}

var _ SecurityDataZoneName = (*_SecurityDataZoneName)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneName)(nil)

// NewSecurityDataZoneName factory function for _SecurityDataZoneName
func NewSecurityDataZoneName(commandTypeContainer SecurityCommandTypeContainer, argument byte, zoneNumber uint8, zoneName string) *_SecurityDataZoneName {
	_result := &_SecurityDataZoneName{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
		ZoneName:             zoneName,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataZoneNameBuilder is a builder for SecurityDataZoneName
type SecurityDataZoneNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneNumber uint8, zoneName string) SecurityDataZoneNameBuilder
	// WithZoneNumber adds ZoneNumber (property field)
	WithZoneNumber(uint8) SecurityDataZoneNameBuilder
	// WithZoneName adds ZoneName (property field)
	WithZoneName(string) SecurityDataZoneNameBuilder
	// Build builds the SecurityDataZoneName or returns an error if something is wrong
	Build() (SecurityDataZoneName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataZoneName
}

// NewSecurityDataZoneNameBuilder() creates a SecurityDataZoneNameBuilder
func NewSecurityDataZoneNameBuilder() SecurityDataZoneNameBuilder {
	return &_SecurityDataZoneNameBuilder{_SecurityDataZoneName: new(_SecurityDataZoneName)}
}

type _SecurityDataZoneNameBuilder struct {
	*_SecurityDataZoneName

	err *utils.MultiError
}

var _ (SecurityDataZoneNameBuilder) = (*_SecurityDataZoneNameBuilder)(nil)

func (m *_SecurityDataZoneNameBuilder) WithMandatoryFields(zoneNumber uint8, zoneName string) SecurityDataZoneNameBuilder {
	return m.WithZoneNumber(zoneNumber).WithZoneName(zoneName)
}

func (m *_SecurityDataZoneNameBuilder) WithZoneNumber(zoneNumber uint8) SecurityDataZoneNameBuilder {
	m.ZoneNumber = zoneNumber
	return m
}

func (m *_SecurityDataZoneNameBuilder) WithZoneName(zoneName string) SecurityDataZoneNameBuilder {
	m.ZoneName = zoneName
	return m
}

func (m *_SecurityDataZoneNameBuilder) Build() (SecurityDataZoneName, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataZoneName.deepCopy(), nil
}

func (m *_SecurityDataZoneNameBuilder) MustBuild() SecurityDataZoneName {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataZoneNameBuilder) DeepCopy() any {
	return m.CreateSecurityDataZoneNameBuilder()
}

// CreateSecurityDataZoneNameBuilder creates a SecurityDataZoneNameBuilder
func (m *_SecurityDataZoneName) CreateSecurityDataZoneNameBuilder() SecurityDataZoneNameBuilder {
	if m == nil {
		return NewSecurityDataZoneNameBuilder()
	}
	return &_SecurityDataZoneNameBuilder{_SecurityDataZoneName: m.deepCopy()}
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

func (m *_SecurityDataZoneName) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneName) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

func (m *_SecurityDataZoneName) GetZoneName() string {
	return m.ZoneName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneName(structType any) SecurityDataZoneName {
	if casted, ok := structType.(SecurityDataZoneName); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneName); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneName) GetTypeName() string {
	return "SecurityDataZoneName"
}

func (m *_SecurityDataZoneName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	// Simple field (zoneName)
	lengthInBits += 88

	return lengthInBits
}

func (m *_SecurityDataZoneName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneName SecurityDataZoneName, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	zoneName, err := ReadSimpleField(ctx, "zoneName", ReadString(readBuffer, uint32(88)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneName' field"))
	}
	m.ZoneName = zoneName

	if closeErr := readBuffer.CloseContext("SecurityDataZoneName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneName")
	}

	return m, nil
}

func (m *_SecurityDataZoneName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneName")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if err := WriteSimpleField[string](ctx, "zoneName", m.GetZoneName(), WriteString(writeBuffer, 88)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneName' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneName")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneName) IsSecurityDataZoneName() {}

func (m *_SecurityDataZoneName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataZoneName) deepCopy() *_SecurityDataZoneName {
	if m == nil {
		return nil
	}
	_SecurityDataZoneNameCopy := &_SecurityDataZoneName{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.ZoneNumber,
		m.ZoneName,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataZoneNameCopy
}

func (m *_SecurityDataZoneName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
