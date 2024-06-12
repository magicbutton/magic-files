drop schema if exists "seed"
-- cascade;

CREATE SCHEMA "seed";
DELETE FROM importdata where name ilike 'seed%'
{{ template "seed0".}} 

