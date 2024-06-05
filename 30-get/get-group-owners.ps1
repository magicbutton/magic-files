<#---
title: Get All Group Owners
tag: group-owners
api: post
connection: sharepoint
input: allgroups.json

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

$tenantdir = join-path $workdir "tenant"
if (-not (Test-Path $tenantdir)) {
    New-Item -Path $tenantdir -ItemType Directory | Out-Null
}
$ownerdir = join-path $tenantdir "owners"
if (-not (Test-Path $ownerdir)) {
    New-Item -Path $ownerdir -ItemType Directory | Out-Null
}

$groups = Get-Content (join-path $tenantdir "allgroups.json") | ConvertFrom-Json

#$groups = Get-PnPAzureADGroup 

foreach ($group in $groups) {
    $ownerFile = Join-Path $ownerdir   "$($group.MailNickname)_owners.json"
    # if ($group.GroupTypes.Length -eq 0) {
    #     write-host "Not a group - Skipping $($group.MailNickname)" -ForegroundColor Gray
    #     continue
    # }
    if ("Unified" -notin $group.GroupTypes) {
        write-host "Not a unified group - Skipping $($group.MailNickname)" -ForegroundColor Gray
        continue
    }

    if ("DynamicMembership" -in $group.GroupTypes) {
        write-host "Dynamic group - Skipping $($group.MailNickname)" -ForegroundColor Gray
        continue
    }
    if (Test-Path $ownerFile) {
        write-host "Existing - Skipping $($group.MailNickname)" -ForegroundColor Gray

        continue
    }
 
    write-host "Getting owners for $($group.MailNickname)" -ForegroundColor Green
    try {
        $owners = Get-PnPAzureADGroupOwner -Identity $group.Id 
        [array]$result = @()
        foreach ($owner in $owners) {
            $result += @{
                Id                = $group.Id 
                DisplayName       = $owner.DisplayName
                UserPrincipalName = $owner.UserPrincipalName
                Mail              = $owner.Mail
               
            }
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

