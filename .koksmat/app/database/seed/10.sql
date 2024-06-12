INSERT INTO
    public.importdata (
        id,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at,
        tenant,
        searchindex,
        name,
        description,
        data
    )
VALUES (
        DEFAULT,
        DEFAULT,
        '',
        DEFAULT,
        '',
        DEFAULT,
        '',
        'seed0',
        'seed0',
        '',
        '{
  "platforms": {
    "sharepoint": {
      "name": "sharepoint"
    },
    "exchange": null
  },
  "containers": [
    {
      "name": "xxxx",
      "url": "https://xx",
      "platform": {
        "name": "sharepoint"
      }
    },
    {
      "name": "shared mail",
      "url": "outlook",
      "platform": null
    }
  ]
}'
    );