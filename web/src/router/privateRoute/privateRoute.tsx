
import { Navigate, useLocation } from 'react-router-dom'
import React from 'react'
function PrivateRoute({ children }: { children: React.ReactNode }) {
  const location = useLocation()
  // 判断用户是否登录
  function checkUserLoggedIn() {
    const token = localStorage.getItem('token')
    return token !== null
  }
  const isLoggedIn = checkUserLoggedIn()
  return isLoggedIn ? children : <Navigate to='/backstage/login' state={{ from: location }} />
}
// 使用PrivateRoute组件
export default PrivateRoute
