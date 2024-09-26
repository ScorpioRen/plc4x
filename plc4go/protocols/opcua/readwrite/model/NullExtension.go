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

// NullExtension is the corresponding interface of NullExtension
type NullExtension interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsNullExtension is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNullExtension()
	// CreateBuilder creates a NullExtensionBuilder
	CreateNullExtensionBuilder() NullExtensionBuilder
}

// _NullExtension is the data-structure of this message
type _NullExtension struct {
	ExtensionObjectDefinitionContract
}

var _ NullExtension = (*_NullExtension)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NullExtension)(nil)

// NewNullExtension factory function for _NullExtension
func NewNullExtension() *_NullExtension {
	_result := &_NullExtension{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NullExtensionBuilder is a builder for NullExtension
type NullExtensionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NullExtensionBuilder
	// Build builds the NullExtension or returns an error if something is wrong
	Build() (NullExtension, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NullExtension
}

// NewNullExtensionBuilder() creates a NullExtensionBuilder
func NewNullExtensionBuilder() NullExtensionBuilder {
	return &_NullExtensionBuilder{_NullExtension: new(_NullExtension)}
}

type _NullExtensionBuilder struct {
	*_NullExtension

	err *utils.MultiError
}

var _ (NullExtensionBuilder) = (*_NullExtensionBuilder)(nil)

func (m *_NullExtensionBuilder) WithMandatoryFields() NullExtensionBuilder {
	return m
}

func (m *_NullExtensionBuilder) Build() (NullExtension, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._NullExtension.deepCopy(), nil
}

func (m *_NullExtensionBuilder) MustBuild() NullExtension {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_NullExtensionBuilder) DeepCopy() any {
	return m.CreateNullExtensionBuilder()
}

// CreateNullExtensionBuilder creates a NullExtensionBuilder
func (m *_NullExtension) CreateNullExtensionBuilder() NullExtensionBuilder {
	if m == nil {
		return NewNullExtensionBuilder()
	}
	return &_NullExtensionBuilder{_NullExtension: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullExtension) GetIdentifier() string {
	return "0"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullExtension) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastNullExtension(structType any) NullExtension {
	if casted, ok := structType.(NullExtension); ok {
		return casted
	}
	if casted, ok := structType.(*NullExtension); ok {
		return *casted
	}
	return nil
}

func (m *_NullExtension) GetTypeName() string {
	return "NullExtension"
}

func (m *_NullExtension) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullExtension) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullExtension) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__nullExtension NullExtension, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullExtension"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullExtension")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullExtension"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullExtension")
	}

	return m, nil
}

func (m *_NullExtension) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullExtension) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullExtension"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullExtension")
		}

		if popErr := writeBuffer.PopContext("NullExtension"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullExtension")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullExtension) IsNullExtension() {}

func (m *_NullExtension) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NullExtension) deepCopy() *_NullExtension {
	if m == nil {
		return nil
	}
	_NullExtensionCopy := &_NullExtension{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NullExtensionCopy
}

func (m *_NullExtension) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
