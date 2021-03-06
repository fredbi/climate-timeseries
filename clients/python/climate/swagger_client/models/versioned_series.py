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


class VersionedSeries(object):
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
        'version': 'Semver',
        'versioned_id': 'str',
        'version_owner': 'str',
        'version_status': 'Vstatus',
        'parent_versioned_id': 'str',
        'formula': 'str',
        'version_status_change_reason': 'Translation',
        'version_title': 'Translation',
        'version_description_short': 'Translation',
        'version_description_long': 'Translation',
        'timeseries': 'TsValues',
        'audit_trail': 'Audit'
    }

    attribute_map = {
        'version': 'version',
        'versioned_id': 'versionedId',
        'version_owner': 'versionOwner',
        'version_status': 'versionStatus',
        'parent_versioned_id': 'parentVersionedId',
        'formula': 'formula',
        'version_status_change_reason': 'versionStatusChangeReason',
        'version_title': 'versionTitle',
        'version_description_short': 'versionDescriptionShort',
        'version_description_long': 'versionDescriptionLong',
        'timeseries': 'timeseries',
        'audit_trail': 'auditTrail'
    }

    def __init__(self, version=None, versioned_id=None, version_owner=None, version_status=None, parent_versioned_id=None, formula=None, version_status_change_reason=None, version_title=None, version_description_short=None, version_description_long=None, timeseries=None, audit_trail=None, _configuration=None):  # noqa: E501
        """VersionedSeries - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._version = None
        self._versioned_id = None
        self._version_owner = None
        self._version_status = None
        self._parent_versioned_id = None
        self._formula = None
        self._version_status_change_reason = None
        self._version_title = None
        self._version_description_short = None
        self._version_description_long = None
        self._timeseries = None
        self._audit_trail = None
        self.discriminator = None

        self.version = version
        self.versioned_id = versioned_id
        if version_owner is not None:
            self.version_owner = version_owner
        self.version_status = version_status
        if parent_versioned_id is not None:
            self.parent_versioned_id = parent_versioned_id
        if formula is not None:
            self.formula = formula
        if version_status_change_reason is not None:
            self.version_status_change_reason = version_status_change_reason
        if version_title is not None:
            self.version_title = version_title
        if version_description_short is not None:
            self.version_description_short = version_description_short
        if version_description_long is not None:
            self.version_description_long = version_description_long
        if timeseries is not None:
            self.timeseries = timeseries
        if audit_trail is not None:
            self.audit_trail = audit_trail

    @property
    def version(self):
        """Gets the version of this VersionedSeries.  # noqa: E501


        :return: The version of this VersionedSeries.  # noqa: E501
        :rtype: Semver
        """
        return self._version

    @version.setter
    def version(self, version):
        """Sets the version of this VersionedSeries.


        :param version: The version of this VersionedSeries.  # noqa: E501
        :type: Semver
        """
        if self._configuration.client_side_validation and version is None:
            raise ValueError("Invalid value for `version`, must not be `None`")  # noqa: E501

        self._version = version

    @property
    def versioned_id(self):
        """Gets the versioned_id of this VersionedSeries.  # noqa: E501


        :return: The versioned_id of this VersionedSeries.  # noqa: E501
        :rtype: str
        """
        return self._versioned_id

    @versioned_id.setter
    def versioned_id(self, versioned_id):
        """Sets the versioned_id of this VersionedSeries.


        :param versioned_id: The versioned_id of this VersionedSeries.  # noqa: E501
        :type: str
        """
        if self._configuration.client_side_validation and versioned_id is None:
            raise ValueError("Invalid value for `versioned_id`, must not be `None`")  # noqa: E501

        self._versioned_id = versioned_id

    @property
    def version_owner(self):
        """Gets the version_owner of this VersionedSeries.  # noqa: E501


        :return: The version_owner of this VersionedSeries.  # noqa: E501
        :rtype: str
        """
        return self._version_owner

    @version_owner.setter
    def version_owner(self, version_owner):
        """Sets the version_owner of this VersionedSeries.


        :param version_owner: The version_owner of this VersionedSeries.  # noqa: E501
        :type: str
        """

        self._version_owner = version_owner

    @property
    def version_status(self):
        """Gets the version_status of this VersionedSeries.  # noqa: E501


        :return: The version_status of this VersionedSeries.  # noqa: E501
        :rtype: Vstatus
        """
        return self._version_status

    @version_status.setter
    def version_status(self, version_status):
        """Sets the version_status of this VersionedSeries.


        :param version_status: The version_status of this VersionedSeries.  # noqa: E501
        :type: Vstatus
        """
        if self._configuration.client_side_validation and version_status is None:
            raise ValueError("Invalid value for `version_status`, must not be `None`")  # noqa: E501

        self._version_status = version_status

    @property
    def parent_versioned_id(self):
        """Gets the parent_versioned_id of this VersionedSeries.  # noqa: E501


        :return: The parent_versioned_id of this VersionedSeries.  # noqa: E501
        :rtype: str
        """
        return self._parent_versioned_id

    @parent_versioned_id.setter
    def parent_versioned_id(self, parent_versioned_id):
        """Sets the parent_versioned_id of this VersionedSeries.


        :param parent_versioned_id: The parent_versioned_id of this VersionedSeries.  # noqa: E501
        :type: str
        """

        self._parent_versioned_id = parent_versioned_id

    @property
    def formula(self):
        """Gets the formula of this VersionedSeries.  # noqa: E501


        :return: The formula of this VersionedSeries.  # noqa: E501
        :rtype: str
        """
        return self._formula

    @formula.setter
    def formula(self, formula):
        """Sets the formula of this VersionedSeries.


        :param formula: The formula of this VersionedSeries.  # noqa: E501
        :type: str
        """

        self._formula = formula

    @property
    def version_status_change_reason(self):
        """Gets the version_status_change_reason of this VersionedSeries.  # noqa: E501


        :return: The version_status_change_reason of this VersionedSeries.  # noqa: E501
        :rtype: Translation
        """
        return self._version_status_change_reason

    @version_status_change_reason.setter
    def version_status_change_reason(self, version_status_change_reason):
        """Sets the version_status_change_reason of this VersionedSeries.


        :param version_status_change_reason: The version_status_change_reason of this VersionedSeries.  # noqa: E501
        :type: Translation
        """

        self._version_status_change_reason = version_status_change_reason

    @property
    def version_title(self):
        """Gets the version_title of this VersionedSeries.  # noqa: E501


        :return: The version_title of this VersionedSeries.  # noqa: E501
        :rtype: Translation
        """
        return self._version_title

    @version_title.setter
    def version_title(self, version_title):
        """Sets the version_title of this VersionedSeries.


        :param version_title: The version_title of this VersionedSeries.  # noqa: E501
        :type: Translation
        """

        self._version_title = version_title

    @property
    def version_description_short(self):
        """Gets the version_description_short of this VersionedSeries.  # noqa: E501


        :return: The version_description_short of this VersionedSeries.  # noqa: E501
        :rtype: Translation
        """
        return self._version_description_short

    @version_description_short.setter
    def version_description_short(self, version_description_short):
        """Sets the version_description_short of this VersionedSeries.


        :param version_description_short: The version_description_short of this VersionedSeries.  # noqa: E501
        :type: Translation
        """

        self._version_description_short = version_description_short

    @property
    def version_description_long(self):
        """Gets the version_description_long of this VersionedSeries.  # noqa: E501


        :return: The version_description_long of this VersionedSeries.  # noqa: E501
        :rtype: Translation
        """
        return self._version_description_long

    @version_description_long.setter
    def version_description_long(self, version_description_long):
        """Sets the version_description_long of this VersionedSeries.


        :param version_description_long: The version_description_long of this VersionedSeries.  # noqa: E501
        :type: Translation
        """

        self._version_description_long = version_description_long

    @property
    def timeseries(self):
        """Gets the timeseries of this VersionedSeries.  # noqa: E501


        :return: The timeseries of this VersionedSeries.  # noqa: E501
        :rtype: TsValues
        """
        return self._timeseries

    @timeseries.setter
    def timeseries(self, timeseries):
        """Sets the timeseries of this VersionedSeries.


        :param timeseries: The timeseries of this VersionedSeries.  # noqa: E501
        :type: TsValues
        """

        self._timeseries = timeseries

    @property
    def audit_trail(self):
        """Gets the audit_trail of this VersionedSeries.  # noqa: E501


        :return: The audit_trail of this VersionedSeries.  # noqa: E501
        :rtype: Audit
        """
        return self._audit_trail

    @audit_trail.setter
    def audit_trail(self, audit_trail):
        """Sets the audit_trail of this VersionedSeries.


        :param audit_trail: The audit_trail of this VersionedSeries.  # noqa: E501
        :type: Audit
        """

        self._audit_trail = audit_trail

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
        if issubclass(VersionedSeries, dict):
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
        if not isinstance(other, VersionedSeries):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, VersionedSeries):
            return True

        return self.to_dict() != other.to_dict()
