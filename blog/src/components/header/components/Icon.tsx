import Image from 'next/image'
import Link from 'next/link'

import image from '../../../../public/alojunior.png'

const Icon = ({ redirectIndex }: { redirectIndex: () => void }) => {

  return (
    <div className='select-none flex justify-start items-center flex-row w-1/3'>
      <Image src={image} alt='a_lo_junior_icon' width={66} height={66} onClick={redirectIndex} className='cursor-pointer'/>
      <Link className='text-2xl ml-2 font-medium text-emerald-400' href='/'>A Lo Junior</Link>
    </div>
  )
}

export default Icon