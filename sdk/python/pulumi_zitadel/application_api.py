# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationApiArgs', 'ApplicationApi']

@pulumi.input_type
class ApplicationApiArgs:
    def __init__(__self__, *,
                 org_id: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 auth_method_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationApi resource.
        :param pulumi.Input[str] org_id: orgID of the application
        :param pulumi.Input[str] project_id: ID of the project
        :param pulumi.Input[str] auth_method_type: Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        :param pulumi.Input[str] name: Name of the application
        """
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "project_id", project_id)
        if auth_method_type is not None:
            pulumi.set(__self__, "auth_method_type", auth_method_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        orgID of the application
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="authMethodType")
    def auth_method_type(self) -> Optional[pulumi.Input[str]]:
        """
        Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        """
        return pulumi.get(self, "auth_method_type")

    @auth_method_type.setter
    def auth_method_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApplicationApiState:
    def __init__(__self__, *,
                 auth_method_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationApi resources.
        :param pulumi.Input[str] auth_method_type: Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        :param pulumi.Input[str] client_id: generated ID for this config
        :param pulumi.Input[str] client_secret: generated secret for this config
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: orgID of the application
        :param pulumi.Input[str] project_id: ID of the project
        """
        if auth_method_type is not None:
            pulumi.set(__self__, "auth_method_type", auth_method_type)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="authMethodType")
    def auth_method_type(self) -> Optional[pulumi.Input[str]]:
        """
        Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        """
        return pulumi.get(self, "auth_method_type")

    @auth_method_type.setter
    def auth_method_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_type", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        generated ID for this config
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        generated secret for this config
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        orgID of the application
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


class ApplicationApi(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing an API application belonging to a project, with all configuration possibilities.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        application_api = zitadel.ApplicationApi("applicationApi",
            org_id=zitadel_org["org"]["id"],
            project_id=zitadel_project["project"]["id"],
            auth_method_type="API_AUTH_METHOD_TYPE_PRIVATE_KEY_JWT")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_type: Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: orgID of the application
        :param pulumi.Input[str] project_id: ID of the project
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationApiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing an API application belonging to a project, with all configuration possibilities.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        application_api = zitadel.ApplicationApi("applicationApi",
            org_id=zitadel_org["org"]["id"],
            project_id=zitadel_project["project"]["id"],
            auth_method_type="API_AUTH_METHOD_TYPE_PRIVATE_KEY_JWT")
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationApiArgs.__new__(ApplicationApiArgs)

            __props__.__dict__["auth_method_type"] = auth_method_type
            __props__.__dict__["name"] = name
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["client_id"] = None
            __props__.__dict__["client_secret"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientId", "clientSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ApplicationApi, __self__).__init__(
            'zitadel:index/applicationApi:ApplicationApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_method_type: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationApi':
        """
        Get an existing ApplicationApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_type: Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        :param pulumi.Input[str] client_id: generated ID for this config
        :param pulumi.Input[str] client_secret: generated secret for this config
        :param pulumi.Input[str] name: Name of the application
        :param pulumi.Input[str] org_id: orgID of the application
        :param pulumi.Input[str] project_id: ID of the project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationApiState.__new__(_ApplicationApiState)

        __props__.__dict__["auth_method_type"] = auth_method_type
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        return ApplicationApi(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authMethodType")
    def auth_method_type(self) -> pulumi.Output[Optional[str]]:
        """
        Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
        """
        return pulumi.get(self, "auth_method_type")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        generated ID for this config
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        generated secret for this config
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        orgID of the application
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")

