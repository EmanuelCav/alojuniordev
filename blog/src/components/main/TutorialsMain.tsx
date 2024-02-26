import TitleSectionMain from './components/sections/TitleSectionMain'
import SectionMain from './components/sections/SectionMain'
import Contents from './components/sections/Contents'

import image from '../../../public/book.png'

const TutorialsMain = () => {
  return (
    <div className='h-screen'>
      <TitleSectionMain title='Tutoriales' />
      <div className='h-5/6 w-full'>
        <SectionMain
          image={image}
          text='fdsfds fdsfds fdsfdsfds'
          isLeft={false}
        />
        <Contents />
      </div>
    </div>
  )
}

export default TutorialsMain