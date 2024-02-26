import TitleTutorialSection from './components/TutorialSection/TitleTutorialSection'
import ShowTutorials from './components/TutorialSection/ShowTutorials'

import { TutorialSectionPropsType } from '@/types/props.types'

const TutorialSection = ({ tutorialSection, title }: TutorialSectionPropsType) => {
    return (
        <div className='w-full'>
            <TitleTutorialSection title={title} />
            <ShowTutorials tutorialSection={tutorialSection} />
        </div>
    )
}

export default TutorialSection