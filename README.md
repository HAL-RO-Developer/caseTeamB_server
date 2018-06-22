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

## 子供情報設定 [/child]
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

        + success: 子供情報を追加しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報取得[GET]
子供ID,誕生日,ニックネーム,性別の一覧を取得します。

+ Request(application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + data (array)
            + (object)
                + child_id: 1 (number),
                + birthday: `2016-10-01T09:00:00+09:00`
                + nickname: sample
                + sex: 0 (number) - 0:男、1:女

            + (object)
                + child_id: 2 (number),
                + birthday: `2017-03-19T09:00:00+09:00`
                + nickname: index
                + sex: 1 (number) - 0:男、1:女

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報削除[DELETE]
登録されている子どもIDの情報を削除します。

+ Request(application/json)
    + Headers

            Authorization: token

    + Attributes
        + child_id: 1 (number)

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

    + Attributes
        + child_id(number): 1


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

    + Attribute

        + List (array)
            + (object)
                + device_id: sample
                + device_status: true (boolean)

            + (object)
                + device_id: index
                + device_status: false (boolean)

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

## デバイス紐付け [/registration]

### ボタン登録[POST]
ボタンIDと各ボタンデバイスとの紐付けを行います。


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

# Group BOCCO x 学習 API

## ユーザー削除 [/work/user]
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


## ICリーダー [/work/reader]
### 回答データを送信[POST]
デバイス情報と読み取ったタグの情報を送信。

+ Request(application/json)

    + Attributes
        + device_id: sample
        + data
            + book_id(number): 1
            + q_no(number): 1
            + user_answer(number): 1

+ Response 200 (application/json)

    + Attribute

        + success: 送信しました。

+ Response 418 (application/json)

    + Attribute

        + error: データベースエラー

## 記録を取得 [/work/record/{device_id}]
### 回答データの取得[GET]
指定されたリーダーの記録情報を取得

+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + data (array)
            + (object)
                + date: `2018-06-21T13:35:08+09:00`
                + book_id: 1 (number)
                + q_no: 1 (number)
                + user_answer: 1 (number)
                + correct: 2 (number)

            + (object)
                + date: `2018-06-22T13:35:08+09:00`
                + book_id: 2 (number)
                + q_no: 3 (number)
                + user_answer: 3 (number)
                + correct: 1 (number)

+ Response 400 (application/json)

    + Attribute

        + error: 回答情報が見つかりませんでした。

# Group BOCCO x 目標ボタン API

## ユーザー削除 [/goal/user]
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

## プッシュ回数増加 [/goal/push]
### ボタンプッシュ[POST]
目標ボタンが押された回数を記録します。

+ Request (applicaition/json)

    + Attributes
        + device_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: プッシュ回数を追加しました。


+ Response 400 (application/json)

    + Attribute

        + error: ボタンIDが見つかりません。


## 目標 [/goal/goal/{device_id}]

### 目標登録[POST]
目標の新規追加を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attributes
        + device_id: sample
        + goal: practice

+ Response 200 (application/json)

    + Attribute

        + success: 目標を追加しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


### 目標取得[GET]
登録されている目標と実行回数を取得します。


+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + created_at: `2018-06-21T13:35:08+09:00`
        + run : 5 (number)
        + goal: practice
        + updated_at: `2018-06-21T13:37:21+09:00`,


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     
### 目標削除[DELETE]
登録されている目標を削除します。


+ Parameters
    + device_id: sample

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

### 達成数変更 [POST]
目標実行数を変更します。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attributes
        + device_id: sample
        + approval : 1 (number) - 増減値

+ Response 200 (application/json)

    + Attribute

        + success: 目標達成を承認しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## メッセージ [/goal/message/{device_id}]

### メッセージ登録[POST]
メッセージの新規追加を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attributes

        + device_id: sample
        + condition : 5 (number)
        + message: practice

+ Response 200 (application/json)

    + Attribute

        + success: メッセージを登録しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


### メッセージ取得[GET]
登録されているメッセージと承認済み実行回数を取得します。

+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token
+ Response 200 (application/json)

    + Attribute

       + practice: 5 (number)

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     
### メッセージ削除[DELETE]
登録されているメッセージを削除します。


+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token
+ Response 200 (application/json)

    + Attribute

        + success: 目標を削除しました。


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー