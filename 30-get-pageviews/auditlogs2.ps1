<#---
title: Get audit logs 2
connection: exchange
---


https://learn.microsoft.com/en-us/purview/audit-log-search-script

#>
param (
    #$record = "SharePoint"
    $record = "SharePointFileOperation"
)
if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir
# Set the start and end dates for the audit log query



[DateTime]$start = [DateTime]::UtcNow.AddDays(-1)
[DateTime]$end = [DateTime]::UtcNow
# https://learn.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-schema#auditlogrecordtype

$resultSize = 5000
$intervalMinutes = 15

#Modify the values for the following variables to configure the audit log search.
$logFile = join-path $workdir  "AuditLogSearchLog.txt"
$outputFile = join-path $workdir  "AuditLogRecords-$record.csv"

#Start script
[DateTime]$currentStart = "2024-05-23T10:00:00.000000Z"
[DateTime]$currentEnd = "2024-05-23T11:00:00.000000Z"

Function Write-LogFile ([String]$Message) {
    $final = [DateTime]::Now.ToUniversalTime().ToString("s") + ":" + $Message
    $final | Out-File $logFile -Append
    koksmat trace log $Message
}

Write-LogFile "BEGIN: Retrieving audit records between $($start) and $($end), RecordType=$record, PageSize=$resultSize."
Write-Host "Retrieving audit records for the date range between $($start) and $($end), RecordType=$record, ResultsSize=$resultSize"

$totalCount = 0
while ($true) {
    $currentEnd = $currentStart.AddMinutes($intervalMinutes)
    if ($currentEnd -gt $end) {
        $currentEnd = $end
    }

    if ($currentStart -eq $currentEnd) {
        break
    }

    $sessionID = [Guid]::NewGuid().ToString() + "_" + "ExtractLogs" + (Get-Date).ToString("yyyyMMddHHmmssfff")
    Write-LogFile "INFO: Retrieving audit records for activities performed between $($currentStart) and $($currentEnd)"
    Write-Host "Retrieving audit records for activities performed between $($currentStart) and $($currentEnd)"
    $currentCount = 0

    $sw = [Diagnostics.StopWatch]::StartNew()
    do {
        $results = Search-UnifiedAuditLog -StartDate $currentStart -EndDate $currentEnd -RecordType $record -SessionId $sessionID -SessionCommand ReturnLargeSet -ResultSize $resultSize

        if (($results | Measure-Object).Count -ne 0) {
            $results | export-csv -Path $outputFile -Append -NoTypeInformation

            $currentTotal = $results[0].ResultCount
            $totalCount += $results.Count
            $currentCount += $results.Count
            Write-LogFile "INFO: Retrieved $($currentCount) audit records out of the total $($currentTotal)"

            if ($currentTotal -eq $results[$results.Count - 1].ResultIndex) {
                $message = "INFO: Successfully retrieved $($currentTotal) audit records for the current time range. Moving on!"
                Write-LogFile $message
                Write-Host "Successfully retrieved $($currentTotal) audit records for the current time range. Moving on to the next interval." -foregroundColor Yellow
                ""
                break
            }
        }
    }
    while (($results | Measure-Object).Count -ne 0)

    $currentStart = $currentEnd
}

Write-LogFile "END: Retrieving audit records between $($start) and $($end), RecordType=$record, PageSize=$resultSize, total count: $totalCount."
Write-Host "Script complete! Finished retrieving audit records for the date range between $($start) and $($end). Total count: $totalCount" -foregroundColor Green