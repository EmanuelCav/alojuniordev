'use client'

import { useState } from 'react'
import { IoMdClose } from "react-icons/io";

import Input from "../general/Input";
import Sumbit from "../general/Sumbit";
import Check from "./components/forms/Check";
import InputPassword from "./components/forms/InputPassword";

const Register = ({ handleRegister }: { handleRegister: () => void }) => {

    const [isPassword, setIsPassword] = useState<boolean>(true)
    const [isConfirm, setIsConfirm] = useState<boolean>(true)

    const handleViewPassword = () => {
        setIsPassword(!isPassword)
    }

    const handleViewConfirm = () => {
        setIsConfirm(!isConfirm)
    }

    return (
        <div className="absolute top-0 left-0 z-20 h-screen w-full items-center flex justify-center background-payload ">
            <form className='shadow-md shadow-emerald-300 w-1/4 bg-white rounded px-8 py-6 relative'>
                <IoMdClose onClick={handleRegister} size={'1.5em'} className="cursor-pointer absolute position-icon hover:bg-gray-300 active:bg-white"/>
                <Input text="Username" type="text" autofocus={true} autocomplete={false} />
                <Input text="Email" type="text" autofocus={false} autocomplete={false} />
                <InputPassword name="password" text="Password" isPassword={isPassword} func={handleViewPassword} />
                <InputPassword name="confirm" text="Confirm password" isPassword={isConfirm} func={handleViewConfirm} />
                <Check />
                <Sumbit text="Aceptar" />
            </form>
        </div>
    )
}

export default Register