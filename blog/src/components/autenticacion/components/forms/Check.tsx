import React from 'react'

const Check = () => {
  return (
    <div className='w-full inline-flex mt-4'>
        <input type="checkbox" name="status" className='cursor-pointer' />
        <p className='text-slate-950 text-md ml-3 cursor-pointer hover:underline active:no-underline'>Aceptar términos y condiciones</p>
    </div>
  )
}

export default Check