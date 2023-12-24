import React, { useState, useEffect } from 'react'
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  UserOutlined,
  VideoCameraOutlined,
  DownOutlined,
  ApartmentOutlined,
  TeamOutlined
} from '@ant-design/icons'
import { Layout, Menu, Button, theme, Dropdown, MenuProps } from 'antd'
import { Route, Routes, Navigate, useNavigate } from 'react-router-dom'
import GeneralView from '../generalView/generalView'
import './home.scss'
import UserManage from '../userManage/userManage'
import RoleManagement from '../roleManagement/roleManagement'
const { Header, Sider, Content } = Layout
const Backstagehome: React.FC = () => {
  // 侧边栏折叠
  const [collapsed, setCollapsed] = useState(false)
  // 监听窗口大小
  useEffect(() => {
    const handler = () => {
      setCollapsed(window.innerWidth <= 768)
    }
    window.addEventListener('resize', handler)
    return () => window.removeEventListener('resize', handler)
  }, [])

  const {
    token: { colorBgContainer, borderRadiusLG }
  } = theme.useToken()
  const navigate = useNavigate()
  // 下拉菜单
  const items: MenuProps['items'] = [
    {
      key: '1',
      label: (
        <a target='_blank' rel='noopener noreferrer' href='https://www.antgroup.com'>
          系统管理员
        </a>
      )
    },
    {
      key: '4',
      danger: true,
      label: 'a danger item'
    }
  ]
  return (
    <Layout style={{ height: '100vh' }}>
      <Sider trigger={null} collapsible collapsed={collapsed} style={{ display: 'flex', flexDirection: 'column' }}>
        {
          collapsed ? null : <div className='demo-logo-vertical'>管理员系统</div>
        }
        <Menu
          theme='dark'
          mode='inline'
          defaultSelectedKeys={['1']}
          onClick={({ key }) => {
            switch (key) {
              case '1':
                navigate('/backstage/generalview')
                break
              case '2':
                navigate('/largedatascreen')
                break
              case '3':
                navigate('/backstage/usermanage')
                break
              case '4':
                navigate('/backstage/rolemanage')
                break
              default:
                return
            }
          }}
          items={[
            {
              key: '1',
              icon: <ApartmentOutlined />,
              label: '后台总览'
            },
            {
              key: '2',
              icon: <VideoCameraOutlined />,
              label: '数据大屏'
            },
            {
              key: '3',
              icon: <UserOutlined />,
              label: '用户管理'
            },
            {
              key: '4',
              icon: <TeamOutlined />,
              label: '角色管理'
            },
            {
              key: '5',
              icon: <UploadOutlined />,
              label: '角色管理'
            },
            {
              key: '6',
              icon: <UploadOutlined />,
              label: '权限管理'
            }
          ]}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <Button
            type='text'
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
              float: 'left'
            }}
          />
          {/* 返回首页按钮 */}
          <Button
            type='text'
            onClick={() => navigate('/home')}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
              float: 'left'
            }}>首页</Button>
          <Dropdown menu={{ items }} className='welcome-message'>
            <div >
              <div>欢迎admin <DownOutlined />
              </div>
              <div></div>
            </div>
          </Dropdown>

        </Header>
        <Content
          style={{
            margin: '24px 16px',
            padding: 24,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG
          }}
        >
          <Routes>
            <Route path='' element={<Navigate to='/backstage/generalview' />}></Route>
            <Route path='generalview' element={<GeneralView></GeneralView>}></Route>
            <Route path='usermanage' element={<UserManage></UserManage>}></Route>
            <Route path='rolemanage' element={<RoleManagement></RoleManagement>}></Route>
          </Routes>
        </Content>
      </Layout>
    </Layout>
  )
}

export default Backstagehome
