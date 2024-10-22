import axios from "axios";
import { GetCookies } from "../cookies";

export const api = axios.create({
	baseURL: import.meta.env.VITE_API_BASE_URL,
	headers: {
		Accept: "application/json",
		"Content-Type": "application/json",
	},
});

api.interceptors.request.use(
	(config) => {
		const token = GetCookies("ACCESS_TOKEN");

		if (token) {
			const parsedToken = JSON.parse(token);

			config.headers.Authorization = `Bearer ${parsedToken}`;
		}

		return config;
	},
	(error) => {
		throw new Promise.reject(error);
	}
);
