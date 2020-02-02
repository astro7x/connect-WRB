#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from functools import partial
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


def enum_field(enum_type):
    def encode_enum(value):
        return value.value

    def decode_enum(t, value):
        return t(value)

    return field(
        metadata={
            "dataclasses_json": {
                "encoder": encode_enum,
                "decoder": partial(decode_enum, enum_type),
            }
        }
    )


class ServiceStatus(Enum):
    PENDING = "PENDING"
    IN_SERVICE = "IN_SERVICE"
    MAINTENANCE = "MAINTENANCE"
    DISCONNECTED = "DISCONNECTED"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value):
        return cls.MISSING_ENUM


class ServiceEndpointRole(Enum):
    CONSUMER = "CONSUMER"
    PROVIDER = "PROVIDER"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value):
        return cls.MISSING_ENUM


@dataclass_json
@dataclass
class ServiceCreateData:
    @dataclass_json
    @dataclass
    class PropertyInput:
        propertyTypeID: str
        id: Optional[str] = None
        stringValue: Optional[str] = None
        intValue: Optional[int] = None
        booleanValue: Optional[bool] = None
        floatValue: Optional[float] = None
        latitudeValue: Optional[float] = None
        longitudeValue: Optional[float] = None
        rangeFromValue: Optional[float] = None
        rangeToValue: Optional[float] = None
        equipmentIDValue: Optional[str] = None
        locationIDValue: Optional[str] = None
        serviceIDValue: Optional[str] = None
        isEditable: Optional[bool] = None
        isInstanceProperty: Optional[bool] = None

    name: str
    serviceTypeId: str
    upstreamServiceIds: List[str]
    externalId: Optional[str] = None
    status: Optional[ServiceStatus] = None
    customerId: Optional[str] = None
    properties: Optional[List[PropertyInput]] = None


@dataclass_json
@dataclass
class AddServiceMutation:
    __QUERY__ = """
    mutation AddServiceMutation($data: ServiceCreateData!) {
  addService(data: $data) {
    id
    name
    externalId
    customer {
      id
      name
      externalId
    }
    endpoints {
      id
      port {
        id
      }
      role
    }
    links {
      id
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class AddServiceMutationData:
        @dataclass_json
        @dataclass
        class Service:
            @dataclass_json
            @dataclass
            class Customer:
                id: str
                name: str
                externalId: Optional[str] = None

            @dataclass_json
            @dataclass
            class ServiceEndpoint:
                @dataclass_json
                @dataclass
                class EquipmentPort:
                    id: str

                id: str
                port: EquipmentPort
                role: ServiceEndpointRole = enum_field(ServiceEndpointRole)

            @dataclass_json
            @dataclass
            class Link:
                id: str

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str] = None
            customer: Optional[Customer] = None

        addService: Optional[Service] = None

    data: Optional[AddServiceMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client, data: ServiceCreateData):
        # fmt: off
        variables = {"data": data}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data