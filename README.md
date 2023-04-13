kmeango
===
k-means法で、スペースで区切られた数値をクラスタリングするコマンドです。

## install
```
$ go install github.com/amanoese/kmeango@latest
```

## Usage
```bash
# 2次元データのクラスタリング(デフォルトでは3個のクラスタに分類)
$ seq 10|xargs -n2 | kmeango
Center : Value
     1 2 :  1 2
     4 5 :  3 4
     4 5 :  5 6
     8 9 :  7 8
     8 9 :  9 10

# 2次元データのクラスタリング(デフォルトでは2個のクラスタに分類)
$ seq 10|xargs -n2 | kmeango -k2
Center : Value
     3 4 :  1 2
     3 4 :  3 4
     3 4 :  5 6
     8 9 :  7 8
     8 9 :  9 10

# n次元データもクラスタリングできます
$ seq 24|xargs -n3 | kmeango -k2
Center : Value
   4 5 6 :  1 2 3
   4 5 6 :  4 5 6
   4 5 6 :  7 8 9
 16 17 18 :  10 11 12
 16 17 18 :  13 14 15
 16 17 18 :  16 17 18
 16 17 18 :  19 20 21
 16 17 18 :  22 23 24

$ seq 24|xargs -n5 | kmeango -k2
Center : Value
 3 4 5 6 7 :  1 2 3 4 5
 3 4 5 6 7 :  6 7 8 9 10
 16 17 18 19 11 :  11 12 13 14 15
 16 17 18 19 11 :  16 17 18 19 20
 16 17 18 19 11 :  21 22 23 24 0
```
