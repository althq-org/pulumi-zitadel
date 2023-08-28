# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTriggerActionsResult',
    'AwaitableGetTriggerActionsResult',
    'get_trigger_actions',
    'get_trigger_actions_output',
]

@pulumi.output_type
class GetTriggerActionsResult:
    """
    A collection of values returned by getTriggerActions.
    """
    def __init__(__self__, action_ids=None, flow_type=None, id=None, org_id=None, trigger_type=None):
        if action_ids and not isinstance(action_ids, list):
            raise TypeError("Expected argument 'action_ids' to be a list")
        pulumi.set(__self__, "action_ids", action_ids)
        if flow_type and not isinstance(flow_type, str):
            raise TypeError("Expected argument 'flow_type' to be a str")
        pulumi.set(__self__, "flow_type", flow_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if trigger_type and not isinstance(trigger_type, str):
            raise TypeError("Expected argument 'trigger_type' to be a str")
        pulumi.set(__self__, "trigger_type", trigger_type)

    @property
    @pulumi.getter(name="actionIds")
    def action_ids(self) -> Sequence[str]:
        """
        IDs of the triggered actions
        """
        return pulumi.get(self, "action_ids")

    @property
    @pulumi.getter(name="flowType")
    def flow_type(self) -> str:
        """
        Type of the flow to which the action triggers belong
        """
        return pulumi.get(self, "flow_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> str:
        """
        Trigger type on when the actions get triggered
        """
        return pulumi.get(self, "trigger_type")


class AwaitableGetTriggerActionsResult(GetTriggerActionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTriggerActionsResult(
            action_ids=self.action_ids,
            flow_type=self.flow_type,
            id=self.id,
            org_id=self.org_id,
            trigger_type=self.trigger_type)


def get_trigger_actions(flow_type: Optional[str] = None,
                        org_id: Optional[str] = None,
                        trigger_type: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTriggerActionsResult:
    """
    Resource representing triggers, when actions get started

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_trigger_actions(org_id=data["zitadel_org"]["default"]["id"],
        flow_type="FLOW_TYPE_EXTERNAL_AUTHENTICATION",
        trigger_type="TRIGGER_TYPE_POST_AUTHENTICATION")
    pulumi.export("triggerActions", default)
    ```


    :param str flow_type: Type of the flow to which the action triggers belong
    :param str org_id: ID of the organization
    :param str trigger_type: Trigger type on when the actions get triggered
    """
    __args__ = dict()
    __args__['flowType'] = flow_type
    __args__['orgId'] = org_id
    __args__['triggerType'] = trigger_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getTriggerActions:getTriggerActions', __args__, opts=opts, typ=GetTriggerActionsResult).value

    return AwaitableGetTriggerActionsResult(
        action_ids=__ret__.action_ids,
        flow_type=__ret__.flow_type,
        id=__ret__.id,
        org_id=__ret__.org_id,
        trigger_type=__ret__.trigger_type)


@_utilities.lift_output_func(get_trigger_actions)
def get_trigger_actions_output(flow_type: Optional[pulumi.Input[str]] = None,
                               org_id: Optional[pulumi.Input[Optional[str]]] = None,
                               trigger_type: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTriggerActionsResult]:
    """
    Resource representing triggers, when actions get started

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    default = zitadel.get_trigger_actions(org_id=data["zitadel_org"]["default"]["id"],
        flow_type="FLOW_TYPE_EXTERNAL_AUTHENTICATION",
        trigger_type="TRIGGER_TYPE_POST_AUTHENTICATION")
    pulumi.export("triggerActions", default)
    ```


    :param str flow_type: Type of the flow to which the action triggers belong
    :param str org_id: ID of the organization
    :param str trigger_type: Trigger type on when the actions get triggered
    """
    ...
