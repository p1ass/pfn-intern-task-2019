# pfn-intern-task-2019

## 概要

PFN夏季インターンシップのコーディング課題公開されているので、バックエンドの課題をやってみました。

[2019年 PFN夏季インターンシップのコーディング課題公開](https://research.preferred.jp/2019/06/internship-coding-task-2019/)

## 環境
- OS: MacOS Mojave 10.14.5
- 言語: Go 1.12

## 事前準備
```bash
git clone https://github.com/naoki-kishi/pfn-intern-task-2019.git

```

## 実行方法

### サーバ

#### 起動方法

```bash
cd pfn-intern-task-2019/job_server 
go run main.go -d "./tests/sample_data" -p 8080
```

- `-d` : Jobデータが保存されているディレクトリ
- `-p` : サーバーのポート

#### 確認

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

#### 起動

```bash
cd pfn-intern-task-2019/worker
go run main.go -p 8080 -c 15 > output/executing_point.csv
```

- `-p` : サーバーのポート
- `-c` : キャパシティ

### グラフ生成
```bash
cd pfn-intern-task-2019/worker/output
python generate_graph.py executing_point.csv
```