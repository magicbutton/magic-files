-- DROP PROCEDURE IF EXISTS public.insert_sharepoint_sites;

CREATE OR REPLACE PROCEDURE public.insert_sharepoint_sites()
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE
    record RECORD;
    platform_id INT;

BEGIN
    -- Fetch the platform_id for 'sharepoint'
    SELECT id INTO platform_id FROM platform WHERE name = 'sharepoint';

    FOR record IN
        SELECT
        
        data_element -> 'data' ->> 'siteid' AS siteid,
            data_element -> 'data' ->> 'title' AS title,
            data_element -> 'data' ->> 'weburl' AS weburl,
            data_element -> 'data' ->> 'groupid' AS groupid,
            data_element -> 'data' ->> 'localeid' AS localeid,
            data_element -> 'data' ->> 'batchname' AS batchname,
            data_element -> 'data' ->> 'hubsiteid' AS hubsiteid,
            data_element -> 'data' ->> 'ishubsite' AS ishubsite,
            data_element -> 'data' ->> 'relatedgroupid' AS relatedgroupid,
            data_element -> 'data' ->> 'isteamsconnected' AS isteamsconnected,
            data_element -> 'data' ->> 'sharingcapability' AS sharingcapability,
            data_element -> 'data' ->> 'isteamschannelconnected' AS isteamschannelconnected,
            data_element -> 'data' ->> 'lastcontentmodifieddate' AS lastcontentmodifieddate,
            name AS batchname
        FROM
            importdata,
            jsonb_array_elements(data) AS data_element
        WHERE
            name ILIKE 'sharepoint/allsites%'
            and  data_element -> 'data' ->> 'siteid' IS NOT NULL
            AND name NOT IN (
                SELECT DISTINCT batchname
                FROM container
                WHERE batchname IS NOT NULL
            )
    LOOP
        INSERT INTO public.container(
            id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, 
            name, description, url, searchindex, type, batchname, platform_id
        )
        VALUES (
            DEFAULT, -- Assuming id is a serial or generated column
            NOW(), -- Assuming created_at is current timestamp
            '', -- Assuming created_by is not provided
            DEFAULT, -- Assuming updated_at is not set initially
            '', -- Assuming updated_by is not set initially
            DEFAULT, -- Assuming deleted_at is not set initially
            '', -- Assuming tenant is not provided
            record.siteid, -- Using title from record
            '', -- Assuming description is not set initially
            record.weburl, -- Using weburl from record
            record.siteid, -- Assuming searchindex is not set initially
            'site', -- Assuming type is 'site'
            record.batchname, -- Using batchname from record
            platform_id -- Using fetched platform_id
        );
    END LOOP;
END;
$BODY$;
ALTER PROCEDURE public.insert_sharepoint_sites()
    OWNER TO pgadmin;
