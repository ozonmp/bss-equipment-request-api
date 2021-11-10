-- +goose Up

INSERT INTO equipment_request (employee_id, equipment_id, created_at, updated_at, equipment_request_status, deleted_at)
SELECT
    (random() * 10 + 5)::int4,
    (random() * 10 + 5)::int4,
    now() - '2 years'::interval * random(),
    now() - '2 years'::interval * random(),
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_STATUS_DO'
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_STATUS_IN_PROGRESS'
         WHEN random() < 0.7 THEN 'EQUIPMENT_REQUEST_STATUS_DONE'
        ELSE 'EQUIPMENT_REQUEST_STATUS_CANCELLED'
    END),
    (CASE
        WHEN random() < 0.3 THEN now() - '2 years'::interval * random()
        ELSE null
    END)
FROM
    generate_series(1, 10000);

UPDATE equipment_request
SET
    done_at = now() - '2 years'::interval * random(),
WHERE
    equipment_request_status = 'EQUIPMENT_REQUEST_STATUS_DONE';


INSERT INTO equipment_request_event (equipment_request_id, type, status, created_at, updated_at, payload)
SELECT
    (random() * 10 + 5)::int4,
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_CREATED'
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_EQUIPMENT_ID'
         WHEN random() < 0.7 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_STATUS'
         ELSE 'EQUIPMENT_REQUEST_EVENT_TYPE_DELETED'
        END),
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_EVENT_STATUS_UNLOCKED'
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_EVENT_STATUS_LOCKED'
         ELSE 'EQUIPMENT_REQUEST_EVENT_STATUS_PROCESSED'
        END),
    now() - '2 years'::interval * random(),
    now() - '2 years'::interval * random(),
    null
FROM
    generate_series(1, 10000);

-- +goose Down
DROP TABLE equipment_request;
DROP TABLE equipment_request_event;

DROP TYPE equipment_request_status;
DROP TYPE equipment_request_event_type;
DROP TYPE equipment_request_event_status;