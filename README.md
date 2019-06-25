# pfn-intern-task-2019

## 概要

PFN夏季インターンシップのコーディング課題公開されているので、バックエンドの課題をやってみました。

[2019年 PFN夏季インターンシップのコーディング課題公開](https://research.preferred.jp/2019/06/internship-coding-task-2019/)

## 事前準備
```bash
git clone https://github.com/naoki-kishi/pfn-intern-task-2019.git

```

## 実行方法

### サーバ
```bash
cd pfn-intern-task-2019/job_server
go run main.go
```

```bash
$ curl "localhost:8080?time=00:00:00" -i
HTTP/1.1 200 OK
Date: Tue, 25 Jun 2019 13:19:35 GMT
Content-Length: 63
Content-Type: text/plain; charset=utf-8

[JobID]
0

[Created]
00:00:00

[Priority]
Low

[Tasks]
7
3
6
6

```

### ワーカー
```bash
cd pfn-intern-task-2019/worker
go run main.go > executing_point.csv
```

### グラフ生成
```bash
cd pfn-intern-task-2019/worker
python generate_graph.py
```