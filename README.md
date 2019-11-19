# talk-with-kiritan
Discordのボイスチャットに東北きりたんが参加してる風に演出するBot

(実はきりたんじゃなくても動くし、キャラクター性も必要ない)
## 必要なもの
* golang
* gomodule
* DiscordBotのトークン(config/config.jsonに書く)
* soundsフォルダ(再生したい音声ファイルを入れておく(現状.wavのみ))

## 使い方(いまのところ)
### 音声認識
1. ```$ go run main.go```で起動
2. http://localhost:8080/recognition で音声認識
3. コンソールに認識結果が出る
4. おわり

### Discord
1. Discordで事前にBotを登録し、config.jsonにトークンを書いておく
2. チャットにきりたん砲全門斉射！と書く
    * きりたんが入ってくる
3. チャットにおつかれさまですと書く
    * きりたんが出ていく

### 備考
* 音声認識とDiscordはまだ繋がっていない
