import { NotificationsProvider } from "reapop";
import { Loader } from "./components/ui";
import Routes from "./routes";

const App = () => {
	return (
		<NotificationsProvider>
			<Loader />
			<Routes />
		</NotificationsProvider>
	);
};

export default App;
