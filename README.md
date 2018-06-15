FORMAT: 1A

# BOCCO x 目標ボタンAPI [/api]

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


## ボタン [/goal/button/{button_id}]

### ボタンID発行[POST]
新規登録するボタンIDの発行を行います。

+ Request
    + Headers

            Authorization: token


+ Response 200 (application/json)

    + Attribute

        + pin: 0000

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー
     

### ボタン一覧取得[GET]
現在登録されているボタンIDの一覧を取得します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + button_id (array)
            + sample,
            + test

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     
     
### ボタンID削除[DELETE]
登録されているボタンIDを削除します。

+ Parameters
    + button_id: sample

+ Request
    + Headers

            Authorization: token


+ Response 200 (application/json)

    + Attribute

        + success: ボタンIDを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     

## デバイス [/goal/device]

### ボタン登録[POST]
ボタンIDと各ボタンデバイスとの紐付けを行います。


+ Request (applicaition/json)
 
    + Attribute
        + pin: 0000
        + mac: abc123


+ Response 200 (application/json)

 + Attribute

      + button_id: sample

+ Response 400 (application/json)

    + Attribute

        + error: pinが見つかりません。

### ボタンプッシュ[PUT]
目標ボタンが押された回数を記録します。

+ Request (applicaition/json)

    + Attributes
        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: プッシュ回数を追加しました。


+ Response 400 (application/json)

    + Attribute

        + error: ボタンIDが見つかりません。


     

## 目標 [/goal/goal/{button_id}]

### 目標登録[POST]
目標の新規追加を行います。

+ Request (application/json)

    + Attributes
        + button_id: sample
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
    + button_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + run : 5 (number)
        + goal: practice


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

     
### 目標削除[DELETE]
登録されている目標を削除します。


+ Parameters
    + button_id: sample

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

### 達成承認 [POST]
目標実行数を変更します。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attributes
        + button_id: sample
        + approval : 1 (number) - 増減値

+ Response 200 (application/json)

    + Attribute

        + success: 目標達成を承認しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## メッセージ [/goal/message/{button_id}]

### メッセージ登録[POST]
メッセージの新規追加を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attributes

        + button_id: sample
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
    + button_id: sample

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
    + button_id: sample

+ Request
    + Headers

            Authorization: token
+ Response 200 (application/json)

    + Attribute

        + success: 目標を削除しました。


+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー