6期

# ToDo

settingとprofileのusecase層の記法を統一

json.Rawmessageの渡し方について

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


# 進捗

## 1121
### UpdateUserProfileの成功
更新のためにidを引数に渡し忘れるというミスに気づいて、わりかしすんなりと実装できた（1hくらい）

### 明日以降の課題
バリデーションを設置してみる

**最適な設計の提案**

推奨方法: バリデーションをUsecase層で行い、一部をController層で補完する
- Controller 層:

     - リクエスト形式や必須フィールドのチェックなど、基本的なバリデーションを実施。

    - 例: リクエストボディが空でないか、JSONの構造が正しいか。

    - 入力が不正である場合、すぐにエラーを返す。

- Usecase 層:

    - ビジネスロジックに基づく詳細なバリデーションを実施。

    - 例: フィールドの内容や長さ、正規表現チェック、ドメイン固有の制約。

**具体例**

1. ProfileImageUrl (プロフィール画像URL)

バリデーション例:
- 必須: プロフィール画像を必須とする場合、空でないことを検証。
- URL形式の確認: 画像URLが有効なURLであることを確認。
- プロトコル制限: HTTPまたはHTTPSで始まることを検証。
- ファイル形式の確認: .jpg, .png, .gif などの拡張子を許可。
実装例:

```
if profileImageUrl != "" {
    if !isValidURL(profileImageUrl) {
        return errors.New("プロフィール画像のURLが無効です")
    }
    if !strings.HasPrefix(profileImageUrl, "http://") && !strings.HasPrefix(profileImageUrl, "https://") {
        return errors.New("プロフィール画像のURLはHTTPまたはHTTPSである必要があります")
    }
    if !strings.HasSuffix(profileImageUrl, ".jpg") && !strings.HasSuffix(profileImageUrl, ".png") && !strings.HasSuffix(profileImageUrl, ".gif") {
        return errors.New("プロフィール画像はJPG, PNG, GIF形式である必要があります")
    }
}
```
2. Bio (自己紹介)

バリデーション例:
- 最大長: 文字数制限（例: 最大500文字）。
- 特殊文字制限: 禁止された文字やスクリプトを含んでいないか確認。
- 空白のみの内容を禁止: 空白文字のみの入力を禁止。

実装例:

```
if len(bio) > 500 {
    return errors.New("自己紹介は500文字以内で入力してください")
}
if strings.TrimSpace(bio) == "" {
    return errors.New("自己紹介に空白以外の内容を入力してください")
}
```

3. Location (居住地)

バリデーション例:
- 最大長: 文字数制限（例: 最大100文字）。
- フォーマット確認: 入力値が一般的な地名の形式に準拠しているか。
- 禁止ワード: 不適切な単語や攻撃的な内容を含まないか確認。

実装例:
```
if len(location) > 100 {
    return errors.New("居住地は100文字以内で入力してください")
}
if containsProhibitedWords(location) {
    return errors.New("不適切な内容が含まれています")
}
```
4. Website (ウェブサイトURL)

バリデーション例:
- 必須: ウェブサイトを必須とする場合、空でないことを検証。
- URL形式の確認: 有効なURLであることを確認。
- プロトコル制限: HTTPまたはHTTPSで始まることを検証。
- 文字数制限: URLの長さが適切であることを確認（例: 最大255文字）。

実装例:
```
if website != "" {
    if len(website) > 255 {
        return errors.New("ウェブサイトのURLは255文字以内で入力してください")
    }
    if !isValidURL(website) {
        return errors.New("ウェブサイトのURLが無効です")
    }
    if !strings.HasPrefix(website, "http://") && !strings.HasPrefix(website, "https://") {
        return errors.New("ウェブサイトのURLはHTTPまたはHTTPSである必要があります")
    }
}
```
補助関数の例
URLの検証関数:

```
import "net/url"

func isValidURL(u string) bool {
    _, err := url.ParseRequestURI(u)
    return err == nil
}
```

禁止ワードの検証関数:

```func containsProhibitedWords(input string) bool {
    prohibitedWords := []string{"inappropriate", "bannedword"} // 不適切な単語リスト
    for _, word := range prohibitedWords {
        if strings.Contains(strings.ToLower(input), word) {
            return true
        }
    }
    return false
}
```

## 1122

### 成果
- UpdateUserProfileのバリデーションをかけた
- UpdateUserSettingsを実装した（バリデーションこみ）

### 大変だったところ
- 誕生日を扱うところ
    - Time.parse関数で好きなフォーマットになるよう処理しなければいけない（2006-01-02というふうにした）