# PGBOUNCER

## init project

```bash
docker compose up
docker compose restart django
docker exec -it django ./manage.py migrate
```

```bash
# config apisix http://localhost:9080/django

curl -i "http://127.0.0.1:9180/apisix/admin/routes?api_key=edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
{
  "id": "django",
  "uri": "/api/blogs",
  "upstream" : {
    "type": "roundrobin",
    "nodes": {
      "fiber-00:8000": 1,
      "fiber-01:8000": 1
    },
    "pass_host": "node",
    "scheme": "http"
  }
}'

```
