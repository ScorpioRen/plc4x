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

// BACnetEventLogRecord is the corresponding interface of BACnetEventLogRecord
type BACnetEventLogRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogDatum returns LogDatum (property field)
	GetLogDatum() BACnetEventLogRecordLogDatum
	// IsBACnetEventLogRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecord()
	// CreateBuilder creates a BACnetEventLogRecordBuilder
	CreateBACnetEventLogRecordBuilder() BACnetEventLogRecordBuilder
}

// _BACnetEventLogRecord is the data-structure of this message
type _BACnetEventLogRecord struct {
	Timestamp BACnetDateTimeEnclosed
	LogDatum  BACnetEventLogRecordLogDatum
}

var _ BACnetEventLogRecord = (*_BACnetEventLogRecord)(nil)

// NewBACnetEventLogRecord factory function for _BACnetEventLogRecord
func NewBACnetEventLogRecord(timestamp BACnetDateTimeEnclosed, logDatum BACnetEventLogRecordLogDatum) *_BACnetEventLogRecord {
	if timestamp == nil {
		panic("timestamp of type BACnetDateTimeEnclosed for BACnetEventLogRecord must not be nil")
	}
	if logDatum == nil {
		panic("logDatum of type BACnetEventLogRecordLogDatum for BACnetEventLogRecord must not be nil")
	}
	return &_BACnetEventLogRecord{Timestamp: timestamp, LogDatum: logDatum}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventLogRecordBuilder is a builder for BACnetEventLogRecord
type BACnetEventLogRecordBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timestamp BACnetDateTimeEnclosed, logDatum BACnetEventLogRecordLogDatum) BACnetEventLogRecordBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(BACnetDateTimeEnclosed) BACnetEventLogRecordBuilder
	// WithTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithTimestampBuilder(func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetEventLogRecordBuilder
	// WithLogDatum adds LogDatum (property field)
	WithLogDatum(BACnetEventLogRecordLogDatum) BACnetEventLogRecordBuilder
	// WithLogDatumBuilder adds LogDatum (property field) which is build by the builder
	WithLogDatumBuilder(func(BACnetEventLogRecordLogDatumBuilder) BACnetEventLogRecordLogDatumBuilder) BACnetEventLogRecordBuilder
	// Build builds the BACnetEventLogRecord or returns an error if something is wrong
	Build() (BACnetEventLogRecord, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventLogRecord
}

// NewBACnetEventLogRecordBuilder() creates a BACnetEventLogRecordBuilder
func NewBACnetEventLogRecordBuilder() BACnetEventLogRecordBuilder {
	return &_BACnetEventLogRecordBuilder{_BACnetEventLogRecord: new(_BACnetEventLogRecord)}
}

type _BACnetEventLogRecordBuilder struct {
	*_BACnetEventLogRecord

	err *utils.MultiError
}

var _ (BACnetEventLogRecordBuilder) = (*_BACnetEventLogRecordBuilder)(nil)

func (m *_BACnetEventLogRecordBuilder) WithMandatoryFields(timestamp BACnetDateTimeEnclosed, logDatum BACnetEventLogRecordLogDatum) BACnetEventLogRecordBuilder {
	return m.WithTimestamp(timestamp).WithLogDatum(logDatum)
}

func (m *_BACnetEventLogRecordBuilder) WithTimestamp(timestamp BACnetDateTimeEnclosed) BACnetEventLogRecordBuilder {
	m.Timestamp = timestamp
	return m
}

func (m *_BACnetEventLogRecordBuilder) WithTimestampBuilder(builderSupplier func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetEventLogRecordBuilder {
	builder := builderSupplier(m.Timestamp.CreateBACnetDateTimeEnclosedBuilder())
	var err error
	m.Timestamp, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateTimeEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetEventLogRecordBuilder) WithLogDatum(logDatum BACnetEventLogRecordLogDatum) BACnetEventLogRecordBuilder {
	m.LogDatum = logDatum
	return m
}

func (m *_BACnetEventLogRecordBuilder) WithLogDatumBuilder(builderSupplier func(BACnetEventLogRecordLogDatumBuilder) BACnetEventLogRecordLogDatumBuilder) BACnetEventLogRecordBuilder {
	builder := builderSupplier(m.LogDatum.CreateBACnetEventLogRecordLogDatumBuilder())
	var err error
	m.LogDatum, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEventLogRecordLogDatumBuilder failed"))
	}
	return m
}

func (m *_BACnetEventLogRecordBuilder) Build() (BACnetEventLogRecord, error) {
	if m.Timestamp == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timestamp' not set"))
	}
	if m.LogDatum == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'logDatum' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEventLogRecord.deepCopy(), nil
}

func (m *_BACnetEventLogRecordBuilder) MustBuild() BACnetEventLogRecord {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventLogRecordBuilder) DeepCopy() any {
	return m.CreateBACnetEventLogRecordBuilder()
}

// CreateBACnetEventLogRecordBuilder creates a BACnetEventLogRecordBuilder
func (m *_BACnetEventLogRecord) CreateBACnetEventLogRecordBuilder() BACnetEventLogRecordBuilder {
	if m == nil {
		return NewBACnetEventLogRecordBuilder()
	}
	return &_BACnetEventLogRecordBuilder{_BACnetEventLogRecord: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetEventLogRecord) GetLogDatum() BACnetEventLogRecordLogDatum {
	return m.LogDatum
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecord(structType any) BACnetEventLogRecord {
	if casted, ok := structType.(BACnetEventLogRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecord) GetTypeName() string {
	return "BACnetEventLogRecord"
}

func (m *_BACnetEventLogRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (logDatum)
	lengthInBits += m.LogDatum.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventLogRecordParse(ctx context.Context, theBytes []byte) (BACnetEventLogRecord, error) {
	return BACnetEventLogRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventLogRecordParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventLogRecord, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventLogRecord, error) {
		return BACnetEventLogRecordParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventLogRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventLogRecord, error) {
	v, err := (&_BACnetEventLogRecord{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventLogRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventLogRecord BACnetEventLogRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	logDatum, err := ReadSimpleField[BACnetEventLogRecordLogDatum](ctx, "logDatum", ReadComplex[BACnetEventLogRecordLogDatum](BACnetEventLogRecordLogDatumParseWithBufferProducer[BACnetEventLogRecordLogDatum]((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logDatum' field"))
	}
	m.LogDatum = logDatum

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecord")
	}

	return m, nil
}

func (m *_BACnetEventLogRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventLogRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventLogRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecord")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetEventLogRecordLogDatum](ctx, "logDatum", m.GetLogDatum(), WriteComplex[BACnetEventLogRecordLogDatum](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'logDatum' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventLogRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventLogRecord")
	}
	return nil
}

func (m *_BACnetEventLogRecord) IsBACnetEventLogRecord() {}

func (m *_BACnetEventLogRecord) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventLogRecord) deepCopy() *_BACnetEventLogRecord {
	if m == nil {
		return nil
	}
	_BACnetEventLogRecordCopy := &_BACnetEventLogRecord{
		m.Timestamp.DeepCopy().(BACnetDateTimeEnclosed),
		m.LogDatum.DeepCopy().(BACnetEventLogRecordLogDatum),
	}
	return _BACnetEventLogRecordCopy
}

func (m *_BACnetEventLogRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
