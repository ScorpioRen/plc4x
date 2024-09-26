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

// SecurityDataMainsRestoredOrApplied is the corresponding interface of SecurityDataMainsRestoredOrApplied
type SecurityDataMainsRestoredOrApplied interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataMainsRestoredOrApplied is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataMainsRestoredOrApplied()
	// CreateBuilder creates a SecurityDataMainsRestoredOrAppliedBuilder
	CreateSecurityDataMainsRestoredOrAppliedBuilder() SecurityDataMainsRestoredOrAppliedBuilder
}

// _SecurityDataMainsRestoredOrApplied is the data-structure of this message
type _SecurityDataMainsRestoredOrApplied struct {
	SecurityDataContract
}

var _ SecurityDataMainsRestoredOrApplied = (*_SecurityDataMainsRestoredOrApplied)(nil)
var _ SecurityDataRequirements = (*_SecurityDataMainsRestoredOrApplied)(nil)

// NewSecurityDataMainsRestoredOrApplied factory function for _SecurityDataMainsRestoredOrApplied
func NewSecurityDataMainsRestoredOrApplied(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataMainsRestoredOrApplied {
	_result := &_SecurityDataMainsRestoredOrApplied{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataMainsRestoredOrAppliedBuilder is a builder for SecurityDataMainsRestoredOrApplied
type SecurityDataMainsRestoredOrAppliedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataMainsRestoredOrAppliedBuilder
	// Build builds the SecurityDataMainsRestoredOrApplied or returns an error if something is wrong
	Build() (SecurityDataMainsRestoredOrApplied, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataMainsRestoredOrApplied
}

// NewSecurityDataMainsRestoredOrAppliedBuilder() creates a SecurityDataMainsRestoredOrAppliedBuilder
func NewSecurityDataMainsRestoredOrAppliedBuilder() SecurityDataMainsRestoredOrAppliedBuilder {
	return &_SecurityDataMainsRestoredOrAppliedBuilder{_SecurityDataMainsRestoredOrApplied: new(_SecurityDataMainsRestoredOrApplied)}
}

type _SecurityDataMainsRestoredOrAppliedBuilder struct {
	*_SecurityDataMainsRestoredOrApplied

	err *utils.MultiError
}

var _ (SecurityDataMainsRestoredOrAppliedBuilder) = (*_SecurityDataMainsRestoredOrAppliedBuilder)(nil)

func (m *_SecurityDataMainsRestoredOrAppliedBuilder) WithMandatoryFields() SecurityDataMainsRestoredOrAppliedBuilder {
	return m
}

func (m *_SecurityDataMainsRestoredOrAppliedBuilder) Build() (SecurityDataMainsRestoredOrApplied, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataMainsRestoredOrApplied.deepCopy(), nil
}

func (m *_SecurityDataMainsRestoredOrAppliedBuilder) MustBuild() SecurityDataMainsRestoredOrApplied {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataMainsRestoredOrAppliedBuilder) DeepCopy() any {
	return m.CreateSecurityDataMainsRestoredOrAppliedBuilder()
}

// CreateSecurityDataMainsRestoredOrAppliedBuilder creates a SecurityDataMainsRestoredOrAppliedBuilder
func (m *_SecurityDataMainsRestoredOrApplied) CreateSecurityDataMainsRestoredOrAppliedBuilder() SecurityDataMainsRestoredOrAppliedBuilder {
	if m == nil {
		return NewSecurityDataMainsRestoredOrAppliedBuilder()
	}
	return &_SecurityDataMainsRestoredOrAppliedBuilder{_SecurityDataMainsRestoredOrApplied: m.deepCopy()}
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

func (m *_SecurityDataMainsRestoredOrApplied) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataMainsRestoredOrApplied(structType any) SecurityDataMainsRestoredOrApplied {
	if casted, ok := structType.(SecurityDataMainsRestoredOrApplied); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataMainsRestoredOrApplied); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataMainsRestoredOrApplied) GetTypeName() string {
	return "SecurityDataMainsRestoredOrApplied"
}

func (m *_SecurityDataMainsRestoredOrApplied) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataMainsRestoredOrApplied) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataMainsRestoredOrApplied) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataMainsRestoredOrApplied SecurityDataMainsRestoredOrApplied, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataMainsRestoredOrApplied"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataMainsRestoredOrApplied")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataMainsRestoredOrApplied"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataMainsRestoredOrApplied")
	}

	return m, nil
}

func (m *_SecurityDataMainsRestoredOrApplied) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataMainsRestoredOrApplied) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataMainsRestoredOrApplied"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataMainsRestoredOrApplied")
		}

		if popErr := writeBuffer.PopContext("SecurityDataMainsRestoredOrApplied"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataMainsRestoredOrApplied")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataMainsRestoredOrApplied) IsSecurityDataMainsRestoredOrApplied() {}

func (m *_SecurityDataMainsRestoredOrApplied) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataMainsRestoredOrApplied) deepCopy() *_SecurityDataMainsRestoredOrApplied {
	if m == nil {
		return nil
	}
	_SecurityDataMainsRestoredOrAppliedCopy := &_SecurityDataMainsRestoredOrApplied{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataMainsRestoredOrAppliedCopy
}

func (m *_SecurityDataMainsRestoredOrApplied) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
