
CREATE OR REPLACE FUNCTION get_or_create_user(lookupemail TEXT)
RETURNS INT AS $$
DECLARE
    user_id INT;
BEGIN
    -- Check if user exists
    SELECT id INTO user_id FROM public."user" WHERE email = lookupemail;

    -- If user not found, insert a new user with default values
    IF NOT FOUND THEN
    INSERT INTO public."user"(
	id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, name, description, url, searchindex, email, batchname)

      VALUES (
            DEFAULT, -- Assuming id is a serial or generated column
            NOW(), -- Assuming created_at is current timestamp
            '', -- Assuming created_by is userid
            DEFAULT, -- Assuming updated_at is not set initially
            '', -- Assuming updated_by is not set initially
            DEFAULT, -- Assuming deleted_at is not set initially
            '', -- Assuming tenant is not set initially
            lookupemail, -- Assuming name is not set initially
            '', -- Assuming description is not set initially
            '', -- Assuming url is not set initially
            lookupemail, 
            lookupemail, 
            ''
        )
        RETURNING id INTO user_id;
    END IF; 

    -- Return the user id
    RETURN user_id;
END;
$$ LANGUAGE plpgsql;
