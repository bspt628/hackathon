"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { loginUser } from "../actions/login";
import { useAuth } from "@/contexts/auth-context"; // useAuthをインポート

export default function LoginPage() {
	const router = useRouter();
	const { setCurrentUser, setIdToken } = useAuth(); // setCurrentUserとsetIdTokenを取得
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(false);

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setIsLoading(true);
		setError(null);
		console.log("User is logging in...");

		const formData = new FormData(event.currentTarget);
		const result = await loginUser(formData); // サーバーサイドのloginUserを呼び出す

		// ログイン処理の結果に基づいてユーザー情報を更新
		if (result.success) {
			console.log("User logged in successfully:");

			// AuthContextのuserとidTokenを更新
			setCurrentUser(JSON.parse(result.user)); // userを更新
			setIdToken(result.id_token); // IDトークンを更新
			console.log("ID token:", result.id_token);

			router.push("/home"); // ログイン成功時にホームページにリダイレクト
		} else {
			setError(result.error ?? "ログインに失敗しました。");
		}

		setIsLoading(false);
	}

	return (
		<div className="min-h-screen bg-black text-white flex items-center justify-center p-4">
			<div className="w-full max-w-md bg-black rounded-2xl p-8 relative">
				<button
					onClick={() => router.back()}
					className="absolute top-4 left-4 p-2 hover:bg-white/10 rounded-full"
				></button>

				<div className="mb-8 text-center">
					<h1 className="text-2xl font-bold mb-2">ログイン</h1>
				</div>

				<form onSubmit={handleSubmit} className="space-y-6">
					<div className="space-y-2">
						<Label htmlFor="username">ユーザーネーム</Label>
						<Input
							id="username"
							name="username"
							required
							className="bg-black border-[#536471] focus:border-[#1d9bf0] text-white"
							placeholder="@username"
						/>
					</div>

					<div className="space-y-2">
						<Label htmlFor="password">パスワード</Label>
						<Input
							id="password"
							name="password"
							type="password"
							required
							className="bg-black border-[#536471] focus:border-[#1d9bf0] text-white"
							placeholder="••••••••"
						/>
					</div>

					{error && (
						<div className="text-red-500 text-sm text-center">{error}</div>
					)}

					<Button
						type="submit"
						disabled={isLoading}
						className="w-full bg-white hover:bg-white/90 text-black"
					>
						{isLoading ? "ログイン中..." : "ログイン"}
					</Button>
				</form>
			</div>
		</div>
	);
}
