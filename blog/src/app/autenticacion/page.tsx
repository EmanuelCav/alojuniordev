'use client'

import { useState } from "react";

import Forms from '@/components/autenticacion/Forms'
import Side from '@/components/autenticacion/Side'

const Auth = () => {

  const [isRegister, setIsRegister] = useState<boolean>(false)

  const handleViewForm = () => {
    setIsRegister(!isRegister)
  }

  return (
    <div className='w-full flex items-center justify-center flex-row height-screen'>
      <Forms isRegister={isRegister} handleViewForm={handleViewForm} />
      {
        !isRegister && <Side />
      }
    </div>
  )
}

export default Auth