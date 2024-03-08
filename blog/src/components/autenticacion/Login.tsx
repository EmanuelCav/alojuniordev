'use client'

import { useForm } from "react-hook-form";
import { yupResolver } from '@hookform/resolvers/yup';

import { LoginPropsType } from "@/types/props.types";

import Input from "@/components/general/Input"
import Sumbit from "@/components/general/Sumbit"

import { registerValidationSchema } from "@/validation/register.valid";

import { loginAction } from "@/server/actions/user.actions";

const Login = ({ login, router }: LoginPropsType) => {

    const { register, handleSubmit, formState: { errors } } = useForm({
        resolver: yupResolver(registerValidationSchema)
    });

    return (
        <div onSubmit={handleSubmit((data) => loginAction(data, login, router))} className="w-1/2 h-full justify-center flex items-center">
            <form className='shadow border-2 border-emerald-300 shadow-slate-300 w-1/2 rounded px-8 py-6'>
                <Input text="email" type="text" autofocus={true} autocomplete={true} register={register} errors={errors.email} />
                <Input text="password" type="password" autofocus={false} autocomplete={false} register={register} errors={errors.password} />
                <p className='mt-4 text-slate-950 text-md cursor-pointer hover:underline active:no-underline'>¿Olvidaste tu contraseña?</p>
                <Sumbit text="Iniciar sesión" />
            </form>
        </div>
    )
}

export default Login