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
from swagger_client.api.series_api import SeriesApi  # noqa: E501
from swagger_client.rest import ApiException


class TestSeriesApi(unittest.TestCase):
    """SeriesApi unit test stubs"""

    def setUp(self):
        self.api = swagger_client.api.series_api.SeriesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_series_get(self):
        """Test case for series_get

        Get all available series, with search filters  # noqa: E501
        """
        pass

    def test_series_post(self):
        """Test case for series_post

        Creates a new time series  # noqa: E501
        """
        pass

    def test_series_series_id_delete(self):
        """Test case for series_series_id_delete

        Deletes a time series  # noqa: E501
        """
        pass

    def test_series_series_id_get(self):
        """Test case for series_series_id_get

        Get all versions of a time series.  # noqa: E501
        """
        pass

    def test_series_series_id_latest_get(self):
        """Test case for series_series_id_latest_get

        Get the latest version of the description of a time series  # noqa: E501
        """
        pass

    def test_series_series_id_latest_values_get(self):
        """Test case for series_series_id_latest_values_get

        Get the values from the latest version of a time series  # noqa: E501
        """
        pass

    def test_series_series_id_post(self):
        """Test case for series_series_id_post

        Creates a new version of a time series  # noqa: E501
        """
        pass

    def test_series_series_id_put(self):
        """Test case for series_series_id_put

        Updates metadata about a time series  # noqa: E501
        """
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

    def test_series_versions_versioned_series_id_delete(self):
        """Test case for series_versions_versioned_series_id_delete

        Deletes a version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_get(self):
        """Test case for series_versions_versioned_series_id_get

        Get a version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_post(self):
        """Test case for series_versions_versioned_series_id_post

        creates a new version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_put(self):
        """Test case for series_versions_versioned_series_id_put

        Updates the metadata of version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_values_delete(self):
        """Test case for series_versions_versioned_series_id_values_delete

        Deletes the values attached to a version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_values_get(self):
        """Test case for series_versions_versioned_series_id_values_get

        Get the values of version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_values_post(self):
        """Test case for series_versions_versioned_series_id_values_post

        Uploads values to the version of a time series  # noqa: E501
        """
        pass

    def test_series_versions_versioned_series_id_values_put(self):
        """Test case for series_versions_versioned_series_id_values_put

        Replaces the values of version of a time series  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
