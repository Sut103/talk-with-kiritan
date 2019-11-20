# talk-with-kiritan
Discordのボイスチャットに東北きりたんが参加してる風に演出するBot

(実はきりたんじゃなくても動くし、キャラクター性も必要ない)
## 必要なもの
* golang
    * gomodule
* ffmpeg
* DiscordBotのトークン
* soundsフォルダ
    * 再生したい音声ファイル.wav

## 使い方
1. Discordで事前にBotを登録し、config.jsonにトークンを書いておく
2. ./soundsに音声ファイルを入れておく
3. ```$ go run main.go```で起動
4. DiscordでVCに入る
5. チャットにきりたん砲全門斉射！と書く
    * Botが入ってくる
6. http://localhost:8080/recognition で音声認識
    * ファイル名(記号や拡張子は除く)をしゃべってみる
    * 適切な単語が入るとBotが対応する音声ファイルを再生する
7. チャットにおつかれさまですと書く
    * Botが出ていく
