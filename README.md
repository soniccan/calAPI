# calAPI
音楽スタジオを予約するシステムを作りたいと考え、予約をデータベースで管理するAPIを作成している。

DBには、sqlite3を用い、以下の構造体で予約データを保存している。

## テーブル名: reservations

このテーブルは、予約情報を保持するために使用されている

| 列名          | データ型      | 制約                      | 説明                             |
|---------------|--------------|---------------------------|----------------------------------|
| ID            | integer      | not null, primary key     | 予約レコードの一意の識別子       |
| USER_ID       | integer      |                           | 予約を行ったユーザーの識別子     |
| RESERVATION_ID| integer      |                           | 予約の識別子                     |
| PAYMENT_ID    | integer      |                           | 支払い情報の識別子               |
| START_AT      | datetime     |                           | 予約開始日時                     |
| END_AT        | datetime     |                           | 予約終了日時                     |

### 列の説明

- **ID**: 予約レコードの一意の識別子。自動インクリメントされます。
- **USER_ID**: 予約を行ったユーザーの識別子。
- **RESERVATION_ID**: 予約の識別子。
- **PAYMENT_ID**: 支払い情報の識別子。
- **START_AT**: 予約の開始日時。
- **END_AT**: 予約の終了日時。

## スキーマ
`UPDATE` `DELETE` `CREATE`を行う際に、クエリパラメータとしていかをbodyに含める必要がある。

クエリパラメータ
```

{
  "type": "object",
  "properties": {
    "userId": {
      "type": "integer"
    },
    "paymentId": {
      "type": "integer"
    },
    "startAt": {
      "type": "string",
      "format": "date-time"
    },
    "endAt": {
      "type": "string",
      "format": "date-time"
    }
  }
}
```
# GET /reservation

このエンドポイントは、データベースから予約情報を全て取り出す。

## リクエスト

- メソッド: `GET`
- エンドポイント: `/reservation`



## レスポンス

**成功時:**

- ステータスコード: `200 OK`
- レスポンスボディ:

```json
{
    "reservationId" 
}

```

# PUT /reservation

このエンドポイントは、特定の予約情報を更新します。

## リクエスト

- メソッド: `PUT`
- エンドポイント: `/reservation`


```
**成功時:**

- ステータスコード: `200 OK`
- レスポンスボディ:

```json
{
    "reservationId" 
}

```


# POST /reservation

このエンドポイントは、予約データをデータベースに送信する。

## リクエスト

- メソッド: `POST`
- エンドポイント: `/reservation`

上記のスキーマを使用して、送信する予約データをリクエストボディに含めてください。



# DELETE /reservation

このエンドポイントは、予約データをデータベースから削除する。

## リクエスト

- メソッド: `DELETE`
- エンドポイント: `/reservation`
- レスポンスボディ:
```json
{
    "reservationId" 
}

```

