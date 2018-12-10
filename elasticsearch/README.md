#elasticsearch

elasticsearch on docker.

## API

## GET /_cat/health?v

Elasticsearch のヘルスチェックを行う

```bash
curl "localhost:9200/_cat/health?v"
```

## GET /_cat/nodes?v

Node の状態確認

```bash
curl "localhost:9200/_cat/nodes?v"
```

## GET /_cat/indices?v

index のリストアップ

```bash
curl "localhost:9200/_cat/indices?v"
```

## PUT /customer?pretty

index の作成

```bash
curl -X PUT "localhost:9200/customer?pretty"
```

## PUT /customer/_doc/1?pretty

index にデータ突っ込む。index作成前に実行も可能。リプレイスも可能。

```bash
curl -X PUT -H 'Content-Type:application/json' "localhost:9200/customer/_doc/1?pretty" -d '{"name": "John Doe"}'
```

## GET /customer/_doc/1?pretty

id 指定でデータ所得

```bash
curl "localhost:9200/customer/_doc/1?pretty"
```

## DELETE /customer?pretty

index 削除

```bash
 curl -X DELETE "localhost:9200/customer?pretty"
```

## POST /customer/_doc?pretty

id ランダム生成でデータ突っ込む

```bash
curl -X POST -H 'Content-Type:application/json' "localhost:9200/customer/_doc/?pretty" -d '{"name": "ddddd"}'
```

## POST /customer/_doc/1/_update?pretty

document のアップデート

```bash
curl -X POST -H 'Content-Type:application/json' "localhost:9200/customer/_doc/1/_update?pretty" -d '{"doc": {"name": "Jane Doe"}}'
```

```bash
curl -X POST -H 'Content-Type:application/json' "localhost:9200/customer/_doc/1/_update?pretty" -d '{"doc": {"name": "Jane Doe", "age": 20}}'
```

現在のソースドキュメントを参照してスクリプトを実行する

```bash
curl -X POST -H 'Content-Type:application/json' "localhost:9200/customer/_doc/1/_update?pretty" -d '{"script" : "ctx._source.age += 5"}'
```

## DELETE /customer/_doc/2?pretty

document の削除

```bash
curl -X DELETE "localhost:9200/customer/_doc/2?pretty"
```

## POST /customer/_doc/_bulk?pretty

bulk を使用して複数ドキュメントをindex
データの最後の行は、改行文字\ nで終わる必要がある

```bash
curl -X POST "localhost:9200/customer/_doc/_bulk?pretty" -H 'Content-Type: application/json' -d'
{"index":{"_id":"1"}}
{"name": "John Doe" }
{"index":{"_id":"2"}}
{"name": "Jane Doe" }
'

```

```bash
curl -X POST "localhost:9200/customer/_doc/_bulk?pretty" -H 'Content-Type: application/json' -d'
{"update":{"_id":"1"}}
{"doc": { "name": "John Doe becomes Jane Doe" } }
{"delete":{"_id":"2"}}
'

```

## GET /bank/_search?q=*&sort=account_number:asc&pretty

URLによる検索API(default10件)

```bash
curl "localhost:9200/bank/_search?q=*&sort=account_number:asc&pretty"
```

jsonによる検索API(default10件)

```bash
curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}
'

```

検索数とどこから検索かを指定

```bash
curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "from": 10,
  "size": 10,
  "sort": [
    { "account_number": "asc" }
  ]
}
'

```

返却する JSON のフィールドを選択

```bash
curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "_source": ["account_number", "balance"]
}
'

```

queryによる検索は下記参照
> https://www.elastic.co/guide/en/elasticsearch/reference/current/_executing_searches.html

filterする際は下記
gte は greater than or equal の略？


```bash
curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "balance": {
            "gte": 20000,
            "lte": 30000
          }
        }
      }
    }
  }
}
'

```
