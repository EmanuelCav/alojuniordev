import Input from "@/components/general/Input"
import Sumbit from "@/components/general/Sumbit"

const Login = () => {
    return (
        <div className="w-1/2 h-full justify-center flex items-center">
            <form className='shadow border-2 border-emerald-300 shadow-slate-300 w-1/2 rounded px-8 py-6'>
                <Input text="Email" type="email" autofocus={true} autocomplete={true} />
                <Input text="Password" type="password" autofocus={false} autocomplete={false} />
                <p className='mt-4 text-slate-950 text-md cursor-pointer hover:underline active:no-underline'>¿Olvidaste tu contraseña?</p>
                <Sumbit text="Iniciar sesión" />
            </form>
        </div>
    )
}

export default Login