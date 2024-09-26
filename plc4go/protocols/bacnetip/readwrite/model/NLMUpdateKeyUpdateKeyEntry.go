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

// NLMUpdateKeyUpdateKeyEntry is the corresponding interface of NLMUpdateKeyUpdateKeyEntry
type NLMUpdateKeyUpdateKeyEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetKeyIdentifier returns KeyIdentifier (property field)
	GetKeyIdentifier() uint16
	// GetKeySize returns KeySize (property field)
	GetKeySize() uint8
	// GetKey returns Key (property field)
	GetKey() []byte
	// IsNLMUpdateKeyUpdateKeyEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMUpdateKeyUpdateKeyEntry()
	// CreateBuilder creates a NLMUpdateKeyUpdateKeyEntryBuilder
	CreateNLMUpdateKeyUpdateKeyEntryBuilder() NLMUpdateKeyUpdateKeyEntryBuilder
}

// _NLMUpdateKeyUpdateKeyEntry is the data-structure of this message
type _NLMUpdateKeyUpdateKeyEntry struct {
	KeyIdentifier uint16
	KeySize       uint8
	Key           []byte
}

var _ NLMUpdateKeyUpdateKeyEntry = (*_NLMUpdateKeyUpdateKeyEntry)(nil)

// NewNLMUpdateKeyUpdateKeyEntry factory function for _NLMUpdateKeyUpdateKeyEntry
func NewNLMUpdateKeyUpdateKeyEntry(keyIdentifier uint16, keySize uint8, key []byte) *_NLMUpdateKeyUpdateKeyEntry {
	return &_NLMUpdateKeyUpdateKeyEntry{KeyIdentifier: keyIdentifier, KeySize: keySize, Key: key}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMUpdateKeyUpdateKeyEntryBuilder is a builder for NLMUpdateKeyUpdateKeyEntry
type NLMUpdateKeyUpdateKeyEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(keyIdentifier uint16, keySize uint8, key []byte) NLMUpdateKeyUpdateKeyEntryBuilder
	// WithKeyIdentifier adds KeyIdentifier (property field)
	WithKeyIdentifier(uint16) NLMUpdateKeyUpdateKeyEntryBuilder
	// WithKeySize adds KeySize (property field)
	WithKeySize(uint8) NLMUpdateKeyUpdateKeyEntryBuilder
	// WithKey adds Key (property field)
	WithKey(...byte) NLMUpdateKeyUpdateKeyEntryBuilder
	// Build builds the NLMUpdateKeyUpdateKeyEntry or returns an error if something is wrong
	Build() (NLMUpdateKeyUpdateKeyEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMUpdateKeyUpdateKeyEntry
}

// NewNLMUpdateKeyUpdateKeyEntryBuilder() creates a NLMUpdateKeyUpdateKeyEntryBuilder
func NewNLMUpdateKeyUpdateKeyEntryBuilder() NLMUpdateKeyUpdateKeyEntryBuilder {
	return &_NLMUpdateKeyUpdateKeyEntryBuilder{_NLMUpdateKeyUpdateKeyEntry: new(_NLMUpdateKeyUpdateKeyEntry)}
}

type _NLMUpdateKeyUpdateKeyEntryBuilder struct {
	*_NLMUpdateKeyUpdateKeyEntry

	err *utils.MultiError
}

var _ (NLMUpdateKeyUpdateKeyEntryBuilder) = (*_NLMUpdateKeyUpdateKeyEntryBuilder)(nil)

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) WithMandatoryFields(keyIdentifier uint16, keySize uint8, key []byte) NLMUpdateKeyUpdateKeyEntryBuilder {
	return m.WithKeyIdentifier(keyIdentifier).WithKeySize(keySize).WithKey(key...)
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) WithKeyIdentifier(keyIdentifier uint16) NLMUpdateKeyUpdateKeyEntryBuilder {
	m.KeyIdentifier = keyIdentifier
	return m
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) WithKeySize(keySize uint8) NLMUpdateKeyUpdateKeyEntryBuilder {
	m.KeySize = keySize
	return m
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) WithKey(key ...byte) NLMUpdateKeyUpdateKeyEntryBuilder {
	m.Key = key
	return m
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) Build() (NLMUpdateKeyUpdateKeyEntry, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._NLMUpdateKeyUpdateKeyEntry.deepCopy(), nil
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) MustBuild() NLMUpdateKeyUpdateKeyEntry {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_NLMUpdateKeyUpdateKeyEntryBuilder) DeepCopy() any {
	return m.CreateNLMUpdateKeyUpdateKeyEntryBuilder()
}

// CreateNLMUpdateKeyUpdateKeyEntryBuilder creates a NLMUpdateKeyUpdateKeyEntryBuilder
func (m *_NLMUpdateKeyUpdateKeyEntry) CreateNLMUpdateKeyUpdateKeyEntryBuilder() NLMUpdateKeyUpdateKeyEntryBuilder {
	if m == nil {
		return NewNLMUpdateKeyUpdateKeyEntryBuilder()
	}
	return &_NLMUpdateKeyUpdateKeyEntryBuilder{_NLMUpdateKeyUpdateKeyEntry: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKeyIdentifier() uint16 {
	return m.KeyIdentifier
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKeySize() uint8 {
	return m.KeySize
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKey() []byte {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyUpdateKeyEntry(structType any) NLMUpdateKeyUpdateKeyEntry {
	if casted, ok := structType.(NLMUpdateKeyUpdateKeyEntry); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyUpdateKeyEntry); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetTypeName() string {
	return "NLMUpdateKeyUpdateKeyEntry"
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (keyIdentifier)
	lengthInBits += 16

	// Simple field (keySize)
	lengthInBits += 8

	// Array field
	if len(m.Key) > 0 {
		lengthInBits += 8 * uint16(len(m.Key))
	}

	return lengthInBits
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMUpdateKeyUpdateKeyEntryParse(ctx context.Context, theBytes []byte) (NLMUpdateKeyUpdateKeyEntry, error) {
	return NLMUpdateKeyUpdateKeyEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NLMUpdateKeyUpdateKeyEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMUpdateKeyUpdateKeyEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMUpdateKeyUpdateKeyEntry, error) {
		return NLMUpdateKeyUpdateKeyEntryParseWithBuffer(ctx, readBuffer)
	}
}

func NLMUpdateKeyUpdateKeyEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NLMUpdateKeyUpdateKeyEntry, error) {
	v, err := (&_NLMUpdateKeyUpdateKeyEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__nLMUpdateKeyUpdateKeyEntry NLMUpdateKeyUpdateKeyEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMUpdateKeyUpdateKeyEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyUpdateKeyEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	keyIdentifier, err := ReadSimpleField(ctx, "keyIdentifier", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyIdentifier' field"))
	}
	m.KeyIdentifier = keyIdentifier

	keySize, err := ReadSimpleField(ctx, "keySize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keySize' field"))
	}
	m.KeySize = keySize

	key, err := readBuffer.ReadByteArray("key", int(keySize))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'key' field"))
	}
	m.Key = key

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyUpdateKeyEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyUpdateKeyEntry")
	}

	return m, nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLMUpdateKeyUpdateKeyEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyUpdateKeyEntry")
	}

	if err := WriteSimpleField[uint16](ctx, "keyIdentifier", m.GetKeyIdentifier(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'keyIdentifier' field")
	}

	if err := WriteSimpleField[uint8](ctx, "keySize", m.GetKeySize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'keySize' field")
	}

	if err := WriteByteArrayField(ctx, "key", m.GetKey(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'key' field")
	}

	if popErr := writeBuffer.PopContext("NLMUpdateKeyUpdateKeyEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLMUpdateKeyUpdateKeyEntry")
	}
	return nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) IsNLMUpdateKeyUpdateKeyEntry() {}

func (m *_NLMUpdateKeyUpdateKeyEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMUpdateKeyUpdateKeyEntry) deepCopy() *_NLMUpdateKeyUpdateKeyEntry {
	if m == nil {
		return nil
	}
	_NLMUpdateKeyUpdateKeyEntryCopy := &_NLMUpdateKeyUpdateKeyEntry{
		m.KeyIdentifier,
		m.KeySize,
		utils.DeepCopySlice[byte, byte](m.Key),
	}
	return _NLMUpdateKeyUpdateKeyEntryCopy
}

func (m *_NLMUpdateKeyUpdateKeyEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
