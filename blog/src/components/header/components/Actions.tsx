import Link from 'next/link'

const Actions = () => {
  return (
    <div className='flex flex-row justify-center items-center'>
      <Link className='mx-6 text-xl cursor-pointer active:no-underline text-emerald-400 text-header' href='/proyectos'>Proyectos</Link>
      <Link className='mx-6 text-xl cursor-pointer active:no-underline text-emerald-400 text-header' href='/tutoriales'>Tutoriales</Link>
    </div >
  )
}

export default Actions