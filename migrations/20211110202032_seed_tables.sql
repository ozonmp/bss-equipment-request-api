-- +goose Up

INSERT INTO equipment_request (employee_id, equipment_id, created_at, updated_at, equipment_request_status, deleted_at)
SELECT
    (random() * 10 + 5)::int4,
    (random() * 10 + 5)::int4,
    now() - '2 years'::interval * random(),
    (CASE
         WHEN random() < 0.3 THEN now() - '2 years'::interval * random()
    ELSE NULL
    END),
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_STATUS_DO'::equipment_request_status
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_STATUS_IN_PROGRESS'::equipment_request_status
         WHEN random() < 0.7 THEN 'EQUIPMENT_REQUEST_STATUS_DONE'::equipment_request_status
        ELSE 'EQUIPMENT_REQUEST_STATUS_CANCELLED'::equipment_request_status
    END),
    (CASE
        WHEN random() < 0.3 THEN now() - '2 years'::interval * random()
        ELSE NULL
    END)
FROM
    generate_series(1, 10000);

UPDATE equipment_request
SET
    done_at = now() - '2 years'::interval * random()
WHERE
    equipment_request_status = 'EQUIPMENT_REQUEST_STATUS_DONE'::equipment_request_status;


INSERT INTO equipment_request_event (equipment_request_id, type, status, created_at, updated_at, payload)
SELECT
    (random() * 10 + 5)::int4,
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_CREATED'::equipment_request_event_type
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_EQUIPMENT_ID'::equipment_request_event_type
         WHEN random() < 0.7 THEN 'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_STATUS'::equipment_request_event_type
         ELSE 'EQUIPMENT_REQUEST_EVENT_TYPE_DELETED'::equipment_request_event_type
        END),
    (CASE
         WHEN random() < 0.3 THEN 'EQUIPMENT_REQUEST_EVENT_STATUS_UNLOCKED'::equipment_request_event_status
         WHEN random() < 0.5 THEN 'EQUIPMENT_REQUEST_EVENT_STATUS_LOCKED'::equipment_request_event_status
         ELSE 'EQUIPMENT_REQUEST_EVENT_STATUS_PROCESSED'::equipment_request_event_status
        END),
    now() - '2 years'::interval * random(),
    (CASE
         WHEN random() < 0.3 THEN now() - '2 years'::interval * random()
    ELSE NULL
    END),
    null
FROM
    generate_series(1, 10000);

-- +goose Down
DROP TABLE equipment_request;
DROP TABLE equipment_request_event;

DROP TYPE equipment_request_status;
DROP TYPE equipment_request_event_type;
DROP TYPE equipment_request_event_status;