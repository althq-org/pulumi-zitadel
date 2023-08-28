# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MachineUserArgs', 'MachineUser']

@pulumi.input_type
class MachineUserArgs:
    def __init__(__self__, *,
                 user_name: pulumi.Input[str],
                 access_token_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MachineUser resource.
        :param pulumi.Input[str] user_name: Username
        :param pulumi.Input[str] access_token_type: Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        :param pulumi.Input[str] description: Description of the user
        :param pulumi.Input[str] name: Name of the machine user
        :param pulumi.Input[str] org_id: ID of the organization
        """
        pulumi.set(__self__, "user_name", user_name)
        if access_token_type is not None:
            pulumi.set(__self__, "access_token_type", access_token_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="accessTokenType")
    def access_token_type(self) -> Optional[pulumi.Input[str]]:
        """
        Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        """
        return pulumi.get(self, "access_token_type")

    @access_token_type.setter
    def access_token_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the user
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the machine user
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _MachineUserState:
    def __init__(__self__, *,
                 access_token_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 login_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 preferred_login_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MachineUser resources.
        :param pulumi.Input[str] access_token_type: Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        :param pulumi.Input[str] description: Description of the user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] login_names: Loginnames
        :param pulumi.Input[str] name: Name of the machine user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] preferred_login_name: Preferred login name
        :param pulumi.Input[str] state: State of the user
        :param pulumi.Input[str] user_name: Username
        """
        if access_token_type is not None:
            pulumi.set(__self__, "access_token_type", access_token_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if login_names is not None:
            pulumi.set(__self__, "login_names", login_names)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if preferred_login_name is not None:
            pulumi.set(__self__, "preferred_login_name", preferred_login_name)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessTokenType")
    def access_token_type(self) -> Optional[pulumi.Input[str]]:
        """
        Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        """
        return pulumi.get(self, "access_token_type")

    @access_token_type.setter
    def access_token_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the user
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="loginNames")
    def login_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Loginnames
        """
        return pulumi.get(self, "login_names")

    @login_names.setter
    def login_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "login_names", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the machine user
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="preferredLoginName")
    def preferred_login_name(self) -> Optional[pulumi.Input[str]]:
        """
        Preferred login name
        """
        return pulumi.get(self, "preferred_login_name")

    @preferred_login_name.setter
    def preferred_login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_login_name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the user
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class MachineUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        default = zitadel.MachineUser("default",
            org_id=data["zitadel_org"]["default"]["id"],
            user_name="machine@example.com",
            description="a machine user")
        ```

        ## Import

        terraform The resource can be imported using the ID format `<id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/machineUser:MachineUser imported '123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token_type: Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        :param pulumi.Input[str] description: Description of the user
        :param pulumi.Input[str] name: Name of the machine user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] user_name: Username
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MachineUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        default = zitadel.MachineUser("default",
            org_id=data["zitadel_org"]["default"]["id"],
            user_name="machine@example.com",
            description="a machine user")
        ```

        ## Import

        terraform The resource can be imported using the ID format `<id[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/machineUser:MachineUser imported '123456789012345678:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param MachineUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MachineUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MachineUserArgs.__new__(MachineUserArgs)

            __props__.__dict__["access_token_type"] = access_token_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["login_names"] = None
            __props__.__dict__["preferred_login_name"] = None
            __props__.__dict__["state"] = None
        super(MachineUser, __self__).__init__(
            'zitadel:index/machineUser:MachineUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            login_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            preferred_login_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'MachineUser':
        """
        Get an existing MachineUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token_type: Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        :param pulumi.Input[str] description: Description of the user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] login_names: Loginnames
        :param pulumi.Input[str] name: Name of the machine user
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] preferred_login_name: Preferred login name
        :param pulumi.Input[str] state: State of the user
        :param pulumi.Input[str] user_name: Username
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MachineUserState.__new__(_MachineUserState)

        __props__.__dict__["access_token_type"] = access_token_type
        __props__.__dict__["description"] = description
        __props__.__dict__["login_names"] = login_names
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["preferred_login_name"] = preferred_login_name
        __props__.__dict__["state"] = state
        __props__.__dict__["user_name"] = user_name
        return MachineUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessTokenType")
    def access_token_type(self) -> pulumi.Output[Optional[str]]:
        """
        Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        """
        return pulumi.get(self, "access_token_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the user
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="loginNames")
    def login_names(self) -> pulumi.Output[Sequence[str]]:
        """
        Loginnames
        """
        return pulumi.get(self, "login_names")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the machine user
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="preferredLoginName")
    def preferred_login_name(self) -> pulumi.Output[str]:
        """
        Preferred login name
        """
        return pulumi.get(self, "preferred_login_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the user
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Username
        """
        return pulumi.get(self, "user_name")

