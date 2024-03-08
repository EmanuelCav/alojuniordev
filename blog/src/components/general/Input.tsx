import { InputLoginTextTypes } from '@/types/key.types'
import { InputPropsType } from '@/types/props.types'

const Input = ({ text, type, autofocus, autocomplete, register, errors }: InputPropsType) => {
    return (
        <div className="my-4">
            <label className="block text-slate-950 text-sm font-bold mb-2">{text}</label>
            {
                errors && <p className='text-red-500 text-sm mb-1'>{errors.message}</p>   
            }
            <input {...register(`${text}` as InputLoginTextTypes)} 
            className="border-emerald-300 appearance-none border rounded w-full py-3 px-3 text-slate-950 leading-tight focus:outline-none focus:shadow-outline focus:shadow focus:shadow-emerald-300" 
            type={type} style={errors ? { border: '1px solid #f00' } : {}} placeholder={text[0].toUpperCase() + text.slice(1, text.length)} autoFocus={autofocus} autoComplete={autocomplete ? 'on' : 'off'} />
        </div>
    )
}

export default Input