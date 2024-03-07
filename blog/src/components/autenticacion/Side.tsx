import Image from 'next/image'

import image from '../../../public/alojunior.png'

const Side = () => {
    return (
        <div className='h-full items-center justify-center flex flex-1 flex-col select-none'>
            <Image src={image} alt='image-icon' width={312} height={312} />
            <h1 className='text-4xl text-emerald-300 mt-4 font-medium'>A Lo Junior</h1>
            <p className='text-2xl mt-4'>Inspira tu futuro en el desarrollo</p>
        </div>
    )
}

export default Side