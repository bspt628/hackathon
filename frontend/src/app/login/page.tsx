"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuth } from "@/contexts/auth-context";

export default function LoginPage() {
	const router = useRouter();
	const { login } = useAuth();
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(false);

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setIsLoading(true);
		setError(null);

		const formData = new FormData(event.currentTarget);
		const username = formData.get("username") as string;
		const password = formData.get("password") as string;

		try {
			// First, fetch the email for the username
			const emailResponse = await fetch(
				`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/users/email/${username}`
			);

			if (!emailResponse.ok) {
				throw new Error("Failed to fetch email");
			}

			const emailData = await emailResponse.json();

			if (!emailData.email) {
				throw new Error("Email not found for username");
			}

			// Then sign in with the email and password
			await login(emailData.email, password);
			router.push("/home");
		} catch (error) {
			console.error("Login error:", error);
			setError(
				error instanceof Error ? error.message : "ログインに失敗しました。"
			);
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<div className="min-h-screen bg-[#F8FAFF] text-foreground bg-gradient-to-br from-[#E6EFFF] to-[#F8FAFF] flex items-center justify-center p-4">
			<div className="w-full max-w-md bg-background rounded-2xl p-8 relative">
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
							className="bg-white border-[#536471] focus:border-primary text-black"
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
							className="bg-white border-[#536471] focus:border-primary text-black"
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
