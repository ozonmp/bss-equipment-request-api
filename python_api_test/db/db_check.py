import psycopg2

from python_api_test.settings import HOST, DB_PORT, POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD

equipmentRequestTable = "equipment_request"
equipmentRequestIDColumn = "id"
equipmentRequestEquipmentIDColumn = "equipment_id"
equipmentRequestEmployeeIDColumn = "employee_id"
equipmentRequestStatusColumn = "equipment_request_status"
equipmentRequestUpdatedAtColumn = "updated_at"
equipmentRequestCreatedAtColumn = "created_at"
equipmentRequestDoneAtColumn = "done_at"
equipmentRequestDeletedAtAtColumn = "deleted_at"


class DBSelect:
    connection = psycopg2.connect(user=POSTGRES_USER, password=POSTGRES_PASSWORD, host=HOST, port=DB_PORT,
                                  database=POSTGRES_DB)
    cursor = connection.cursor()

    def get_last_equipment_request_id(self, employee_id, equipment_id):
        self.cursor.execute(
            f"""select {equipmentRequestIDColumn},{equipmentRequestEmployeeIDColumn},
            {equipmentRequestEquipmentIDColumn},{equipmentRequestCreatedAtColumn},
            {equipmentRequestUpdatedAtColumn},{equipmentRequestDoneAtColumn},{equipmentRequestDeletedAtAtColumn},
            {equipmentRequestStatusColumn}
            from {equipmentRequestTable} 
            where {equipmentRequestEmployeeIDColumn}={employee_id}
            and {equipmentRequestEquipmentIDColumn}={equipment_id}order by {equipmentRequestIDColumn} desc limit 1""")
        equipment_id = self.cursor.fetchone()
        return equipment_id

    def get_all_equipment_request_id(self, limit, offset):
        self.cursor.execute(f"""select * from {equipmentRequestTable} where {equipmentRequestDeletedAtAtColumn} is null 
                                order by {equipmentRequestIDColumn} asc limit {limit} offset {offset}""")
        equipment_request_id = self.cursor.fetchall()
        return equipment_request_id

    def get_remove_equipment_request(self, id):
        self.cursor.execute(
            f"""select {equipmentRequestIDColumn}, {equipmentRequestDeletedAtAtColumn} from {equipmentRequestTable} 
            where {equipmentRequestIDColumn}={id}""")
        remove_equipment_request = self.cursor.fetchall()
        return remove_equipment_request

    def get_update_equipment_request(self, equipment_request_id, equipment_id):
        self.cursor.execute(
            f"""select {equipmentRequestIDColumn},  {equipmentRequestEquipmentIDColumn}, 
            {equipmentRequestUpdatedAtColumn} 
            from {equipmentRequestTable} 
            where {equipmentRequestIDColumn}={equipment_request_id} 
            and {equipmentRequestEquipmentIDColumn}={equipment_id}""")
        update_equipment_request = self.cursor.fetchall()
        return update_equipment_request

    def get_update_status_equipment_request(self, id):
        self.cursor.execute(
            f"""select {equipmentRequestIDColumn}, {equipmentRequestStatusColumn}, {equipmentRequestUpdatedAtColumn} 
            from {equipmentRequestTable}
            where {equipmentRequestIDColumn}={id}""")
        equipment_request_update_status = self.cursor.fetchall()
        return equipment_request_update_status
