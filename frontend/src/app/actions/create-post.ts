"use server";

interface CreatePostRequest {
	user_id: string;
	content: string;
	media_urls: string[];
	visibility: string;
	original_post_id: string | null;
	reply_to_id: string | null;
	root_post_id: string | null;
	is_repost: boolean;
	is_reply: boolean;
}

export async function createPost(content: string, idToken: string) {
	try {
		// First, get the user ID from the token
		const userResponse = await fetch(
            `https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/users/firebase`,
			{
				headers: {
					Authorization: `Bearer ${idToken}`,
				},
			}
		);

		if (!userResponse.ok) {
			throw new Error("Failed to fetch user data");
		}

		const userData = await userResponse.json();
		const userId = userData.id;

		const postData: CreatePostRequest = {
			user_id: userId,
			content,
			media_urls: [],
			visibility: "public",
			original_post_id: null,
			reply_to_id: null,
			root_post_id: null,
			is_repost: false,
			is_reply: false,
		};

		if (content.length > 140) {
			throw new Error("Post content is too long");
		}

		const response = await fetch(
			"https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts",
			{
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${idToken}`,
				},
				body: JSON.stringify(postData),
			}
		);

		if (!response.ok) {
			throw new Error("Failed to create post");
		}

		return { success: true };
	} catch (error) {
		console.error("Create post error:", error);
		return {
			success: false,
			error:
				error instanceof Error ? error.message : "ポストの作成に失敗しました。",
		};
	}
}
