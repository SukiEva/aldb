import http from "~/utils/request";

export const getData = (data: object) => {
    return http.request({
        url: "/api/data",
        method: "get",
        data,
    });
};
