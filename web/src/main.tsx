import React from 'react'
import ReactDOM from 'react-dom/client'
// 样式初始化
import 'reset-css'
// 全局样式
import '@/assets/styles/global.module.scss'
import App from './App.tsx'
// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
