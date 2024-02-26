import React from 'react'

import { SearchTagTitlePropsType } from '@/types/props.types'

const SearchTagTitle = ({ title, text }: SearchTagTitlePropsType) => {
  return (
    <div className='bg-emerald-300 p-6'>
        <h1 className='text-2xl text-white font-medium'>{title}</h1>
        <p className='text-2xl text-slate-950 mt-4'>{text}</p>
    </div>
  )
}

export default SearchTagTitle