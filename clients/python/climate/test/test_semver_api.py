# coding: utf-8

"""
    Climate time series API

    The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins.   # noqa: E501

    OpenAPI spec version: v0.0.1
    Contact: fredbi@yahoo.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import unittest

import swagger_client
from swagger_client.api.semver_api import SemverApi  # noqa: E501
from swagger_client.rest import ApiException


class TestSemverApi(unittest.TestCase):
    """SemverApi unit test stubs"""

    def setUp(self):
        self.api = swagger_client.api.semver_api.SemverApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_series_series_id_semver_get(self):
        """Test case for series_series_id_semver_get

        Get all semver tags associated to a series  # noqa: E501
        """
        pass

    def test_series_series_id_semver_semver_get(self):
        """Test case for series_series_id_semver_semver_get

        Get a version of a time series with a semver tag  # noqa: E501
        """
        pass

    def test_series_series_id_semver_semver_values_get(self):
        """Test case for series_series_id_semver_semver_values_get

        Get the values of version of a time series with a semver tag  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
