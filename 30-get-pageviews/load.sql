CREATE TABLE IF NOT EXISTS public.pageviews(
    CreationTime timestamp,
    Workload varchar(50),
    ClientIP varchar(50),
    UserId varchar(255),
    BrowserName varchar(50),
    BrowserVersion varchar(50),
    Platform varchar(50),
    Site varchar(50),
    PageUrl varchar(255)
);

INSERT INTO public.pageviews(CreationTime, Workload, ClientIP, UserId, BrowserName, BrowserVersion, Platform, Site, PageUrl)
SELECT
    creationtime,
    workload,
    clientip,
    userid,
    browsername,
    browserversion,
    platform,
    site,
    objectid
FROM
    import.filtered_auditlogs;

