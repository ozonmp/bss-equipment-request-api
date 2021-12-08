import requests

from python_api_test.settings import HOST, API_PORT

URL = f"http://{HOST}:{API_PORT}/api/v1"

HEADER = {
    "accept": "application/json"
}


class APIClient:

    def create(self, params, path="/equipment_requests/create"):
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER,
                                 json=params)
        return response

    def get_list(self, params, path="/equipment_requests/list"):
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER,
                                 json=params)
        return response

    def remove_request(self, params, path="/equipment_requests/remove"):
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER,
                                 json=params)
        return response

    def describe_request(self, params, path="/equipment_requests/{equipmentRequestId}"):
        path = path.format(**params)
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER)
        return response

    def update_equipment_id(self, params, path="/update/equipment_id"):
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER,
                                 json=params)
        return response

    def update_status(self, params, path="/update/status"):
        response = requests.post(url=f"{URL}{path}",
                                 headers=HEADER,
                                 json=params)
        return response