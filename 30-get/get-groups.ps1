<#---
title: Get All Sites in Tenant
tag: all-sites-in-tenant
api: post
connection: sharepoint
output: allgroups.json
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

#$groups = Get-PnPAzureADGroup 
$allgroups = @()
foreach ($group in $groups) {
    $allgroups += @{
     
        Id              = $group.Id
        MailNickname    = $group.MailNickname
        DisplayName     = $group.DisplayName
        Owners          = $group.Owners
        Members         = $group.Members
        Description     = $group.Description
        MailEnabled     = $group.MailEnabled
        SecurityEnabled = $group.SecurityEnabled
        Visibility      = $group.Visibility
        GroupTypes      = $group.GroupTypes


        
    }
    <# $currentItemName is the current item #>
}


$outputfile = Join-Path $tenantdir  "allgroups.json"

$allgroups | ConvertTo-Json -Depth 10 | Out-File -FilePath $outputfile -Encoding utf8NoBOM
