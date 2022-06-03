"""
    Spyderbat API UI & Public APIs

    Restful APIs for use by UI & customers.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import unittest

import sbapi
from sbapi.api.source_data_api import SourceDataApi  # noqa: E501


class TestSourceDataApi(unittest.TestCase):
    """SourceDataApi unit test stubs"""

    def setUp(self):
        self.api = SourceDataApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_src_data_query(self):
        """Test case for src_data_query

        Query source data  # noqa: E501
        """
        pass

    def test_src_data_query_v2(self):
        """Test case for src_data_query_v2

        Query source data  # noqa: E501
        """
        pass

    def test_src_send_data(self):
        """Test case for src_send_data

        Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
