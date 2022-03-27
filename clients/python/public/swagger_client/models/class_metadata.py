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


class ClassMetadata(object):
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
        'from_template': 'bool',
        'tag_search': 'bool',
        'extra_fields': 'list[str]',
        'has_one_class': 'dict(str, str)',
        'has_zero_one_class': 'dict(str, str)',
        'has_zero_many_class': 'dict(str, str)'
    }

    attribute_map = {
        'from_template': 'fromTemplate',
        'tag_search': 'tagSearch',
        'extra_fields': 'extraFields',
        'has_one_class': 'hasOneClass',
        'has_zero_one_class': 'hasZeroOneClass',
        'has_zero_many_class': 'hasZeroManyClass'
    }

    def __init__(self, from_template=None, tag_search=None, extra_fields=None, has_one_class=None, has_zero_one_class=None, has_zero_many_class=None, _configuration=None):  # noqa: E501
        """ClassMetadata - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._from_template = None
        self._tag_search = None
        self._extra_fields = None
        self._has_one_class = None
        self._has_zero_one_class = None
        self._has_zero_many_class = None
        self.discriminator = None

        if from_template is not None:
            self.from_template = from_template
        if tag_search is not None:
            self.tag_search = tag_search
        if extra_fields is not None:
            self.extra_fields = extra_fields
        if has_one_class is not None:
            self.has_one_class = has_one_class
        if has_zero_one_class is not None:
            self.has_zero_one_class = has_zero_one_class
        if has_zero_many_class is not None:
            self.has_zero_many_class = has_zero_many_class

    @property
    def from_template(self):
        """Gets the from_template of this ClassMetadata.  # noqa: E501


        :return: The from_template of this ClassMetadata.  # noqa: E501
        :rtype: bool
        """
        return self._from_template

    @from_template.setter
    def from_template(self, from_template):
        """Sets the from_template of this ClassMetadata.


        :param from_template: The from_template of this ClassMetadata.  # noqa: E501
        :type: bool
        """

        self._from_template = from_template

    @property
    def tag_search(self):
        """Gets the tag_search of this ClassMetadata.  # noqa: E501


        :return: The tag_search of this ClassMetadata.  # noqa: E501
        :rtype: bool
        """
        return self._tag_search

    @tag_search.setter
    def tag_search(self, tag_search):
        """Sets the tag_search of this ClassMetadata.


        :param tag_search: The tag_search of this ClassMetadata.  # noqa: E501
        :type: bool
        """

        self._tag_search = tag_search

    @property
    def extra_fields(self):
        """Gets the extra_fields of this ClassMetadata.  # noqa: E501


        :return: The extra_fields of this ClassMetadata.  # noqa: E501
        :rtype: list[str]
        """
        return self._extra_fields

    @extra_fields.setter
    def extra_fields(self, extra_fields):
        """Sets the extra_fields of this ClassMetadata.


        :param extra_fields: The extra_fields of this ClassMetadata.  # noqa: E501
        :type: list[str]
        """

        self._extra_fields = extra_fields

    @property
    def has_one_class(self):
        """Gets the has_one_class of this ClassMetadata.  # noqa: E501


        :return: The has_one_class of this ClassMetadata.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._has_one_class

    @has_one_class.setter
    def has_one_class(self, has_one_class):
        """Sets the has_one_class of this ClassMetadata.


        :param has_one_class: The has_one_class of this ClassMetadata.  # noqa: E501
        :type: dict(str, str)
        """

        self._has_one_class = has_one_class

    @property
    def has_zero_one_class(self):
        """Gets the has_zero_one_class of this ClassMetadata.  # noqa: E501


        :return: The has_zero_one_class of this ClassMetadata.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._has_zero_one_class

    @has_zero_one_class.setter
    def has_zero_one_class(self, has_zero_one_class):
        """Sets the has_zero_one_class of this ClassMetadata.


        :param has_zero_one_class: The has_zero_one_class of this ClassMetadata.  # noqa: E501
        :type: dict(str, str)
        """

        self._has_zero_one_class = has_zero_one_class

    @property
    def has_zero_many_class(self):
        """Gets the has_zero_many_class of this ClassMetadata.  # noqa: E501


        :return: The has_zero_many_class of this ClassMetadata.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._has_zero_many_class

    @has_zero_many_class.setter
    def has_zero_many_class(self, has_zero_many_class):
        """Sets the has_zero_many_class of this ClassMetadata.


        :param has_zero_many_class: The has_zero_many_class of this ClassMetadata.  # noqa: E501
        :type: dict(str, str)
        """

        self._has_zero_many_class = has_zero_many_class

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
        if issubclass(ClassMetadata, dict):
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
        if not isinstance(other, ClassMetadata):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ClassMetadata):
            return True

        return self.to_dict() != other.to_dict()
