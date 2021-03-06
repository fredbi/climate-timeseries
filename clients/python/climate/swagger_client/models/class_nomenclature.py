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


class ClassNomenclature(object):
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
        'id': 'int',
        'short_code': 'ClassNomenclatureName',
        'title': 'Translation',
        'description_short': 'Translation',
        'description_long': 'Translation',
        'audit_trail': 'Audit'
    }

    attribute_map = {
        'id': 'id',
        'short_code': 'shortCode',
        'title': 'title',
        'description_short': 'descriptionShort',
        'description_long': 'descriptionLong',
        'audit_trail': 'auditTrail'
    }

    discriminator_value_class_map = {
        'measurement': 'Measurement',
        'owner': 'Owner',
        'constant': 'Constant',
        'source': 'Source',
        'multiplier': 'Multiplier',
        'ostatus': 'Ostatus',
        'status': 'Status',
        'mdomain': 'Mdomain',
        'munit': 'Munit',
        'zone': 'Zone',
        'musystem': 'Musystem',
        'period': 'Period',
        'mdimension': 'Mdimension',
        'role': 'Role',
        'theme': 'Theme',
        'ztype': 'Ztype'
    }

    def __init__(self, id=None, short_code=None, title=None, description_short=None, description_long=None, audit_trail=None, _configuration=None):  # noqa: E501
        """ClassNomenclature - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._id = None
        self._short_code = None
        self._title = None
        self._description_short = None
        self._description_long = None
        self._audit_trail = None
        self.discriminator = 'shortCode'

        self.id = id
        self.short_code = short_code
        self.title = title
        if description_short is not None:
            self.description_short = description_short
        if description_long is not None:
            self.description_long = description_long
        if audit_trail is not None:
            self.audit_trail = audit_trail

    @property
    def id(self):
        """Gets the id of this ClassNomenclature.  # noqa: E501


        :return: The id of this ClassNomenclature.  # noqa: E501
        :rtype: int
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this ClassNomenclature.


        :param id: The id of this ClassNomenclature.  # noqa: E501
        :type: int
        """
        if self._configuration.client_side_validation and id is None:
            raise ValueError("Invalid value for `id`, must not be `None`")  # noqa: E501

        self._id = id

    @property
    def short_code(self):
        """Gets the short_code of this ClassNomenclature.  # noqa: E501


        :return: The short_code of this ClassNomenclature.  # noqa: E501
        :rtype: ClassNomenclatureName
        """
        return self._short_code

    @short_code.setter
    def short_code(self, short_code):
        """Sets the short_code of this ClassNomenclature.


        :param short_code: The short_code of this ClassNomenclature.  # noqa: E501
        :type: ClassNomenclatureName
        """
        if self._configuration.client_side_validation and short_code is None:
            raise ValueError("Invalid value for `short_code`, must not be `None`")  # noqa: E501

        self._short_code = short_code

    @property
    def title(self):
        """Gets the title of this ClassNomenclature.  # noqa: E501


        :return: The title of this ClassNomenclature.  # noqa: E501
        :rtype: Translation
        """
        return self._title

    @title.setter
    def title(self, title):
        """Sets the title of this ClassNomenclature.


        :param title: The title of this ClassNomenclature.  # noqa: E501
        :type: Translation
        """
        if self._configuration.client_side_validation and title is None:
            raise ValueError("Invalid value for `title`, must not be `None`")  # noqa: E501

        self._title = title

    @property
    def description_short(self):
        """Gets the description_short of this ClassNomenclature.  # noqa: E501


        :return: The description_short of this ClassNomenclature.  # noqa: E501
        :rtype: Translation
        """
        return self._description_short

    @description_short.setter
    def description_short(self, description_short):
        """Sets the description_short of this ClassNomenclature.


        :param description_short: The description_short of this ClassNomenclature.  # noqa: E501
        :type: Translation
        """

        self._description_short = description_short

    @property
    def description_long(self):
        """Gets the description_long of this ClassNomenclature.  # noqa: E501


        :return: The description_long of this ClassNomenclature.  # noqa: E501
        :rtype: Translation
        """
        return self._description_long

    @description_long.setter
    def description_long(self, description_long):
        """Sets the description_long of this ClassNomenclature.


        :param description_long: The description_long of this ClassNomenclature.  # noqa: E501
        :type: Translation
        """

        self._description_long = description_long

    @property
    def audit_trail(self):
        """Gets the audit_trail of this ClassNomenclature.  # noqa: E501


        :return: The audit_trail of this ClassNomenclature.  # noqa: E501
        :rtype: Audit
        """
        return self._audit_trail

    @audit_trail.setter
    def audit_trail(self, audit_trail):
        """Sets the audit_trail of this ClassNomenclature.


        :param audit_trail: The audit_trail of this ClassNomenclature.  # noqa: E501
        :type: Audit
        """

        self._audit_trail = audit_trail

    def get_real_child_model(self, data):
        """Returns the real base class specified by the discriminator"""
        discriminator_value = data[self.discriminator].lower()
        return self.discriminator_value_class_map.get(discriminator_value)

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
        if issubclass(ClassNomenclature, dict):
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
        if not isinstance(other, ClassNomenclature):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ClassNomenclature):
            return True

        return self.to_dict() != other.to_dict()
