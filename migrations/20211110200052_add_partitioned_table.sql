-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION equipment_request_event_insert_trigger()
RETURNS TRIGGER AS $$
DECLARE table_master varchar(255) := 'equipment_request_event';
table_part varchar(255) := '';
BEGIN
	table_part := table_master || '_y' || date_part( 'year', NEW.created_at )::text || '_m' || date_part( 'month', NEW.created_at )::text;
	PERFORM 1 FROM pg_class WHERE relname = table_part LIMIT 1;
	IF NOT FOUND THEN
		EXECUTE ' CREATE TABLE ' || table_part || ' (LIKE ' || table_master || ' INCLUDING ALL)';
EXECUTE 'ALTER TABLE ' || table_part || ' INHERIT ' || table_master ;
END IF;
execute 'INSERT INTO ' || table_part || ' VALUES ( ($1).* ) ' USING NEW;
RETURN NULL;
END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd


CREATE TRIGGER equipment_request_event_insert_trigger
    BEFORE INSERT ON equipment_request_event
    FOR EACH ROW EXECUTE PROCEDURE equipment_request_event_insert_trigger();

-- CREATE TABLE equipment_request_event_y2021_m12 PARTITION OF measurement
--     FOR VALUES FROM ('2021-12-01') TO ('2022-01-01');
--
-- CREATE INDEX equipment_request_event_y2021_m12_created_at ON equipment_request_event_y2021_m12 (created_at);

-- +goose Down
DROP TRIGGER equipment_request_event_insert_trigger ON equipment_request_event;