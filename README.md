lgtm
====

コマンドラインでlgtmのurlを得るためのツールです。

lgtm.in か、Github上に上がっている任意のリポジトリ(Privateも可)を指定することが可能です。

## lgtm.in からランダム画像を取得する場合

```
$ lgtm in
> https://lgtm.in/p/8QCAXr0Tp
```

## 任意のリポジトリから画像を取得する場合

環境変数で指定されたGithub上の任意のリポジトリ上のフォルダの存在する画像一覧から、ランダムな画像URLを取得します。

```
$ lgtm rand
> https://YOUR_GITHUB_RESPOSITORY/XXX.png
```

なお、以下の環境変数を設定してください。

| 変数 | 意味 |
|-------------------|----------------------|
| LGTM_GITHUB_TOKEN | リポジトリへのアクセスTOKEN |
| LGTM_GITHUB_OWNER | リポジトリ所有者(このリポジトリであれば `hikaruworld` ) |
| LGTM_GITHUB_REPO | 画像が格納されているリポジトリ名 |
| LGTM_GITHUB_ROOT_PATH | フォルダを切っている場合はそのパス |

