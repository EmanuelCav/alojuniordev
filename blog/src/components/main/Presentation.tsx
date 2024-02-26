import MainTitle from './components/presentation/MainTitle'
import SectionMain from './components/sections/SectionMain'

import image from '../../../public/developer.png'

const Presentation = () => {
    return (
        <div className='h-screen w-full flex flex-col justify-center items-center'>
            <MainTitle />
            <SectionMain
                image={image}
                text='fdsfds dsfds fdsfds'
                isLeft={false}
            />
        </div>
    )
}

export default Presentation