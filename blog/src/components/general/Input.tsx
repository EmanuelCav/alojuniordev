import { InputPropsType } from '@/types/props.types'

const Input = ({ text, type, autofocus }: InputPropsType) => {
    return (
        <div className="mb-4">
            <label className="block text-slate-950 text-sm font-bold mb-2">{text}</label>
            <input className="border-emerald-300 appearance-none border rounded w-full py-3 px-3 text-slate-950 leading-tight 
            focus:outline-none focus:shadow-outline focus:shadow focus:shadow-emerald-300" type={type} placeholder={text} autoFocus={autofocus} />
        </div>
    )
}

export default Input