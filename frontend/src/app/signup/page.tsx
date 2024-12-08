"use client";

import { useRouter } from "next/navigation";
import { useState, useEffect, useCallback } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { signupUser } from "../actions/signup";

export default function SignupPage() {
	const router = useRouter();
	const [isLoading, setIsLoading] = useState(false);
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [confirmPassword, setConfirmPassword] = useState("");
	const [username, setUsername] = useState("");
	const [displayName, setDisplayName] = useState("");
	const [touched, setTouched] = useState({
		email: false,
		password: false,
		confirmPassword: false,
		username: false,
		displayName: false,
	});
	const [errors, setErrors] = useState({
		email: "",
		password: "",
		confirmPassword: "",
		username: "",
		displayName: "",
		form: "",
	});

	const validateEmail = (email: string) => {
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		return emailRegex.test(email)
			? ""
			: "メールアドレスの形式が正しくありません。";
	};

	const validatePassword = (password: string) => {
		return password.length > 5 ? "" : "パスワードは6文字以上にしてください。";
	};

	const validateConfirmPassword = useCallback(
		(confirmPassword: string) => {
			return confirmPassword === password ? "" : "パスワードが一致しません。";
		},
		[password]
	);

	const validateUsername = (username: string) => {
		return username ? "" : "ユーザーネームは必須です。";
	};

	const validateDisplayName = (displayName: string) => {
		return displayName ? "" : "表示名は必須です。";
	};

	useEffect(() => {
		setErrors({
			email: touched.email ? validateEmail(email) : "",
			password: touched.password ? validatePassword(password) : "",
			confirmPassword: touched.confirmPassword
				? validateConfirmPassword(confirmPassword)
				: "",
			username: touched.username ? validateUsername(username) : "",
			displayName: touched.displayName ? validateDisplayName(displayName) : "",
			form: "",
		});
	}, [
		email,
		password,
		confirmPassword,
		username,
		displayName,
		validateConfirmPassword,
		touched,
	]);

	const handleBlur = (field: keyof typeof touched) => {
		setTouched((prev) => ({ ...prev, [field]: true }));
	};

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setIsLoading(true);

		// すべてのフィールドをタッチ済みにする
		setTouched({
			email: true,
			password: true,
			confirmPassword: true,
			username: true,
			displayName: true,
		});

		if (Object.values(errors).some((error) => error !== "")) {
			setIsLoading(false);
			return;
		}

		const formData = new FormData(event.currentTarget);
		const result = await signupUser(formData);

		if (result.success) {
			router.push("/home");
		} else {
			setErrors({
				...errors,
				form: result.error ?? "アカウントの作成に失敗しました。",
			});
		}
		setIsLoading(false);
	}

	return (
		<div className="min-h-screen bg-[#F8FAFF] text-foreground bg-gradient-to-br from-[#E6EFFF] to-[#F8FAFF] flex items-center justify-center p-4">
			<div className="w-full max-w-md bg-background rounded-2xl p-8 relative">
				<button
					onClick={() => router.back()}
					className="absolute top-4 left-4 p-2 hover:bg-white/10 rounded-full"
				></button>

				<div className="mb-8 text-center">
					<h1 className="text-2xl font-bold mb-2">アカウントを作成</h1>
				</div>

				<form onSubmit={handleSubmit} className="space-y-6">
					<div className="space-y-2">
						<Label htmlFor="email">メールアドレス</Label>
						<Input
							id="email"
							name="email"
							type="email"
							required
							className="bg-white border-[#536471] focus:border-primary text-black"
							placeholder="example@email.com"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
							onBlur={() => handleBlur("email")}
						/>
						{touched.email && errors.email && (
							<div className="text-red-500 text-sm">{errors.email}</div>
						)}
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
							value={password}
							onChange={(e) => setPassword(e.target.value)}
							onBlur={() => handleBlur("password")}
						/>
						{touched.password && errors.password && (
							<div className="text-red-500 text-sm">{errors.password}</div>
						)}
					</div>

					<div className="space-y-2">
						<Label htmlFor="confirm_password">パスワード（確認）</Label>
						<Input
							id="confirm_password"
							name="confirm_password"
							type="password"
							required
							className="bg-white border-[#536471] focus:border-primary text-black"
							placeholder="••••••••"
							value={confirmPassword}
							onChange={(e) => setConfirmPassword(e.target.value)}
							onBlur={() => handleBlur("confirmPassword")}
						/>
						{touched.confirmPassword && errors.confirmPassword && (
							<div className="text-red-500 text-sm">
								{errors.confirmPassword}
							</div>
						)}
					</div>

					<div className="space-y-2">
						<Label htmlFor="username">ユーザーネーム</Label>
						<Input
							id="username"
							name="username"
							required
							className="bg-white border-[#536471] focus:border-primary text-black"
							placeholder="@username"
							value={username}
							onChange={(e) => setUsername(e.target.value)}
							onBlur={() => handleBlur("username")}
						/>
						{touched.username && errors.username && (
							<div className="text-red-500 text-sm">{errors.username}</div>
						)}
					</div>

					<div className="space-y-2">
						<Label htmlFor="display_name">表示名</Label>
						<Input
							id="display_name"
							name="display_name"
							required
							className="bg-white border-[#536471] focus:border-primary text-black"
							placeholder="表示名"
							value={displayName}
							onChange={(e) => setDisplayName(e.target.value)}
							onBlur={() => handleBlur("displayName")}
						/>
						{touched.displayName && errors.displayName && (
							<div className="text-red-500 text-sm">{errors.displayName}</div>
						)}
					</div>

					{errors.form && (
						<div className="text-red-500 text-sm text-center">
							{errors.form}
						</div>
					)}

					<Button
						type="submit"
						disabled={
							isLoading || Object.values(errors).some((error) => error !== "")
						}
						className="w-full bg-white hover:bg-white/90 text-black"
					>
						{isLoading ? "処理中..." : "次へ"}
					</Button>
				</form>
			</div>
		</div>
	);
}
