import request from '../utils/request';

// @Tags user
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/v1/user [get]
// {
//  page     int
//	pageSize int
// }
export const getUserByIDInContext = (data) => {
    return request({
        url: '/api/v1/user',
        method: 'get',
        data
    })
}

// @Tags user
// @Summary 登录用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data 
// @Success 200 {string} json "{"success":true,"data":{token:'token'},"msg":"获取成功"}"
// @Router /api/auth/login [post]
export const loginUser = (data) => {
    return request({
        url: "/api/auth/login",
        method: "post",
        data
    })
}