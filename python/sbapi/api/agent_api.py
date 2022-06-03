"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from sbapi.api_client import ApiClient, Endpoint as _Endpoint
from sbapi.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from sbapi.model.agent import Agent


class AgentApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.agent_list_endpoint = _Endpoint(
            settings={
                'response_type': ([Agent],),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/org/{orgUID}/agent/',
                'operation_id': 'agent_list',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'org_uid',
                    'agent_registration_uid_equals',
                    'original_association',
                    'page',
                    'page_size',
                    'source_uid_equals',
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
                    'agent_registration_uid_equals',
                    'page_size',
                    'source_uid_equals',
                ]
            },
            root_map={
                'validations': {
                    ('org_uid',): {
                        'max_length': 64,
                        'min_length': 10,
                    },
                    ('agent_registration_uid_equals',): {
                        'max_length': 64,
                    },
                    ('page_size',): {

                        'inclusive_maximum': 100,
                    },
                    ('source_uid_equals',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'org_uid':
                        (str,),
                    'agent_registration_uid_equals':
                        (str,),
                    'original_association':
                        (bool,),
                    'page':
                        (int,),
                    'page_size':
                        (int,),
                    'source_uid_equals':
                        (str,),
                },
                'attribute_map': {
                    'org_uid': 'orgUID',
                    'agent_registration_uid_equals': 'agent_registration_uid_equals',
                    'original_association': 'original_association',
                    'page': 'page',
                    'page_size': 'page_size',
                    'source_uid_equals': 'source_uid_equals',
                },
                'location_map': {
                    'org_uid': 'path',
                    'agent_registration_uid_equals': 'query',
                    'original_association': 'query',
                    'page': 'query',
                    'page_size': 'query',
                    'source_uid_equals': 'query',
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
        self.agent_load_endpoint = _Endpoint(
            settings={
                'response_type': (Agent,),
                'auth': [
                    'apiToken'
                ],
                'endpoint_path': '/api/v1/org/{orgUID}/agent/{agentUID}',
                'operation_id': 'agent_load',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'agent_uid',
                    'org_uid',
                ],
                'required': [
                    'agent_uid',
                    'org_uid',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'agent_uid',
                    'org_uid',
                ]
            },
            root_map={
                'validations': {
                    ('agent_uid',): {
                        'max_length': 64,
                    },
                    ('org_uid',): {
                        'max_length': 64,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'agent_uid':
                        (str,),
                    'org_uid':
                        (str,),
                },
                'attribute_map': {
                    'agent_uid': 'agentUID',
                    'org_uid': 'orgUID',
                },
                'location_map': {
                    'agent_uid': 'path',
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

    def agent_list(
        self,
        org_uid,
        **kwargs
    ):
        """List agents  # noqa: E501

         Lists the agents associated with an organization  * Requires the action  *agent:List* on the organization    # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.agent_list(org_uid, async_req=True)
        >>> result = thread.get()

        Args:
            org_uid (str):

        Keyword Args:
            agent_registration_uid_equals (str): [optional]
            original_association (bool): [optional]
            page (int): [optional]
            page_size (int): [optional]
            source_uid_equals (str): [optional]
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
            [Agent]
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
        return self.agent_list_endpoint.call_with_http_info(**kwargs)

    def agent_load(
        self,
        agent_uid,
        org_uid,
        **kwargs
    ):
        """Load an agent  # noqa: E501

         Load a specified agent  * Requires the action  *agent:Load* on the organization    # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.agent_load(agent_uid, org_uid, async_req=True)
        >>> result = thread.get()

        Args:
            agent_uid (str):
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
            Agent
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
        kwargs['agent_uid'] = \
            agent_uid
        kwargs['org_uid'] = \
            org_uid
        return self.agent_load_endpoint.call_with_http_info(**kwargs)

