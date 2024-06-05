/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.importsharepointsite
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,isteamsconnected boolean  
    ,localeid character varying COLLATE pg_catalog."default" 
    ,hubsiteid character varying COLLATE pg_catalog."default" 
    ,status character varying COLLATE pg_catalog."default" 
    ,url character varying COLLATE pg_catalog."default" 
    ,title character varying COLLATE pg_catalog."default" 
    ,lastcontentmodified character varying COLLATE pg_catalog."default"  
    ,releatedgroupid character varying COLLATE pg_catalog."default" 
    ,groupid character varying COLLATE pg_catalog."default" 
    ,isteamschannelconnected boolean  
    ,ishubsite boolean  
    ,sharingcapability character varying COLLATE pg_catalog."default" 
    ,owner character varying COLLATE pg_catalog."default" 


);




---- create above / drop below ----

DROP TABLE public.importsharepointsite;

