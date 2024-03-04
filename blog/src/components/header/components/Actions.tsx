import Link from 'next/link'

const Actions = ({ redirectAuth }: { redirectAuth: () => void }) => {
  return (
    <div className='flex flex-row justify-center items-center w-1/3'>
      <Link className='mx-6 text-xl cursor-pointer active:no-underline text-emerald-400 text-header' href='/proyectos'>Proyectos</Link>
      <Link className='mx-6 text-xl cursor-pointer active:no-underline text-emerald-400 text-header' href='/tutoriales'>Tutoriales</Link>
      <Link className='mx-6 text-xl cursor-pointer active:no-underline text-emerald-400 text-header' href='/tutoriales'>Usuario</Link>
      <button className='mx-6 text-xl cursor-pointer bg-transparent hover:bg-emerald-400 text-emerald-400 font-medium hover:text-white py-2 px-4 border border-emerald-400 hover:border-transparent rounded active:bg-transparent'
      onClick={redirectAuth}>
        Acceder
      </button>
    </div >
  )
}

export default Actions