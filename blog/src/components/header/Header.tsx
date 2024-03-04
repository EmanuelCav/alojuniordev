'use client'

import { usePathname, useRouter } from 'next/navigation'

import Icon from './components/Icon'
import Actions from './components/Actions'
import Search from './components/Search'

const Header = () => {

    const router = useRouter()
    const pathname = usePathname()

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
                pathname !== '/autenticacion' &&
                <>
                    <Search />
                    <Actions redirectAuth={redirectAuth} />
                </>
            }
        </div>
    )
}

export default Header