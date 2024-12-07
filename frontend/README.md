- URI設計

いいねをつける
- https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/likes/{id}
- POSTリクエスト

リポスト
引用するかしないかを選ぶ「引用リポスト」or「リポスト」
- https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/repost
- POSTリクエスト
- 引数に、
    -    IsQuoteRepost bool `json:"is_quote_repost"`
	-	AdditionalComment string `json:"additional_comment"`
をとる。
引用する場合は、
- 既存のポスト欄に、"repost from {username}"と付記してリポストする
- is_repostがtrue、root_post_idがリポスト元のidになる
- ポストの送信でfetchするURIは通常ポストの時と一緒

返信
- 既存のポスト欄に、"reply to {username}"と付記して返信をする
- is_replyがtrue, original_post_idが返信元のid, root_post_idが、
- 返信元のidのroot_post_idがNULLでなかったらその値を継承
- 返信元のidのroot_post_idがNULLだったら返信元のidを継承
- ポストの送信でfetchするURIは通常ポストの時と一緒



- https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api//follow/{id}/status




This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
