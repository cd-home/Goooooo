import request from "@/utils/request";

export const login = (data) => {
    return request({
        method: 'POST',
        url: "/admin/user/login",
        data: data
    })
}

export const register = (data) => {
    return request({
        method: 'POST',
        url: "/admin/user/register",
        data: data
    })
}