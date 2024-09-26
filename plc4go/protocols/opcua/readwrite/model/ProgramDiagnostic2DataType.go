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

// ProgramDiagnostic2DataType is the corresponding interface of ProgramDiagnostic2DataType
type ProgramDiagnostic2DataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetCreateSessionId returns CreateSessionId (property field)
	GetCreateSessionId() NodeId
	// GetCreateClientName returns CreateClientName (property field)
	GetCreateClientName() PascalString
	// GetInvocationCreationTime returns InvocationCreationTime (property field)
	GetInvocationCreationTime() int64
	// GetLastTransitionTime returns LastTransitionTime (property field)
	GetLastTransitionTime() int64
	// GetLastMethodCall returns LastMethodCall (property field)
	GetLastMethodCall() PascalString
	// GetLastMethodSessionId returns LastMethodSessionId (property field)
	GetLastMethodSessionId() NodeId
	// GetNoOfLastMethodInputArguments returns NoOfLastMethodInputArguments (property field)
	GetNoOfLastMethodInputArguments() int32
	// GetLastMethodInputArguments returns LastMethodInputArguments (property field)
	GetLastMethodInputArguments() []ExtensionObjectDefinition
	// GetNoOfLastMethodOutputArguments returns NoOfLastMethodOutputArguments (property field)
	GetNoOfLastMethodOutputArguments() int32
	// GetLastMethodOutputArguments returns LastMethodOutputArguments (property field)
	GetLastMethodOutputArguments() []ExtensionObjectDefinition
	// GetNoOfLastMethodInputValues returns NoOfLastMethodInputValues (property field)
	GetNoOfLastMethodInputValues() int32
	// GetLastMethodInputValues returns LastMethodInputValues (property field)
	GetLastMethodInputValues() []Variant
	// GetNoOfLastMethodOutputValues returns NoOfLastMethodOutputValues (property field)
	GetNoOfLastMethodOutputValues() int32
	// GetLastMethodOutputValues returns LastMethodOutputValues (property field)
	GetLastMethodOutputValues() []Variant
	// GetLastMethodCallTime returns LastMethodCallTime (property field)
	GetLastMethodCallTime() int64
	// GetLastMethodReturnStatus returns LastMethodReturnStatus (property field)
	GetLastMethodReturnStatus() StatusCode
	// IsProgramDiagnostic2DataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsProgramDiagnostic2DataType()
	// CreateBuilder creates a ProgramDiagnostic2DataTypeBuilder
	CreateProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder
}

// _ProgramDiagnostic2DataType is the data-structure of this message
type _ProgramDiagnostic2DataType struct {
	ExtensionObjectDefinitionContract
	CreateSessionId               NodeId
	CreateClientName              PascalString
	InvocationCreationTime        int64
	LastTransitionTime            int64
	LastMethodCall                PascalString
	LastMethodSessionId           NodeId
	NoOfLastMethodInputArguments  int32
	LastMethodInputArguments      []ExtensionObjectDefinition
	NoOfLastMethodOutputArguments int32
	LastMethodOutputArguments     []ExtensionObjectDefinition
	NoOfLastMethodInputValues     int32
	LastMethodInputValues         []Variant
	NoOfLastMethodOutputValues    int32
	LastMethodOutputValues        []Variant
	LastMethodCallTime            int64
	LastMethodReturnStatus        StatusCode
}

var _ ProgramDiagnostic2DataType = (*_ProgramDiagnostic2DataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ProgramDiagnostic2DataType)(nil)

// NewProgramDiagnostic2DataType factory function for _ProgramDiagnostic2DataType
func NewProgramDiagnostic2DataType(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, noOfLastMethodInputArguments int32, lastMethodInputArguments []ExtensionObjectDefinition, noOfLastMethodOutputArguments int32, lastMethodOutputArguments []ExtensionObjectDefinition, noOfLastMethodInputValues int32, lastMethodInputValues []Variant, noOfLastMethodOutputValues int32, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) *_ProgramDiagnostic2DataType {
	if createSessionId == nil {
		panic("createSessionId of type NodeId for ProgramDiagnostic2DataType must not be nil")
	}
	if createClientName == nil {
		panic("createClientName of type PascalString for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodCall == nil {
		panic("lastMethodCall of type PascalString for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodSessionId == nil {
		panic("lastMethodSessionId of type NodeId for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodReturnStatus == nil {
		panic("lastMethodReturnStatus of type StatusCode for ProgramDiagnostic2DataType must not be nil")
	}
	_result := &_ProgramDiagnostic2DataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CreateSessionId:                   createSessionId,
		CreateClientName:                  createClientName,
		InvocationCreationTime:            invocationCreationTime,
		LastTransitionTime:                lastTransitionTime,
		LastMethodCall:                    lastMethodCall,
		LastMethodSessionId:               lastMethodSessionId,
		NoOfLastMethodInputArguments:      noOfLastMethodInputArguments,
		LastMethodInputArguments:          lastMethodInputArguments,
		NoOfLastMethodOutputArguments:     noOfLastMethodOutputArguments,
		LastMethodOutputArguments:         lastMethodOutputArguments,
		NoOfLastMethodInputValues:         noOfLastMethodInputValues,
		LastMethodInputValues:             lastMethodInputValues,
		NoOfLastMethodOutputValues:        noOfLastMethodOutputValues,
		LastMethodOutputValues:            lastMethodOutputValues,
		LastMethodCallTime:                lastMethodCallTime,
		LastMethodReturnStatus:            lastMethodReturnStatus,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ProgramDiagnostic2DataTypeBuilder is a builder for ProgramDiagnostic2DataType
type ProgramDiagnostic2DataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, noOfLastMethodInputArguments int32, lastMethodInputArguments []ExtensionObjectDefinition, noOfLastMethodOutputArguments int32, lastMethodOutputArguments []ExtensionObjectDefinition, noOfLastMethodInputValues int32, lastMethodInputValues []Variant, noOfLastMethodOutputValues int32, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder
	// WithCreateSessionId adds CreateSessionId (property field)
	WithCreateSessionId(NodeId) ProgramDiagnostic2DataTypeBuilder
	// WithCreateSessionIdBuilder adds CreateSessionId (property field) which is build by the builder
	WithCreateSessionIdBuilder(func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithCreateClientName adds CreateClientName (property field)
	WithCreateClientName(PascalString) ProgramDiagnostic2DataTypeBuilder
	// WithCreateClientNameBuilder adds CreateClientName (property field) which is build by the builder
	WithCreateClientNameBuilder(func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithInvocationCreationTime adds InvocationCreationTime (property field)
	WithInvocationCreationTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastTransitionTime adds LastTransitionTime (property field)
	WithLastTransitionTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCall adds LastMethodCall (property field)
	WithLastMethodCall(PascalString) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCallBuilder adds LastMethodCall (property field) which is build by the builder
	WithLastMethodCallBuilder(func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodSessionId adds LastMethodSessionId (property field)
	WithLastMethodSessionId(NodeId) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodSessionIdBuilder adds LastMethodSessionId (property field) which is build by the builder
	WithLastMethodSessionIdBuilder(func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithNoOfLastMethodInputArguments adds NoOfLastMethodInputArguments (property field)
	WithNoOfLastMethodInputArguments(int32) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodInputArguments adds LastMethodInputArguments (property field)
	WithLastMethodInputArguments(...ExtensionObjectDefinition) ProgramDiagnostic2DataTypeBuilder
	// WithNoOfLastMethodOutputArguments adds NoOfLastMethodOutputArguments (property field)
	WithNoOfLastMethodOutputArguments(int32) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodOutputArguments adds LastMethodOutputArguments (property field)
	WithLastMethodOutputArguments(...ExtensionObjectDefinition) ProgramDiagnostic2DataTypeBuilder
	// WithNoOfLastMethodInputValues adds NoOfLastMethodInputValues (property field)
	WithNoOfLastMethodInputValues(int32) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodInputValues adds LastMethodInputValues (property field)
	WithLastMethodInputValues(...Variant) ProgramDiagnostic2DataTypeBuilder
	// WithNoOfLastMethodOutputValues adds NoOfLastMethodOutputValues (property field)
	WithNoOfLastMethodOutputValues(int32) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodOutputValues adds LastMethodOutputValues (property field)
	WithLastMethodOutputValues(...Variant) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCallTime adds LastMethodCallTime (property field)
	WithLastMethodCallTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodReturnStatus adds LastMethodReturnStatus (property field)
	WithLastMethodReturnStatus(StatusCode) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodReturnStatusBuilder adds LastMethodReturnStatus (property field) which is build by the builder
	WithLastMethodReturnStatusBuilder(func(StatusCodeBuilder) StatusCodeBuilder) ProgramDiagnostic2DataTypeBuilder
	// Build builds the ProgramDiagnostic2DataType or returns an error if something is wrong
	Build() (ProgramDiagnostic2DataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ProgramDiagnostic2DataType
}

// NewProgramDiagnostic2DataTypeBuilder() creates a ProgramDiagnostic2DataTypeBuilder
func NewProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder {
	return &_ProgramDiagnostic2DataTypeBuilder{_ProgramDiagnostic2DataType: new(_ProgramDiagnostic2DataType)}
}

type _ProgramDiagnostic2DataTypeBuilder struct {
	*_ProgramDiagnostic2DataType

	err *utils.MultiError
}

var _ (ProgramDiagnostic2DataTypeBuilder) = (*_ProgramDiagnostic2DataTypeBuilder)(nil)

func (m *_ProgramDiagnostic2DataTypeBuilder) WithMandatoryFields(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, noOfLastMethodInputArguments int32, lastMethodInputArguments []ExtensionObjectDefinition, noOfLastMethodOutputArguments int32, lastMethodOutputArguments []ExtensionObjectDefinition, noOfLastMethodInputValues int32, lastMethodInputValues []Variant, noOfLastMethodOutputValues int32, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder {
	return m.WithCreateSessionId(createSessionId).WithCreateClientName(createClientName).WithInvocationCreationTime(invocationCreationTime).WithLastTransitionTime(lastTransitionTime).WithLastMethodCall(lastMethodCall).WithLastMethodSessionId(lastMethodSessionId).WithNoOfLastMethodInputArguments(noOfLastMethodInputArguments).WithLastMethodInputArguments(lastMethodInputArguments...).WithNoOfLastMethodOutputArguments(noOfLastMethodOutputArguments).WithLastMethodOutputArguments(lastMethodOutputArguments...).WithNoOfLastMethodInputValues(noOfLastMethodInputValues).WithLastMethodInputValues(lastMethodInputValues...).WithNoOfLastMethodOutputValues(noOfLastMethodOutputValues).WithLastMethodOutputValues(lastMethodOutputValues...).WithLastMethodCallTime(lastMethodCallTime).WithLastMethodReturnStatus(lastMethodReturnStatus)
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithCreateSessionId(createSessionId NodeId) ProgramDiagnostic2DataTypeBuilder {
	m.CreateSessionId = createSessionId
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithCreateSessionIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(m.CreateSessionId.CreateNodeIdBuilder())
	var err error
	m.CreateSessionId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithCreateClientName(createClientName PascalString) ProgramDiagnostic2DataTypeBuilder {
	m.CreateClientName = createClientName
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithCreateClientNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(m.CreateClientName.CreatePascalStringBuilder())
	var err error
	m.CreateClientName, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithInvocationCreationTime(invocationCreationTime int64) ProgramDiagnostic2DataTypeBuilder {
	m.InvocationCreationTime = invocationCreationTime
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastTransitionTime(lastTransitionTime int64) ProgramDiagnostic2DataTypeBuilder {
	m.LastTransitionTime = lastTransitionTime
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCall(lastMethodCall PascalString) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodCall = lastMethodCall
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCallBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(m.LastMethodCall.CreatePascalStringBuilder())
	var err error
	m.LastMethodCall, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodSessionId(lastMethodSessionId NodeId) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodSessionId = lastMethodSessionId
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodSessionIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(m.LastMethodSessionId.CreateNodeIdBuilder())
	var err error
	m.LastMethodSessionId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithNoOfLastMethodInputArguments(noOfLastMethodInputArguments int32) ProgramDiagnostic2DataTypeBuilder {
	m.NoOfLastMethodInputArguments = noOfLastMethodInputArguments
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodInputArguments(lastMethodInputArguments ...ExtensionObjectDefinition) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodInputArguments = lastMethodInputArguments
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithNoOfLastMethodOutputArguments(noOfLastMethodOutputArguments int32) ProgramDiagnostic2DataTypeBuilder {
	m.NoOfLastMethodOutputArguments = noOfLastMethodOutputArguments
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodOutputArguments(lastMethodOutputArguments ...ExtensionObjectDefinition) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodOutputArguments = lastMethodOutputArguments
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithNoOfLastMethodInputValues(noOfLastMethodInputValues int32) ProgramDiagnostic2DataTypeBuilder {
	m.NoOfLastMethodInputValues = noOfLastMethodInputValues
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodInputValues(lastMethodInputValues ...Variant) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodInputValues = lastMethodInputValues
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithNoOfLastMethodOutputValues(noOfLastMethodOutputValues int32) ProgramDiagnostic2DataTypeBuilder {
	m.NoOfLastMethodOutputValues = noOfLastMethodOutputValues
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodOutputValues(lastMethodOutputValues ...Variant) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodOutputValues = lastMethodOutputValues
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCallTime(lastMethodCallTime int64) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodCallTime = lastMethodCallTime
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodReturnStatus(lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder {
	m.LastMethodReturnStatus = lastMethodReturnStatus
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodReturnStatusBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(m.LastMethodReturnStatus.CreateStatusCodeBuilder())
	var err error
	m.LastMethodReturnStatus, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return m
}

func (m *_ProgramDiagnostic2DataTypeBuilder) Build() (ProgramDiagnostic2DataType, error) {
	if m.CreateSessionId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'createSessionId' not set"))
	}
	if m.CreateClientName == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'createClientName' not set"))
	}
	if m.LastMethodCall == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lastMethodCall' not set"))
	}
	if m.LastMethodSessionId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lastMethodSessionId' not set"))
	}
	if m.LastMethodReturnStatus == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lastMethodReturnStatus' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ProgramDiagnostic2DataType.deepCopy(), nil
}

func (m *_ProgramDiagnostic2DataTypeBuilder) MustBuild() ProgramDiagnostic2DataType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ProgramDiagnostic2DataTypeBuilder) DeepCopy() any {
	return m.CreateProgramDiagnostic2DataTypeBuilder()
}

// CreateProgramDiagnostic2DataTypeBuilder creates a ProgramDiagnostic2DataTypeBuilder
func (m *_ProgramDiagnostic2DataType) CreateProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder {
	if m == nil {
		return NewProgramDiagnostic2DataTypeBuilder()
	}
	return &_ProgramDiagnostic2DataTypeBuilder{_ProgramDiagnostic2DataType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ProgramDiagnostic2DataType) GetIdentifier() string {
	return "24035"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ProgramDiagnostic2DataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ProgramDiagnostic2DataType) GetCreateSessionId() NodeId {
	return m.CreateSessionId
}

func (m *_ProgramDiagnostic2DataType) GetCreateClientName() PascalString {
	return m.CreateClientName
}

func (m *_ProgramDiagnostic2DataType) GetInvocationCreationTime() int64 {
	return m.InvocationCreationTime
}

func (m *_ProgramDiagnostic2DataType) GetLastTransitionTime() int64 {
	return m.LastTransitionTime
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodCall() PascalString {
	return m.LastMethodCall
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodSessionId() NodeId {
	return m.LastMethodSessionId
}

func (m *_ProgramDiagnostic2DataType) GetNoOfLastMethodInputArguments() int32 {
	return m.NoOfLastMethodInputArguments
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodInputArguments() []ExtensionObjectDefinition {
	return m.LastMethodInputArguments
}

func (m *_ProgramDiagnostic2DataType) GetNoOfLastMethodOutputArguments() int32 {
	return m.NoOfLastMethodOutputArguments
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodOutputArguments() []ExtensionObjectDefinition {
	return m.LastMethodOutputArguments
}

func (m *_ProgramDiagnostic2DataType) GetNoOfLastMethodInputValues() int32 {
	return m.NoOfLastMethodInputValues
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodInputValues() []Variant {
	return m.LastMethodInputValues
}

func (m *_ProgramDiagnostic2DataType) GetNoOfLastMethodOutputValues() int32 {
	return m.NoOfLastMethodOutputValues
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodOutputValues() []Variant {
	return m.LastMethodOutputValues
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodCallTime() int64 {
	return m.LastMethodCallTime
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodReturnStatus() StatusCode {
	return m.LastMethodReturnStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastProgramDiagnostic2DataType(structType any) ProgramDiagnostic2DataType {
	if casted, ok := structType.(ProgramDiagnostic2DataType); ok {
		return casted
	}
	if casted, ok := structType.(*ProgramDiagnostic2DataType); ok {
		return *casted
	}
	return nil
}

func (m *_ProgramDiagnostic2DataType) GetTypeName() string {
	return "ProgramDiagnostic2DataType"
}

func (m *_ProgramDiagnostic2DataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (createSessionId)
	lengthInBits += m.CreateSessionId.GetLengthInBits(ctx)

	// Simple field (createClientName)
	lengthInBits += m.CreateClientName.GetLengthInBits(ctx)

	// Simple field (invocationCreationTime)
	lengthInBits += 64

	// Simple field (lastTransitionTime)
	lengthInBits += 64

	// Simple field (lastMethodCall)
	lengthInBits += m.LastMethodCall.GetLengthInBits(ctx)

	// Simple field (lastMethodSessionId)
	lengthInBits += m.LastMethodSessionId.GetLengthInBits(ctx)

	// Simple field (noOfLastMethodInputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodInputArguments) > 0 {
		for _curItem, element := range m.LastMethodInputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodInputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLastMethodOutputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodOutputArguments) > 0 {
		for _curItem, element := range m.LastMethodOutputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodOutputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLastMethodInputValues)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodInputValues) > 0 {
		for _curItem, element := range m.LastMethodInputValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodInputValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLastMethodOutputValues)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodOutputValues) > 0 {
		for _curItem, element := range m.LastMethodOutputValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodOutputValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (lastMethodCallTime)
	lengthInBits += 64

	// Simple field (lastMethodReturnStatus)
	lengthInBits += m.LastMethodReturnStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ProgramDiagnostic2DataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ProgramDiagnostic2DataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__programDiagnostic2DataType ProgramDiagnostic2DataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ProgramDiagnostic2DataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ProgramDiagnostic2DataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	createSessionId, err := ReadSimpleField[NodeId](ctx, "createSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createSessionId' field"))
	}
	m.CreateSessionId = createSessionId

	createClientName, err := ReadSimpleField[PascalString](ctx, "createClientName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createClientName' field"))
	}
	m.CreateClientName = createClientName

	invocationCreationTime, err := ReadSimpleField(ctx, "invocationCreationTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invocationCreationTime' field"))
	}
	m.InvocationCreationTime = invocationCreationTime

	lastTransitionTime, err := ReadSimpleField(ctx, "lastTransitionTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastTransitionTime' field"))
	}
	m.LastTransitionTime = lastTransitionTime

	lastMethodCall, err := ReadSimpleField[PascalString](ctx, "lastMethodCall", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCall' field"))
	}
	m.LastMethodCall = lastMethodCall

	lastMethodSessionId, err := ReadSimpleField[NodeId](ctx, "lastMethodSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodSessionId' field"))
	}
	m.LastMethodSessionId = lastMethodSessionId

	noOfLastMethodInputArguments, err := ReadSimpleField(ctx, "noOfLastMethodInputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodInputArguments' field"))
	}
	m.NoOfLastMethodInputArguments = noOfLastMethodInputArguments

	lastMethodInputArguments, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "lastMethodInputArguments", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("298")), readBuffer), uint64(noOfLastMethodInputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodInputArguments' field"))
	}
	m.LastMethodInputArguments = lastMethodInputArguments

	noOfLastMethodOutputArguments, err := ReadSimpleField(ctx, "noOfLastMethodOutputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodOutputArguments' field"))
	}
	m.NoOfLastMethodOutputArguments = noOfLastMethodOutputArguments

	lastMethodOutputArguments, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "lastMethodOutputArguments", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("298")), readBuffer), uint64(noOfLastMethodOutputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodOutputArguments' field"))
	}
	m.LastMethodOutputArguments = lastMethodOutputArguments

	noOfLastMethodInputValues, err := ReadSimpleField(ctx, "noOfLastMethodInputValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodInputValues' field"))
	}
	m.NoOfLastMethodInputValues = noOfLastMethodInputValues

	lastMethodInputValues, err := ReadCountArrayField[Variant](ctx, "lastMethodInputValues", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfLastMethodInputValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodInputValues' field"))
	}
	m.LastMethodInputValues = lastMethodInputValues

	noOfLastMethodOutputValues, err := ReadSimpleField(ctx, "noOfLastMethodOutputValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodOutputValues' field"))
	}
	m.NoOfLastMethodOutputValues = noOfLastMethodOutputValues

	lastMethodOutputValues, err := ReadCountArrayField[Variant](ctx, "lastMethodOutputValues", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfLastMethodOutputValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodOutputValues' field"))
	}
	m.LastMethodOutputValues = lastMethodOutputValues

	lastMethodCallTime, err := ReadSimpleField(ctx, "lastMethodCallTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCallTime' field"))
	}
	m.LastMethodCallTime = lastMethodCallTime

	lastMethodReturnStatus, err := ReadSimpleField[StatusCode](ctx, "lastMethodReturnStatus", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodReturnStatus' field"))
	}
	m.LastMethodReturnStatus = lastMethodReturnStatus

	if closeErr := readBuffer.CloseContext("ProgramDiagnostic2DataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ProgramDiagnostic2DataType")
	}

	return m, nil
}

func (m *_ProgramDiagnostic2DataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ProgramDiagnostic2DataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ProgramDiagnostic2DataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ProgramDiagnostic2DataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "createSessionId", m.GetCreateSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createSessionId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "createClientName", m.GetCreateClientName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createClientName' field")
		}

		if err := WriteSimpleField[int64](ctx, "invocationCreationTime", m.GetInvocationCreationTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'invocationCreationTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastTransitionTime", m.GetLastTransitionTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastTransitionTime' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "lastMethodCall", m.GetLastMethodCall(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCall' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "lastMethodSessionId", m.GetLastMethodSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodSessionId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodInputArguments", m.GetNoOfLastMethodInputArguments(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodInputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodInputArguments", m.GetLastMethodInputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodInputArguments' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodOutputArguments", m.GetNoOfLastMethodOutputArguments(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodOutputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodOutputArguments", m.GetLastMethodOutputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodOutputArguments' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodInputValues", m.GetNoOfLastMethodInputValues(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodInputValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodInputValues", m.GetLastMethodInputValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodInputValues' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodOutputValues", m.GetNoOfLastMethodOutputValues(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodOutputValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodOutputValues", m.GetLastMethodOutputValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodOutputValues' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastMethodCallTime", m.GetLastMethodCallTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCallTime' field")
		}

		if err := WriteSimpleField[StatusCode](ctx, "lastMethodReturnStatus", m.GetLastMethodReturnStatus(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodReturnStatus' field")
		}

		if popErr := writeBuffer.PopContext("ProgramDiagnostic2DataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ProgramDiagnostic2DataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ProgramDiagnostic2DataType) IsProgramDiagnostic2DataType() {}

func (m *_ProgramDiagnostic2DataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ProgramDiagnostic2DataType) deepCopy() *_ProgramDiagnostic2DataType {
	if m == nil {
		return nil
	}
	_ProgramDiagnostic2DataTypeCopy := &_ProgramDiagnostic2DataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.CreateSessionId.DeepCopy().(NodeId),
		m.CreateClientName.DeepCopy().(PascalString),
		m.InvocationCreationTime,
		m.LastTransitionTime,
		m.LastMethodCall.DeepCopy().(PascalString),
		m.LastMethodSessionId.DeepCopy().(NodeId),
		m.NoOfLastMethodInputArguments,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.LastMethodInputArguments),
		m.NoOfLastMethodOutputArguments,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.LastMethodOutputArguments),
		m.NoOfLastMethodInputValues,
		utils.DeepCopySlice[Variant, Variant](m.LastMethodInputValues),
		m.NoOfLastMethodOutputValues,
		utils.DeepCopySlice[Variant, Variant](m.LastMethodOutputValues),
		m.LastMethodCallTime,
		m.LastMethodReturnStatus.DeepCopy().(StatusCode),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ProgramDiagnostic2DataTypeCopy
}

func (m *_ProgramDiagnostic2DataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
