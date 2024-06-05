do $$
    declare
    arow record;
BEGIN
    FOR arow IN 
    
          select title,
          siteurl,
          HubSiteId,
          (page ->> 'PageName') AS PageName,
          (page ->> 'Editor') AS Editor,
          to_date((page ->> 'CreatedOn'),'YYYY-MM-DD') AS CreatedOn,
          to_date((page ->> 'ModifiedOn'),'YYYY-MM-DD') AS ModifiedOn,
          (page ->> 'RelativeURL') AS RelativeURL
          from (
          SELECT (sitedata ->> 'title') AS title,
                (sitedata ->> 'siteurl') AS siteurl,
                (sitedata ->> 'HubSiteId') AS HubSiteId,
                jsonb_array_elements(sitedata->'pages') AS page
              FROM
                (SELECT name,
                        DATA AS sitedata
                FROM public.importdata
                WHERE name LIKE '%/hubsite-spokes-pages%-sites-%.json' ) AS sites
              ) as pages
    LOOP
      INSERT INTO public.file(id, created_at, updated_at, deleted_at, tenant,type,size, extension,name,description,  url, modified, modifiedby, level1, level2, level3, level4)


        VALUES (DEFAULT, DEFAULT, DEFAULT, DEFAULT, '',
        'page',
        -1,
        '',
        arow.RelativeURL ,
        arow.PageName  ,
        arow.siteurl, -- url
        arow.ModifiedOn, -- modified
        arow.Editor, -- modifiedby

        'intranet', -- level1

        arow.HubSiteId, -- level2
        arow.siteurl, -- level3
        arow.RelativeURL -- level4

        );
      END LOOP;
END;
$$;