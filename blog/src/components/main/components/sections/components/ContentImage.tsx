import Image, { StaticImageData } from 'next/image'

const ContentImage = ({ image }: { image: StaticImageData }) => {
    return (
        <div className='h-full w-1/2 flex justify-center items-center select-none'>
            <Image src={image} alt='Junior Image' width={312} height={312} />
        </div>
    )
}

export default ContentImage