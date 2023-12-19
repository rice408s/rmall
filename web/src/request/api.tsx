/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import axios from "axios";
import * as apiUrl from "@/request/api-url";
import { notification } from "antd";
const key = "keepOnlyOne";
/**
 *  接口请求数据时执行的方法
 *  接受参数为请求的路径apiUrl、请求接口配置参数configObj
 * @param {String} apiUrl            用户传入的请求路径
 * @param {Object} configObj        用户传入的接口参数
 */
async function getDataFromServer(apiUrl: any, configObj: any): Promise<unknown> {
	// 用户传入的接口配置参数 （默认）
	const { method = "POST", params = {}, data = {}, timeout = 100, headers = {} } = configObj;
	return await new Promise(function (resolve, reject) {
		axios({
			url: apiUrl,
			method,
			params,
			data,
			timeout,
			headers: {
				...headers
			}
		})
			.then(function (response) {
				notification.success({
					key,
					message: "成功"
				});
				resolve(response);
			})
			.catch(function (error) {
				notification.error({
					key,
					message: "错误"
				});
				reject(error);
			});
	});
}
// 用户注册
export async function userRegisterApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userRegister, configObj);
}
// 用户登录
export async function userLoginApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userLogin, configObj);
}
// 用户信息
export async function userInfoApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userInfo, configObj);
}
// 用户登出
export async function userLogoutApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userLogout, configObj);
}
// 更新用户信息
export async function userUpdateApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userUpdate, configObj);
}
// 删除用户
export async function userDeleteApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userDelete, configObj);
}
// 获取所有用户
export async function userAllApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.userAll, configObj);
}
// 管理员登陆
export async function adminLoginApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.adminLogin, configObj);
}
// 管理员注册
export async function adminRegisterApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.adminRegister, configObj);
}
// 管理员信息
export async function adminUpdateApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.adminUpdate, configObj);
}
// 管理员删除
export async function adminDeleteApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.adminDelete, configObj);
}
// 管理员信息
export async function adminAllApi(configObj: any): Promise<unknown> {
	return await getDataFromServer(apiUrl.adminAll, configObj);
}
