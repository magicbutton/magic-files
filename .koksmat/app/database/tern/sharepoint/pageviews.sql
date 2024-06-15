create view "sharepoint"."pageviews" as
select * from events    
where operation = 'PageViewed';