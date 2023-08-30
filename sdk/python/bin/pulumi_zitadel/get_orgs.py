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
    'GetOrgsResult',
    'AwaitableGetOrgsResult',
    'get_orgs',
    'get_orgs_output',
]

@pulumi.output_type
class GetOrgsResult:
    """
    A collection of values returned by getOrgs.
    """
    def __init__(__self__, domain=None, domain_method=None, id=None, ids=None, name=None, name_method=None, primary_domain=None, state=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domain_method and not isinstance(domain_method, str):
            raise TypeError("Expected argument 'domain_method' to be a str")
        pulumi.set(__self__, "domain_method", domain_method)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_method and not isinstance(name_method, str):
            raise TypeError("Expected argument 'name_method' to be a str")
        pulumi.set(__self__, "name_method", name_method)
        if primary_domain and not isinstance(primary_domain, str):
            raise TypeError("Expected argument 'primary_domain' to be a str")
        pulumi.set(__self__, "primary_domain", primary_domain)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        A domain of the org.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainMethod")
    def domain_method(self) -> Optional[str]:
        """
        Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        """
        return pulumi.get(self, "domain_method")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        A set of all organization IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the org.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameMethod")
    def name_method(self) -> Optional[str]:
        """
        Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        """
        return pulumi.get(self, "name_method")

    @property
    @pulumi.getter(name="primaryDomain")
    def primary_domain(self) -> str:
        """
        Primary domain of the org
        """
        return pulumi.get(self, "primary_domain")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
        """
        return pulumi.get(self, "state")


class AwaitableGetOrgsResult(GetOrgsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgsResult(
            domain=self.domain,
            domain_method=self.domain_method,
            id=self.id,
            ids=self.ids,
            name=self.name,
            name_method=self.name_method,
            primary_domain=self.primary_domain,
            state=self.state)


def get_orgs(domain: Optional[str] = None,
             domain_method: Optional[str] = None,
             name: Optional[str] = None,
             name_method: Optional[str] = None,
             state: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgsResult:
    """
    Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.


    :param str domain: A domain of the org.
    :param str domain_method: Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str name: Name of the org.
    :param str name_method: Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str state: State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['domainMethod'] = domain_method
    __args__['name'] = name
    __args__['nameMethod'] = name_method
    __args__['state'] = state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getOrgs:getOrgs', __args__, opts=opts, typ=GetOrgsResult).value

    return AwaitableGetOrgsResult(
        domain=__ret__.domain,
        domain_method=__ret__.domain_method,
        id=__ret__.id,
        ids=__ret__.ids,
        name=__ret__.name,
        name_method=__ret__.name_method,
        primary_domain=__ret__.primary_domain,
        state=__ret__.state)


@_utilities.lift_output_func(get_orgs)
def get_orgs_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                    domain_method: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    name_method: Optional[pulumi.Input[Optional[str]]] = None,
                    state: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgsResult]:
    """
    Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.


    :param str domain: A domain of the org.
    :param str domain_method: Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str name: Name of the org.
    :param str name_method: Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str state: State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
    """
    ...