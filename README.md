go-rainbow
=========================

CSVをカラムごとに色分けして表示するCLIコマンド.

ビルド
-------------------------

```
$ go get -u github.com/val00274/go-rainbow
$ cd /path/to/go-rainbow
$ make
$ cp ./rainbow /path/to/bindir
```

使い方
-------------------------

```
$ cat data.csv | rainbow
$ cat data.tsv | rainbow --tsv
$ cat data.txt | rainbow --comma ';'
```

`--tsv` オプションでTSV（タブ区切り）ファイルを扱う.

`--comma CHAR` オプションで区切り文字（１文字のみ）を指定可能.

参考
-------------------------

アイデアはここから https://github.com/mechatroner/rainbow_csv
