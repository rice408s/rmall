import { Navigate, useRoutes } from "react-router-dom";
import { RouteObject } from "@/routers/interface";
import Login from "@/views/login/index";
import Home from "@/pages/home/home";
import Foregroundlogin from "@/pages/login/login";
// * 导入所有router
const metaRouters = import.meta.globEager("./modules/*.tsx");

// * 处理路由
export const routerArray: RouteObject[] = [];

Object.keys(metaRouters).forEach(item => {
	Object.keys(metaRouters[item]).forEach((key: any) => {
		routerArray.push(...metaRouters[item][key]);
	});
});

export const rootRouter: RouteObject[] = [
	{
		path: "/",
		element: <Navigate to="/login" />
	},
	{
		path: "/login",
		element: <Login />,
		meta: {
			requiresAuth: false,
			title: "登录页",
			key: "login"
		}
	},
	{
		// 电商主页
		path: "/homemain",
		element: <Home />,
		meta: {
			requiresAuth: false,
			title: "电商主页",
			key: "homemain"
		}
	},
	{
		// 电商登录页
		path: "/foregroundlogin",
		element: <Foregroundlogin />,
		meta: {
			requiresAuth: false,
			title: "前台登录页",
			key: "foregroundlogin"
		}
	},
	...routerArray,
	{
		path: "*",
		element: <Navigate to="/404" />
	}
];

const Router = () => {
	const routes = useRoutes(rootRouter);
	return routes;
};

export default Router;
