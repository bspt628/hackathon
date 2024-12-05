"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { loginUser } from "../actions/login";

export default function LoginPage() {
	const router = useRouter();
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(false);

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setIsLoading(true);
		setError(null);

		const result = await loginUser(new FormData(event.currentTarget));

		if (result.success) {
			router.push("/home"); // Redirect to home page after successful login
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
				>
				</button>

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