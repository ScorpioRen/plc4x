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

// BVLCForeignDeviceTableEntry is the corresponding interface of BVLCForeignDeviceTableEntry
type BVLCForeignDeviceTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetTtl returns Ttl (property field)
	GetTtl() uint16
	// GetSecondRemainingBeforePurge returns SecondRemainingBeforePurge (property field)
	GetSecondRemainingBeforePurge() uint16
	// IsBVLCForeignDeviceTableEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCForeignDeviceTableEntry()
	// CreateBuilder creates a BVLCForeignDeviceTableEntryBuilder
	CreateBVLCForeignDeviceTableEntryBuilder() BVLCForeignDeviceTableEntryBuilder
}

// _BVLCForeignDeviceTableEntry is the data-structure of this message
type _BVLCForeignDeviceTableEntry struct {
	Ip                         []uint8
	Port                       uint16
	Ttl                        uint16
	SecondRemainingBeforePurge uint16
}

var _ BVLCForeignDeviceTableEntry = (*_BVLCForeignDeviceTableEntry)(nil)

// NewBVLCForeignDeviceTableEntry factory function for _BVLCForeignDeviceTableEntry
func NewBVLCForeignDeviceTableEntry(ip []uint8, port uint16, ttl uint16, secondRemainingBeforePurge uint16) *_BVLCForeignDeviceTableEntry {
	return &_BVLCForeignDeviceTableEntry{Ip: ip, Port: port, Ttl: ttl, SecondRemainingBeforePurge: secondRemainingBeforePurge}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCForeignDeviceTableEntryBuilder is a builder for BVLCForeignDeviceTableEntry
type BVLCForeignDeviceTableEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ip []uint8, port uint16, ttl uint16, secondRemainingBeforePurge uint16) BVLCForeignDeviceTableEntryBuilder
	// WithIp adds Ip (property field)
	WithIp(...uint8) BVLCForeignDeviceTableEntryBuilder
	// WithPort adds Port (property field)
	WithPort(uint16) BVLCForeignDeviceTableEntryBuilder
	// WithTtl adds Ttl (property field)
	WithTtl(uint16) BVLCForeignDeviceTableEntryBuilder
	// WithSecondRemainingBeforePurge adds SecondRemainingBeforePurge (property field)
	WithSecondRemainingBeforePurge(uint16) BVLCForeignDeviceTableEntryBuilder
	// Build builds the BVLCForeignDeviceTableEntry or returns an error if something is wrong
	Build() (BVLCForeignDeviceTableEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCForeignDeviceTableEntry
}

// NewBVLCForeignDeviceTableEntryBuilder() creates a BVLCForeignDeviceTableEntryBuilder
func NewBVLCForeignDeviceTableEntryBuilder() BVLCForeignDeviceTableEntryBuilder {
	return &_BVLCForeignDeviceTableEntryBuilder{_BVLCForeignDeviceTableEntry: new(_BVLCForeignDeviceTableEntry)}
}

type _BVLCForeignDeviceTableEntryBuilder struct {
	*_BVLCForeignDeviceTableEntry

	err *utils.MultiError
}

var _ (BVLCForeignDeviceTableEntryBuilder) = (*_BVLCForeignDeviceTableEntryBuilder)(nil)

func (m *_BVLCForeignDeviceTableEntryBuilder) WithMandatoryFields(ip []uint8, port uint16, ttl uint16, secondRemainingBeforePurge uint16) BVLCForeignDeviceTableEntryBuilder {
	return m.WithIp(ip...).WithPort(port).WithTtl(ttl).WithSecondRemainingBeforePurge(secondRemainingBeforePurge)
}

func (m *_BVLCForeignDeviceTableEntryBuilder) WithIp(ip ...uint8) BVLCForeignDeviceTableEntryBuilder {
	m.Ip = ip
	return m
}

func (m *_BVLCForeignDeviceTableEntryBuilder) WithPort(port uint16) BVLCForeignDeviceTableEntryBuilder {
	m.Port = port
	return m
}

func (m *_BVLCForeignDeviceTableEntryBuilder) WithTtl(ttl uint16) BVLCForeignDeviceTableEntryBuilder {
	m.Ttl = ttl
	return m
}

func (m *_BVLCForeignDeviceTableEntryBuilder) WithSecondRemainingBeforePurge(secondRemainingBeforePurge uint16) BVLCForeignDeviceTableEntryBuilder {
	m.SecondRemainingBeforePurge = secondRemainingBeforePurge
	return m
}

func (m *_BVLCForeignDeviceTableEntryBuilder) Build() (BVLCForeignDeviceTableEntry, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BVLCForeignDeviceTableEntry.deepCopy(), nil
}

func (m *_BVLCForeignDeviceTableEntryBuilder) MustBuild() BVLCForeignDeviceTableEntry {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BVLCForeignDeviceTableEntryBuilder) DeepCopy() any {
	return m.CreateBVLCForeignDeviceTableEntryBuilder()
}

// CreateBVLCForeignDeviceTableEntryBuilder creates a BVLCForeignDeviceTableEntryBuilder
func (m *_BVLCForeignDeviceTableEntry) CreateBVLCForeignDeviceTableEntryBuilder() BVLCForeignDeviceTableEntryBuilder {
	if m == nil {
		return NewBVLCForeignDeviceTableEntryBuilder()
	}
	return &_BVLCForeignDeviceTableEntryBuilder{_BVLCForeignDeviceTableEntry: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCForeignDeviceTableEntry) GetTtl() uint16 {
	return m.Ttl
}

func (m *_BVLCForeignDeviceTableEntry) GetSecondRemainingBeforePurge() uint16 {
	return m.SecondRemainingBeforePurge
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCForeignDeviceTableEntry(structType any) BVLCForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCForeignDeviceTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCForeignDeviceTableEntry"
}

func (m *_BVLCForeignDeviceTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (ttl)
	lengthInBits += 16

	// Simple field (secondRemainingBeforePurge)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCForeignDeviceTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCForeignDeviceTableEntryParse(ctx context.Context, theBytes []byte) (BVLCForeignDeviceTableEntry, error) {
	return BVLCForeignDeviceTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BVLCForeignDeviceTableEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCForeignDeviceTableEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCForeignDeviceTableEntry, error) {
		return BVLCForeignDeviceTableEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BVLCForeignDeviceTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCForeignDeviceTableEntry, error) {
	v, err := (&_BVLCForeignDeviceTableEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BVLCForeignDeviceTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bVLCForeignDeviceTableEntry BVLCForeignDeviceTableEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	ttl, err := ReadSimpleField(ctx, "ttl", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ttl' field"))
	}
	m.Ttl = ttl

	secondRemainingBeforePurge, err := ReadSimpleField(ctx, "secondRemainingBeforePurge", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondRemainingBeforePurge' field"))
	}
	m.SecondRemainingBeforePurge = secondRemainingBeforePurge

	if closeErr := readBuffer.CloseContext("BVLCForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForeignDeviceTableEntry")
	}

	return m, nil
}

func (m *_BVLCForeignDeviceTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCForeignDeviceTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BVLCForeignDeviceTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCForeignDeviceTableEntry")
	}

	if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'ip' field")
	}

	if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'port' field")
	}

	if err := WriteSimpleField[uint16](ctx, "ttl", m.GetTtl(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'ttl' field")
	}

	if err := WriteSimpleField[uint16](ctx, "secondRemainingBeforePurge", m.GetSecondRemainingBeforePurge(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'secondRemainingBeforePurge' field")
	}

	if popErr := writeBuffer.PopContext("BVLCForeignDeviceTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCForeignDeviceTableEntry")
	}
	return nil
}

func (m *_BVLCForeignDeviceTableEntry) IsBVLCForeignDeviceTableEntry() {}

func (m *_BVLCForeignDeviceTableEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCForeignDeviceTableEntry) deepCopy() *_BVLCForeignDeviceTableEntry {
	if m == nil {
		return nil
	}
	_BVLCForeignDeviceTableEntryCopy := &_BVLCForeignDeviceTableEntry{
		utils.DeepCopySlice[uint8, uint8](m.Ip),
		m.Port,
		m.Ttl,
		m.SecondRemainingBeforePurge,
	}
	return _BVLCForeignDeviceTableEntryCopy
}

func (m *_BVLCForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
