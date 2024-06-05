/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.job
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,url character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,status character varying COLLATE pg_catalog."default"  NOT NULL
    ,prefix character varying COLLATE pg_catalog."default"  NOT NULL
    ,type character varying COLLATE pg_catalog."default"  NOT NULL
    ,data JSONB  


);




---- create above / drop below ----

DROP TABLE public.job;

