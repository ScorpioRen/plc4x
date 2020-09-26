/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "s7_payload_user_data_item.h"

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants so we can use the
// enum constant to directly access a given types discriminator values)
const plc4c_s7_read_write_s7_payload_user_data_item_discriminator plc4c_s7_read_write_s7_payload_user_data_item_discriminators[] = {
  {/* plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_request */
   .cpuFunctionType = 0x04},
  {/* plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_response */
   .cpuFunctionType = 0x08}
};

// Function returning the discriminator values for a given type constant.
plc4c_s7_read_write_s7_payload_user_data_item_discriminator plc4c_s7_read_write_s7_payload_user_data_item_get_discriminator(plc4c_s7_read_write_s7_payload_user_data_item_type type) {
  return plc4c_s7_read_write_s7_payload_user_data_item_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_s7_read_write_s7_payload_user_data_item plc4c_s7_read_write_s7_payload_user_data_item_null_const;

plc4c_s7_read_write_s7_payload_user_data_item plc4c_s7_read_write_s7_payload_user_data_item_null() {
  return plc4c_s7_read_write_s7_payload_user_data_item_null_const;
}


// Constant values.
static const uint16_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH_const = 28;
uint16_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH() {
  return PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_s7_payload_user_data_item_parse(plc4c_spi_read_buffer* io, unsigned int cpuFunctionType, plc4c_s7_read_write_s7_payload_user_data_item** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(io);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_s7_payload_user_data_item));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Enum field (returnCode)
  plc4c_s7_read_write_data_transport_error_code returnCode = plc4c_s7_read_write_data_transport_error_code_null();
  _res = plc4c_spi_read_signed_byte(io, 8, (int8_t*) &returnCode);
  if(_res != OK) {
    return _res;
  }
  (*_message)->return_code = returnCode;

  // Enum field (transportSize)
  plc4c_s7_read_write_data_transport_size transportSize = plc4c_s7_read_write_data_transport_size_null();
  _res = plc4c_spi_read_signed_byte(io, 8, (int8_t*) &transportSize);
  if(_res != OK) {
    return _res;
  }
  (*_message)->transport_size = transportSize;

  // Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint16_t dataLength = 0;
  _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &dataLength);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (szlId)
  plc4c_s7_read_write_szl_id* szlId;
  _res = plc4c_s7_read_write_szl_id_parse(io, (void*) &szlId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->szl_id = szlId;

  // Simple Field (szlIndex)
  uint16_t szlIndex = 0;
  _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &szlIndex);
  if(_res != OK) {
    return _res;
  }
  (*_message)->szl_index = szlIndex;

  // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
  if(cpuFunctionType == 0x04) { /* S7PayloadUserDataItemCpuFunctionReadSzlRequest */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_request;
  } else 
  if(cpuFunctionType == 0x08) { /* S7PayloadUserDataItemCpuFunctionReadSzlResponse */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_response;
                    
    // Const Field (szlItemLength)
    uint16_t szlItemLength = 0;
    _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &szlItemLength);
    if(_res != OK) {
      return _res;
    }
    if(szlItemLength != PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH()) {
      return PARSE_ERROR;
      // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH + " but got " + szlItemLength);
    }


                    
    // Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    uint16_t szlItemCount = 0;
    _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &szlItemCount);
    if(_res != OK) {
      return _res;
    }


                    
    // Array field (items)
    plc4c_list* items = NULL;
    plc4c_utils_list_create(&items);
    if(items == NULL) {
      return NO_MEMORY;
    }
    {
      // Count array
      uint8_t itemCount = szlItemCount;
      for(int curItem = 0; curItem < itemCount; curItem++) {
        bool lastItem = curItem == (itemCount - 1);
        plc4c_s7_read_write_szl_data_tree_item* _value = NULL;
        _res = plc4c_s7_read_write_szl_data_tree_item_parse(io, (void*) &_value);
        if(_res != OK) {
          return _res;
        }
        plc4c_utils_list_insert_head_value(items, _value);
      }
    }
    (*_message)->s7_payload_user_data_item_cpu_function_read_szl_response_items = items;

  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_s7_payload_user_data_item_serialize(plc4c_spi_write_buffer* io, plc4c_s7_read_write_s7_payload_user_data_item* _message) {
  plc4c_return_code _res = OK;

  // Enum field (returnCode)
  _res = plc4c_spi_write_signed_byte(io, 8, _message->return_code);
  if(_res != OK) {
    return _res;
  }

  // Enum field (transportSize)
  _res = plc4c_spi_write_signed_byte(io, 8, _message->transport_size);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_short(io, 16, (plc4c_s7_read_write_s7_payload_user_data_item_length_in_bytes(_message)) - (4));
  if(_res != OK) {
    return _res;
  }

  // Simple Field (szlId)
  _res = plc4c_s7_read_write_szl_id_serialize(io, _message->szl_id);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (szlIndex)
  _res = plc4c_spi_write_unsigned_short(io, 16, _message->szl_index);
  if(_res != OK) {
    return _res;
  }

  // Switch Field (Depending of the current type, serialize the sub-type elements)
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_request: {

      break;
    }
    case plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_response: {

      // Const Field (szlItemLength)
      plc4c_spi_write_unsigned_short(io, 16, PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH());

      // Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
      _res = plc4c_spi_write_unsigned_short(io, 16, plc4c_spi_evaluation_helper_count(_message->s7_payload_user_data_item_cpu_function_read_szl_response_items));
      if(_res != OK) {
        return _res;
      }

      // Array field (items)
      {
        uint8_t itemCount = plc4c_utils_list_size(_message->s7_payload_user_data_item_cpu_function_read_szl_response_items);
        for(int curItem = 0; curItem < itemCount; curItem++) {
          bool lastItem = curItem == (itemCount - 1);
          plc4c_s7_read_write_szl_data_tree_item* _value = (plc4c_s7_read_write_szl_data_tree_item*) plc4c_utils_list_get_value(_message->s7_payload_user_data_item_cpu_function_read_szl_response_items, curItem);
          _res = plc4c_s7_read_write_szl_data_tree_item_serialize(io, (void*) _value);
          if(_res != OK) {
            return _res;
          }
        }
      }

      break;
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_s7_payload_user_data_item_length_in_bytes(plc4c_s7_read_write_s7_payload_user_data_item* _message) {
  return plc4c_s7_read_write_s7_payload_user_data_item_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_s7_payload_user_data_item_length_in_bits(plc4c_s7_read_write_s7_payload_user_data_item* _message) {
  uint16_t lengthInBits = 0;

  // Enum Field (returnCode)
  lengthInBits += 8;

  // Enum Field (transportSize)
  lengthInBits += 8;

  // Implicit Field (dataLength)
  lengthInBits += 16;

  // Simple field (szlId)
  lengthInBits += plc4c_s7_read_write_szl_id_length_in_bits(_message->szl_id);

  // Simple field (szlIndex)
  lengthInBits += 16;

  // Depending of the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_request: {

      break;
    }
    case plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_response: {

      // Const Field (szlItemLength)
      lengthInBits += 16;


      // Implicit Field (szlItemCount)
      lengthInBits += 16;


      // Array field
      if(_message->s7_payload_user_data_item_cpu_function_read_szl_response_items != NULL) {
        plc4c_list_element* curElement = _message->s7_payload_user_data_item_cpu_function_read_szl_response_items->tail;
        while (curElement != NULL) {
          lengthInBits += plc4c_s7_read_write_szl_data_tree_item_length_in_bits((plc4c_s7_read_write_szl_data_tree_item*) curElement->value);
          curElement = curElement->next;
        }
      }

      break;
    }
  }

  return lengthInBits;
}

