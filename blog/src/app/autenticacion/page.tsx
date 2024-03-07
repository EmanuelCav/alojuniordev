import Login from '@/components/autenticacion/Login'
import Side from '@/components/autenticacion/Side'

const Auth = () => {

  return (
    <div className='w-full flex items-center justify-center flex-row height-screen'>
      <Login />
      <Side />
    </div>
  )
}

export default Auth