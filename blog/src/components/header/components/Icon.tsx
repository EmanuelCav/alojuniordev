'use client'

import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

import image from '../../../../public/alojunior.png'

const Icon = () => {

  const router = useRouter()

  const redirectIndex = () => {
    router.push('/')
  }

  return (
    <div className='flex justify-center items-center flex-row cursor-pointer select-none' onClick={redirectIndex}>
      <Image src={image} alt='a_lo_junior_icon' width={66} height={66} />
      <Link className='text-2xl ml-2 font-medium text-emerald-400' href='/'>A Lo Junior</Link>
    </div>
  )
}

export default Icon