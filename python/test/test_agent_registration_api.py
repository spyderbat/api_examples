"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import unittest

import sbapi
from sbapi.api.agent_registration_api import AgentRegistrationApi  # noqa: E501


class TestAgentRegistrationApi(unittest.TestCase):
    """AgentRegistrationApi unit test stubs"""

    def setUp(self):
        self.api = AgentRegistrationApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_agent_registration_create(self):
        """Test case for agent_registration_create

        Create an agent registration  # noqa: E501
        """
        pass

    def test_agent_registration_download_link(self):
        """Test case for agent_registration_download_link

        Get a download link for this registration  # noqa: E501
        """
        pass

    def test_agent_registration_get_log(self):
        """Test case for agent_registration_get_log

        Get log of recent agent registration activity  # noqa: E501
        """
        pass

    def test_agent_registration_list(self):
        """Test case for agent_registration_list

        List agent registrations  # noqa: E501
        """
        pass

    def test_agent_registration_load(self):
        """Test case for agent_registration_load

        Load an agent registration  # noqa: E501
        """
        pass

    def test_agent_registration_update(self):
        """Test case for agent_registration_update

        Update an agent registration  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
