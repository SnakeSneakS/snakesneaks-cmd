# snakesneaks-cmd
- 暇つぶしに作った、ヘビが出てくるだけのコマンド

# 使い方
- バイナリの取得
    - 次のどちらかの方法でバイナリを取得してください
        1. github UI上のreleaseからバイナリ(`bin/`配下の適切なバイナリ)をダウンロードして実行
        2. `go build -o snakesneaks`を実行
- `/usr/local/bin/snakesneaks`などのように、PATHが通っているところへバイナリをおいてください

<!--
# CLI interface
PS1="snakesneaks ~ $ "

# tag
git tag vx.x.x
git push origin --tags or git push origin ${tag}
git tag -d TAGNAME
git push --delete origin tag TAGNAME


# refs
- [good github action blog](https://toranoana-lab.hatenablog.com/entry/2022/12/09/000000)
-->