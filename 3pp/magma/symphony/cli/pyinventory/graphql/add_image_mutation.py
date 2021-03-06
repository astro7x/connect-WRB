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


class ImageEntity(Enum):
    LOCATION = "LOCATION"
    WORK_ORDER = "WORK_ORDER"
    SITE_SURVEY = "SITE_SURVEY"
    EQUIPMENT = "EQUIPMENT"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value):
        return cls.MISSING_ENUM


@dataclass_json
@dataclass
class AddImageInput:
    entityType: ImageEntity = enum_field(ImageEntity)
    entityId: str
    imgKey: str
    fileName: str
    fileSize: int
    modified: datetime = DATETIME_FIELD
    contentType: str
    category: Optional[str] = None


@dataclass_json
@dataclass
class AddImageMutation:
    __QUERY__ = """
    mutation AddImageMutation($input: AddImageInput!) {
  addImage(input: $input) {
    id
    fileName
  }
}

    """

    @dataclass_json
    @dataclass
    class AddImageMutationData:
        @dataclass_json
        @dataclass
        class File:
            id: str
            fileName: str

        addImage: Optional[File] = None

    data: Optional[AddImageMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client, input: AddImageInput):
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
