import { create } from "zustand";

export const useLoader = create((set) => ({
	isLoading: false,
	setIsLoading: (isLoading) =>
		set({
			isLoading: isLoading,
		}),
}));
