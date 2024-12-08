"use server";

export async function deletePost(postId: string, idToken: string) {
	try {
		const response = await fetch(
			`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/${postId}`,
			{
				method: "DELETE",
				headers: {
					Authorization: `Bearer ${idToken}`,
				},
			}
		);

		if (!response.ok) {
			throw new Error("Failed to delete post");
		}

		return { success: true };
	} catch (error) {
		console.error("Delete post error:", error);
		return {
			success: false,
			error:
				error instanceof Error ? error.message : "投稿の削除に失敗しました。",
		};
	}
}
