import { Routes, Route, BrowserRouter, Navigate } from 'react-router-dom'
import Home from '@/pages/home/home'
import Login from '@/pages/login/login'
import Register from '@/pages/register/register'
import React from 'react'
// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function Router () {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/home" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/" element={<Navigate to="/home"></Navigate>}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  )
}
