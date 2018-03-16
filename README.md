# ペッパー受付システムを題材に gRPC やってみたンゴ

1.  ペッパーが受付予約 ID みたいなのをリクエスト送ります。
2.  その KEY から情報取得します。（とりま訪問者名、予約者名、部屋 ID、時間帯）
3.  とりまクライアントは普通の WEB ページでリクエストしたら、情報が表示されるだけにしよ

## PHASE1 とりあえず固定値でもいいから値を返せるようにしてみた

### 気づいたこと

結構めんどい。
grpc-gateway 難しそう。JS から気軽に取得できるようにしたいが
JS で client 書くコストと JSON で取得できるようにするコストのどちらが高いだろう？
[twirp](https://github.com/twitchtv/twirp)で書いてみて、これと言って問題なければそっちでいこう

gRPC では proto で渡す引数、返り値は message で定期しないといけない。たとえ空でも

### GCP Endpoint の話

https://cloud.google.com/endpoints/docs/grpc/transcoding?hl=ja

### Go たのしい

実行ファイルは package が main で main 関数がないといけない？ちょっとハマった
[Go 言語: var, init, main が実行される順番](https://qiita.com/suin/items/ab2db295742afcf02334)

使用言語
API サーバ：Go

WEB サーバ：Node.js
