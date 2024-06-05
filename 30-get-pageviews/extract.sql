CREATE SCHEMA IF NOT EXISTS import;

CREATE TABLE IF NOT EXISTS import.auditlogs(
    RecordType varchar(50),
    CreationDate timestamp,
    UserIds varchar(255),
    Operations varchar(50),
    AuditData jsonb,
    ResultIndex integer,
    ResultCount integer,
    Identity UUID,
    IsValid boolean,
    ObjectState varchar(50)
);

\COPY import.auditlogs FROM '/Users/nielsgregersjohansen/kitchens/magic-files/.koksmat/workdir/AuditLogRecords.csv' WITH (FORMAT csv, HEADER true);
