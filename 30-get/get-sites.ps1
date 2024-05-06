<#---
title: Get All Sites in Tenant
tag: all-sites-in-tenant
api: post
connection: sharepoint
output: allsites.json
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

$sites = Get-PnPTenantSite # -Filter "Url -like 'nexiintra'"
$allsites = @()
foreach ($site in $sites) {
    $allsites += @{
        Url                     = $site.Url
        Title                   = $site.Title
        Owner                   = $site.Owner
        StorageQuota            = $site.StorageQuota
        StorageUsed             = $site.StorageUsed
        LastContentModifiedDate = $site.LastContentModifiedDate
        Created                 = $site.Created
        SharingCapability       = $site.SharingCapability.value__
        Status                  = $site.Status
        IsTeamsConnected        = $site.IsTeamsConnected
        IsTeamsChannelConnected = $site.IsTeamsChannelConnected
        HusSiteId               = $site.HubSiteId
        GroupId                 = $site.GroupId
        IsHubsite               = $site.IsHubsite
        LocaleId                = $site.LocaleId
        RelatedGroupId          = $site.RelatedGroupId
    }
    <# $currentItemName is the current item #>
}


$outputfile = Join-Path $tenantdir  "allsites.json"

$allsites | ConvertTo-Json -Depth 10 | Out-File -FilePath $outputfile -Encoding utf8NoBOM
