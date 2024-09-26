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

// SecurityDataExitDelayStarted is the corresponding interface of SecurityDataExitDelayStarted
type SecurityDataExitDelayStarted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataExitDelayStarted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataExitDelayStarted()
	// CreateBuilder creates a SecurityDataExitDelayStartedBuilder
	CreateSecurityDataExitDelayStartedBuilder() SecurityDataExitDelayStartedBuilder
}

// _SecurityDataExitDelayStarted is the data-structure of this message
type _SecurityDataExitDelayStarted struct {
	SecurityDataContract
}

var _ SecurityDataExitDelayStarted = (*_SecurityDataExitDelayStarted)(nil)
var _ SecurityDataRequirements = (*_SecurityDataExitDelayStarted)(nil)

// NewSecurityDataExitDelayStarted factory function for _SecurityDataExitDelayStarted
func NewSecurityDataExitDelayStarted(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataExitDelayStarted {
	_result := &_SecurityDataExitDelayStarted{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataExitDelayStartedBuilder is a builder for SecurityDataExitDelayStarted
type SecurityDataExitDelayStartedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataExitDelayStartedBuilder
	// Build builds the SecurityDataExitDelayStarted or returns an error if something is wrong
	Build() (SecurityDataExitDelayStarted, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataExitDelayStarted
}

// NewSecurityDataExitDelayStartedBuilder() creates a SecurityDataExitDelayStartedBuilder
func NewSecurityDataExitDelayStartedBuilder() SecurityDataExitDelayStartedBuilder {
	return &_SecurityDataExitDelayStartedBuilder{_SecurityDataExitDelayStarted: new(_SecurityDataExitDelayStarted)}
}

type _SecurityDataExitDelayStartedBuilder struct {
	*_SecurityDataExitDelayStarted

	err *utils.MultiError
}

var _ (SecurityDataExitDelayStartedBuilder) = (*_SecurityDataExitDelayStartedBuilder)(nil)

func (m *_SecurityDataExitDelayStartedBuilder) WithMandatoryFields() SecurityDataExitDelayStartedBuilder {
	return m
}

func (m *_SecurityDataExitDelayStartedBuilder) Build() (SecurityDataExitDelayStarted, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataExitDelayStarted.deepCopy(), nil
}

func (m *_SecurityDataExitDelayStartedBuilder) MustBuild() SecurityDataExitDelayStarted {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataExitDelayStartedBuilder) DeepCopy() any {
	return m.CreateSecurityDataExitDelayStartedBuilder()
}

// CreateSecurityDataExitDelayStartedBuilder creates a SecurityDataExitDelayStartedBuilder
func (m *_SecurityDataExitDelayStarted) CreateSecurityDataExitDelayStartedBuilder() SecurityDataExitDelayStartedBuilder {
	if m == nil {
		return NewSecurityDataExitDelayStartedBuilder()
	}
	return &_SecurityDataExitDelayStartedBuilder{_SecurityDataExitDelayStarted: m.deepCopy()}
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

func (m *_SecurityDataExitDelayStarted) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataExitDelayStarted(structType any) SecurityDataExitDelayStarted {
	if casted, ok := structType.(SecurityDataExitDelayStarted); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataExitDelayStarted); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataExitDelayStarted) GetTypeName() string {
	return "SecurityDataExitDelayStarted"
}

func (m *_SecurityDataExitDelayStarted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataExitDelayStarted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataExitDelayStarted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataExitDelayStarted SecurityDataExitDelayStarted, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataExitDelayStarted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataExitDelayStarted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataExitDelayStarted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataExitDelayStarted")
	}

	return m, nil
}

func (m *_SecurityDataExitDelayStarted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataExitDelayStarted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataExitDelayStarted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataExitDelayStarted")
		}

		if popErr := writeBuffer.PopContext("SecurityDataExitDelayStarted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataExitDelayStarted")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataExitDelayStarted) IsSecurityDataExitDelayStarted() {}

func (m *_SecurityDataExitDelayStarted) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataExitDelayStarted) deepCopy() *_SecurityDataExitDelayStarted {
	if m == nil {
		return nil
	}
	_SecurityDataExitDelayStartedCopy := &_SecurityDataExitDelayStarted{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataExitDelayStartedCopy
}

func (m *_SecurityDataExitDelayStarted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
