'use client'

import { useRouter } from 'next/navigation'

import Login from '@/components/autenticacion/Login'
import Side from '@/components/autenticacion/Side'

import { userStore } from '@/server/store/user.store'

const Auth = () => {

  const router = useRouter()

  const { login } = userStore()

  return (
    <div className='w-full flex items-center justify-center flex-row height-screen'>
      <Login login={login} router={router} />
      <Side />
    </div>
  )
}

export default Auth