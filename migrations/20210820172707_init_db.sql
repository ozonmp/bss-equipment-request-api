-- +goose Up
CREATE TABLE equipment_request (
    id BIGSERIAL PRIMARY KEY,
    employee_id bigint NOT NULL, -- other service responsible for employees
    equipment_id bigint NOT NULL,  -- other service responsible for equipments (equipment_type moved to this service)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    done_at TIMESTAMPTZ,
    equipment_request_status_id bigint,
    CONSTRAINT fk_equipment_request_status
       FOREIGN KEY(equipment_request_status_id)
           REFERENCES equipment_request_status(id)
           ON DELETE SET NULL
);

CREATE TABLE equipment_request_status (
    id BIGSERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

-- +goose Down
DROP TABLE equipment_request;
DROP TABLE equipment_request_status;