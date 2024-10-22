import react from "@vitejs/plugin-react-swc";
import { resolve } from "path";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	resolve: {
		alias: {
			"@components": resolve(__dirname, "src/components"),
			"@pages": resolve(__dirname, "src/pages"),
			"@assets": resolve(__dirname, "src/assets"),
			"@hooks": resolve(__dirname, "src/hooks"),
			"@utils": resolve(__dirname, "src/utils"),
			"@store": resolve(__dirname, "src/store"),
		},
	},
});
