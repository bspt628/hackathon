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

export async function replyToPost(postId: string, content: string, idToken: string) {
	try {
		// First, get the user ID from the token
		console.log("idToken", idToken);
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
		const userId = userData.user_id;
		console.log("userId", userId);

		console.log("reply to postId", postId);

		
		
		// post_idの元ポストのIDを取得する
		const originalPostResponse = await fetch(
			`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/timeline/one/${postId}`,
			{
				headers: {
					Authorization: `Bearer ${idToken}`,
				},
			}
		);
		if (!originalPostResponse.ok) {
			throw new Error("Failed to fetch original post data");
		}
		const originalPostData = await originalPostResponse.json();
		console.log("originalPostData", originalPostData);

		const rootPostId = (originalPostData.root_post_id == "") ? postId : originalPostData.root_post_id;
		console.log("rootPostId", rootPostId);
			
		const postData: CreatePostRequest = {
			user_id: userId,
			content,
			media_urls: [],
			visibility: "public",
			original_post_id: null,
			reply_to_id: postId,
			root_post_id: rootPostId,
			is_repost: false,
			is_reply: true,
		};

		console.log("[HTTP Request] Replying to post:", postId, content);
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
			throw new Error("Failed to reply to post");
		}

		return { success: true };
	} catch (error) {
		console.error("Reply error:", error);
		return {
			success: false,
			error: error instanceof Error ? error.message : "返信に失敗しました。",
		};
	}
}
