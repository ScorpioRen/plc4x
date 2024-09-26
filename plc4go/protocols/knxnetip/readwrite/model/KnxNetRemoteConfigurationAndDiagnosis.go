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

// KnxNetRemoteConfigurationAndDiagnosis is the corresponding interface of KnxNetRemoteConfigurationAndDiagnosis
type KnxNetRemoteConfigurationAndDiagnosis interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetRemoteConfigurationAndDiagnosis is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetRemoteConfigurationAndDiagnosis()
	// CreateBuilder creates a KnxNetRemoteConfigurationAndDiagnosisBuilder
	CreateKnxNetRemoteConfigurationAndDiagnosisBuilder() KnxNetRemoteConfigurationAndDiagnosisBuilder
}

// _KnxNetRemoteConfigurationAndDiagnosis is the data-structure of this message
type _KnxNetRemoteConfigurationAndDiagnosis struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetRemoteConfigurationAndDiagnosis = (*_KnxNetRemoteConfigurationAndDiagnosis)(nil)
var _ ServiceIdRequirements = (*_KnxNetRemoteConfigurationAndDiagnosis)(nil)

// NewKnxNetRemoteConfigurationAndDiagnosis factory function for _KnxNetRemoteConfigurationAndDiagnosis
func NewKnxNetRemoteConfigurationAndDiagnosis(version uint8) *_KnxNetRemoteConfigurationAndDiagnosis {
	_result := &_KnxNetRemoteConfigurationAndDiagnosis{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// KnxNetRemoteConfigurationAndDiagnosisBuilder is a builder for KnxNetRemoteConfigurationAndDiagnosis
type KnxNetRemoteConfigurationAndDiagnosisBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint8) KnxNetRemoteConfigurationAndDiagnosisBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint8) KnxNetRemoteConfigurationAndDiagnosisBuilder
	// Build builds the KnxNetRemoteConfigurationAndDiagnosis or returns an error if something is wrong
	Build() (KnxNetRemoteConfigurationAndDiagnosis, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() KnxNetRemoteConfigurationAndDiagnosis
}

// NewKnxNetRemoteConfigurationAndDiagnosisBuilder() creates a KnxNetRemoteConfigurationAndDiagnosisBuilder
func NewKnxNetRemoteConfigurationAndDiagnosisBuilder() KnxNetRemoteConfigurationAndDiagnosisBuilder {
	return &_KnxNetRemoteConfigurationAndDiagnosisBuilder{_KnxNetRemoteConfigurationAndDiagnosis: new(_KnxNetRemoteConfigurationAndDiagnosis)}
}

type _KnxNetRemoteConfigurationAndDiagnosisBuilder struct {
	*_KnxNetRemoteConfigurationAndDiagnosis

	err *utils.MultiError
}

var _ (KnxNetRemoteConfigurationAndDiagnosisBuilder) = (*_KnxNetRemoteConfigurationAndDiagnosisBuilder)(nil)

func (m *_KnxNetRemoteConfigurationAndDiagnosisBuilder) WithMandatoryFields(version uint8) KnxNetRemoteConfigurationAndDiagnosisBuilder {
	return m.WithVersion(version)
}

func (m *_KnxNetRemoteConfigurationAndDiagnosisBuilder) WithVersion(version uint8) KnxNetRemoteConfigurationAndDiagnosisBuilder {
	m.Version = version
	return m
}

func (m *_KnxNetRemoteConfigurationAndDiagnosisBuilder) Build() (KnxNetRemoteConfigurationAndDiagnosis, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._KnxNetRemoteConfigurationAndDiagnosis.deepCopy(), nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosisBuilder) MustBuild() KnxNetRemoteConfigurationAndDiagnosis {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_KnxNetRemoteConfigurationAndDiagnosisBuilder) DeepCopy() any {
	return m.CreateKnxNetRemoteConfigurationAndDiagnosisBuilder()
}

// CreateKnxNetRemoteConfigurationAndDiagnosisBuilder creates a KnxNetRemoteConfigurationAndDiagnosisBuilder
func (m *_KnxNetRemoteConfigurationAndDiagnosis) CreateKnxNetRemoteConfigurationAndDiagnosisBuilder() KnxNetRemoteConfigurationAndDiagnosisBuilder {
	if m == nil {
		return NewKnxNetRemoteConfigurationAndDiagnosisBuilder()
	}
	return &_KnxNetRemoteConfigurationAndDiagnosisBuilder{_KnxNetRemoteConfigurationAndDiagnosis: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetServiceType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetRemoteConfigurationAndDiagnosis(structType any) KnxNetRemoteConfigurationAndDiagnosis {
	if casted, ok := structType.(KnxNetRemoteConfigurationAndDiagnosis); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetRemoteConfigurationAndDiagnosis); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetTypeName() string {
	return "KnxNetRemoteConfigurationAndDiagnosis"
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetRemoteConfigurationAndDiagnosis KnxNetRemoteConfigurationAndDiagnosis, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetRemoteConfigurationAndDiagnosis"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetRemoteConfigurationAndDiagnosis")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteConfigurationAndDiagnosis"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetRemoteConfigurationAndDiagnosis")
	}

	return m, nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteConfigurationAndDiagnosis"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetRemoteConfigurationAndDiagnosis")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteConfigurationAndDiagnosis"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetRemoteConfigurationAndDiagnosis")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) IsKnxNetRemoteConfigurationAndDiagnosis() {}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) deepCopy() *_KnxNetRemoteConfigurationAndDiagnosis {
	if m == nil {
		return nil
	}
	_KnxNetRemoteConfigurationAndDiagnosisCopy := &_KnxNetRemoteConfigurationAndDiagnosis{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	m.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetRemoteConfigurationAndDiagnosisCopy
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
