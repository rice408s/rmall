/* eslint-disable @typescript-eslint/no-unused-vars */
import React from "react";
// import type { CSSProperties } from 'react'
import { useNavigate } from "react-router-dom";
import { adminLoginApi } from "@/request/api";
import { notification } from "antd";
import _ from "lodash";
import { AlipayOutlined, LockOutlined, MobileOutlined, TaobaoOutlined, UserOutlined, WeiboOutlined } from "@ant-design/icons";
import { LoginFormPage, ProConfigProvider, ProFormCaptcha, ProFormCheckbox, ProFormText } from "@ant-design/pro-components";
import { Button, Divider, Space, Tabs, message, theme } from "antd";
import type { CSSProperties } from "react";
import { useState } from "react";
type LoginType = "phone" | "account";
const iconStyles: CSSProperties = {
	color: "rgba(0, 0, 0, 0.2)",
	fontSize: "18px",
	verticalAlign: "middle",
	cursor: "pointer"
};
// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function Foregroundlogin() {
	const [loginType, setLoginType] = useState<LoginType>("phone");
	const navigate = useNavigate();
	// 创建防抖函数
	const debouncedLogin = _.debounce(async values => {
		console.log(values);
		void adminLoginApi({
			data: values
		})
			.then((res: any) => {
				console.log(res);
				notification.success({
					message: "登录成功",
					description: "欢迎回来"
				});
			})
			.catch((err: any) => {
				console.log(err);
				notification.error({
					message: "密码或用户名错误",
					description: "请重新输入"
				});
			});
	}, 1000);
	return (
		<div
			style={{
				backgroundColor: "white",
				height: "100vh"
			}}
		>
			<LoginFormPage
				backgroundImageUrl="https://pic.imgdb.cn/item/658189fdc458853aef694884.jpg"
				logo="https://github.githubassets.com/images/modules/logos_page/Octocat.png"
				// backgroundVideoUrl="https://gw.alipayobjects.com/v/huamei_gcee1x/afts/video/jXRBRK_VAwoAAAAAAAAAAAAAK4eUAQBr"
				title="宏伟商城"
				onFinish={values => {
					return Promise.resolve(debouncedLogin(values));
				}}
				containerStyle={{
					backgroundColor: "rgba(255, 255, 255,0.25)",
					backdropFilter: "blur(4px)"
				}}
				subTitle="全球最大的商城平台"
				actions={
					<div
						style={{
							display: "flex",
							justifyContent: "center",
							alignItems: "center",
							flexDirection: "column"
						}}
					></div>
				}
			>
				<Tabs centered activeKey={loginType} onChange={activeKey => setLoginType(activeKey as LoginType)}>
					<Tabs.TabPane key={"account"} tab={"账号密码登录"} />
					<Tabs.TabPane key={"phone"} tab={"手机号登录"} />
				</Tabs>
				{loginType === "account" && (
					<>
						<ProFormText
							name="username"
							fieldProps={{
								size: "large",
								prefix: (
									<UserOutlined
										style={
											{
												// color: token.colorText
											}
										}
										className={"prefixIcon"}
										rev={undefined}
									/>
								)
							}}
							placeholder={"用户名: admin or user"}
							rules={[
								{
									required: true,
									message: "请输入用户名!"
								}
							]}
						/>
						<ProFormText.Password
							name="password"
							fieldProps={{
								size: "large",
								prefix: (
									<LockOutlined
										style={
											{
												// color: token.colorText
											}
										}
										className={"prefixIcon"}
										rev={undefined}
									/>
								)
							}}
							placeholder={"密码: ant.design"}
							rules={[
								{
									required: true,
									message: "请输入密码！"
								}
							]}
						/>
					</>
				)}
				{loginType === "phone" && (
					<>
						<ProFormText
							fieldProps={{
								size: "large",
								prefix: (
									<MobileOutlined
										style={
											{
												// color: token.colorText
											}
										}
										className={"prefixIcon"}
										rev={undefined}
									/>
								)
							}}
							name="mobile"
							placeholder={"手机号"}
							rules={[
								{
									required: true,
									message: "请输入手机号！"
								},
								{
									pattern: /^1\d{10}$/,
									message: "手机号格式错误！"
								}
							]}
						/>
						<ProFormCaptcha
							fieldProps={{
								size: "large",
								prefix: (
									<LockOutlined
										style={
											{
												// color: token.colorText
											}
										}
										className={"prefixIcon"}
										rev={undefined}
									/>
								)
							}}
							captchaProps={{
								size: "large"
							}}
							placeholder={"请输入验证码"}
							captchaTextRender={(timing, count) => {
								if (timing) {
									return `${count} ${"获取验证码"}`;
								}
								return "获取验证码";
							}}
							name="captcha"
							rules={[
								{
									required: true,
									message: "请输入验证码！"
								}
							]}
							onGetCaptcha={async () => {
								message.success("获取验证码成功！验证码为：1234");
							}}
						/>
					</>
				)}
				<div
					style={{
						marginBlockEnd: 24
					}}
				>
					<ProFormCheckbox noStyle name="autoLogin">
						自动登录
					</ProFormCheckbox>
					<a
						style={{
							float: "right"
						}}
					>
						忘记密码
					</a>
				</div>
			</LoginFormPage>
		</div>
	);
}
