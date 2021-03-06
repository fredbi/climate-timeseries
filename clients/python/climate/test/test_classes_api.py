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
from swagger_client.api.classes_api import ClassesApi  # noqa: E501
from swagger_client.rest import ApiException


class TestClassesApi(unittest.TestCase):
    """ClassesApi unit test stubs"""

    def setUp(self):
        self.api = swagger_client.api.classes_api.ClassesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_classes_class_id_get(self):
        """Test case for classes_class_id_get

        Get metadata about a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_class_id_members_class_member_id_delete(self):
        """Test case for classes_class_id_members_class_member_id_delete

        Removes a member from a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_class_id_members_class_member_id_put(self):
        """Test case for classes_class_id_members_class_member_id_put

        Update a member of a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_class_id_members_get(self):
        """Test case for classes_class_id_members_get

        Get all the members of a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_class_id_members_post(self):
        """Test case for classes_class_id_members_post

        Add a new member in a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_class_id_put(self):
        """Test case for classes_class_id_put

        Update descriptive metadata about a nomenclature class  # noqa: E501
        """
        pass

    def test_classes_get(self):
        """Test case for classes_get

        Get all valid nomenclature classes  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
