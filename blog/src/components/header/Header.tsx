import React from 'react'

import Icon from './components/Icon'
import Actions from './components/Actions'
import Search from './components/Search'

const Header = () => {
    return (
        <div className='flex flex-row justify-between items-center py-4 px-10 bg-white w-full border-solid border-b-emerald-300 border-b-2 fixed top-0'>
            <Icon />
            <Search />
            <Actions />
        </div>
    )
}

export default Header