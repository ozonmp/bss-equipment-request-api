from python_api_test.api_method.api_method import APIClient


def create_equipment_request(employee_id, equipment_id, created_at,
                             equipment_request_status="EQUIPMENT_REQUEST_STATUS_UNSPECIFIED",
                             updated_at=None, deleted_at=None, done_at=None):

    data = {
        "employeeId": employee_id,
        "equipmentId": equipment_id,
        "createdAt": created_at,
        "updatedAt": updated_at,
        "deletedAt": deleted_at,
        "doneAt": done_at,
        "equipmentRequestStatus": equipment_request_status
    }
    equipment_request_id = APIClient().create(params=data).json()["equipmentRequestId"]
    return equipment_request_id
