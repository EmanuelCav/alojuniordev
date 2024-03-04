
const Register = () => {
    return (
        <form className='shadow-md shadow-emerald-300 w-1/2 rounded px-8 py-6'>
            <div className="mb-4">
                <label className="block text-slate-950 text-sm font-bold mb-2">Username</label>
                <input className="shadow-emerald-300 shadow-sm border-emerald-300 appearance-none border rounded w-full py-2 px-3 text-slate-950 leading-tight 
                    focus:outline-none focus:shadow-outline focus:shadow focus:shadow-emerald-300" type="text" placeholder="Username" />
            </div>
        </form>
    )
}

export default Register