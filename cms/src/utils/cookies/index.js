import cookies from "js-cookie";

export const SetCookies = (tokenString) => {
	const token = JSON.stringify(tokenString);

	return cookies.set("ACCESS_TOKEN", token, {
		expires: new Date(Date.now() + 3600 * 24000),
		sameSite: true,
	});
};

export const GetCookies = (cookieName) => {
	return cookies.get(cookieName);
};

export const DeleteCookies = (cookieName) => {
	return cookies.remove(cookieName);
};
