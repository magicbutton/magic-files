/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.relation
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
    ,user_id int   NOT NULL
    ,date character varying COLLATE pg_catalog."default"   NOT NULL
    ,type character varying COLLATE pg_catalog."default"  NOT NULL


);

                ALTER TABLE IF EXISTS public.relation
                ADD FOREIGN KEY (user_id)
                REFERENCES public.user (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                -- lollipop
                CREATE TABLE public.relation_m2m_file (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                created_by character varying COLLATE pg_catalog."default"  ,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_by character varying COLLATE pg_catalog."default",
                deleted_at timestamp with time zone
                
                    ,relation_id int  
 
                    ,file_id int  
 

                );
            

                ALTER TABLE public.relation_m2m_file
                ADD FOREIGN KEY (relation_id)
                REFERENCES public.relation (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.relation_m2m_file
                ADD FOREIGN KEY (file_id)
                REFERENCES public.file (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----
DROP TABLE IF EXISTS public.relation_m2m_file;
DROP TABLE public.relation;

