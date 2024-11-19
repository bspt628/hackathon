CREATE TABLE password_reset_tokens (
    id SERIAL PRIMARY KEY,  -- トークンの一意なID
    email VARCHAR(255) NOT NULL,  -- トークンが関連するユーザーのメールアドレス
    token VARCHAR(255) NOT NULL,  -- パスワードリセット用のトークン
    expiry TIMESTAMP NOT NULL,  -- トークンの有効期限
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- トークンの作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 更新日時
);
