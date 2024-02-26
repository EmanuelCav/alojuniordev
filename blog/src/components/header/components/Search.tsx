import { HiOutlineSearch } from "react-icons/hi";

const Search = () => {
    return (
        <form className="max-w-md mx-auto">
            <label className="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
            <div className="relative">
                <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                    <HiOutlineSearch size={24} className="text-emerald-200" />
                </div>
                <input type="search" className="border border-solid border-emerald-200 shadow appearance-none block w-full py-2 px-3 ps-10 text-slate-950 rounded bg-gray-50 leading-tight focus:outline-none focus:ring-emerald-200 focus:shadow-emerald-200"
                    placeholder="Busca algua herramienta" autoComplete="off" />
            </div>
        </form>
    )
}

export default Search