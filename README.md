# ペッパー受付システムを題材にgRPCやってみたンゴ

1. ペッパーが受付予約IDみたいなのをリクエスト送ります。 
2. そのKEYから情報取得します。（とりま訪問者名、予約者名、部屋ID、時間帯）
3. とりまクライアントは普通のWEBページでリクエストしたら、情報が表示されるだけにしよ

## PHASE1 とりあえず固定値でもいいから値を返せるようにしてみた
### 気づいたこと
結構めんどい。
grpc-gateway難しそう。JSから気軽に取得できるようにしたいが
JSでclient書くコストとJSONで取得できるようにするコストのどちらが高いだろう？
[twirp](https://github.com/twitchtv/twirp)で書いてみて、これと言って問題なければそっちでいこう

### Goたのしい
実行ファイルはpackageがmainでmain関数がないといけない？ちょっとハマった
[Go言語: var, init, mainが実行される順番](https://qiita.com/suin/items/ab2db295742afcf02334)

使用言語
APIサーバ：Go

WEBサーバ：Node.js