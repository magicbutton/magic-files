<#---
title: Get shared mailboxes
connection: exchange
---
#>



if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"

$sharedmailboxesdir = join-path $workdir "sharedmailboxes"
if (-not (Test-Path $sharedmailboxesdir)) {
    New-Item -Path $sharedmailboxesdir -ItemType Directory | Out-Null
}

Get-EXOMailbox -RecipientTypeDetails SharedMailbox -Resultsize Unlimited 
| select DisplayName, UserPrincipalName , PrimarySmtpAddress, Identity
| ConvertTo-Json 
| Out-File -FilePath (join-path $sharedmailboxesdir  "sharedmailboxes.json") -Encoding utf8NoBOM