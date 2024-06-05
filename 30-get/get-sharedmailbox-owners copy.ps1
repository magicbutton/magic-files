<#---
title: Get All Group Owners
tag: group-owners
api: post
connection: exchange
input: sharedmailboxes.json

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

$tenantdir = join-path $workdir "sharedmailboxes"
if (-not (Test-Path $tenantdir)) {
    New-Item -Path $tenantdir -ItemType Directory | Out-Null
}
$ownerdir = join-path $tenantdir "owners"
if (-not (Test-Path $ownerdir)) {
    New-Item -Path $ownerdir -ItemType Directory | Out-Null
}

$sharedmailboxes = Get-Content (join-path $workdir "sharedmailboxes" "sharedmailboxes.json") | ConvertFrom-Json

#$groups = Get-PnPAzureADGroup 

foreach ($sharedmailbox in $sharedmailboxes) {
    $ownerFile = Join-Path $ownerdir   "$($sharedmailbox.PrimarySmtpAddress)_owners.json"

    if (Test-Path $ownerFile) {
        write-host "Existing - Skipping $($sharedmailbox.PrimarySmtpAddress)" -ForegroundColor Gray

        continue
    }
 
    write-host "Getting owners for $($sharedmailbox.Identity)" -ForegroundColor Green
    try {
        
        $owners = Get-EXOMailboxPermission -Identity $sharedmailbox.Identity # -User "Owners" -ResultSize Unlimited | select DisplayName, UserPrincipalName, Mail
        $stats = Get-EXOMailboxFolderStatistics -Identity $sharedmailbox.Identity # | select DisplayName, UserPrincipalName, Mail
        [array]$result = @()
        foreach ($owner in $owners) {
            $result += $owner
        }
        
        $result | ConvertTo-Json -Depth 10 | Out-File -FilePath $ownerFile -Encoding utf8NoBOM
    
    }
    catch {
        write-host "Error getting owners for $($group.MailNickname)" -ForegroundColor Red
    }
    finally {
        <#Do this after the try block regardless of whether an exception occurred or not#>
    }

   
}

