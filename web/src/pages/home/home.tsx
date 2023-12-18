import { Button } from 'antd'
import React from 'react'
import { adminLoginApi } from '@/request/api'
// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function Home () {
  return (
    <div>
      Home
      <Button type="primary" onClick={() => {
        console.log('点击了按钮')
        void adminLoginApi({
          data: {
            username: 'rice',
            password: '123456'
          }
        }).then((res: any) => {
          console.log(res)
        })
      }}>按钮</Button>
    </div>
  )
}
