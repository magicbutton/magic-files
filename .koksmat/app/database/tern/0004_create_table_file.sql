/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.file
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
    ,type character varying COLLATE pg_catalog."default"  NOT NULL
    ,size character varying COLLATE pg_catalog."default"   NOT NULL
    ,extension character varying COLLATE pg_catalog."default"  NOT NULL
    ,createdby character varying COLLATE pg_catalog."default" 
    ,created character varying COLLATE pg_catalog."default"  
    ,modified character varying COLLATE pg_catalog."default"  
    ,modifiedby character varying COLLATE pg_catalog."default" 
    ,version character varying COLLATE pg_catalog."default" 
    ,published boolean  
    ,level1 character varying COLLATE pg_catalog."default" 
    ,level2 character varying COLLATE pg_catalog."default" 
    ,level3 character varying COLLATE pg_catalog."default" 
    ,level4 character varying COLLATE pg_catalog."default" 
    ,level5 character varying COLLATE pg_catalog."default" 
    ,level6 character varying COLLATE pg_catalog."default" 
    ,level7 character varying COLLATE pg_catalog."default" 
    ,level8 character varying COLLATE pg_catalog."default" 
    ,level9 character varying COLLATE pg_catalog."default" 
    ,level10 character varying COLLATE pg_catalog."default" 
    ,level11 character varying COLLATE pg_catalog."default" 
    ,level12 character varying COLLATE pg_catalog."default" 
    ,level13 character varying COLLATE pg_catalog."default" 
    ,level14 character varying COLLATE pg_catalog."default" 
    ,level15 character varying COLLATE pg_catalog."default" 
    ,level16 character varying COLLATE pg_catalog."default" 
    ,batchname character varying COLLATE pg_catalog."default" 


);




---- create above / drop below ----

DROP TABLE public.file;

