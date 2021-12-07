import datetime

import pytest
from hamcrest.core import assert_that, equal_to
from requests import codes

from python_api_test.api_method.api_method import APIClient
from python_api_test.helpers.create_equipment_request import create_equipment_request
from python_api_test.db.db_check import DBSelect

api_client = APIClient()
db_client = DBSelect()


@pytest.mark.parametrize(
    "employee_id, equipment_id, created_at, updated_at, deleted_at, done_at, equipment_request_status",
    [("2147483647", "15", "2021-12-31T23:44:02.361Z", None, None, None,
      "EQUIPMENT_REQUEST_STATUS_UNSPECIFIED"),
     ("1", "2", "2021-12-01T00:00:00.000Z", "2021-12-02T00:00:00.000Z", None, None, "EQUIPMENT_REQUEST_STATUS_DO"),
     ("15", "100", "2021-01-01T23:44:02.361Z", None, None, None, "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     ("15", "15", "2021-01-01T23:44:02.361Z", "2021-12-02T00:00:00.000Z", None, "2021-12-03T12:00:00.000Z",
      "EQUIPMENT_REQUEST_STATUS_DONE"),
     ("3", "15", "2021-01-01T23:44:02.361Z", None, "2021-01-31T23:44:02.361Z", None,
      "EQUIPMENT_REQUEST_STATUS_CANCELLED")])
def test_create(employee_id, equipment_id, created_at, updated_at, deleted_at, done_at, equipment_request_status):
    data = {
        "employeeId": employee_id,
        "equipmentId": equipment_id,
        "createdAt": created_at,
        "updatedAt": updated_at,
        "deletedAt": deleted_at,
        "doneAt": done_at,
        "equipmentRequestStatus": equipment_request_status
    }
    response = api_client.create(params=data)
    status_code = response.status_code
    response = response.json()

    equipment_request_id = db_client.get_last_equipment_request_id(employee_id=employee_id, equipment_id=equipment_id)

    assert_that(status_code, equal_to(codes.ok))
    assert_that(response["equipmentRequestId"], equal_to(str(equipment_request_id[0])))
    assert_that(str(equipment_request_id[1]), equal_to(employee_id))
    assert_that(str(equipment_request_id[2]), equal_to(equipment_id))
    assert_that(equipment_request_id[3].replace(tzinfo=None),
                equal_to(datetime.datetime.strptime(created_at, "%Y-%m-%dT%H:%M:%S.%fZ")))
    if equipment_request_id[4] is not None:
        assert_that(equipment_request_id[4].replace(tzinfo=None),
                    equal_to(datetime.datetime.strptime(updated_at, "%Y-%m-%dT%H:%M:%S.%fZ")))
    if equipment_request_id[5] is not None:
        assert_that(equipment_request_id[5].replace(tzinfo=None),
                    equal_to(datetime.datetime.strptime(done_at, "%Y-%m-%dT%H:%M:%S.%fZ")))
    if equipment_request_id[6] is not None:
        assert_that(equipment_request_id[6].replace(tzinfo=None),
                    equal_to(datetime.datetime.strptime(deleted_at, "%Y-%m-%dT%H:%M:%S.%fZ")))
    assert_that(equipment_request_id[7], equal_to(equipment_request_status))


@pytest.mark.parametrize("limit, offset", [("5", "5"), ("10", "0"), ("1", "200")])
def test_should_all_equipment_requests(limit, offset):
    data = {
        "limit": limit,
        "offset": offset
    }
    response = api_client.get_list(params=data)
    status_code = response.status_code
    response = response.json()["items"]
    all_equipment_request = db_client.get_all_equipment_request_id(limit=limit, offset=offset)

    assert_that(status_code, equal_to(codes.ok))
    assert_that(len(response), equal_to(int(limit)))
    for idx in range(len(response)):
        assert int(response[idx]["id"]) == int(all_equipment_request[idx][0])
        assert int(response[idx]["employeeId"]) == int(all_equipment_request[idx][1])
        assert int(response[idx]["equipmentId"]) == int(all_equipment_request[idx][2])
        assert response[idx]["deletedAt"] is None


@pytest.mark.parametrize("equipment_request_id", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_DO"))])
def test_should_remove(equipment_request_id):
    data = {
        "equipmentRequestId": equipment_request_id
    }
    response = api_client.remove_request(params=data)
    status_code = response.status_code
    time = response.headers["Date"]
    response = response.json()
    remove_equipment_request = db_client.get_remove_equipment_request(id=equipment_request_id)

    assert_that(status_code, equal_to(codes.ok))
    assert_that(response["removed"], equal_to(True))
    for elem in remove_equipment_request:
        assert_that(str(elem[0]), equal_to(equipment_request_id))
        assert_that(elem[1].replace(tzinfo=None, microsecond=0),
                    equal_to(datetime.datetime.strptime(time, "%a, %d %b %Y %H:%M:%S %Z")))


@pytest.mark.parametrize("equipment_request_id", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_DO"))])
def test_should_get_equipment_request(equipment_request_id):
    data = {
        "equipmentRequestId": equipment_request_id
    }
    response = api_client.describe_request(params=data)
    status_code = response.status_code
    response = response.json()["equipmentRequest"]

    assert_that(status_code, equal_to(codes.ok))
    assert_that(equipment_request_id, equal_to(response["id"]))


@pytest.mark.parametrize("equipment_request_id", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), "30")])
def test_should_update_equipment_id(equipment_request_id, new_equipment_id):
    data = {
        "equipmentRequestId": equipment_request_id,
        "equipmentId": new_equipment_id
    }
    response = api_client.update_equipment_id(params=data)
    status_code = response.status_code
    time = response.headers["Date"]
    response = response.json()
    update_equipment_request = db_client.get_update_equipment_request(equipment_request_id=equipment_request_id,
                                                                      equipment_id=new_equipment_id)
    assert_that(status_code, equal_to(codes.ok))
    assert_that(response["updated"], equal_to(True))
    for elem in update_equipment_request:
        assert_that(str(elem[0]), equal_to(equipment_request_id))
        assert_that(str(elem[1]), equal_to(new_equipment_id))
        assert_that(elem[2].replace(tzinfo=None, microsecond=0),
                    equal_to(datetime.datetime.strptime(time, "%a, %d %b %Y %H:%M:%S %Z")))


@pytest.mark.parametrize("equipment_request_id, equipment_request_status", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"),
     "EQUIPMENT_REQUEST_STATUS_DONE")])
def test_should_update_status(equipment_request_id, equipment_request_status):
    data = {
        "equipmentRequestId": equipment_request_id,
        "equipmentRequestStatus": equipment_request_status
    }
    response = api_client.update_status(params=data)
    status_code = response.status_code
    time = response.headers["Date"]
    response = response.json()
    update_status_equipment_request = db_client.get_update_status_equipment_request(id=equipment_request_id)
    assert_that(status_code, equal_to(codes.ok))
    assert_that(response["updated"], equal_to(True))
    for elem in update_status_equipment_request:
        assert_that(str(elem[0]), equal_to(equipment_request_id))
        assert_that(elem[1], equal_to(equipment_request_status))
        assert_that(elem[2].replace(tzinfo=None, microsecond=0),
                    equal_to(datetime.datetime.strptime(time, "%a, %d %b %Y %H:%M:%S %Z")))


@pytest.mark.parametrize("equipment_request_id, equipment_request_status", [
    (create_equipment_request(
        employee_id="10", equipment_id="15", created_at="2021-12-31T23:59:02.361Z",
        equipment_request_status="EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"), 1)])
def test_should_update_status_enum(equipment_request_id, equipment_request_status):
    data = {
        "equipmentRequestId": equipment_request_id,
        "equipmentRequestStatus": equipment_request_status
    }
    response = api_client.update_status(params=data)
    status = {1: "EQUIPMENT_REQUEST_STATUS_DO"}
    status_code = response.status_code
    response = response.json()
    update_status_equipment_request = db_client.get_update_status_equipment_request(id=equipment_request_id)
    assert_that(status_code, equal_to(codes.ok))
    assert_that(response["updated"], equal_to(True))
    for elem in update_status_equipment_request:
        assert_that(str(elem[0]), equal_to(equipment_request_id))
        assert_that(elem[1], equal_to(status[1]))
