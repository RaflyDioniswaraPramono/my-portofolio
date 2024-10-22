import { AppbarLink } from "@components/ui";
import { useFetch } from "@hooks/useFetch";

const Appbar = () => {
	const { responseData } = useFetch({
		url: "/content",
		method: "GET",
	});

	return (
		<div className="bg-zinc-800">
			<div className="flex justify-start items-center">
				{responseData?.data?.map((content) => {
					const { id, content_text, link_path } = content;

					return (
						<AppbarLink key={id} end={id && 1} contentText={content_text} linkPath={link_path} />
					);
				})}
			</div>
		</div>
	);
};

export default Appbar;
