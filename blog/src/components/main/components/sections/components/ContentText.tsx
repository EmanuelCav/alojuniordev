import React from 'react'

const ContentText = ({ text }: { text: string }) => {
    return (
        <div className='flex-1 h-full flex justify-center items-center p-12'>
            <p className='text-2xl text-slate-950 mt-4'>{text}</p>
        </div>
    )
}

export default ContentText