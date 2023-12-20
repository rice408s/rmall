import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom'
import React from 'react'
import Home from '../pages/home/home'
import { Foregroundlogin } from '../pages/login/login'
import Backstagelogin from '../backstage-pages/login/login'
export default function Routemain() {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Navigate to='/home' />} />
          <Route path='/home' element={<Home></Home>} />
          <Route path='/foregroundlogin' element={<Foregroundlogin></Foregroundlogin>} />
          <Route path='/backstagelogin' element={<Backstagelogin></Backstagelogin>} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}
