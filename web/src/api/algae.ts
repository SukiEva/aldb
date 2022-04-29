import http from "~/utils/request";

export const getArticleList = (data:object) =>{
    return http.request({
        url: '/app/blog/article/page',
        method: 'POST',
        data
    })
}