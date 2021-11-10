-- +goose Up
CREATE TYPE equipment_request_status AS ENUM (
    'EQUIPMENT_REQUEST_STATUS_DO',
    'EQUIPMENT_REQUEST_STATUS_IN_PROGRESS',
    'EQUIPMENT_REQUEST_STATUS_DONE',
    'EQUIPMENT_REQUEST_STATUS_CANCELLED'
);

CREATE TYPE equipment_request_event_type AS ENUM (
    'EQUIPMENT_REQUEST_EVENT_TYPE_CREATED',
    'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_EQUIPMENT_ID',
    'EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_STATUS',
    'EQUIPMENT_REQUEST_EVENT_TYPE_DELETED'
);

CREATE TYPE equipment_request_event_status AS ENUM (
    'EQUIPMENT_REQUEST_EVENT_STATUS_UNLOCKED',
    'EQUIPMENT_REQUEST_EVENT_STATUS_LOCKED',
    'EQUIPMENT_REQUEST_EVENT_STATUS_PROCESSED'
);

CREATE TABLE equipment_request (
    id BIGSERIAL PRIMARY KEY,
    employee_id bigint NOT NULL, -- assumption, that other service will be responsible for employees
    equipment_id bigint NOT NULL,  -- assumption, that other service will be responsible for equipments (equipment_type moved to this service)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    done_at TIMESTAMPTZ,
    equipment_request_status equipment_request_status NOT NULL DEFAULT 'EQUIPMENT_REQUEST_STATUS_DO',
    deleted_at TIMESTAMPTZ
);

CREATE INDEX equipment_request_employee_id_idx ON equipment_request USING btree (employee_id);
CREATE INDEX equipment_request_equipment_id_idx ON equipment_request USING btree (equipment_id);
CREATE INDEX equipment_request_equipment_request_status_idx ON equipment_request USING btree (equipment_request_status);

CREATE TABLE equipment_request_event (
    id BIGSERIAL PRIMARY KEY,
    equipment_request_id bigint NOT NULL,
    type equipment_request_event_type NOT NULL,
    status equipment_request_event_status NOT NULL DEFAULT 'EQUIPMENT_REQUEST_EVENT_STATUS_UNLOCKED',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    payload jsonb,
    CONSTRAINT fk_equipment_request
        FOREIGN KEY(equipment_request_id)
            REFERENCES equipment_request(id)
            ON DELETE RESTRICT
);

CREATE INDEX equipment_request_event_created_at_idx ON equipment_request_event USING btree (created_at);

-- +goose Down
DROP TABLE equipment_request;
DROP TABLE equipment_request_event;

DROP TYPE equipment_request_status;
DROP TYPE equipment_request_event_type;
DROP TYPE equipment_request_event_status;