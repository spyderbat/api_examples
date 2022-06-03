"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import unittest

import sbapi
from sbapi.api.org_type_api import OrgTypeApi  # noqa: E501


class TestOrgTypeApi(unittest.TestCase):
    """OrgTypeApi unit test stubs"""

    def setUp(self):
        self.api = OrgTypeApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_org_type_limit_active_sources(self):
        """Test case for org_type_limit_active_sources

        Loads limits regarding active sources  # noqa: E501
        """
        pass

    def test_org_type_limit_org_roles(self):
        """Test case for org_type_limit_org_roles

        Loads limits regarding org roles  # noqa: E501
        """
        pass

    def test_org_type_load(self):
        """Test case for org_type_load

        Load the org type for the organization  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
