# talk-with-kiritan
Discordのボイスチャットに東北きりたんが参加してる風に演出するBot

(実はきりたんじゃなくても動くし、キャラクター性も必要ない)
## 必要なもの
* golang
* gomodule
* DiscordBotのトークン(config/config.jsonに書く)
* soundsフォルダ(再生したいセリフを入れておく(現状.wavのみ))
* ffmpeg

## 使い方
1. Discordで事前にBotを登録し、config.jsonにトークンを書いておく
2. ```$ go run main.go```で起動
3. DiscordでVCに入る
3. チャットにきりたん砲全門斉射！と書く
    * きりたんが入ってくる
4. http://localhost:8080/recognition で音声認識
    * ファイル名(記号や拡張子は除く)をしゃべってみる。
    * 適切な単語が入るときりたんが ~~対応する音声ファイルを再生する~~ しゃべる。
5. チャットにおつかれさまですと書く
    * きりたんが出ていく
