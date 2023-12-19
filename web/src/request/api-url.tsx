import { publicIp } from "@/request/apilp";
// 用户注册的API端点
export const userRegister = `${publicIp}/user/register`;

// 用户登录的API端点
export const userLogin = `${publicIp}/user/login`;

// 获取用户信息的API端点
export const userInfo = `${publicIp}/user/info`;

// 获取管理员信息的API端点
export const adminInfo = `${publicIp}/admin/info`;

// 管理员登出的API端点
export const adminLogout = `${publicIp}/admin/logout`;

// 用户登出的API端点
export const userLogout = `${publicIp}/user/logout`;

// 更新用户信息的API端点
export const userUpdate = `${publicIp}/user/update`;

// 删除用户的API端点
export const userDelete = `${publicIp}/user/delete`;

// 获取所有用户的API端点
export const userAll = `${publicIp}/user/all`;
// 管理员登陆
export const adminLogin = `${publicIp}/admin/login`;
// 管理员注册
export const adminRegister = `${publicIp}/admin/register`;
// 管理员信息
export const adminUpdate = `${publicIp}/admin/update`;
// 管理员删除
export const adminDelete = `${publicIp}/admin/delete`;
// 管理员信息
export const adminAll = `${publicIp}/admin/all`;
