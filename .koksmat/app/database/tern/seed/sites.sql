

-- do $$
--     declare
--     arow record;
-- BEGIN
--     FOR arow IN SELECT 
--  (page_element ->> 'DisplayName') AS DisplayName,
--   (page_element ->> 'PrimarySmtpAddress') AS PrimarySmtpAddress,
--   (page_element ->> 'Identity') AS Identity

-- FROM
--     (SELECT jsonb_array_elements(DATA) AS page_element
--      FROM public.importdata
--      WHERE name LIKE '%/sharedmailboxes/sharedmailboxes.json'
--          AND jsonb_typeof(DATA) = 'array' ) AS owners
--     LOOP
INSERT INTO public.importsharedmailbox(
	id, created_at, updated_at, deleted_at, tenant, name, description, url, email)
    
 VALUES (DEFAULT, DEFAULT, DEFAULT, DEFAULT, '', arow.DisplayName  , '', arow.Identity, arow.PrimarySmtpAddress);
    END LOOP;
-- END;
-- $$;