"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: apisupport@spyderbat.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import sbapi
from sbapi.api.api_key_api import APIKeyApi  # noqa: E501


class TestAPIKeyApi(unittest.TestCase):
    """APIKeyApi unit test stubs"""

    def setUp(self):
        self.api = APIKeyApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_api_key_create(self):
        """Test case for api_key_create

        Creates a new API key  # noqa: E501
        """
        pass

    def test_api_key_delete(self):
        """Test case for api_key_delete

        Delete an API key  # noqa: E501
        """
        pass

    def test_api_key_list(self):
        """Test case for api_key_list

        Lists an API key  # noqa: E501
        """
        pass

    def test_api_key_update(self):
        """Test case for api_key_update

        Updates an API key  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
