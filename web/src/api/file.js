import request from "@/utils/request";

export const downloadFileStream = (data) => {
    return request({
        method: 'GET',
        url: "/admin/file/stream",
        data: data,
        responseType: 'blob'
    })
}