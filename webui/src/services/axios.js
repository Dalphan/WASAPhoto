import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

const setAuth = () => {
	instance.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage['token']
}

export default instance;
