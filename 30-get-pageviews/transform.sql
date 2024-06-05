CREATE SCHEMA IF NOT EXISTS import;

CREATE TABLE IF NOT EXISTS import.filtered_auditlogs(
    RecordType varchar(50),
    CreationDate timestamp,
    UserIds varchar(255),
    Operations varchar(50),
    CreationTime timestamp,
    Operation varchar(50),
    Workload varchar(50),
    ClientIP varchar(50),
    UserId varchar(255),
    BrowserName varchar(50),
    BrowserVersion varchar(50),
    Platform varchar(50),
    Site varchar(50),
    ObjectId varchar(255)
);

INSERT INTO import.filtered_auditlogs(RecordType, CreationDate, UserIds, Operations, CreationTime, Operation, Workload, ClientIP, UserId, BrowserName, BrowserVersion, Platform, Site, ObjectId)
SELECT
    RecordType,
    CreationDate,
    UserIds,
    Operations,
    to_timestamp(CreationTime, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
    Operation,
    Workload,
    ClientIP,
    UserId,
    BrowserName,
    BrowserVersion,
    Platform,
    Site,
    ObjectId
FROM (
    SELECT
        RecordType,
        CreationDate,
        UserIds,
        Operations,
        --AuditData->>'AppAccessContext' AS AppAccessContext,
        AuditData ->> 'CreationTime' AS CreationTime,
        --AuditData->>'Id' AS Id,
        AuditData ->> 'Operation' AS Operation,
        --AuditData->>'OrganizationId' AS OrganizationId,
        --AuditData->>'UserKey' AS UserKey,
        --AuditData->>'UserType' AS UserType,
        --AuditData->>'Version' AS Version,
        AuditData ->> 'Workload' AS Workload,
        AuditData ->> 'ClientIP' AS ClientIP,
        AuditData ->> 'UserId' AS UserId,
        --AuditData->>'AuthenticationType' AS AuthenticationType,
        AuditData ->> 'BrowserName' AS BrowserName,
        AuditData ->> 'BrowserVersion' AS BrowserVersion,
        AuditData ->> 'Platform' AS Platform,
        AuditData ->> 'Site' AS Site,
        --AuditData->>'UserAgent' AS UserAgent,
        AuditData ->> 'ObjectId' AS ObjectId
        --AuditData->>'AssociatedAdminUnits' AS AssociatedAdminUnits
    FROM
        import.auditlogs
    WHERE
        operations = 'PageViewed') AS pages
WHERE
    ObjectId LIKE '%sharepoint.com/sites/%'
    AND objectId NOT LIKE '%/_layouts/15/%'
