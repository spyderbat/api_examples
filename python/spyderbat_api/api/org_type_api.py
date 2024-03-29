"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: apisupport@spyderbat.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from spyderbat_api.api_client import ApiClient, Endpoint as _Endpoint
from spyderbat_api.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from spyderbat_api.model.dao_org_type import DaoOrgType
from spyderbat_api.model.session_org_type_max_limit import SessionOrgTypeMaxLimit


class OrgTypeApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.org_type_limit_active_sources_endpoint = _Endpoint(
            settings={
                'response_type': (SessionOrgTypeMaxLimit,),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/org/{orgUID}/type/limit/active_sources',
                'operation_id': 'org_type_limit_active_sources',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'org_uid',
                ],
                'required': [
                    'org_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'org_uid',
                ]
            },
            root_map={
                'validations': {
                    ('org_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'org_uid':
                        (str,),
                },
                'attribute_map': {
                    'org_uid': 'orgUID',
                },
                'location_map': {
                    'org_uid': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.org_type_limit_org_roles_endpoint = _Endpoint(
            settings={
                'response_type': (SessionOrgTypeMaxLimit,),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/org/{orgUID}/type/limit/org_roles',
                'operation_id': 'org_type_limit_org_roles',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'org_uid',
                ],
                'required': [
                    'org_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'org_uid',
                ]
            },
            root_map={
                'validations': {
                    ('org_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'org_uid':
                        (str,),
                },
                'attribute_map': {
                    'org_uid': 'orgUID',
                },
                'location_map': {
                    'org_uid': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.org_type_load_endpoint = _Endpoint(
            settings={
                'response_type': (DaoOrgType,),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/org/{orgUID}/type',
                'operation_id': 'org_type_load',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'org_uid',
                ],
                'required': [
                    'org_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'org_uid',
                ]
            },
            root_map={
                'validations': {
                    ('org_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'org_uid':
                        (str,),
                },
                'attribute_map': {
                    'org_uid': 'orgUID',
                },
                'location_map': {
                    'org_uid': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def org_type_limit_active_sources(
        self,
        org_uid,
        **kwargs
    ):
        """Loads limits regarding active sources  # noqa: E501

         Loads the limits regarding active sources allowed on the organization, the active sources in an org are calculated on a 5m0s basis.    * Requires action *org:Load*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.org_type_limit_active_sources(org_uid, async_req=True)
        >>> result = thread.get()

        Args:
            org_uid (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            SessionOrgTypeMaxLimit
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['org_uid'] = \
            org_uid
        return self.org_type_limit_active_sources_endpoint.call_with_http_info(**kwargs)

    def org_type_limit_org_roles(
        self,
        org_uid,
        **kwargs
    ):
        """Loads limits regarding org roles  # noqa: E501

         Loads the limit information regarding the number of associated roles allowed per an organization   * Requires action *org:Load*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.org_type_limit_org_roles(org_uid, async_req=True)
        >>> result = thread.get()

        Args:
            org_uid (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            SessionOrgTypeMaxLimit
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['org_uid'] = \
            org_uid
        return self.org_type_limit_org_roles_endpoint.call_with_http_info(**kwargs)

    def org_type_load(
        self,
        org_uid,
        **kwargs
    ):
        """Load the org type for the organization  # noqa: E501

         Loads the org type for the organiation   * Requires action *org:Load*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.org_type_load(org_uid, async_req=True)
        >>> result = thread.get()

        Args:
            org_uid (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            DaoOrgType
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['org_uid'] = \
            org_uid
        return self.org_type_load_endpoint.call_with_http_info(**kwargs)

