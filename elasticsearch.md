## 创建es中article_index
```go
{
    "properties": {
        "article_id": {
            "type": "keyword",
            "index": true
        },
        "article_title": {
            "type": "text",
            "index": true
        },
        "article_des": {
            "type": "text",
            "index": true
        },
        "created_at": {
            "type": "date",
            "index": true
        },
        "updated_at": {
            "type": "date",
            "index": true
        },
        "deleted_at": {
            "type": "date",
            "index": true
        }
    }
}
```