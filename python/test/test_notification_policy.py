"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: apisupport@spyderbat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import spyderbat_api
from spyderbat_api.model.notification_policy_destination import NotificationPolicyDestination
from spyderbat_api.model.notification_policy_routes_inner import NotificationPolicyRoutesInner
globals()['NotificationPolicyDestination'] = NotificationPolicyDestination
globals()['NotificationPolicyRoutesInner'] = NotificationPolicyRoutesInner
from spyderbat_api.model.notification_policy import NotificationPolicy


class TestNotificationPolicy(unittest.TestCase):
    """NotificationPolicy unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testNotificationPolicy(self):
        """Test NotificationPolicy"""
        # FIXME: construct object with mandatory attributes with example values
        # model = NotificationPolicy()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()