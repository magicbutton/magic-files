create view "sharepoint"."sites" as
select * from container where
 platform_id in (select id from platform where name='sharepoint')