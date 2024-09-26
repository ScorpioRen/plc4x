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

// Annotation is the corresponding interface of Annotation
type Annotation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetMessage returns Message (property field)
	GetMessage() PascalString
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// GetAnnotationTime returns AnnotationTime (property field)
	GetAnnotationTime() int64
	// IsAnnotation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAnnotation()
	// CreateBuilder creates a AnnotationBuilder
	CreateAnnotationBuilder() AnnotationBuilder
}

// _Annotation is the data-structure of this message
type _Annotation struct {
	ExtensionObjectDefinitionContract
	Message        PascalString
	UserName       PascalString
	AnnotationTime int64
}

var _ Annotation = (*_Annotation)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Annotation)(nil)

// NewAnnotation factory function for _Annotation
func NewAnnotation(message PascalString, userName PascalString, annotationTime int64) *_Annotation {
	if message == nil {
		panic("message of type PascalString for Annotation must not be nil")
	}
	if userName == nil {
		panic("userName of type PascalString for Annotation must not be nil")
	}
	_result := &_Annotation{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Message:                           message,
		UserName:                          userName,
		AnnotationTime:                    annotationTime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AnnotationBuilder is a builder for Annotation
type AnnotationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(message PascalString, userName PascalString, annotationTime int64) AnnotationBuilder
	// WithMessage adds Message (property field)
	WithMessage(PascalString) AnnotationBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(PascalStringBuilder) PascalStringBuilder) AnnotationBuilder
	// WithUserName adds UserName (property field)
	WithUserName(PascalString) AnnotationBuilder
	// WithUserNameBuilder adds UserName (property field) which is build by the builder
	WithUserNameBuilder(func(PascalStringBuilder) PascalStringBuilder) AnnotationBuilder
	// WithAnnotationTime adds AnnotationTime (property field)
	WithAnnotationTime(int64) AnnotationBuilder
	// Build builds the Annotation or returns an error if something is wrong
	Build() (Annotation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Annotation
}

// NewAnnotationBuilder() creates a AnnotationBuilder
func NewAnnotationBuilder() AnnotationBuilder {
	return &_AnnotationBuilder{_Annotation: new(_Annotation)}
}

type _AnnotationBuilder struct {
	*_Annotation

	err *utils.MultiError
}

var _ (AnnotationBuilder) = (*_AnnotationBuilder)(nil)

func (m *_AnnotationBuilder) WithMandatoryFields(message PascalString, userName PascalString, annotationTime int64) AnnotationBuilder {
	return m.WithMessage(message).WithUserName(userName).WithAnnotationTime(annotationTime)
}

func (m *_AnnotationBuilder) WithMessage(message PascalString) AnnotationBuilder {
	m.Message = message
	return m
}

func (m *_AnnotationBuilder) WithMessageBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) AnnotationBuilder {
	builder := builderSupplier(m.Message.CreatePascalStringBuilder())
	var err error
	m.Message, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_AnnotationBuilder) WithUserName(userName PascalString) AnnotationBuilder {
	m.UserName = userName
	return m
}

func (m *_AnnotationBuilder) WithUserNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) AnnotationBuilder {
	builder := builderSupplier(m.UserName.CreatePascalStringBuilder())
	var err error
	m.UserName, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_AnnotationBuilder) WithAnnotationTime(annotationTime int64) AnnotationBuilder {
	m.AnnotationTime = annotationTime
	return m
}

func (m *_AnnotationBuilder) Build() (Annotation, error) {
	if m.Message == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if m.UserName == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'userName' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._Annotation.deepCopy(), nil
}

func (m *_AnnotationBuilder) MustBuild() Annotation {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AnnotationBuilder) DeepCopy() any {
	return m.CreateAnnotationBuilder()
}

// CreateAnnotationBuilder creates a AnnotationBuilder
func (m *_Annotation) CreateAnnotationBuilder() AnnotationBuilder {
	if m == nil {
		return NewAnnotationBuilder()
	}
	return &_AnnotationBuilder{_Annotation: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Annotation) GetIdentifier() string {
	return "893"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Annotation) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Annotation) GetMessage() PascalString {
	return m.Message
}

func (m *_Annotation) GetUserName() PascalString {
	return m.UserName
}

func (m *_Annotation) GetAnnotationTime() int64 {
	return m.AnnotationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAnnotation(structType any) Annotation {
	if casted, ok := structType.(Annotation); ok {
		return casted
	}
	if casted, ok := structType.(*Annotation); ok {
		return *casted
	}
	return nil
}

func (m *_Annotation) GetTypeName() string {
	return "Annotation"
}

func (m *_Annotation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	// Simple field (annotationTime)
	lengthInBits += 64

	return lengthInBits
}

func (m *_Annotation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Annotation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__annotation Annotation, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Annotation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Annotation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	message, err := ReadSimpleField[PascalString](ctx, "message", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	userName, err := ReadSimpleField[PascalString](ctx, "userName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	annotationTime, err := ReadSimpleField(ctx, "annotationTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'annotationTime' field"))
	}
	m.AnnotationTime = annotationTime

	if closeErr := readBuffer.CloseContext("Annotation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Annotation")
	}

	return m, nil
}

func (m *_Annotation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Annotation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Annotation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Annotation")
		}

		if err := WriteSimpleField[PascalString](ctx, "message", m.GetMessage(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "userName", m.GetUserName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if err := WriteSimpleField[int64](ctx, "annotationTime", m.GetAnnotationTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'annotationTime' field")
		}

		if popErr := writeBuffer.PopContext("Annotation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Annotation")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Annotation) IsAnnotation() {}

func (m *_Annotation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Annotation) deepCopy() *_Annotation {
	if m == nil {
		return nil
	}
	_AnnotationCopy := &_Annotation{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Message.DeepCopy().(PascalString),
		m.UserName.DeepCopy().(PascalString),
		m.AnnotationTime,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AnnotationCopy
}

func (m *_Annotation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
