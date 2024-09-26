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

// BACnetConstructedDataBACnetIPNATTraversal is the corresponding interface of BACnetConstructedDataBACnetIPNATTraversal
type BACnetConstructedDataBACnetIPNATTraversal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBacnetIPNATTraversal returns BacnetIPNATTraversal (property field)
	GetBacnetIPNATTraversal() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataBACnetIPNATTraversal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPNATTraversal()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPNATTraversalBuilder
	CreateBACnetConstructedDataBACnetIPNATTraversalBuilder() BACnetConstructedDataBACnetIPNATTraversalBuilder
}

// _BACnetConstructedDataBACnetIPNATTraversal is the data-structure of this message
type _BACnetConstructedDataBACnetIPNATTraversal struct {
	BACnetConstructedDataContract
	BacnetIPNATTraversal BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataBACnetIPNATTraversal = (*_BACnetConstructedDataBACnetIPNATTraversal)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPNATTraversal)(nil)

// NewBACnetConstructedDataBACnetIPNATTraversal factory function for _BACnetConstructedDataBACnetIPNATTraversal
func NewBACnetConstructedDataBACnetIPNATTraversal(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, bacnetIPNATTraversal BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPNATTraversal {
	if bacnetIPNATTraversal == nil {
		panic("bacnetIPNATTraversal of type BACnetApplicationTagBoolean for BACnetConstructedDataBACnetIPNATTraversal must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPNATTraversal{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BacnetIPNATTraversal:          bacnetIPNATTraversal,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPNATTraversalBuilder is a builder for BACnetConstructedDataBACnetIPNATTraversal
type BACnetConstructedDataBACnetIPNATTraversalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bacnetIPNATTraversal BACnetApplicationTagBoolean) BACnetConstructedDataBACnetIPNATTraversalBuilder
	// WithBacnetIPNATTraversal adds BacnetIPNATTraversal (property field)
	WithBacnetIPNATTraversal(BACnetApplicationTagBoolean) BACnetConstructedDataBACnetIPNATTraversalBuilder
	// WithBacnetIPNATTraversalBuilder adds BacnetIPNATTraversal (property field) which is build by the builder
	WithBacnetIPNATTraversalBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataBACnetIPNATTraversalBuilder
	// Build builds the BACnetConstructedDataBACnetIPNATTraversal or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPNATTraversal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPNATTraversal
}

// NewBACnetConstructedDataBACnetIPNATTraversalBuilder() creates a BACnetConstructedDataBACnetIPNATTraversalBuilder
func NewBACnetConstructedDataBACnetIPNATTraversalBuilder() BACnetConstructedDataBACnetIPNATTraversalBuilder {
	return &_BACnetConstructedDataBACnetIPNATTraversalBuilder{_BACnetConstructedDataBACnetIPNATTraversal: new(_BACnetConstructedDataBACnetIPNATTraversal)}
}

type _BACnetConstructedDataBACnetIPNATTraversalBuilder struct {
	*_BACnetConstructedDataBACnetIPNATTraversal

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPNATTraversalBuilder) = (*_BACnetConstructedDataBACnetIPNATTraversalBuilder)(nil)

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) WithMandatoryFields(bacnetIPNATTraversal BACnetApplicationTagBoolean) BACnetConstructedDataBACnetIPNATTraversalBuilder {
	return m.WithBacnetIPNATTraversal(bacnetIPNATTraversal)
}

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) WithBacnetIPNATTraversal(bacnetIPNATTraversal BACnetApplicationTagBoolean) BACnetConstructedDataBACnetIPNATTraversalBuilder {
	m.BacnetIPNATTraversal = bacnetIPNATTraversal
	return m
}

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) WithBacnetIPNATTraversalBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataBACnetIPNATTraversalBuilder {
	builder := builderSupplier(m.BacnetIPNATTraversal.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.BacnetIPNATTraversal, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) Build() (BACnetConstructedDataBACnetIPNATTraversal, error) {
	if m.BacnetIPNATTraversal == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'bacnetIPNATTraversal' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataBACnetIPNATTraversal.deepCopy(), nil
}

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) MustBuild() BACnetConstructedDataBACnetIPNATTraversal {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataBACnetIPNATTraversalBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataBACnetIPNATTraversalBuilder()
}

// CreateBACnetConstructedDataBACnetIPNATTraversalBuilder creates a BACnetConstructedDataBACnetIPNATTraversalBuilder
func (m *_BACnetConstructedDataBACnetIPNATTraversal) CreateBACnetConstructedDataBACnetIPNATTraversalBuilder() BACnetConstructedDataBACnetIPNATTraversalBuilder {
	if m == nil {
		return NewBACnetConstructedDataBACnetIPNATTraversalBuilder()
	}
	return &_BACnetConstructedDataBACnetIPNATTraversalBuilder{_BACnetConstructedDataBACnetIPNATTraversal: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_NAT_TRAVERSAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetBacnetIPNATTraversal() BACnetApplicationTagBoolean {
	return m.BacnetIPNATTraversal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetBacnetIPNATTraversal())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPNATTraversal(structType any) BACnetConstructedDataBACnetIPNATTraversal {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPNATTraversal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPNATTraversal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPNATTraversal"
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (bacnetIPNATTraversal)
	lengthInBits += m.BacnetIPNATTraversal.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPNATTraversal BACnetConstructedDataBACnetIPNATTraversal, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPNATTraversal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPNATTraversal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetIPNATTraversal, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "bacnetIPNATTraversal", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetIPNATTraversal' field"))
	}
	m.BacnetIPNATTraversal = bacnetIPNATTraversal

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), bacnetIPNATTraversal)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPNATTraversal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPNATTraversal")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPNATTraversal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPNATTraversal")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "bacnetIPNATTraversal", m.GetBacnetIPNATTraversal(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bacnetIPNATTraversal' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPNATTraversal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPNATTraversal")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) IsBACnetConstructedDataBACnetIPNATTraversal() {}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) deepCopy() *_BACnetConstructedDataBACnetIPNATTraversal {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPNATTraversalCopy := &_BACnetConstructedDataBACnetIPNATTraversal{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.BacnetIPNATTraversal.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPNATTraversalCopy
}

func (m *_BACnetConstructedDataBACnetIPNATTraversal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
