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

// IdentifyReplyCommandExtendedDiagnosticSummary is the corresponding interface of IdentifyReplyCommandExtendedDiagnosticSummary
type IdentifyReplyCommandExtendedDiagnosticSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetLowApplication returns LowApplication (property field)
	GetLowApplication() ApplicationIdContainer
	// GetHighApplication returns HighApplication (property field)
	GetHighApplication() ApplicationIdContainer
	// GetArea returns Area (property field)
	GetArea() byte
	// GetCrc returns Crc (property field)
	GetCrc() uint16
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() uint32
	// GetNetworkVoltage returns NetworkVoltage (property field)
	GetNetworkVoltage() byte
	// GetUnitInLearnMode returns UnitInLearnMode (property field)
	GetUnitInLearnMode() bool
	// GetNetworkVoltageLow returns NetworkVoltageLow (property field)
	GetNetworkVoltageLow() bool
	// GetNetworkVoltageMarginal returns NetworkVoltageMarginal (property field)
	GetNetworkVoltageMarginal() bool
	// GetEnableChecksumAlarm returns EnableChecksumAlarm (property field)
	GetEnableChecksumAlarm() bool
	// GetOutputUnit returns OutputUnit (property field)
	GetOutputUnit() bool
	// GetInstallationMMIError returns InstallationMMIError (property field)
	GetInstallationMMIError() bool
	// GetEEWriteError returns EEWriteError (property field)
	GetEEWriteError() bool
	// GetEEChecksumError returns EEChecksumError (property field)
	GetEEChecksumError() bool
	// GetEEDataError returns EEDataError (property field)
	GetEEDataError() bool
	// GetMicroReset returns MicroReset (property field)
	GetMicroReset() bool
	// GetCommsTxError returns CommsTxError (property field)
	GetCommsTxError() bool
	// GetInternalStackOverflow returns InternalStackOverflow (property field)
	GetInternalStackOverflow() bool
	// GetMicroPowerReset returns MicroPowerReset (property field)
	GetMicroPowerReset() bool
	// GetNetworkVoltageInVolts returns NetworkVoltageInVolts (virtual field)
	GetNetworkVoltageInVolts() float32
	// IsIdentifyReplyCommandExtendedDiagnosticSummary is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandExtendedDiagnosticSummary()
	// CreateBuilder creates a IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	CreateIdentifyReplyCommandExtendedDiagnosticSummaryBuilder() IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
}

// _IdentifyReplyCommandExtendedDiagnosticSummary is the data-structure of this message
type _IdentifyReplyCommandExtendedDiagnosticSummary struct {
	IdentifyReplyCommandContract
	LowApplication         ApplicationIdContainer
	HighApplication        ApplicationIdContainer
	Area                   byte
	Crc                    uint16
	SerialNumber           uint32
	NetworkVoltage         byte
	UnitInLearnMode        bool
	NetworkVoltageLow      bool
	NetworkVoltageMarginal bool
	EnableChecksumAlarm    bool
	OutputUnit             bool
	InstallationMMIError   bool
	EEWriteError           bool
	EEChecksumError        bool
	EEDataError            bool
	MicroReset             bool
	CommsTxError           bool
	InternalStackOverflow  bool
	MicroPowerReset        bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
	reservedField2 *uint8
}

var _ IdentifyReplyCommandExtendedDiagnosticSummary = (*_IdentifyReplyCommandExtendedDiagnosticSummary)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandExtendedDiagnosticSummary)(nil)

// NewIdentifyReplyCommandExtendedDiagnosticSummary factory function for _IdentifyReplyCommandExtendedDiagnosticSummary
func NewIdentifyReplyCommandExtendedDiagnosticSummary(lowApplication ApplicationIdContainer, highApplication ApplicationIdContainer, area byte, crc uint16, serialNumber uint32, networkVoltage byte, unitInLearnMode bool, networkVoltageLow bool, networkVoltageMarginal bool, enableChecksumAlarm bool, outputUnit bool, installationMMIError bool, EEWriteError bool, EEChecksumError bool, EEDataError bool, microReset bool, commsTxError bool, internalStackOverflow bool, microPowerReset bool, numBytes uint8) *_IdentifyReplyCommandExtendedDiagnosticSummary {
	_result := &_IdentifyReplyCommandExtendedDiagnosticSummary{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		LowApplication:               lowApplication,
		HighApplication:              highApplication,
		Area:                         area,
		Crc:                          crc,
		SerialNumber:                 serialNumber,
		NetworkVoltage:               networkVoltage,
		UnitInLearnMode:              unitInLearnMode,
		NetworkVoltageLow:            networkVoltageLow,
		NetworkVoltageMarginal:       networkVoltageMarginal,
		EnableChecksumAlarm:          enableChecksumAlarm,
		OutputUnit:                   outputUnit,
		InstallationMMIError:         installationMMIError,
		EEWriteError:                 EEWriteError,
		EEChecksumError:              EEChecksumError,
		EEDataError:                  EEDataError,
		MicroReset:                   microReset,
		CommsTxError:                 commsTxError,
		InternalStackOverflow:        internalStackOverflow,
		MicroPowerReset:              microPowerReset,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandExtendedDiagnosticSummaryBuilder is a builder for IdentifyReplyCommandExtendedDiagnosticSummary
type IdentifyReplyCommandExtendedDiagnosticSummaryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lowApplication ApplicationIdContainer, highApplication ApplicationIdContainer, area byte, crc uint16, serialNumber uint32, networkVoltage byte, unitInLearnMode bool, networkVoltageLow bool, networkVoltageMarginal bool, enableChecksumAlarm bool, outputUnit bool, installationMMIError bool, EEWriteError bool, EEChecksumError bool, EEDataError bool, microReset bool, commsTxError bool, internalStackOverflow bool, microPowerReset bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithLowApplication adds LowApplication (property field)
	WithLowApplication(ApplicationIdContainer) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithHighApplication adds HighApplication (property field)
	WithHighApplication(ApplicationIdContainer) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithArea adds Area (property field)
	WithArea(byte) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithCrc adds Crc (property field)
	WithCrc(uint16) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithSerialNumber adds SerialNumber (property field)
	WithSerialNumber(uint32) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithNetworkVoltage adds NetworkVoltage (property field)
	WithNetworkVoltage(byte) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithUnitInLearnMode adds UnitInLearnMode (property field)
	WithUnitInLearnMode(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithNetworkVoltageLow adds NetworkVoltageLow (property field)
	WithNetworkVoltageLow(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithNetworkVoltageMarginal adds NetworkVoltageMarginal (property field)
	WithNetworkVoltageMarginal(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithEnableChecksumAlarm adds EnableChecksumAlarm (property field)
	WithEnableChecksumAlarm(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithOutputUnit adds OutputUnit (property field)
	WithOutputUnit(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithInstallationMMIError adds InstallationMMIError (property field)
	WithInstallationMMIError(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithEEWriteError adds EEWriteError (property field)
	WithEEWriteError(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithEEChecksumError adds EEChecksumError (property field)
	WithEEChecksumError(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithEEDataError adds EEDataError (property field)
	WithEEDataError(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithMicroReset adds MicroReset (property field)
	WithMicroReset(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithCommsTxError adds CommsTxError (property field)
	WithCommsTxError(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithInternalStackOverflow adds InternalStackOverflow (property field)
	WithInternalStackOverflow(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// WithMicroPowerReset adds MicroPowerReset (property field)
	WithMicroPowerReset(bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
	// Build builds the IdentifyReplyCommandExtendedDiagnosticSummary or returns an error if something is wrong
	Build() (IdentifyReplyCommandExtendedDiagnosticSummary, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandExtendedDiagnosticSummary
}

// NewIdentifyReplyCommandExtendedDiagnosticSummaryBuilder() creates a IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
func NewIdentifyReplyCommandExtendedDiagnosticSummaryBuilder() IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	return &_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder{_IdentifyReplyCommandExtendedDiagnosticSummary: new(_IdentifyReplyCommandExtendedDiagnosticSummary)}
}

type _IdentifyReplyCommandExtendedDiagnosticSummaryBuilder struct {
	*_IdentifyReplyCommandExtendedDiagnosticSummary

	err *utils.MultiError
}

var _ (IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) = (*_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder)(nil)

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithMandatoryFields(lowApplication ApplicationIdContainer, highApplication ApplicationIdContainer, area byte, crc uint16, serialNumber uint32, networkVoltage byte, unitInLearnMode bool, networkVoltageLow bool, networkVoltageMarginal bool, enableChecksumAlarm bool, outputUnit bool, installationMMIError bool, EEWriteError bool, EEChecksumError bool, EEDataError bool, microReset bool, commsTxError bool, internalStackOverflow bool, microPowerReset bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	return m.WithLowApplication(lowApplication).WithHighApplication(highApplication).WithArea(area).WithCrc(crc).WithSerialNumber(serialNumber).WithNetworkVoltage(networkVoltage).WithUnitInLearnMode(unitInLearnMode).WithNetworkVoltageLow(networkVoltageLow).WithNetworkVoltageMarginal(networkVoltageMarginal).WithEnableChecksumAlarm(enableChecksumAlarm).WithOutputUnit(outputUnit).WithInstallationMMIError(installationMMIError).WithEEWriteError(EEWriteError).WithEEChecksumError(EEChecksumError).WithEEDataError(EEDataError).WithMicroReset(microReset).WithCommsTxError(commsTxError).WithInternalStackOverflow(internalStackOverflow).WithMicroPowerReset(microPowerReset)
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithLowApplication(lowApplication ApplicationIdContainer) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.LowApplication = lowApplication
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithHighApplication(highApplication ApplicationIdContainer) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.HighApplication = highApplication
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithArea(area byte) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.Area = area
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithCrc(crc uint16) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.Crc = crc
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithSerialNumber(serialNumber uint32) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.SerialNumber = serialNumber
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithNetworkVoltage(networkVoltage byte) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.NetworkVoltage = networkVoltage
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithUnitInLearnMode(unitInLearnMode bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.UnitInLearnMode = unitInLearnMode
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithNetworkVoltageLow(networkVoltageLow bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.NetworkVoltageLow = networkVoltageLow
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithNetworkVoltageMarginal(networkVoltageMarginal bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.NetworkVoltageMarginal = networkVoltageMarginal
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithEnableChecksumAlarm(enableChecksumAlarm bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.EnableChecksumAlarm = enableChecksumAlarm
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithOutputUnit(outputUnit bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.OutputUnit = outputUnit
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithInstallationMMIError(installationMMIError bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.InstallationMMIError = installationMMIError
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithEEWriteError(EEWriteError bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.EEWriteError = EEWriteError
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithEEChecksumError(EEChecksumError bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.EEChecksumError = EEChecksumError
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithEEDataError(EEDataError bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.EEDataError = EEDataError
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithMicroReset(microReset bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.MicroReset = microReset
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithCommsTxError(commsTxError bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.CommsTxError = commsTxError
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithInternalStackOverflow(internalStackOverflow bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.InternalStackOverflow = internalStackOverflow
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) WithMicroPowerReset(microPowerReset bool) IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	m.MicroPowerReset = microPowerReset
	return m
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) Build() (IdentifyReplyCommandExtendedDiagnosticSummary, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._IdentifyReplyCommandExtendedDiagnosticSummary.deepCopy(), nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) MustBuild() IdentifyReplyCommandExtendedDiagnosticSummary {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder) DeepCopy() any {
	return m.CreateIdentifyReplyCommandExtendedDiagnosticSummaryBuilder()
}

// CreateIdentifyReplyCommandExtendedDiagnosticSummaryBuilder creates a IdentifyReplyCommandExtendedDiagnosticSummaryBuilder
func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) CreateIdentifyReplyCommandExtendedDiagnosticSummaryBuilder() IdentifyReplyCommandExtendedDiagnosticSummaryBuilder {
	if m == nil {
		return NewIdentifyReplyCommandExtendedDiagnosticSummaryBuilder()
	}
	return &_IdentifyReplyCommandExtendedDiagnosticSummaryBuilder{_IdentifyReplyCommandExtendedDiagnosticSummary: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetAttribute() Attribute {
	return Attribute_ExtendedDiagnosticSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLowApplication() ApplicationIdContainer {
	return m.LowApplication
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetHighApplication() ApplicationIdContainer {
	return m.HighApplication
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetArea() byte {
	return m.Area
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetCrc() uint16 {
	return m.Crc
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetSerialNumber() uint32 {
	return m.SerialNumber
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltage() byte {
	return m.NetworkVoltage
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetUnitInLearnMode() bool {
	return m.UnitInLearnMode
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageLow() bool {
	return m.NetworkVoltageLow
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageMarginal() bool {
	return m.NetworkVoltageMarginal
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEnableChecksumAlarm() bool {
	return m.EnableChecksumAlarm
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetOutputUnit() bool {
	return m.OutputUnit
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetInstallationMMIError() bool {
	return m.InstallationMMIError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEWriteError() bool {
	return m.EEWriteError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEChecksumError() bool {
	return m.EEChecksumError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEDataError() bool {
	return m.EEDataError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroReset() bool {
	return m.MicroReset
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetCommsTxError() bool {
	return m.CommsTxError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetInternalStackOverflow() bool {
	return m.InternalStackOverflow
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroPowerReset() bool {
	return m.MicroPowerReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageInVolts() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetNetworkVoltage()) / float32(float32(6.375)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandExtendedDiagnosticSummary(structType any) IdentifyReplyCommandExtendedDiagnosticSummary {
	if casted, ok := structType.(IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetTypeName() string {
	return "IdentifyReplyCommandExtendedDiagnosticSummary"
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (lowApplication)
	lengthInBits += 8

	// Simple field (highApplication)
	lengthInBits += 8

	// Simple field (area)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	// Simple field (serialNumber)
	lengthInBits += 32

	// Simple field (networkVoltage)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (unitInLearnMode)
	lengthInBits += 1

	// Simple field (networkVoltageLow)
	lengthInBits += 1

	// Simple field (networkVoltageMarginal)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (enableChecksumAlarm)
	lengthInBits += 1

	// Simple field (outputUnit)
	lengthInBits += 1

	// Simple field (installationMMIError)
	lengthInBits += 1

	// Simple field (EEWriteError)
	lengthInBits += 1

	// Simple field (EEChecksumError)
	lengthInBits += 1

	// Simple field (EEDataError)
	lengthInBits += 1

	// Simple field (microReset)
	lengthInBits += 1

	// Simple field (commsTxError)
	lengthInBits += 1

	// Simple field (internalStackOverflow)
	lengthInBits += 1

	// Simple field (microPowerReset)
	lengthInBits += 1

	return lengthInBits
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandExtendedDiagnosticSummary IdentifyReplyCommandExtendedDiagnosticSummary, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowApplication, err := ReadEnumField[ApplicationIdContainer](ctx, "lowApplication", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowApplication' field"))
	}
	m.LowApplication = lowApplication

	highApplication, err := ReadEnumField[ApplicationIdContainer](ctx, "highApplication", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highApplication' field"))
	}
	m.HighApplication = highApplication

	area, err := ReadSimpleField(ctx, "area", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'area' field"))
	}
	m.Area = area

	crc, err := ReadSimpleField(ctx, "crc", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	m.Crc = crc

	serialNumber, err := ReadSimpleField(ctx, "serialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serialNumber' field"))
	}
	m.SerialNumber = serialNumber

	networkVoltage, err := ReadSimpleField(ctx, "networkVoltage", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkVoltage' field"))
	}
	m.NetworkVoltage = networkVoltage

	networkVoltageInVolts, err := ReadVirtualField[float32](ctx, "networkVoltageInVolts", (*float32)(nil), float32(networkVoltage)/float32(float32(6.375)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkVoltageInVolts' field"))
	}
	_ = networkVoltageInVolts

	unitInLearnMode, err := ReadSimpleField(ctx, "unitInLearnMode", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitInLearnMode' field"))
	}
	m.UnitInLearnMode = unitInLearnMode

	networkVoltageLow, err := ReadSimpleField(ctx, "networkVoltageLow", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkVoltageLow' field"))
	}
	m.NetworkVoltageLow = networkVoltageLow

	networkVoltageMarginal, err := ReadSimpleField(ctx, "networkVoltageMarginal", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkVoltageMarginal' field"))
	}
	m.NetworkVoltageMarginal = networkVoltageMarginal

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(1)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(1)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(1)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField2 = reservedField2

	enableChecksumAlarm, err := ReadSimpleField(ctx, "enableChecksumAlarm", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableChecksumAlarm' field"))
	}
	m.EnableChecksumAlarm = enableChecksumAlarm

	outputUnit, err := ReadSimpleField(ctx, "outputUnit", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'outputUnit' field"))
	}
	m.OutputUnit = outputUnit

	installationMMIError, err := ReadSimpleField(ctx, "installationMMIError", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'installationMMIError' field"))
	}
	m.InstallationMMIError = installationMMIError

	EEWriteError, err := ReadSimpleField(ctx, "EEWriteError", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'EEWriteError' field"))
	}
	m.EEWriteError = EEWriteError

	EEChecksumError, err := ReadSimpleField(ctx, "EEChecksumError", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'EEChecksumError' field"))
	}
	m.EEChecksumError = EEChecksumError

	EEDataError, err := ReadSimpleField(ctx, "EEDataError", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'EEDataError' field"))
	}
	m.EEDataError = EEDataError

	microReset, err := ReadSimpleField(ctx, "microReset", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'microReset' field"))
	}
	m.MicroReset = microReset

	commsTxError, err := ReadSimpleField(ctx, "commsTxError", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commsTxError' field"))
	}
	m.CommsTxError = commsTxError

	internalStackOverflow, err := ReadSimpleField(ctx, "internalStackOverflow", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'internalStackOverflow' field"))
	}
	m.InternalStackOverflow = internalStackOverflow

	microPowerReset, err := ReadSimpleField(ctx, "microPowerReset", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'microPowerReset' field"))
	}
	m.MicroPowerReset = microPowerReset

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandExtendedDiagnosticSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandExtendedDiagnosticSummary")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandExtendedDiagnosticSummary")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "lowApplication", "ApplicationIdContainer", m.GetLowApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'lowApplication' field")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "highApplication", "ApplicationIdContainer", m.GetHighApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'highApplication' field")
		}

		if err := WriteSimpleField[byte](ctx, "area", m.GetArea(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'area' field")
		}

		if err := WriteSimpleField[uint16](ctx, "crc", m.GetCrc(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if err := WriteSimpleField[uint32](ctx, "serialNumber", m.GetSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'serialNumber' field")
		}

		if err := WriteSimpleField[byte](ctx, "networkVoltage", m.GetNetworkVoltage(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkVoltage' field")
		}
		// Virtual field
		networkVoltageInVolts := m.GetNetworkVoltageInVolts()
		_ = networkVoltageInVolts
		if _networkVoltageInVoltsErr := writeBuffer.WriteVirtual(ctx, "networkVoltageInVolts", m.GetNetworkVoltageInVolts()); _networkVoltageInVoltsErr != nil {
			return errors.Wrap(_networkVoltageInVoltsErr, "Error serializing 'networkVoltageInVolts' field")
		}

		if err := WriteSimpleField[bool](ctx, "unitInLearnMode", m.GetUnitInLearnMode(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitInLearnMode' field")
		}

		if err := WriteSimpleField[bool](ctx, "networkVoltageLow", m.GetNetworkVoltageLow(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkVoltageLow' field")
		}

		if err := WriteSimpleField[bool](ctx, "networkVoltageMarginal", m.GetNetworkVoltageMarginal(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkVoltageMarginal' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 1)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 1)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 1)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 3")
		}

		if err := WriteSimpleField[bool](ctx, "enableChecksumAlarm", m.GetEnableChecksumAlarm(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enableChecksumAlarm' field")
		}

		if err := WriteSimpleField[bool](ctx, "outputUnit", m.GetOutputUnit(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'outputUnit' field")
		}

		if err := WriteSimpleField[bool](ctx, "installationMMIError", m.GetInstallationMMIError(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'installationMMIError' field")
		}

		if err := WriteSimpleField[bool](ctx, "EEWriteError", m.GetEEWriteError(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'EEWriteError' field")
		}

		if err := WriteSimpleField[bool](ctx, "EEChecksumError", m.GetEEChecksumError(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'EEChecksumError' field")
		}

		if err := WriteSimpleField[bool](ctx, "EEDataError", m.GetEEDataError(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'EEDataError' field")
		}

		if err := WriteSimpleField[bool](ctx, "microReset", m.GetMicroReset(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'microReset' field")
		}

		if err := WriteSimpleField[bool](ctx, "commsTxError", m.GetCommsTxError(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'commsTxError' field")
		}

		if err := WriteSimpleField[bool](ctx, "internalStackOverflow", m.GetInternalStackOverflow(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'internalStackOverflow' field")
		}

		if err := WriteSimpleField[bool](ctx, "microPowerReset", m.GetMicroPowerReset(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'microPowerReset' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandExtendedDiagnosticSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandExtendedDiagnosticSummary")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) IsIdentifyReplyCommandExtendedDiagnosticSummary() {
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) deepCopy() *_IdentifyReplyCommandExtendedDiagnosticSummary {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandExtendedDiagnosticSummaryCopy := &_IdentifyReplyCommandExtendedDiagnosticSummary{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		m.LowApplication,
		m.HighApplication,
		m.Area,
		m.Crc,
		m.SerialNumber,
		m.NetworkVoltage,
		m.UnitInLearnMode,
		m.NetworkVoltageLow,
		m.NetworkVoltageMarginal,
		m.EnableChecksumAlarm,
		m.OutputUnit,
		m.InstallationMMIError,
		m.EEWriteError,
		m.EEChecksumError,
		m.EEDataError,
		m.MicroReset,
		m.CommsTxError,
		m.InternalStackOverflow,
		m.MicroPowerReset,
		m.reservedField0,
		m.reservedField1,
		m.reservedField2,
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandExtendedDiagnosticSummaryCopy
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
