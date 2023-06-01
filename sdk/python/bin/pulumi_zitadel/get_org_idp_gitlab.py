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
    'GetOrgIdpGitlabResult',
    'AwaitableGetOrgIdpGitlabResult',
    'get_org_idp_gitlab',
    'get_org_idp_gitlab_output',
]

@pulumi.output_type
class GetOrgIdpGitlabResult:
    """
    A collection of values returned by getOrgIdpGitlab.
    """
    def __init__(__self__, client_id=None, client_secret=None, id=None, is_auto_creation=None, is_auto_update=None, is_creation_allowed=None, is_linking_allowed=None, name=None, org_id=None, scopes=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_auto_creation and not isinstance(is_auto_creation, bool):
            raise TypeError("Expected argument 'is_auto_creation' to be a bool")
        pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        if is_auto_update and not isinstance(is_auto_update, bool):
            raise TypeError("Expected argument 'is_auto_update' to be a bool")
        pulumi.set(__self__, "is_auto_update", is_auto_update)
        if is_creation_allowed and not isinstance(is_creation_allowed, bool):
            raise TypeError("Expected argument 'is_creation_allowed' to be a bool")
        pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        if is_linking_allowed and not isinstance(is_linking_allowed, bool):
            raise TypeError("Expected argument 'is_linking_allowed' to be a bool")
        pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> bool:
        """
        enabled if a new account in ZITADEL are created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> bool:
        """
        enabled if a the ZITADEL account fields are updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> bool:
        """
        enabled if users are able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> bool:
        """
        enabled if users are able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the IDP
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
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")


class AwaitableGetOrgIdpGitlabResult(GetOrgIdpGitlabResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgIdpGitlabResult(
            client_id=self.client_id,
            client_secret=self.client_secret,
            id=self.id,
            is_auto_creation=self.is_auto_creation,
            is_auto_update=self.is_auto_update,
            is_creation_allowed=self.is_creation_allowed,
            is_linking_allowed=self.is_linking_allowed,
            name=self.name,
            org_id=self.org_id,
            scopes=self.scopes)


def get_org_idp_gitlab(id: Optional[str] = None,
                       org_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgIdpGitlabResult:
    """
    Datasource representing a GitLab IdP of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    gitlab = zitadel.get_org_idp_gitlab(id="177073614158299139")
    ```


    :param str id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getOrgIdpGitlab:getOrgIdpGitlab', __args__, opts=opts, typ=GetOrgIdpGitlabResult).value

    return AwaitableGetOrgIdpGitlabResult(
        client_id=__ret__.client_id,
        client_secret=__ret__.client_secret,
        id=__ret__.id,
        is_auto_creation=__ret__.is_auto_creation,
        is_auto_update=__ret__.is_auto_update,
        is_creation_allowed=__ret__.is_creation_allowed,
        is_linking_allowed=__ret__.is_linking_allowed,
        name=__ret__.name,
        org_id=__ret__.org_id,
        scopes=__ret__.scopes)


@_utilities.lift_output_func(get_org_idp_gitlab)
def get_org_idp_gitlab_output(id: Optional[pulumi.Input[str]] = None,
                              org_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgIdpGitlabResult]:
    """
    Datasource representing a GitLab IdP of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    gitlab = zitadel.get_org_idp_gitlab(id="177073614158299139")
    ```


    :param str id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    ...
