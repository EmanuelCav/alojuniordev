import { AiFillEye, AiFillEyeInvisible  } from "react-icons/ai";

import { InputPasswordPropsType } from '@/types/props.types'

const InputPassword = ({ text, name, isPassword, func }: InputPasswordPropsType) => {
  return (
    <div className="my-4">
      <label className="block text-slate-950 text-sm font-bold mb-2">{text}</label>
      <div className="relative justify-center items-center flex">
        {
          isPassword ? <AiFillEye size={'1.5em'} className="cursor-pointer absolute text-emerald-300 right" style={{ left: '92%' }} onClick={func} />
          : <AiFillEyeInvisible size={'1.5em'} className="cursor-pointer absolute text-emerald-300" style={{ left: '92%' }} onClick={func} />
        }
        <input name={name} className="border-emerald-300 appearance-none border rounded w-full py-3 px-3 text-slate-950 leading-tight 
              focus:outline-none focus:shadow-outline focus:shadow focus:shadow-emerald-300" type={isPassword ? 'password': 'text'} placeholder={text} />

      </div>
    </div>
  )
}

export default InputPassword