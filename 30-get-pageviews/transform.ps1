$env:PGPASSWORD = "SecurePassword!"
psql -h magicbox.postgres.database.azure.com -U pgadmin   -d files -f "$PSScriptRoot/transform.sql"
