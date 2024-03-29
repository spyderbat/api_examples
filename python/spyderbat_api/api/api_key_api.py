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
from spyderbat_api.model.api_key import APIKey
from spyderbat_api.model.api_key_create_input import ApiKeyCreateInput
from spyderbat_api.model.api_key_update_input import ApiKeyUpdateInput
from spyderbat_api.model.validation_error import ValidationError


class APIKeyApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.api_key_create_endpoint = _Endpoint(
            settings={
                'response_type': (str,),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/user/{userUID}/apikey/',
                'operation_id': 'api_key_create',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'user_uid',
                    'api_key_create_input',
                ],
                'required': [
                    'user_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'user_uid',
                ]
            },
            root_map={
                'validations': {
                    ('user_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'user_uid':
                        (str,),
                    'api_key_create_input':
                        (ApiKeyCreateInput,),
                },
                'attribute_map': {
                    'user_uid': 'userUID',
                },
                'location_map': {
                    'user_uid': 'path',
                    'api_key_create_input': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )
        self.api_key_delete_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/user/{userUID}/apikey/{uid}',
                'operation_id': 'api_key_delete',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'uid',
                    'user_uid',
                ],
                'required': [
                    'uid',
                    'user_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'uid',
                    'user_uid',
                ]
            },
            root_map={
                'validations': {
                    ('uid',): {
                        'max_length': 64,
                    },
                    ('user_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'uid':
                        (str,),
                    'user_uid':
                        (str,),
                },
                'attribute_map': {
                    'uid': 'uid',
                    'user_uid': 'userUID',
                },
                'location_map': {
                    'uid': 'path',
                    'user_uid': 'path',
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
        self.api_key_list_endpoint = _Endpoint(
            settings={
                'response_type': ([APIKey],),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/user/{userUID}/apikey/',
                'operation_id': 'api_key_list',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'user_uid',
                ],
                'required': [
                    'user_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'user_uid',
                ]
            },
            root_map={
                'validations': {
                    ('user_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'user_uid':
                        (str,),
                },
                'attribute_map': {
                    'user_uid': 'userUID',
                },
                'location_map': {
                    'user_uid': 'path',
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
        self.api_key_update_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/user/{userUID}/apikey/{uid}',
                'operation_id': 'api_key_update',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'uid',
                    'user_uid',
                    'api_key_update_input',
                ],
                'required': [
                    'uid',
                    'user_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'uid',
                    'user_uid',
                ]
            },
            root_map={
                'validations': {
                    ('uid',): {
                        'max_length': 64,
                    },
                    ('user_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'uid':
                        (str,),
                    'user_uid':
                        (str,),
                    'api_key_update_input':
                        (ApiKeyUpdateInput,),
                },
                'attribute_map': {
                    'uid': 'uid',
                    'user_uid': 'userUID',
                },
                'location_map': {
                    'uid': 'path',
                    'user_uid': 'path',
                    'api_key_update_input': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )

    def api_key_create(
        self,
        user_uid,
        **kwargs
    ):
        """Creates a new API key  # noqa: E501

         Creates a new API key which is associated with the user    * Requires global action *user:APIKeyCreate*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.api_key_create(user_uid, async_req=True)
        >>> result = thread.get()

        Args:
            user_uid (str): User UID

        Keyword Args:
            api_key_create_input (ApiKeyCreateInput): [optional]
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
            str
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
        kwargs['user_uid'] = \
            user_uid
        return self.api_key_create_endpoint.call_with_http_info(**kwargs)

    def api_key_delete(
        self,
        uid,
        user_uid,
        **kwargs
    ):
        """Delete an API key  # noqa: E501

         Deletes a specific API key   * Requires global action *user:APIKeyDelete*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.api_key_delete(uid, user_uid, async_req=True)
        >>> result = thread.get()

        Args:
            uid (str): API Key UID
            user_uid (str): User UID

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
            None
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
        kwargs['uid'] = \
            uid
        kwargs['user_uid'] = \
            user_uid
        return self.api_key_delete_endpoint.call_with_http_info(**kwargs)

    def api_key_list(
        self,
        user_uid,
        **kwargs
    ):
        """Lists an API key  # noqa: E501

         Lists API keys associated with a user   * Requires global action *user:APIKeyList*    # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.api_key_list(user_uid, async_req=True)
        >>> result = thread.get()

        Args:
            user_uid (str): User UID

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
            [APIKey]
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
        kwargs['user_uid'] = \
            user_uid
        return self.api_key_list_endpoint.call_with_http_info(**kwargs)

    def api_key_update(
        self,
        uid,
        user_uid,
        **kwargs
    ):
        """Updates an API key  # noqa: E501

         Updates a specific API Key, the only fields which can be updated are description and enabled   * Requires global action *user:APIKeyUpdate*   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.api_key_update(uid, user_uid, async_req=True)
        >>> result = thread.get()

        Args:
            uid (str): API Key UID
            user_uid (str): User UID

        Keyword Args:
            api_key_update_input (ApiKeyUpdateInput): [optional]
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
            None
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
        kwargs['uid'] = \
            uid
        kwargs['user_uid'] = \
            user_uid
        return self.api_key_update_endpoint.call_with_http_info(**kwargs)

