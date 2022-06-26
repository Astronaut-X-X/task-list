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
export const getWeekPlan = (data) => {
    return request({
        url: '/api/v1/weekplan',
        method: 'get',
        data
    })
}

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
export const addWeekPlan = (data) => {
    return request({
        url: '/api/v1/weekplan',
        method: 'post',
        data
    })
}

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
export const updateWeekPlan = (data) => {
    return request({
        url: '/api/v1/weekplan',
        method: 'put',
        data
    })
}

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
export const deleteWeekPlan = (data) => {
    return request({
        url: '/api/v1/weekplan',
        method: 'delete',
        data
    })
}