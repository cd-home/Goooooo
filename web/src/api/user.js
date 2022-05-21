import request from "@/utils/request";

export const login = (data) => {
    console.log(111111)
    return request({
        method: 'POST',
        url: "/user/register",
        data: data
    })
}