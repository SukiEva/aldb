import http from "~/utils/request";

export const getData = (data: object) => {
    return http.request({
        url: "/api/data",
        method: "get",
        data,
    });
};

export const addRiver = (data: object) => {
    return http.request({
        url: "/api/river/add",
        method: "post",
        data,
    });
};

export const getRivers = (data: object) => {
    return http.request({
        url: "/api/river/all",
        method: "get",
        data,
    });
};
