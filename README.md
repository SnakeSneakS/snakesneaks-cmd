# snakesneaks-cmd
- 暇つぶしに作った、ヘビが出てくるだけのコマンド

# 使い方
- バイナリの取得
    - 次のどちらかの方法でバイナリを取得してください
        1. github UI上のreleaseからバイナリ`bin/snakesneaks-(OS)-$(ARCH)`をダウンロードして実行
        2. `go build -o snakesneaks`を実行
- `/usr/local/bin/snakesneaks`などのように、PATHが通っているところへバイナリをおいてください

<!--
PS1="snakesneaks ~ $ "
git tag vx.x.x && git push origin --tags
-->