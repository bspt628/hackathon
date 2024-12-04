import { MessageSquare, Repeat2, Heart } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ja } from "date-fns/locale";

interface PostProps {
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
}

export function Post({
	display_name,
	username,
	created_at,
	content,
	replies_count,
	reposts_count,
	likes_count,
}: PostProps) {
	const timeAgo = formatDistanceToNow(new Date(created_at), {
		addSuffix: true,
		locale: ja,
	});

	return (
		<div className="border-b border-[#2f3336] p-4 hover:bg-white/[0.03]">
			<div className="flex gap-4">
				<div className="w-10 h-10 rounded-full bg-[#2f3336]" />
				<div className="flex-1">
					<div className="flex items-center gap-2">
						<span className="font-bold">{display_name}</span>
						<span className="text-[#71767b]">@{username}</span>
						<span className="text-[#71767b]">·</span>
						<span className="text-[#71767b]">{timeAgo}</span>
					</div>
					<p className="mt-2 break-words">{content}</p>
					<div className="flex justify-between mt-4 max-w-md text-[#71767b]">
						<button className="flex items-center gap-2 hover:text-[#1d9bf0]">
							<MessageSquare className="w-5 h-5" />
							<span>{replies_count}</span>
						</button>
						<button className="flex items-center gap-2 hover:text-[#00ba7c]">
							<Repeat2 className="w-5 h-5" />
							<span>{reposts_count}</span>
						</button>
						<button className="flex items-center gap-2 hover:text-[#f91880]">
							<Heart className="w-5 h-5" />
							<span>{likes_count}</span>
						</button>
					</div>
				</div>
			</div>
		</div>
	);
}
