'use client'

import Image from 'next/image'
import { useRouter } from 'next/navigation'

import { ITutorial } from '@/interface/Content'

const Tutorial = ({ tutorial }: { tutorial: ITutorial }) => {

    const router = useRouter()

    const redirectTutorial = () => {
        router.push(`/tutoriales/${tutorial.tool}`)
    }

    return (
        <div className="m-6 cursor-pointer select-none hvr-bounce-in flex flex-row items-center justify-center h-96" onClick={redirectTutorial}>
            <div className="h-full w-1/2 flex items-center justify-center">
                <Image src={tutorial.image} alt={`${tutorial.tool}-image`} height={0} width={150} style={{ objectFit: 'contain' }} />
            </div>
            <div className="flex-1 flex items-start justify-start flex-col p-4">
                <h1 className='text-2xl font-medium text-slate-950 mt-4'>{tutorial.tool}</h1>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quidem, eligendi similique obcaecati distinctio repudiandae magnam aliquam sequi placeat doloribus fuga quisquam, molestias id, corporis sit dolor voluptatem repellat alias explicabo.</p>
            </div>
        </div>
    )
}

export default Tutorial