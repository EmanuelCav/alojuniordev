import Input from "@/components/general/Input"
import Sumbit from "@/components/general/Sumbit"

const Login = () => {
    return (
        <form className='shadow border-2 border-emerald-300 shadow-slate-300 w-1/2 rounded px-8 py-6'>
            <Input text="Username" type="text" autofocus={true} />
            <Input text="Password" type="password" autofocus={false} />
            <Sumbit text="Iniciar sesión" />
        </form>
    )
}

export default Login