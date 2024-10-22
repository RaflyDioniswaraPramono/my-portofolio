import { createBrowserRouter } from "react-router-dom";
import { Signin, Dashboard, NotFound } from "../pages";
import {
	DashboardContent,
	BiographyContent,
	PortofolioContent,
	ExperienceContent,
	SosmedContent,
	MessageContent,
} from "@components/layout";

export const router = createBrowserRouter([
	{
		errorElement: <NotFound />,
	},
	{
		path: "/",
		element: "",
	},
	{
		path: "/dashboard",
		element: <Dashboard />,
		children: [
			{
				index: true,
				element: <DashboardContent />,
			},
			{
				path: "biography",
				element: <BiographyContent />,
			},
			{
				path: "portofolio",
				element: <PortofolioContent />,
			},
			{
				path: "experience",
				element: <ExperienceContent />,
			},
			{
				path: "sosmed",
				element: <SosmedContent />,
			},
			{
				path: "message",
				element: <MessageContent />,
			},
		],
	},
	{
		path: "/signin",
		element: <Signin />,
	},
]);
