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

// AirConditioningDataSetZoneGroupOn is the corresponding interface of AirConditioningDataSetZoneGroupOn
type AirConditioningDataSetZoneGroupOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// IsAirConditioningDataSetZoneGroupOn is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetZoneGroupOn()
	// CreateBuilder creates a AirConditioningDataSetZoneGroupOnBuilder
	CreateAirConditioningDataSetZoneGroupOnBuilder() AirConditioningDataSetZoneGroupOnBuilder
}

// _AirConditioningDataSetZoneGroupOn is the data-structure of this message
type _AirConditioningDataSetZoneGroupOn struct {
	AirConditioningDataContract
	ZoneGroup byte
}

var _ AirConditioningDataSetZoneGroupOn = (*_AirConditioningDataSetZoneGroupOn)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetZoneGroupOn)(nil)

// NewAirConditioningDataSetZoneGroupOn factory function for _AirConditioningDataSetZoneGroupOn
func NewAirConditioningDataSetZoneGroupOn(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte) *_AirConditioningDataSetZoneGroupOn {
	_result := &_AirConditioningDataSetZoneGroupOn{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetZoneGroupOnBuilder is a builder for AirConditioningDataSetZoneGroupOn
type AirConditioningDataSetZoneGroupOnBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte) AirConditioningDataSetZoneGroupOnBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetZoneGroupOnBuilder
	// Build builds the AirConditioningDataSetZoneGroupOn or returns an error if something is wrong
	Build() (AirConditioningDataSetZoneGroupOn, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetZoneGroupOn
}

// NewAirConditioningDataSetZoneGroupOnBuilder() creates a AirConditioningDataSetZoneGroupOnBuilder
func NewAirConditioningDataSetZoneGroupOnBuilder() AirConditioningDataSetZoneGroupOnBuilder {
	return &_AirConditioningDataSetZoneGroupOnBuilder{_AirConditioningDataSetZoneGroupOn: new(_AirConditioningDataSetZoneGroupOn)}
}

type _AirConditioningDataSetZoneGroupOnBuilder struct {
	*_AirConditioningDataSetZoneGroupOn

	err *utils.MultiError
}

var _ (AirConditioningDataSetZoneGroupOnBuilder) = (*_AirConditioningDataSetZoneGroupOnBuilder)(nil)

func (m *_AirConditioningDataSetZoneGroupOnBuilder) WithMandatoryFields(zoneGroup byte) AirConditioningDataSetZoneGroupOnBuilder {
	return m.WithZoneGroup(zoneGroup)
}

func (m *_AirConditioningDataSetZoneGroupOnBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetZoneGroupOnBuilder {
	m.ZoneGroup = zoneGroup
	return m
}

func (m *_AirConditioningDataSetZoneGroupOnBuilder) Build() (AirConditioningDataSetZoneGroupOn, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AirConditioningDataSetZoneGroupOn.deepCopy(), nil
}

func (m *_AirConditioningDataSetZoneGroupOnBuilder) MustBuild() AirConditioningDataSetZoneGroupOn {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AirConditioningDataSetZoneGroupOnBuilder) DeepCopy() any {
	return m.CreateAirConditioningDataSetZoneGroupOnBuilder()
}

// CreateAirConditioningDataSetZoneGroupOnBuilder creates a AirConditioningDataSetZoneGroupOnBuilder
func (m *_AirConditioningDataSetZoneGroupOn) CreateAirConditioningDataSetZoneGroupOnBuilder() AirConditioningDataSetZoneGroupOnBuilder {
	if m == nil {
		return NewAirConditioningDataSetZoneGroupOnBuilder()
	}
	return &_AirConditioningDataSetZoneGroupOnBuilder{_AirConditioningDataSetZoneGroupOn: m.deepCopy()}
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

func (m *_AirConditioningDataSetZoneGroupOn) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetZoneGroupOn) GetZoneGroup() byte {
	return m.ZoneGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetZoneGroupOn(structType any) AirConditioningDataSetZoneGroupOn {
	if casted, ok := structType.(AirConditioningDataSetZoneGroupOn); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetZoneGroupOn); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetZoneGroupOn) GetTypeName() string {
	return "AirConditioningDataSetZoneGroupOn"
}

func (m *_AirConditioningDataSetZoneGroupOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataSetZoneGroupOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetZoneGroupOn) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetZoneGroupOn AirConditioningDataSetZoneGroupOn, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetZoneGroupOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetZoneGroupOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetZoneGroupOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetZoneGroupOn")
	}

	return m, nil
}

func (m *_AirConditioningDataSetZoneGroupOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetZoneGroupOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetZoneGroupOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetZoneGroupOn")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetZoneGroupOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetZoneGroupOn")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetZoneGroupOn) IsAirConditioningDataSetZoneGroupOn() {}

func (m *_AirConditioningDataSetZoneGroupOn) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetZoneGroupOn) deepCopy() *_AirConditioningDataSetZoneGroupOn {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetZoneGroupOnCopy := &_AirConditioningDataSetZoneGroupOn{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetZoneGroupOnCopy
}

func (m *_AirConditioningDataSetZoneGroupOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
