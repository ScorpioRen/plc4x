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

// SecurityDataMainsFailure is the corresponding interface of SecurityDataMainsFailure
type SecurityDataMainsFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataMainsFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataMainsFailure()
	// CreateBuilder creates a SecurityDataMainsFailureBuilder
	CreateSecurityDataMainsFailureBuilder() SecurityDataMainsFailureBuilder
}

// _SecurityDataMainsFailure is the data-structure of this message
type _SecurityDataMainsFailure struct {
	SecurityDataContract
}

var _ SecurityDataMainsFailure = (*_SecurityDataMainsFailure)(nil)
var _ SecurityDataRequirements = (*_SecurityDataMainsFailure)(nil)

// NewSecurityDataMainsFailure factory function for _SecurityDataMainsFailure
func NewSecurityDataMainsFailure(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataMainsFailure {
	_result := &_SecurityDataMainsFailure{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataMainsFailureBuilder is a builder for SecurityDataMainsFailure
type SecurityDataMainsFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataMainsFailureBuilder
	// Build builds the SecurityDataMainsFailure or returns an error if something is wrong
	Build() (SecurityDataMainsFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataMainsFailure
}

// NewSecurityDataMainsFailureBuilder() creates a SecurityDataMainsFailureBuilder
func NewSecurityDataMainsFailureBuilder() SecurityDataMainsFailureBuilder {
	return &_SecurityDataMainsFailureBuilder{_SecurityDataMainsFailure: new(_SecurityDataMainsFailure)}
}

type _SecurityDataMainsFailureBuilder struct {
	*_SecurityDataMainsFailure

	err *utils.MultiError
}

var _ (SecurityDataMainsFailureBuilder) = (*_SecurityDataMainsFailureBuilder)(nil)

func (m *_SecurityDataMainsFailureBuilder) WithMandatoryFields() SecurityDataMainsFailureBuilder {
	return m
}

func (m *_SecurityDataMainsFailureBuilder) Build() (SecurityDataMainsFailure, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataMainsFailure.deepCopy(), nil
}

func (m *_SecurityDataMainsFailureBuilder) MustBuild() SecurityDataMainsFailure {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataMainsFailureBuilder) DeepCopy() any {
	return m.CreateSecurityDataMainsFailureBuilder()
}

// CreateSecurityDataMainsFailureBuilder creates a SecurityDataMainsFailureBuilder
func (m *_SecurityDataMainsFailure) CreateSecurityDataMainsFailureBuilder() SecurityDataMainsFailureBuilder {
	if m == nil {
		return NewSecurityDataMainsFailureBuilder()
	}
	return &_SecurityDataMainsFailureBuilder{_SecurityDataMainsFailure: m.deepCopy()}
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

func (m *_SecurityDataMainsFailure) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataMainsFailure(structType any) SecurityDataMainsFailure {
	if casted, ok := structType.(SecurityDataMainsFailure); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataMainsFailure); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataMainsFailure) GetTypeName() string {
	return "SecurityDataMainsFailure"
}

func (m *_SecurityDataMainsFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataMainsFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataMainsFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataMainsFailure SecurityDataMainsFailure, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataMainsFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataMainsFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataMainsFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataMainsFailure")
	}

	return m, nil
}

func (m *_SecurityDataMainsFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataMainsFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataMainsFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataMainsFailure")
		}

		if popErr := writeBuffer.PopContext("SecurityDataMainsFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataMainsFailure")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataMainsFailure) IsSecurityDataMainsFailure() {}

func (m *_SecurityDataMainsFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataMainsFailure) deepCopy() *_SecurityDataMainsFailure {
	if m == nil {
		return nil
	}
	_SecurityDataMainsFailureCopy := &_SecurityDataMainsFailure{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataMainsFailureCopy
}

func (m *_SecurityDataMainsFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
