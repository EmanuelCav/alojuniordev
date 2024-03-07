'use client'

import { usePathname, useRouter } from 'next/navigation'
import { useState } from 'react'

import Icon from './components/Icon'
import Actions from './components/Actions'
import Search from './components/Search'
import Register from '../autenticacion/Register'

const Header = () => {

    const router = useRouter()
    const pathname = usePathname()

    const [isRegister, setIsRegister] = useState<boolean>(false)

    const handleRegister = () => {
        setIsRegister(!isRegister)
    }

    const redirectIndex = () => {
        router.push('/')
    }

    const redirectAuth = () => {
        router.push('/autenticacion')
    }

    return (
        <div className='flex flex-row justify-between items-center py-4 px-10 bg-white w-full border-solid border-b-emerald-300 border-b-2 fixed top-0 shadow shadow-emerald-300'>
            <Icon redirectIndex={redirectIndex} />
            {
                isRegister && <Register handleRegister={handleRegister} />
            }
            {
                pathname === '/autenticacion' ? (
                    <p className='text-lg text-slate-950'>
                        ¿No tienes cuenta aún?
                        <span className='text-xl ml-2 text-emerald-300 text-header cursor-pointer active:no-underline' onClick={handleRegister}>Registrate</span>
                    </p>
                ) : (
                    <>
                        <Search />
                        <Actions redirectAuth={redirectAuth} />
                    </>
                )
            }
        </div>
    )
}

export default Header