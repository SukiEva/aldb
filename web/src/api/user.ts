import http from "~/utils/request";

export const registerUser = (data: object) => {
    return http.request({
        url: "/register",
        method: "post",
        data,
    });
};
