import { Button } from 'antd'
import { useNavigate } from 'react-router-dom'
import React from 'react'
export default function Home() {
  const navigate = useNavigate()
  return (
    <div>
      <Button
        onClick={() => {
          navigate('/backstage/login')
        }}
      >
				管理员登录入口
      </Button>
      <Button
        onClick={() => {
          navigate('/foregroundlogin')
        }}
      >
				用户登录入口
      </Button>
    </div>
  )
}
