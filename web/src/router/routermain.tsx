import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom'
import React, { Suspense } from 'react'
import { Foregroundlogin } from '../pages/login/login'
import PrivateRoute from './privateRoute/privateRoute'
import { Spin } from 'antd'
import LargeDataScreen from '../backstage-pages/largeDataScreen/largeDataScreen'
const Register = React.lazy(() => import('../pages/register/register'))
const Home = React.lazy(() => import('../pages/home/home'))
const Backstagelogin = React.lazy(() => import('../backstage-pages/login/login'))
const Backstagehome = React.lazy(() => import('../backstage-pages/home/home'))
export default function Routemain() {
  return (
    <div>
      <BrowserRouter>
        <Suspense fallback={<Spin></Spin>}>
          <Routes>
            <Route path='/' element={<Navigate to='/home' />} />
            <Route path='/home' element={<Home></Home>} />
            <Route path='/foregroundlogin' element={<Foregroundlogin></Foregroundlogin>} />
            <Route path='/backstage/login' element={<Backstagelogin></Backstagelogin>} />
            <Route path='/backstage/*' element={<PrivateRoute><Backstagehome></Backstagehome></PrivateRoute>} />
            <Route path='/largedatascreen' element={<LargeDataScreen></LargeDataScreen>}></Route>
            <Route path='/register' element={<Register></Register>}></Route>
          </Routes>
        </Suspense>
      </BrowserRouter>
    </div>
  )
}
