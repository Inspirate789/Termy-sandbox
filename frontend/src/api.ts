import axios from 'axios'
import config from './config'
import authService from "./services/AuthService";

const api = axios.create({
    baseURL: config.api_url,
});

api.interceptors.request.use((config) => {
    config.headers.Authorization = "Bearer " + authService.getToken();
    return config;
})

// api.interceptors.response.use(response => response, error => {
//     const status = error.response ? error.response.status : null
//
//     if (status === 401) {
//         return authService.refresh();
//     }
//
//     return Promise.reject(error);
// });

export default api;
