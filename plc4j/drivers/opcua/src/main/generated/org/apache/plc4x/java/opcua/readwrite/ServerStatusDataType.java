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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ServerStatusDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 864;
  }

  // Properties.
  protected final long startTime;
  protected final long currentTime;
  protected final ServerState state;
  protected final BuildInfo buildInfo;
  protected final long secondsTillShutdown;
  protected final LocalizedText shutdownReason;

  public ServerStatusDataType(
      long startTime,
      long currentTime,
      ServerState state,
      BuildInfo buildInfo,
      long secondsTillShutdown,
      LocalizedText shutdownReason) {
    super();
    this.startTime = startTime;
    this.currentTime = currentTime;
    this.state = state;
    this.buildInfo = buildInfo;
    this.secondsTillShutdown = secondsTillShutdown;
    this.shutdownReason = shutdownReason;
  }

  public long getStartTime() {
    return startTime;
  }

  public long getCurrentTime() {
    return currentTime;
  }

  public ServerState getState() {
    return state;
  }

  public BuildInfo getBuildInfo() {
    return buildInfo;
  }

  public long getSecondsTillShutdown() {
    return secondsTillShutdown;
  }

  public LocalizedText getShutdownReason() {
    return shutdownReason;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ServerStatusDataType");

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (currentTime)
    writeSimpleField("currentTime", currentTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (state)
    writeSimpleEnumField(
        "state",
        "ServerState",
        state,
        writeEnum(ServerState::getValue, ServerState::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (buildInfo)
    writeSimpleField("buildInfo", buildInfo, writeComplex(writeBuffer));

    // Simple Field (secondsTillShutdown)
    writeSimpleField(
        "secondsTillShutdown", secondsTillShutdown, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (shutdownReason)
    writeSimpleField("shutdownReason", shutdownReason, writeComplex(writeBuffer));

    writeBuffer.popContext("ServerStatusDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ServerStatusDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (startTime)
    lengthInBits += 64;

    // Simple field (currentTime)
    lengthInBits += 64;

    // Simple field (state)
    lengthInBits += 32;

    // Simple field (buildInfo)
    lengthInBits += buildInfo.getLengthInBits();

    // Simple field (secondsTillShutdown)
    lengthInBits += 32;

    // Simple field (shutdownReason)
    lengthInBits += shutdownReason.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("ServerStatusDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long startTime = readSimpleField("startTime", readSignedLong(readBuffer, 64));

    long currentTime = readSimpleField("currentTime", readSignedLong(readBuffer, 64));

    ServerState state =
        readEnumField(
            "state",
            "ServerState",
            readEnum(ServerState::enumForValue, readUnsignedLong(readBuffer, 32)));

    BuildInfo buildInfo =
        readSimpleField(
            "buildInfo",
            readComplex(
                () -> (BuildInfo) ExtensionObjectDefinition.staticParse(readBuffer, (int) (340)),
                readBuffer));

    long secondsTillShutdown =
        readSimpleField("secondsTillShutdown", readUnsignedLong(readBuffer, 32));

    LocalizedText shutdownReason =
        readSimpleField(
            "shutdownReason", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ServerStatusDataType");
    // Create the instance
    return new ServerStatusDataTypeBuilderImpl(
        startTime, currentTime, state, buildInfo, secondsTillShutdown, shutdownReason);
  }

  public static class ServerStatusDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long startTime;
    private final long currentTime;
    private final ServerState state;
    private final BuildInfo buildInfo;
    private final long secondsTillShutdown;
    private final LocalizedText shutdownReason;

    public ServerStatusDataTypeBuilderImpl(
        long startTime,
        long currentTime,
        ServerState state,
        BuildInfo buildInfo,
        long secondsTillShutdown,
        LocalizedText shutdownReason) {
      this.startTime = startTime;
      this.currentTime = currentTime;
      this.state = state;
      this.buildInfo = buildInfo;
      this.secondsTillShutdown = secondsTillShutdown;
      this.shutdownReason = shutdownReason;
    }

    public ServerStatusDataType build() {
      ServerStatusDataType serverStatusDataType =
          new ServerStatusDataType(
              startTime, currentTime, state, buildInfo, secondsTillShutdown, shutdownReason);
      return serverStatusDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ServerStatusDataType)) {
      return false;
    }
    ServerStatusDataType that = (ServerStatusDataType) o;
    return (getStartTime() == that.getStartTime())
        && (getCurrentTime() == that.getCurrentTime())
        && (getState() == that.getState())
        && (getBuildInfo() == that.getBuildInfo())
        && (getSecondsTillShutdown() == that.getSecondsTillShutdown())
        && (getShutdownReason() == that.getShutdownReason())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStartTime(),
        getCurrentTime(),
        getState(),
        getBuildInfo(),
        getSecondsTillShutdown(),
        getShutdownReason());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
