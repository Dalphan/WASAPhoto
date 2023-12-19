import axios from "axios";
import getCurrentId from "./authentication";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

const setAuth = () => {
	// console.log("QUI PASSA")
	instance.defaults.headers.common['Authorization'] = 'Bearer ' + getCurrentId()
}

export {
	setAuth,
	instance as axios,
}