/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable no-unused-vars */
import React, { useState, startTransition, useEffect } from 'react'
import './login.scss'
import {
  AlipayCircleOutlined,
  LockOutlined,
  MobileOutlined,
  TaobaoCircleOutlined,
  UserOutlined,
  WeiboCircleOutlined
} from '@ant-design/icons'
import {
  LoginForm,
  ProConfigProvider,
  ProFormCaptcha,
  ProFormCheckbox,
  ProFormText,
  setAlpha
} from '@ant-design/pro-components'
import { Space, Spin, Tabs, message, theme } from 'antd'
import _ from 'lodash'
import type { CSSProperties } from 'react'
import { adminLoginApi } from '../../request/api'
import { notification } from 'antd'
import { useNavigate } from 'react-router-dom'
const iconStyles: CSSProperties = {
  color: 'rgba(0, 0, 0, 0.2)',
  fontSize: '18px',
  verticalAlign: 'middle',
  cursor: 'pointer'
}
export default function Backstagelogin() {
  const navigate = useNavigate()
  const debouncedLogin = _.debounce(async(values: never) => {
    void adminLoginApi({
      data: values
    })
      .then((res: any) => {
        console.log(res)
        notification.success({
          message: '登录成功',
          description: '欢迎回来'
        })
        localStorage.setItem('token', res.data.token)
        startTransition(() => {
          navigate('/backstage')
        })
      })
      .catch((err: any) => {
        console.log(err)
        notification.error({
          message: '密码或用户名错误',
          description: '请重新输入'
        })
      })
  }, 1000)
  // 实现图片加载显示加载动画
  const [isImageLoaded, setIsImageLoaded] = useState(false)
  const img = new Image()
  useEffect(() => {
    img.src = 'https://pic.imgdb.cn/item/65879e3fc458853aef740537.jpg'
    img.onload = () => setIsImageLoaded(true)
  }, [])
  return (
    <div style={{ backgroundColor: '#eae8ea', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
      <Spin spinning={!isImageLoaded}>
        <div style={{ boxShadow: '0 0 8px rgba(0, 0, 0, 0.5)', height: '80vh', width: '80vw', borderRadius: '20px', display: 'flex' }} >
          <div className='Backstageloginleft' ></div>
          <div className='Backstageloginright'>

            <LoginForm
              title='管理员登录'
              subTitle='全球最大的商城平台'
              onFinish={(values:never) => {
                return Promise.resolve(debouncedLogin(values))
              }}
              actions={
                <Space>
              其他登录方式
                  <AlipayCircleOutlined style={iconStyles} />
                  <TaobaoCircleOutlined style={iconStyles} />
                  <WeiboCircleOutlined style={iconStyles} />
                </Space>
              }
            >
              {true && (
                <>
                  <ProFormText
                    name='username'
                    fieldProps={{
                      size: 'large',
                      prefix: <UserOutlined className={'prefixIcon'} />
                    }}
                    placeholder={'请输入用户名'}
                    rules={[
                      {
                        required: true,
                        message: '请输入用户名!'
                      }
                    ]}
                  />
                  <ProFormText.Password
                    name='password'
                    fieldProps={{
                      size: 'large',
                      prefix: <LockOutlined className={'prefixIcon'} />,
                      strengthText:
                    'Password should contain numbers, letters and special characters, at least 8 characters long.'
                    }}
                    placeholder={'请输入密码'}
                    rules={[
                      {
                        required: true,
                        message: '请输入密码！'
                      }
                    ]}
                  />
                </>
              )}
              <div
                style={{
                  marginBlockEnd: 24
                }}
              >
                <ProFormCheckbox noStyle name='autoLogin'>
              自动登录
                </ProFormCheckbox>
                <a
                  style={{
                    float: 'right'
                  }}
                >
              忘记密码
                </a>
              </div>
            </LoginForm>
          </div>
        </div>
      </Spin>
    </div>
  )
}
