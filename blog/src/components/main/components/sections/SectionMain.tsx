import ContentImage from "./components/ContentImage"
import ContentText from "./components/ContentText"

import { SectionMainPropsType } from "@/types/props.types"

const SectionMain = ({ image, text, isLeft }: SectionMainPropsType) => {
    return (
        <div className='w-full h-1/2'>
            {
                isLeft ? (
                    <div className="w-full h-full flex-row flex justify-center items-center">
                        <ContentText text={text} />
                        <ContentImage image={image} />
                    </div>
                ) : (
                    <div className="w-full h-full flex-row flex justify-center items-center">
                        <ContentImage image={image} />
                        <ContentText text={text} />
                    </div>
                )
            }
        </div>
    )
}

export default SectionMain