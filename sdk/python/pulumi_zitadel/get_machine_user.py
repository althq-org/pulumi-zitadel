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
    'GetMachineUserResult',
    'AwaitableGetMachineUserResult',
    'get_machine_user',
    'get_machine_user_output',
]

@pulumi.output_type
class GetMachineUserResult:
    """
    A collection of values returned by getMachineUser.
    """
    def __init__(__self__, access_token_type=None, description=None, id=None, login_names=None, name=None, org_id=None, preferred_login_name=None, state=None, user_id=None, user_name=None):
        if access_token_type and not isinstance(access_token_type, str):
            raise TypeError("Expected argument 'access_token_type' to be a str")
        pulumi.set(__self__, "access_token_type", access_token_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if login_names and not isinstance(login_names, list):
            raise TypeError("Expected argument 'login_names' to be a list")
        pulumi.set(__self__, "login_names", login_names)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if preferred_login_name and not isinstance(preferred_login_name, str):
            raise TypeError("Expected argument 'preferred_login_name' to be a str")
        pulumi.set(__self__, "preferred_login_name", preferred_login_name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessTokenType")
    def access_token_type(self) -> str:
        """
        Access token type
        """
        return pulumi.get(self, "access_token_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the user
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loginNames")
    def login_names(self) -> Sequence[str]:
        """
        Loginnames
        """
        return pulumi.get(self, "login_names")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the machine user
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="preferredLoginName")
    def preferred_login_name(self) -> str:
        """
        Preferred login name
        """
        return pulumi.get(self, "preferred_login_name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the user
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        Username
        """
        return pulumi.get(self, "user_name")


class AwaitableGetMachineUserResult(GetMachineUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMachineUserResult(
            access_token_type=self.access_token_type,
            description=self.description,
            id=self.id,
            login_names=self.login_names,
            name=self.name,
            org_id=self.org_id,
            preferred_login_name=self.preferred_login_name,
            state=self.state,
            user_id=self.user_id,
            user_name=self.user_name)


def get_machine_user(org_id: Optional[str] = None,
                     user_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMachineUserResult:
    """
    Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    machine_user_machine_user = zitadel.get_machine_user(org_id=data["zitadel_org"]["org"]["id"],
        user_id="177073617463410691")
    pulumi.export("machineUser", machine_user_machine_user)
    ```


    :param str org_id: ID of the organization
    :param str user_id: The ID of this resource.
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getMachineUser:getMachineUser', __args__, opts=opts, typ=GetMachineUserResult).value

    return AwaitableGetMachineUserResult(
        access_token_type=__ret__.access_token_type,
        description=__ret__.description,
        id=__ret__.id,
        login_names=__ret__.login_names,
        name=__ret__.name,
        org_id=__ret__.org_id,
        preferred_login_name=__ret__.preferred_login_name,
        state=__ret__.state,
        user_id=__ret__.user_id,
        user_name=__ret__.user_name)


@_utilities.lift_output_func(get_machine_user)
def get_machine_user_output(org_id: Optional[pulumi.Input[str]] = None,
                            user_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMachineUserResult]:
    """
    Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    machine_user_machine_user = zitadel.get_machine_user(org_id=data["zitadel_org"]["org"]["id"],
        user_id="177073617463410691")
    pulumi.export("machineUser", machine_user_machine_user)
    ```


    :param str org_id: ID of the organization
    :param str user_id: The ID of this resource.
    """
    ...
