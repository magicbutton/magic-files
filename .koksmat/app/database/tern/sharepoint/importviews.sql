CREATE OR REPLACE PROCEDURE insert_audit_records()
LANGUAGE plpgsql
AS $$
DECLARE
    record RECORD;
    user_id INT;
BEGIN
    FOR record IN
        SELECT
            data_element -> 'data' ->> 'site' AS site,
            data_element -> 'data' ->> 'userid' AS userid,
            data_element -> 'data' ->> 'operation' AS operation,
            data_element -> 'data' ->> 'objectid' AS objectid,
            data_element -> 'data' ->> 'platform' AS platform,
            data_element -> 'data' ->> 'creationtime' AS creationtime,
            data_element -> 'data' ->> 'auditdata' AS auditdata,
            name AS batchname
        FROM
            importdata,
            jsonb_array_elements(data) AS data_element
        WHERE
            name ILIKE 'audit/records%'
        LIMIT 10 
    LOOP
        -- Get or create user
        user_id := get_or_create_user(record.userid);
        file_id := find_or_create_file(record.objectid);

        INSERT INTO public.events(
            id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, 
            name, description, url, searchindex, type, data, eventdate, user_id, 
            operation, batchname, file_id)
        VALUES (
            DEFAULT, -- Assuming id is a serial or generated column
            NOW(), -- Assuming created_at is current timestamp
            record.userid, -- Assuming created_by is userid
            DEFAULT, -- Assuming updated_at is not set initially
            '', -- Assuming updated_by is not set initially
            DEFAULT, -- Assuming deleted_at is not set initially
            '', -- Assuming tenant is site
            '', -- Assuming name is not set initially
            '', -- Assuming description is not set initially
            record.objectid, -- Assuming url is not set initially
            '', -- Assuming searchindex is not set initially
            '', -- Assuming type is not set initially
            record.auditdata::jsonb, -- Assuming data is not set initially
            record.creationtime, -- eventdate
            user_id, -- user_id
            record.operation, -- operation
            record.batchname, -- batchname
            file_id 
        );
    END LOOP;
END;
$$;

-- To call the procedure
-- CALL insert_audit_records();