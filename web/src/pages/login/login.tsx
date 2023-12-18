import '@/pages/login/login.module.scss'
import React from 'react'
import {
  LockOutlined,
  UserOutlined
} from '@ant-design/icons'
import {
  LoginForm,
  ProConfigProvider,
  ProFormText
} from '@ant-design/pro-components'
import { Tabs, theme, Button } from 'antd'
// import type { CSSProperties } from 'react'
import styles from './login.module.scss'
import { useNavigate } from 'react-router-dom'
import { adminLoginApi } from '@/request/api'
import _ from 'lodash'
// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function Login () {
  const navigate = useNavigate()
  const { token } = theme.useToken()

  // 创建防抖函数
  const debouncedLogin = _.debounce(async (values) => {
    console.log(values)
    void adminLoginApi({
      data: values
    }).then((res: any) => {
      console.log(res)
    })
  }, 1000) // 延迟1秒
  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'center', alignContent: 'center', marginTop: '15vh' }}>
        <div className={styles.Logincontent}>
         <div style={{ display: 'flex', flexDirection: 'row-reverse' }}>
          <Button
          type="text"
          style={{
            fontSize: '16px',
            width: 64,
            height: 64
          }}
          onClick={() => {
            navigate('/home', { replace: true })
          }}
        >
          首页
        </Button>
          </div>
          <ProConfigProvider hashed={false}>
            <div style={{ backgroundColor: token.colorBgContainer }}>
              <LoginForm
              onFinish={async (values) => {
                void debouncedLogin(values)
              }}
                // logo="https://pic.imgdb.cn/item/657d3b79c458853aefcfaebf.png"
                title="管理员登录入口"
                // actions={
                //   <Space>
                //     其他登录方式
                //     <AlipayCircleOutlined style={iconStyles} />
                //     <TaobaoCircleOutlined style={iconStyles} />
                //     <WeiboCircleOutlined style={iconStyles} />
                //   </Space>
                // }
              >
                <Tabs centered activeKey={'account'}>
                  <Tabs.TabPane key={'account'} tab={'账号密码登录'} />
                </Tabs>
                {
                  <>
                    <ProFormText
                      name="username"
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
                      name="password"
                      fieldProps={{
                        size: 'large',
                        prefix: <LockOutlined className={'prefixIcon'} />,
                        strengthText:
                          'Password should contain numbers, letters and special characters, at least 8 characters long.'
                        // 提示强度的代码
                        // statusRender: (value) => {
                        //   const getStatus = () => {
                        //     if (value && value.length > 12) {
                        //       return 'ok';
                        //     }
                        //     if (value && value.length > 6) {
                        //       return 'pass';
                        //     }
                        //     return 'poor';
                        //   };
                        //   const status = getStatus();
                        //   if (status === 'pass') {
                        //     return (
                        //       <div style={{ color: token.colorWarning }}>
                        //         强度：中
                        //       </div>
                        //     );
                        //   }
                        //   if (status === 'ok') {
                        //     return (
                        //       <div style={{ color: token.colorSuccess }}>
                        //         强度：强
                        //       </div>
                        //     );
                        //   }
                        //   return (
                        //     <div style={{ color: token.colorError }}>强度：弱</div>
                        //   );
                        // },
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
                }
                <div
                  style={{
                    marginBlockEnd: 24
                  }}
                >
                   <a
                    style={{
                      float: 'left'
                    }}
                  >
                    管理员注册
                  </a>
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
          </ProConfigProvider>
        </div>
      </div>
    </div>
  )
}
