"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
from typing import Type

from magma.enodebd.devices.baicells import BaicellsHandler
from magma.enodebd.devices.baicells_old import BaicellsOldHandler
from magma.enodebd.devices.baicells_qafa import BaicellsQAFAHandler
from magma.enodebd.devices.baicells_qafb import BaicellsQAFBHandler
from magma.enodebd.devices.experimental.cavium import CaviumHandler
from magma.enodebd.devices.device_utils import EnodebDeviceName
from magma.enodebd.state_machines.enb_acs import EnodebAcsStateMachine

# This exists only to break a circular dependency. Otherwise there's no
# point of having these names for the devices


DEVICE_HANDLER_BY_NAME = {
    EnodebDeviceName.BAICELLS: BaicellsHandler,
    EnodebDeviceName.BAICELLS_OLD: BaicellsOldHandler,
    EnodebDeviceName.BAICELLS_QAFA: BaicellsQAFAHandler,
    EnodebDeviceName.BAICELLS_QAFB: BaicellsQAFBHandler,
    EnodebDeviceName.CAVIUM: CaviumHandler,
}


def get_device_handler_from_name(
    name: EnodebDeviceName,
) -> Type[EnodebAcsStateMachine]:
    return DEVICE_HANDLER_BY_NAME[name]
