import http from "~/utils/request";

export const registerUser = (data: object) => {
    return http.request({
        url: "/register",
        method: "post",
        data,
    });
};

export const getUser = (query: string) => {
    return http.request({
        url: "/api/user/info?email=" + query,
        method: "get"
    });
};

export const changePassword = (data: object) => {
    return http.request({
        url: "/api/user/pwd",
        method: "post",
        data,
    });
};
