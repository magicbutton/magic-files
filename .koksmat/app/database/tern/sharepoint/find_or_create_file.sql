-- FUNCTION: public.find_or_create_file(character varying)

-- DROP FUNCTION IF EXISTS public.find_or_create_file(character varying);

CREATE OR REPLACE FUNCTION public.find_or_create_file(
	findurl character varying)
    RETURNS integer
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
    file_id INT;
BEGIN
    -- Check if file exists with the given URL
    SELECT id INTO file_id FROM public.file WHERE url = findurl;

    -- If file not found, insert a new file with default values
    IF NOT FOUND THEN
        INSERT INTO public.file (
	id, created_at, created_by, updated_at, updated_by, deleted_at, 
            url, tenant, name, searchindex, type, size, extension,
            version, published,
             batchname)
        VALUES (
			DEFAULT,
			DEFAULT,
			'',
			DEFAULT,
			'',			
			DEFAULT,
			
	
            findurl, -- URL from the parameter
            '', -- Replace with actual default value
            findurl, -- Replace with actual default value
            findurl, -- Replace with actual default value
             '', -- Replace with actual default value
            '', -- Replace with actual default value
            '', -- Replace with actual default value
            '', -- Replace with actual default value
            false,
             '' -- Replace with actual default values
        ) RETURNING id INTO file_id;
    END IF;

    -- Return the file id
    RETURN file_id;
END;
$BODY$;

ALTER FUNCTION public.find_or_create_file(character varying)
    OWNER TO pgadmin;
