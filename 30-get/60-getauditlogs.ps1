<#---

---

Certainly! Here's an example request to retrieve SharePoint audit logs for the last week using Microsoft Graph API, including real date values:

```http
GET https://graph.microsoft.com/v1.0/auditLogs/directoryAudits?$filter=activityDateTime ge 2024-04-29T00:00:00Z and activityDateTime le 2024-05-05T23:59:59Z and category eq 'SharePoint'
```

In this example:

- `activityDateTime ge 2024-04-29T00:00:00Z` specifies that the activity date should be greater than or equal to April 29, 2024, at midnight UTC.
- `activityDateTime le 2024-05-05T23:59:59Z` specifies that the activity date should be less than or equal to May 5, 2024, at 11:59:59 PM UTC.
- `category eq 'SharePoint'` filters the logs to include only SharePoint-related activities.

Make sure to replace these date values with the actual start and end dates you want to query. Additionally, ensure that you have the necessary permissions and authentication set up to access Microsoft Graph API.
#>