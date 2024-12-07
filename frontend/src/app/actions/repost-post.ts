"use server";

export async function repostPost(
	postId: string,
	isQuoteRepost: boolean,
	additionalComment: string,
	idToken: string
) {
	try {
		const response = await fetch(
			"https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/repost",
			{
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${idToken}`,
				},
				body: JSON.stringify({
					original_post_id: postId,
					is_quote_repost: isQuoteRepost,
					additional_comment: additionalComment,
				}),
			}
		);

		if (!response.ok) {
			throw new Error("Failed to repost");
		}

		return { success: true };
	} catch (error) {
		console.error("Repost error:", error);
		return {
			success: false,
			error:
				error instanceof Error ? error.message : "リポストに失敗しました。",
		};
	}
}
