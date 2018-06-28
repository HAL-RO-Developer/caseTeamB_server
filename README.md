FORMAT: 1A
# ケーススタディ HAL大阪　API一覧

# Group A/B共通
## ユーザー登録 [/signup]
ユーザー情報の登録、およびサインインするためのAPI

### サインアップ [POST]
ユーザー情報の登録を行います。

+ Request (applicaition/json)

    + Attribute
        + name: sample
        + pass: password

+ Response 200 (application/json)

    + Attribute

        + success: ユーザー登録を行いました。
+ Response 400 (application/json)

    + Attribute

        + error: 登録済みのユーザー名です。
 
## サインイン [/signin]

### サインイン [POST]
登録されているユーザー情報を元にサインインを行います。

+ Request (applicaition/json)

    + Attribute
        + name: sample
        + pass: password

+ Response 200 (application/json)

    + Attribute

        + token: sample

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## ユーザー削除 [/user]
### ユーザー削除[DELETE]
登録されているユーザー情報を削除します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: ユーザー情報を削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


## 子供情報設定 [/child/{child_id}]
### 子供情報追加[POST]
子供の誕生日と性別、ニックネームを設定します。

+ Request(application/json)
    + Headers

            Authorization: token

    + Attributes
        + nickname: sample
        + birthday : `2000-01-01`
        + sex : 0 (number) - 0:男、1:女

+ Response 200 (application/json)

    + Attribute

        + child_id: 1 (number)

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報取得[GET]
子供ID,誕生日,ニックネーム,性別の一覧を取得します。

+ Request(application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes

        + children (array)
            + (object)
                + child_id: 1 (number)
                + birthday: `2016-10-01T09:00:00+09:00`
                + nickname: sample
                + sex: 0 (number) - 0:男、1:女

            + (object)
                + child_id: 2 (number)
                + birthday: `2017-03-19T09:00:00+09:00`
                + nickname: index
                + sex: 1 (number) - 0:男、1:女

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報削除[DELETE]
登録されている子どもIDの情報を削除します。

+ Parameters
    + child_id: 1

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: 削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## デバイス [/device/{device_id}]

### デバイスID発行[POST]
新規登録するデバイスIDの発行を行います。

+ Request(application/json)
    + Headers

            Authorization: token

    + Attribute
        + child_id: 1 (number)

+ Response 200 (application/json)

    + Attribute

        + pin: 0000

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー
     

### デバイス一覧取得[GET]
現在登録されているデバイスIDの一覧を取得します。

+ Request(application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes

        + devices (array)
            + (object)
                + device_id: sample
                + device_alive: true (boolean)

            + (object)
                + device_id: index
                + device_alive: false (boolean)

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### デバイスID削除[DELETE]
登録されているデバイスIDを削除します。

+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: ボタンIDを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## BOCCOAPI [/bocco]

### BOCCOAPI設定[POST]
BOCCOAPIに登録したメールアドレスと、パスワードの入力

+ Request
    + Headers

            Authorization: token

    + Attribute
        + email: sample@gmail.com
        + key : sample - APIkey
        + password: abc123

+ Response 200 (application/json)

    + Attribute

        + success: メールアドレスとパスワードを登録しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### BOCCOAPI設定の取得[GET]
BOCCOAPIに登録したメールアドレスの取得

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + email: sample@gmail.com

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### BOCCOAPI削除[DELETE]
BOCCOAPIに登録したメールアドレスと、パスワードの削除

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: メールアドレスとパスワードを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


# Group BOCCO x 学習 API

## 記録を取得 [/work/record/{device_id}]
### 回答データの取得[GET]
指定されたリーダーの記録情報を取得

+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes
        + record (array)
            + (object)
                + date: `2018-06-21T13:35:08+09:00` - 回答日時
                + genre_name: 算数 - 問題のジャンル名
                + sentence: 1 + 1は? - 問題文
                + user_answer: 2 - ユーザーの回答
                + correct: 2 - 問題の正解
                + result: true (boolean) - 正解:true,不正解:false

            + (object)
                + date: `2018-06-22T13:35:08+09:00`
                + genre_name: 社会
                + sentence: 兵庫県の県庁所在地は？
                + user_answer: 兵庫市
                + correct: 神戸市
                + result: false (boolean)

+ Response 400 (application/json)

    + Attribute

        + error: 回答情報が見つかりませんでした。

# Group BOCCO x 目標ボタン API

## 目標 [/goal/goal/{child_id}/{goal_id}]

### 目標登録[POST]
目標の新規追加を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attribute
        + child_id: 1 (number)
        + content: practice
        + criteria: 20 (number) - 達成目標数
        + deadline : `2018-07-01` - 達成期日(なければ空)

+ Response 200 (application/json)

    + Attribute

        + goal_id: sample

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### ボタン登録[PUT]
使用するdevice_idを登録します

+ Request (application/json)
    + Headers

            Authorization: token

    + Attribute
        + goal_id: sample
        + device_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: 登録しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 目標取得[GET]
登録されている目標と実行回数を取得します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes
        + goals(array)
            + (object)
                + created_at: `2018-06-21T13:35:08+09:00`
                + child_id: 1 (number)
                + goal_id : test - 目標ID
                + device_id : sample
                + run : 5 (number) - 目標実行数
                + content: practice - 目標名称
                + criteria: 20 (number) - 達成目標数
                + deadline : `2018-07-11T13:35:08+09:00` - 達成期日(なければ空)
                + status : 0 (number) - 達成状況(0:未実行、1:実行中、2:達成済み、3:達成失敗)
                + updated_at: `2018-06-21T13:37:21+09:00`,
            + (object)
                + created_at: `2018-06-21T13:35:08+09:00`
                + child_id: 2 (number)
                + goal_id : sample
                + device_id : index
                + run : 5 (number)
                + content: index
                + criteria: 30 (number)
                + deadline : なし
                + status : 0 (number) - 達成状況(0:未実行、1:実行中、2:達成済み、3:達成失敗)
                + updated_at: `2018-06-21T13:37:21+09:00`,


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     
### 目標削除[DELETE]
登録されている目標を削除します。

+ Parameters
    + goal_id : sample
+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: 目標を削除しました。


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## 目標達成承認/非承認 [/goal/approval]

### 達成数変更 [PUT]
目標実行数を変更します。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attribute
        + goal_id: sample
        + approval : 1 (number) - 増減値

+ Response 200 (application/json)

    + Attribute

        + success: 目標達成を承認しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## メッセージ [/goal/message]

### メッセージ編集[PUT]
メッセージの編集を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attribute

        + goal_id: sample
        + condition : 5 (number) - メッセージ出力条件
        + message: practice

+ Response 200 (application/json)

    + Attribute

        + success: メッセージを編集しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


### メッセージ取得[GET]
登録されているメッセージとメッセージ出力条件を取得します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes

        + messages(array)
            + (object)
                + goal_id: sample
                + content: practice
                + condtion: 5 (number)
                + message: practice

            + (object)
                + goal_id: index
                + content: test
                + condtion: 5 (number)
                + message: sample

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

# Group ICリーダー&ボタン側API

## デバイス [/thing/registration]

### デバイス登録[POST]
デバイスIDと各デバイスとの紐付けを行います。


+ Request (applicaition/json)
 
    + Attribute
        + pin: 0000
        + mac: abc123

+ Response 200 (application/json)

 + Attribute

      + device_id: sample

+ Response 400 (application/json)

    + Attribute

        + error: pinが見つかりません。

## ICリーダー [/thing/reader]
### 回答データを送信[POST]
デバイス情報と読み取ったタグの情報を送信。

+ Request(application/json)

    + Attribute
        + device_id: sample
        + uuid: 1234 (number)
        
+ Response 200 (application/json)

    + Attribute

        + success: 送信しました。

+ Response 418 (application/json)

    + Attribute

        + error: データベースエラー

## プッシュ回数増加 [/thing/button]

### ボタンプッシュ[PUT]
目標ボタンが押された回数を記録します。

+ Request (applicaition/json)

    + Attribute
        + device_id: sample

+ Response 200 (application/json)

    + Attribute

        + achievement: 60 (number) - 目標の達成率

+ Response 400 (application/json)

    + Attribute

        + error: 目標IDが見つかりません。
