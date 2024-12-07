"use server";

export async function likePost(postId: string, idToken: string) {
	try {
		const response = await fetch(
			`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/likes/${postId}`,
			{
				method: "POST",
				headers: {
					Authorization: `Bearer ${idToken}`,
				},
			}
		);

		if (!response.ok) {
			throw new Error("Failed to like post");
		}

		return { success: true };
	} catch (error) {
		console.error("Like post error:", error);
		return {
			success: false,
			error: error instanceof Error ? error.message : "いいねに失敗しました。",
		};
	}
}

export async function unlikePost(postId: string, idToken: string) {
	try {
		const response = await fetch(
			`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/likes/${postId}`,
			{
				method: "DELETE",
				headers: {
					Authorization: `Bearer ${idToken}`,
				},
			}
		);

		if (!response.ok) {
			throw new Error("Failed to unlike post");
		}
	
		return { success: true };
	}
	catch (error) {
		console.error("Unlike post error:", error);
		return {
			success: false,
			error: error instanceof Error ? error.message : "いいねの取り消しに失敗しました。",
		};
	}
}