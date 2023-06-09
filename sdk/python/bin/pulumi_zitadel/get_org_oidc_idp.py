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
    'GetOrgOidcIdpResult',
    'AwaitableGetOrgOidcIdpResult',
    'get_org_oidc_idp',
    'get_org_oidc_idp_output',
]

@pulumi.output_type
class GetOrgOidcIdpResult:
    """
    A collection of values returned by getOrgOidcIdp.
    """
    def __init__(__self__, auto_register=None, client_id=None, client_secret=None, display_name_mapping=None, id=None, idp_id=None, issuer=None, name=None, org_id=None, scopes=None, styling_type=None, username_mapping=None):
        if auto_register and not isinstance(auto_register, bool):
            raise TypeError("Expected argument 'auto_register' to be a bool")
        pulumi.set(__self__, "auto_register", auto_register)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if display_name_mapping and not isinstance(display_name_mapping, str):
            raise TypeError("Expected argument 'display_name_mapping' to be a str")
        pulumi.set(__self__, "display_name_mapping", display_name_mapping)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_id and not isinstance(idp_id, str):
            raise TypeError("Expected argument 'idp_id' to be a str")
        pulumi.set(__self__, "idp_id", idp_id)
        if issuer and not isinstance(issuer, str):
            raise TypeError("Expected argument 'issuer' to be a str")
        pulumi.set(__self__, "issuer", issuer)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if styling_type and not isinstance(styling_type, str):
            raise TypeError("Expected argument 'styling_type' to be a str")
        pulumi.set(__self__, "styling_type", styling_type)
        if username_mapping and not isinstance(username_mapping, str):
            raise TypeError("Expected argument 'username_mapping' to be a str")
        pulumi.set(__self__, "username_mapping", username_mapping)

    @property
    @pulumi.getter(name="autoRegister")
    def auto_register(self) -> bool:
        """
        auto register for users from this idp
        """
        return pulumi.get(self, "auto_register")

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
    @pulumi.getter(name="displayNameMapping")
    def display_name_mapping(self) -> str:
        """
        definition which field is mapped to the display name of the user
        """
        return pulumi.get(self, "display_name_mapping")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter
    def issuer(self) -> str:
        """
        the oidc issuer of the identity provider
        """
        return pulumi.get(self, "issuer")

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

    @property
    @pulumi.getter(name="stylingType")
    def styling_type(self) -> str:
        """
        Some identity providers specify the styling of the button to their login
        """
        return pulumi.get(self, "styling_type")

    @property
    @pulumi.getter(name="usernameMapping")
    def username_mapping(self) -> str:
        """
        definition which field is mapped to the email of the user
        """
        return pulumi.get(self, "username_mapping")


class AwaitableGetOrgOidcIdpResult(GetOrgOidcIdpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgOidcIdpResult(
            auto_register=self.auto_register,
            client_id=self.client_id,
            client_secret=self.client_secret,
            display_name_mapping=self.display_name_mapping,
            id=self.id,
            idp_id=self.idp_id,
            issuer=self.issuer,
            name=self.name,
            org_id=self.org_id,
            scopes=self.scopes,
            styling_type=self.styling_type,
            username_mapping=self.username_mapping)


def get_org_oidc_idp(idp_id: Optional[str] = None,
                     org_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgOidcIdpResult:
    """
    Datasource representing a generic OIDC IdP on the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    org_oidc_idp_org_oidc_idp = zitadel.get_org_oidc_idp(org_id=data["zitadel_org"]["org"]["id"],
        idp_id="177073612581240835")
    pulumi.export("orgOidcIdp", org_oidc_idp_org_oidc_idp)
    ```


    :param str idp_id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    __args__ = dict()
    __args__['idpId'] = idp_id
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getOrgOidcIdp:getOrgOidcIdp', __args__, opts=opts, typ=GetOrgOidcIdpResult).value

    return AwaitableGetOrgOidcIdpResult(
        auto_register=__ret__.auto_register,
        client_id=__ret__.client_id,
        client_secret=__ret__.client_secret,
        display_name_mapping=__ret__.display_name_mapping,
        id=__ret__.id,
        idp_id=__ret__.idp_id,
        issuer=__ret__.issuer,
        name=__ret__.name,
        org_id=__ret__.org_id,
        scopes=__ret__.scopes,
        styling_type=__ret__.styling_type,
        username_mapping=__ret__.username_mapping)


@_utilities.lift_output_func(get_org_oidc_idp)
def get_org_oidc_idp_output(idp_id: Optional[pulumi.Input[str]] = None,
                            org_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgOidcIdpResult]:
    """
    Datasource representing a generic OIDC IdP on the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    org_oidc_idp_org_oidc_idp = zitadel.get_org_oidc_idp(org_id=data["zitadel_org"]["org"]["id"],
        idp_id="177073612581240835")
    pulumi.export("orgOidcIdp", org_oidc_idp_org_oidc_idp)
    ```


    :param str idp_id: The ID of this resource.
    :param str org_id: ID of the organization
    """
    ...
