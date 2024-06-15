drop schema if exists "sharepoint";
-- cascade;

CREATE SCHEMA "sharepoint";

{{ template "pageviews.sql".}} 
{{ template "get_or_create_user.sql".}} 
{{ template "find_or_create_file.sql".}} 


{{ template "importviews.sql".}} 