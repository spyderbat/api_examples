"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: apisupport@spyderbat.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import spyderbat_api
from spyderbat_api.api.org_api import OrgApi  # noqa: E501


class TestOrgApi(unittest.TestCase):
    """OrgApi unit test stubs"""

    def setUp(self):
        self.api = OrgApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_org_assign_role(self):
        """Test case for org_assign_role

        Assign OrgRole  # noqa: E501
        """
        pass

    def test_org_invite_users(self):
        """Test case for org_invite_users

        Invite users to an organization  # noqa: E501
        """
        pass

    def test_org_list(self):
        """Test case for org_list

        List organizations  # noqa: E501
        """
        pass

    def test_org_list_role(self):
        """Test case for org_list_role

        List OrgRoles  # noqa: E501
        """
        pass

    def test_org_load(self):
        """Test case for org_load

        Load an organization  # noqa: E501
        """
        pass

    def test_org_load_notification_policy(self):
        """Test case for org_load_notification_policy

        Load Notification Policy  # noqa: E501
        """
        pass

    def test_org_test_notification_target(self):
        """Test case for org_test_notification_target

        Test Notification Target  # noqa: E501
        """
        pass

    def test_org_unassign_role(self):
        """Test case for org_unassign_role

        Unassign OrgRole  # noqa: E501
        """
        pass

    def test_org_update(self):
        """Test case for org_update

        Update an organization  # noqa: E501
        """
        pass

    def test_org_update_notification_policy(self):
        """Test case for org_update_notification_policy

        Update an organization's notification policy  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
