# coding: utf-8

"""
    Climate time series API

    The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins.   # noqa: E501

    OpenAPI spec version: v0.0.1
    Contact: fredbi@yahoo.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.configuration import Configuration


class Theme(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'tags': 'list[str]',
        'linked_documents': 'Documents'
    }

    attribute_map = {
        'tags': 'tags',
        'linked_documents': 'linkedDocuments'
    }

    def __init__(self, tags=None, linked_documents=None, _configuration=None):  # noqa: E501
        """Theme - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._tags = None
        self._linked_documents = None
        self.discriminator = None

        if tags is not None:
            self.tags = tags
        if linked_documents is not None:
            self.linked_documents = linked_documents

    @property
    def tags(self):
        """Gets the tags of this Theme.  # noqa: E501


        :return: The tags of this Theme.  # noqa: E501
        :rtype: list[str]
        """
        return self._tags

    @tags.setter
    def tags(self, tags):
        """Sets the tags of this Theme.


        :param tags: The tags of this Theme.  # noqa: E501
        :type: list[str]
        """

        self._tags = tags

    @property
    def linked_documents(self):
        """Gets the linked_documents of this Theme.  # noqa: E501


        :return: The linked_documents of this Theme.  # noqa: E501
        :rtype: Documents
        """
        return self._linked_documents

    @linked_documents.setter
    def linked_documents(self, linked_documents):
        """Sets the linked_documents of this Theme.


        :param linked_documents: The linked_documents of this Theme.  # noqa: E501
        :type: Documents
        """

        self._linked_documents = linked_documents

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(Theme, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Theme):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, Theme):
            return True

        return self.to_dict() != other.to_dict()
