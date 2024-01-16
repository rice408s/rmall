/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable no-undef */
/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable @typescript-eslint/explicit-function-return-type */
import React from 'react'
import './register.scss'
import { Button, Form, Input } from 'antd'
import { useNavigate } from 'react-router-dom'
export default function Register() {
  const navigate = useNavigate()

  const onFinish = (values: any) => {
    console.log('Success:', values)
  }

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo)
  }
  type FieldType = {
    username?: string;
    password?: string;
    remember?: string;
  };

  return (
    <div style={{ backgroundColor: '#D7FCFA', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
      <div
        style={{
          height: '80vh',
          width: '80vw',
          backgroundColor: 'white',
          boxShadow: '0px 10px 20px rgba(0, 0, 0, 0.1)',
          borderRadius: '20px',
          display: 'flex'
        }}>
        <div style={{ width: '50vw', height: '80vh', padding: '5vh 3vw', display: 'flex', flexDirection: 'column', boxSizing: 'border-box' }}>
          <div className='Register-left-nav'>
            <div className='Register-left-nav-name'>
                宏伟商城
            </div>
            <div style={{ width: '8vw', display: 'flex', justifyContent: 'space-between' }}>
              <div style={{ color: '#8a8c97' }}><Button type='text'
                onClick={() => {
                  navigate('/foregroundlogin')
                }}
              >返回登录</Button></div>
              <div style={{ color: '#8a8c97' }}><Button type='text'
                onClick={() => {
                  navigate('/home')
                }}
              >回到主页</Button></div>
            </div>
          </div>
          <div className='Register-left-content'>
            <div style={{ marginBottom: '10px' }}><span style={{ fontSize: 24, fontWeight: 'bold' }}>开始</span>填写你的信息</div>
            <Form
              name='basic'
              labelCol={{ span: 8 }}
              layout='vertical'
              wrapperCol={{ span: 16 }}
              style={{ maxWidth: 600 }}
              //   initialValues={{ remember: true }}表单初始值
              onFinish={onFinish}
              onFinishFailed={onFinishFailed}
              autoComplete='off'
            >
              <Form.Item<FieldType>
                label='Username'
                name='username'
                rules={[{ required: true, message: 'Please input your username!' }]}
              >
                <Input />
              </Form.Item>
              <Form.Item<FieldType>
                label='Password'
                name='password'
                rules={[{ required: true, message: 'Please input your password!' }]}
              >
                <Input.Password />
              </Form.Item>
              <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                <Button type='primary' htmlType='submit'>
        Submit
                </Button>
              </Form.Item>
            </Form>
          </div>
        </div>
        <div style={{ width: '30vw', height: '80vh', backgroundColor: 'blue' }}>将来页面展示图片</div>
      </div>
    </div>
  )
}
