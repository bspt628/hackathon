6期

# ToDo

## Userテーブル操作のAPIエンドポイントを作成

ある程度、単一エンドポイントで複数のフィールドを変更する必要がありそうなので、まとめたいと思う

### プロフィール関連
- profile_image_url
- bio
- location
- website
この4つをまとめて更新するAPIエンドポイント

### ユーザー設定関連
- display_name
- birth_date
- language
この3つをまとめて更新するAPIエンドポイント

### 通知関連
- notification_settings
これを更新するAPIエンドポイント

### プライバシー関連
- is_private（設定画面）
これを更新するAPIエンドポイント

### そのほか
- is_banned（特定の操作時に呼び出す）
- username（重複不可であるため）
- email（重複不可であるため、認証する？）
これらについては、1つずつ独立のAPIエンドポイントを作成


- /api/users/{id}/profile
- /api/users/{id}/settings
- /api/users/{id}/notification-settings
- /api/users/{id}/privacy
- /api/users/{id}/ban-status
- /api/users/{id}/username
- /api/users/{id}/email

以下は、各カテゴリに基づいて設計したAPIエンドポイントです。

プロフィール関連

エンドポイント URI:

PATCH /api/users/{id}/profile

対象項目:
	•	profile_image_url
	•	bio
	•	location
	•	website

リクエスト例:

{
  "profile_image_url": "https://example.com/image.jpg",
  "bio": "Software Engineer based in Tokyo.",
  "location": "Tokyo, Japan",
  "website": "https://uchidahiroto.dev"
}

ユーザー設定関連

エンドポイント URI:

PATCH /api/users/{id}/settings

対象項目:
	•	display_name
	•	birth_date
	•	language

リクエスト例:

{
  "display_name": "Hiroto Uchida",
  "birth_date": "1998-05-15",
  "language": "en"
}

通知関連

エンドポイント URI:

PATCH /api/users/{id}/notification-settings

対象項目:
	•	notification_settings (JSON)

リクエスト例:

{
  "notification_settings": {
    "email_notifications": true,
    "push_notifications": false,
    "newsletter_subscribed": true
  }
}

プライバシー関連

エンドポイント URI:

PATCH /api/users/{id}/privacy

対象項目:
	•	is_private

リクエスト例:

{
  "is_private": true
}

その他の独立エンドポイント

1. ユーザーのBAN状態更新

エンドポイント URI:

PATCH /api/users/{id}/ban-status

対象項目:
	•	is_banned

リクエスト例:

{
  "is_banned": true,
  "reason": "Violation of community guidelines"
}

2. ユーザー名更新

エンドポイント URI:

PATCH /api/users/{id}/username

対象項目:
	•	username

リクエスト例:

{
  "username": "hiroto"
}

注意事項:
	•	サーバー側でユニーク性をチェックし、重複がある場合はエラーレスポンスを返す。

3. メールアドレス更新

エンドポイント URI:

PATCH /api/users/{id}/email

対象項目:
	•	email

リクエスト例:

{
  "email": "hiroto@example.com"
}

注意事項:
	•	認証フローの追加: メールアドレスの更新には認証が必要です。新しいメールアドレスに確認用メールを送信し、リンクをクリックした時点で更新が確定する仕組みにします。

設計のポイント

	1.	RESTfulに従う
各カテゴリごとにまとめることで、URI設計がシンプルでわかりやすくなります。
	2.	PATCHメソッドを使用
各エンドポイントは部分的な更新を行うため、PATCHを使用。
	3.	ID指定
URLで{id}を指定することで、対象ユーザーを一意に特定します。
	4.	バリデーション
必要に応じて、サーバー側でリクエスト内容のバリデーションを実施。例えば、usernameやemailはユニーク性を確認。
	5.	将来的な拡張を意識
設定項目や操作内容が増えても、カテゴリ単位でURIを拡張するだけで対応可能です。

この設計を元に実装すれば、保守性や拡張性の高いAPIが構築できます。