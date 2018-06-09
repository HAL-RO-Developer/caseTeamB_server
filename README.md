FORMAT: 1A

# BOCCO x 目標ボタンAPI [/api]

## ユーザー登録 [/api/signup]
ユーザー情報の登録、およびサインインするためのAPI


### サインアップ [POST]
ユーザー情報の登録を行います。

+ Request (applicaition/json)

    + Headers
 
        Accept: application/json
 
    + Attribute
        + name: sample
        + pass: password


+ Response 200 (application/json)

    + Attribute

        + success: ユーザー登録を行いました。
+ Response 400 (application/json)

    + Attribute

        + error: 登録済みのユーザー名です。

 
## サインイン [/api/signin]
  
### サインイン [POST]
登録されているユーザー情報を元にサインインを行います。

+ Request (applicaition/json)

    + Headers
 
         Accept: application/json
 
    + Attribute
        + name: sample
        + pass: password

+ Response 200 (application/json)

    + Attribute

        + token: sample

+ Response 400 (application/json)

    + Attribute

        + error: ユーザー名が不正です。

## ユーザー削除 [/api/user]
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

        + error: アクセストークンが不正です。
     
## ボタン [/api/button]

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

        + error: 発行済のボタンIDです。
     

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

        + error: アクセストークンが不正です。

     
     
### ボタンID削除[DELETE]
登録されているボタンIDを削除します。

+ Request (application/json)
    + Headers

            Authorization: token
    + Attribute

        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: ボタンIDを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。

     

## デバイス [/api/device]

### ボタン登録[POST]
ボタンIDと各ボタンデバイスとの紐付けを行います。


+ Request (applicaition/json)

    + Headers
 
        Accept: application/json
 
    + Attribute
        + button_id: sample
        + mac: abc123


+ Response 200 (application/json)

 + Attribute

      + success: ボタンIDを登録しました。

+ Response 400 (application/json)

    + Attribute

        + error: このボタンIDは登録済みです。

### ボタンプッシュ[PUT]
目標ボタンが押された回数を記録します。

+ Request (applicaition/json)

    + Headers
 
        Accept: application/json
 
    + Attributes
        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: プッシュ回数を追加しました。


+ Response 400 (application/json)

    + Attribute

        + error: ボタンIDが見つかりません。


     

## 目標 [/api/goal]

### 目標登録[POST]
目標の新規追加を行います。

+ Request (application/json)

    + Headers
 
        Accept: application/json
 
    + Attributes
        + goal: practice

+ Response 200 (application/json)

    + Attribute

        + success: 目標を追加しました。

+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。


### 目標取得[GET]
登録されている目標を取得します。


+ Request(application/json)

    + Headers
 
        Accept: application/json
 
    + Attributes
        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + goal: practice
        + archive : 5


+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。

     
### 目標削除[DELETE]
登録されている目標を削除します。


+ Request

    + Headers
 
        Accept: application/json
 
    + Attributes
        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: 目標を削除しました。


+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。


## メッセージ [/api/message]

### メッセージ登録[POST]
メッセージの新規追加を行います。

+ Request (application/json)

    + Headers
 
        Accept: application/json
 
    + Attributes
        + message: practice
        + button_id: sample

+ Response 200 (application/json)

    + Attribute

        + success: メッセージを追加しました。

+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。


### メッセージ取得[GET]
登録されているメッセージを取得します。


+ Request(application/json)

    + Headers
 
        Accept: application/json
 
    + Attributes
        + name: sample

+ Response 200 (application/json)

    + Attribute

       + practice: 5
       + test : 10


+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。

     
### メッセージ削除[DELETE]
登録されているメッセージを削除します。


+ Request

    + Headers
 
        Accept: application/json
 
    + Attributes
        + message: sample

+ Response 200 (application/json)

    + Attribute

        + success: 目標を削除しました。


+ Response 400 (application/json)

    + Attribute

        + error: アクセストークンが不正です。