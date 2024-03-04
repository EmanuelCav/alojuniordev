
const Sumbit = ({ text }: { text: string }) => {
  return (
    <div className="mb-2 mt-6">
        <button className="rounded bg-emerald-300 text-white w-full p-2 block text-lg font-bold hover:bg-emerald-200 focus:outline-none focus:shadow-outline active:bg-emerald-300">
            {text}
        </button>
    </div>
  )
}

export default Sumbit