import pytest
from hamcrest.core import assert_that, equal_to

from python_api_test.api_method.api_method import APIClient
from python_api_test.helpers.create_equipment_request import create_equipment_request

api_client = APIClient()


@pytest.mark.parametrize(
    "employee_id, equipment_id, created_at, updated_at, deleted_at, done_at, equipment_request_status",
    [("0", "1", "2021-12-31T23:44:02.361Z", None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("ba7392eb-1580-4487-89af-0654e9202d92", "1", "2021-12-31T23:44:02.361Z", None, None, None,
      "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("1", "-1", "2021-12-31T23:44:02.361Z", None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("1", None, "2021-12-31T23:44:02.361Z", None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("1", "1", None, None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("1", "1", "2021-12-31", None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("1", "1", "2021-12-31T23:44:02.361Z", None, None, None, "IN_PROGRESS"),
     ("18446744073709551616", "1", "2021-12-31T23:44:02.361Z", None, None, None, "IN_PROGRESS")])
def test_not_create(employee_id, equipment_id, created_at, updated_at, deleted_at, done_at, equipment_request_status):
    data = {
        "employeeId": employee_id,
        "equipmentId": equipment_id,
        "createdAt": created_at,
        "updatedAt": updated_at,
        "deletedAt": deleted_at,
        "doneAt": done_at,
        "equipmentRequestStatus": equipment_request_status
    }
    response = api_client.create(params=data).json()
    assert_that(response["code"], equal_to(3))


@pytest.mark.parametrize("limit, offset", [("0", "5"), ("-1", "5"), ("10", "-1"), ("1", "201"), ("5", "51"),
                                           ("18446744073709551616", "5")])
def test_should_not_get_equipment_requests(limit, offset):
    data = {
        "limit": limit,
        "offset": offset
    }
    response = api_client.get_list(params=data).json()
    assert_that(response["code"], equal_to(3))


@pytest.mark.parametrize("equipment_request_id", [(int(create_equipment_request(
    employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
    equipment_request_status="EQUIPMENT_REQUEST_STATUS_DO")) + 1),
                                                  0, None, -1])
def test_should_not_remove(equipment_request_id):
    data = {
        "equipmentRequestId": equipment_request_id
    }
    response = api_client.remove_request(params=data).json()
    assert_that(response["code"], equal_to(3))


@pytest.mark.parametrize("equipment_request_id, new_equipment_id", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), 0),
    (create_equipment_request(employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
                              equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), None),
    (create_equipment_request(employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
                              equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), -1)])
def test_should_not_update_equipment_id(equipment_request_id, new_equipment_id):
    data = {
        "equipmentRequestId": equipment_request_id,
        "equipmentId": new_equipment_id
    }
    response = api_client.update_equipment_id(params=data).json()
    assert_that(response["code"], equal_to(3))


@pytest.mark.parametrize("equipment_request_id, equipment_request_status", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), None),
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), "DO"),
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), 5)])
def test_should_not_update_status(equipment_request_id, equipment_request_status):
    data = {
        "equipmentRequestId": equipment_request_id,
        "equipmentRequestStatus": equipment_request_status
    }
    response = api_client.update_status(params=data).json()
    assert_that(response["code"], equal_to(3))
