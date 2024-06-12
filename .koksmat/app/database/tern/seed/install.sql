drop schema if exists "azure"
-- cascade;

CREATE SCHEMA "azure";

{{ template "seed0.sql".}} 
{{ template "view-platforms".}} 
{{ template "sites.sql".}} 
