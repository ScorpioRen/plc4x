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

// BACnetConfirmedServiceRequestVTOpen is the corresponding interface of BACnetConfirmedServiceRequestVTOpen
type BACnetConfirmedServiceRequestVTOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetVtClass returns VtClass (property field)
	GetVtClass() BACnetVTClassTagged
	// GetLocalVtSessionIdentifier returns LocalVtSessionIdentifier (property field)
	GetLocalVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestVTOpen is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestVTOpen()
	// CreateBuilder creates a BACnetConfirmedServiceRequestVTOpenBuilder
	CreateBACnetConfirmedServiceRequestVTOpenBuilder() BACnetConfirmedServiceRequestVTOpenBuilder
}

// _BACnetConfirmedServiceRequestVTOpen is the data-structure of this message
type _BACnetConfirmedServiceRequestVTOpen struct {
	BACnetConfirmedServiceRequestContract
	VtClass                  BACnetVTClassTagged
	LocalVtSessionIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestVTOpen = (*_BACnetConfirmedServiceRequestVTOpen)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestVTOpen)(nil)

// NewBACnetConfirmedServiceRequestVTOpen factory function for _BACnetConfirmedServiceRequestVTOpen
func NewBACnetConfirmedServiceRequestVTOpen(vtClass BACnetVTClassTagged, localVtSessionIdentifier BACnetApplicationTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestVTOpen {
	if vtClass == nil {
		panic("vtClass of type BACnetVTClassTagged for BACnetConfirmedServiceRequestVTOpen must not be nil")
	}
	if localVtSessionIdentifier == nil {
		panic("localVtSessionIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestVTOpen must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestVTOpen{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		VtClass:                               vtClass,
		LocalVtSessionIdentifier:              localVtSessionIdentifier,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestVTOpenBuilder is a builder for BACnetConfirmedServiceRequestVTOpen
type BACnetConfirmedServiceRequestVTOpenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(vtClass BACnetVTClassTagged, localVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestVTOpenBuilder
	// WithVtClass adds VtClass (property field)
	WithVtClass(BACnetVTClassTagged) BACnetConfirmedServiceRequestVTOpenBuilder
	// WithVtClassBuilder adds VtClass (property field) which is build by the builder
	WithVtClassBuilder(func(BACnetVTClassTaggedBuilder) BACnetVTClassTaggedBuilder) BACnetConfirmedServiceRequestVTOpenBuilder
	// WithLocalVtSessionIdentifier adds LocalVtSessionIdentifier (property field)
	WithLocalVtSessionIdentifier(BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestVTOpenBuilder
	// WithLocalVtSessionIdentifierBuilder adds LocalVtSessionIdentifier (property field) which is build by the builder
	WithLocalVtSessionIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestVTOpenBuilder
	// Build builds the BACnetConfirmedServiceRequestVTOpen or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestVTOpen, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestVTOpen
}

// NewBACnetConfirmedServiceRequestVTOpenBuilder() creates a BACnetConfirmedServiceRequestVTOpenBuilder
func NewBACnetConfirmedServiceRequestVTOpenBuilder() BACnetConfirmedServiceRequestVTOpenBuilder {
	return &_BACnetConfirmedServiceRequestVTOpenBuilder{_BACnetConfirmedServiceRequestVTOpen: new(_BACnetConfirmedServiceRequestVTOpen)}
}

type _BACnetConfirmedServiceRequestVTOpenBuilder struct {
	*_BACnetConfirmedServiceRequestVTOpen

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestVTOpenBuilder) = (*_BACnetConfirmedServiceRequestVTOpenBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) WithMandatoryFields(vtClass BACnetVTClassTagged, localVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestVTOpenBuilder {
	return m.WithVtClass(vtClass).WithLocalVtSessionIdentifier(localVtSessionIdentifier)
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) WithVtClass(vtClass BACnetVTClassTagged) BACnetConfirmedServiceRequestVTOpenBuilder {
	m.VtClass = vtClass
	return m
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) WithVtClassBuilder(builderSupplier func(BACnetVTClassTaggedBuilder) BACnetVTClassTaggedBuilder) BACnetConfirmedServiceRequestVTOpenBuilder {
	builder := builderSupplier(m.VtClass.CreateBACnetVTClassTaggedBuilder())
	var err error
	m.VtClass, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetVTClassTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) WithLocalVtSessionIdentifier(localVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestVTOpenBuilder {
	m.LocalVtSessionIdentifier = localVtSessionIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) WithLocalVtSessionIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestVTOpenBuilder {
	builder := builderSupplier(m.LocalVtSessionIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.LocalVtSessionIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) Build() (BACnetConfirmedServiceRequestVTOpen, error) {
	if m.VtClass == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'vtClass' not set"))
	}
	if m.LocalVtSessionIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'localVtSessionIdentifier' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestVTOpen.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) MustBuild() BACnetConfirmedServiceRequestVTOpen {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestVTOpenBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestVTOpenBuilder()
}

// CreateBACnetConfirmedServiceRequestVTOpenBuilder creates a BACnetConfirmedServiceRequestVTOpenBuilder
func (m *_BACnetConfirmedServiceRequestVTOpen) CreateBACnetConfirmedServiceRequestVTOpenBuilder() BACnetConfirmedServiceRequestVTOpenBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestVTOpenBuilder()
	}
	return &_BACnetConfirmedServiceRequestVTOpenBuilder{_BACnetConfirmedServiceRequestVTOpen: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_OPEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen) GetVtClass() BACnetVTClassTagged {
	return m.VtClass
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetLocalVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.LocalVtSessionIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTOpen(structType any) BACnetConfirmedServiceRequestVTOpen {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTOpen); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTOpen"
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (vtClass)
	lengthInBits += m.VtClass.GetLengthInBits(ctx)

	// Simple field (localVtSessionIdentifier)
	lengthInBits += m.LocalVtSessionIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestVTOpen) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestVTOpen BACnetConfirmedServiceRequestVTOpen, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vtClass, err := ReadSimpleField[BACnetVTClassTagged](ctx, "vtClass", ReadComplex[BACnetVTClassTagged](BACnetVTClassTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtClass' field"))
	}
	m.VtClass = vtClass

	localVtSessionIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "localVtSessionIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localVtSessionIdentifier' field"))
	}
	m.LocalVtSessionIdentifier = localVtSessionIdentifier

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTOpen")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTOpen")
		}

		if err := WriteSimpleField[BACnetVTClassTagged](ctx, "vtClass", m.GetVtClass(), WriteComplex[BACnetVTClassTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtClass' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "localVtSessionIdentifier", m.GetLocalVtSessionIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'localVtSessionIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTOpen")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestVTOpen) IsBACnetConfirmedServiceRequestVTOpen() {}

func (m *_BACnetConfirmedServiceRequestVTOpen) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestVTOpen) deepCopy() *_BACnetConfirmedServiceRequestVTOpen {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestVTOpenCopy := &_BACnetConfirmedServiceRequestVTOpen{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.VtClass.DeepCopy().(BACnetVTClassTagged),
		m.LocalVtSessionIdentifier.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestVTOpenCopy
}

func (m *_BACnetConfirmedServiceRequestVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
