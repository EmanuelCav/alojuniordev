import React from 'react'

const TitleSectionMain = ({ title }: { title: string }) => {
    return (
        <div className='w-full h-1/6 items-center flex justify-start px-12 border-solid border-b-emerald-300 border-b-2 select-none'>
            <p className='text-emerald-400 text-4xl font-medium'>{title}</p>
        </div>
    )
}

export default TitleSectionMain