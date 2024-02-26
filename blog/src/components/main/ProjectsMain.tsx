import TitleSectionMain from './components/sections/TitleSectionMain'
import SectionMain from './components/sections/SectionMain'
import Contents from './components/sections/Contents'

import image from '../../../public/projects.png'

const ProjectsMain = () => {
  return (
    <div className='h-screen'>
      <TitleSectionMain title='Proyectos' />
      <div className='h-5/6 w-full'>
        <SectionMain
          image={image}
          text='fdsfds fdsfds fdsfdsfds' 
          isLeft={true}
          />
        <Contents />
      </div>
    </div>
  )
}

export default ProjectsMain