import { Backdrop, CircularProgress } from "@mui/material";
import { useLoader } from "@store/useLoader";

const Loader = () => {
	const { isLoading } = useLoader();

	return (
		<Backdrop sx={(theme) => ({ zIndex: theme.zIndex.drawer + 1 })} open={isLoading}>
			<div className="flex flex-col justify-center items-center gap-1">
				<CircularProgress style={{ color: "#fbbf24" }} />
				<p className="text-sm text-amber-400">Loading ...</p>
			</div>
		</Backdrop>
	);
};

export default Loader;
