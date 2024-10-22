import { useState, useEffect, useCallback } from "react";
import { useLoader } from "@store/useLoader";
import { api } from "@utils/api";

export const useFetch = ({ url, method }) => {
	const [responseData, setResponseData] = useState([]);

	const { setIsLoading } = useLoader();

	const fetchContent = useCallback(async () => {
		try {
			setIsLoading(true);

			const response = await api({
				url,
				method,
      });
      
			setResponseData(response?.data);
		} catch (error) {
			console.log(error);
		} finally {
			setIsLoading(false);
		}
	}, [method, setIsLoading, url]);

	useEffect(() => {
		fetchContent();
	}, [fetchContent]);

	return { responseData };
};
