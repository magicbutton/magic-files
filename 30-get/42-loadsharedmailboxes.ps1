<#---
title: Load to importdata table
connection: postgres
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
$importdir = join-path $workdir "sharedmailboxes" 
Get-ChildItem -Path $importdir -Filter "*.json" -Recurse | ForEach-Object {
    write-host  $_.FullName
    magic-files import $_.FullName
}