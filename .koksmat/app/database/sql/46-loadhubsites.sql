--CREATE TABLE sharepoint_sites (LIKE original_table_name INCLUDING ALL);
-- INSERT INTO
--     public.importsharepointsite (
--         created_at,
--         updated_at,
--         deleted_at,
--         tenant,
--         name,
--         description,
--         isteamsconnected,
--         localeid,
--         hubsiteid,
--         status,
--         url,
--         title,
--         lastcontentmodified,
--         releatedgroupid,
--         groupid,
--         isteamschannelconnected,
--         ishubsite,
--         sharingcapability,
--         OWNER
--     )
UPDATE importsharepointsite set ishubsite = true where url in (
SELECT

    
    (page_element ->> 'SiteUrl') AS SiteUrl --,
    -- (page_element ->> 'SiteUrl') AS SiteUrl,
    -- (page_element ->> 'Title') AS Title
    
FROM
    (
        SELECT
            jsonb_array_elements(DATA) AS page_element
        FROM
            public.importdata
        WHERE
            name LIKE '%hubsites.json%'
    ) AS subquery_alias)