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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetListOfCovReferences returns ListOfCovReferences (property field)
	GetListOfCovReferences() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
	// IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry()
	// CreateBuilder creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry struct {
	MonitoredObjectIdentifier BACnetContextTagObjectIdentifier
	ListOfCovReferences       BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
}

var _ BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry)(nil)

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, listOfCovReferences BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry must not be nil")
	}
	if listOfCovReferences == nil {
		panic("listOfCovReferences of type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry must not be nil")
	}
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{MonitoredObjectIdentifier: monitoredObjectIdentifier, ListOfCovReferences: listOfCovReferences}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder is a builder for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, listOfCovReferences BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	// WithMonitoredObjectIdentifier adds MonitoredObjectIdentifier (property field)
	WithMonitoredObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	// WithMonitoredObjectIdentifierBuilder adds MonitoredObjectIdentifier (property field) which is build by the builder
	WithMonitoredObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	// WithListOfCovReferences adds ListOfCovReferences (property field)
	WithListOfCovReferences(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	// WithListOfCovReferencesBuilder adds ListOfCovReferences (property field) which is build by the builder
	WithListOfCovReferencesBuilder(func(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
	// Build builds the BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry or returns an error if something is wrong
	Build() (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
}

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder() creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder{_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry: new(_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry)}
}

type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder struct {
	*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry

	err *utils.MultiError
}

var _ (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder)(nil)

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) WithMandatoryFields(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, listOfCovReferences BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	return m.WithMonitoredObjectIdentifier(monitoredObjectIdentifier).WithListOfCovReferences(listOfCovReferences)
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) WithMonitoredObjectIdentifier(monitoredObjectIdentifier BACnetContextTagObjectIdentifier) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) WithMonitoredObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	builder := builderSupplier(m.MonitoredObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.MonitoredObjectIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) WithListOfCovReferences(listOfCovReferences BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	m.ListOfCovReferences = listOfCovReferences
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) WithListOfCovReferencesBuilder(builderSupplier func(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	builder := builderSupplier(m.ListOfCovReferences.CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder())
	var err error
	m.ListOfCovReferences, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesBuilder failed"))
	}
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) Build() (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	if m.MonitoredObjectIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'monitoredObjectIdentifier' not set"))
	}
	if m.ListOfCovReferences == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'listOfCovReferences' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry.deepCopy(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) MustBuild() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder) DeepCopy() any {
	return m.CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder()
}

// CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder
func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder {
	if m == nil {
		return NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder()
	}
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryBuilder{_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetListOfCovReferences() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences {
	return m.ListOfCovReferences
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(structType any) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (listOfCovReferences)
	lengthInBits += m.ListOfCovReferences.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParse(ctx context.Context, theBytes []byte) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
		return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	v, err := (&_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	listOfCovReferences, err := ReadSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](ctx, "listOfCovReferences", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovReferences' field"))
	}
	m.ListOfCovReferences = listOfCovReferences

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	return m, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](ctx, "listOfCovReferences", m.GetListOfCovReferences(), WriteComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfCovReferences' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry() {
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) deepCopy() *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	if m == nil {
		return nil
	}
	_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryCopy := &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{
		m.MonitoredObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.ListOfCovReferences.DeepCopy().(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences),
	}
	return _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryCopy
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
