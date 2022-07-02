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
export const getTodo = (data) => {
    return request({
        url: '/api/v1/todo',
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
export const getTodayTodo = (data) => {
    return request({
        url: '/api/v1/todo/today',
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
export const addTodo = (data) => {
    return request({
        url: '/api/v1/todo',
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
export const updateTodo = (data) => {
    return request({
        url: '/api/v1/todo',
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
export const deleteTodo = (data) => {
    return request({
        url: '/api/v1/todo',
        method: 'delete',
        data
    })
}