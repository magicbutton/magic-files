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

Get-ChildItem -Path $workdir -Filter "hubsites.json" -Recurse | ForEach-Object {
    magic-files import $_.FullName
}