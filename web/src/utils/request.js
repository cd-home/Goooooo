import axios from "axios";
import ElMessage from "element-plus"

const instance = axios.create({
    baseURL: process.env.VITE_APP_ADDR,
    timeout: process.env.VITE_APP_TIMEOUT,
});

请求拦截器
instance.interceptors.request.use(
    config => {
        // 请求头设置、Token等
        return config
    },
    error => {
        // DEBUG 调试
        console.log(error)
        return Promise.reject(error)
    }
);

// 响应拦截器
instance.interceptors.request.use(
    response => {
        // 调用失败，根据code不同提示不同的业务信息，供参考
        if (resp.code === 1) {
            ElMessage.warn(message);
            return Promise.reject(new Error(resp.message))
        }
        // 调用成功
        return response
    },
    error => {
        console.log(error)
        const status = error.response.status
        if (error.response && status) {
            let message = "请求失败!"
            switch (status) {
                default:
                    message = "请求失败!"
                    break;
                case 400:
                    message = "请求错误!"
                    break;
                case 401:
                    message = "认证失败!"
                    break;
                case 404:
                    message = "资源不存在!"
                    break;
                case 408:
                    message = "请求超时!"
                    break;
                case 500:
                    message = "服务器内部错误!"
                    break;
                case 502:
                    message = "网关错误!"
                    break;
                case 503:
                    message = "服务不可用!"
                    break;
                case 504:
                    message = "网关超时!"
                    break;
            }
            ElMessage.error(message);
            return Promise.reject(error);
        }
        return Promise.reject(error);
    }
);

// 封装请求方法
const request = ({method, url, data, config}) => {
    method = method.toUpperCase();
    switch (method) {
        case "POST":
            return instance.post(url, data, {...config})
        case "GET":
            return instance.get(url, {params: data, ...config})
        case "PUT":
            return instance.put(url, data, {...config})
        case "DELETE":
            return instance.delete(url, {params: data, ...config})
        default:
            console.log("请求方法不被允许")
            return false
    }
}

export default request;
