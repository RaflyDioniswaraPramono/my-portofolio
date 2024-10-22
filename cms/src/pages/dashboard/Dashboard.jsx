import { Outlet } from "react-router-dom";
import { Appbar } from "@components/template";

const Dashboard = () => {
	return (
		<div>
			<Appbar />
			<Outlet />
		</div>
	);
};

export default Dashboard;
