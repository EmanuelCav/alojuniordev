import tutorials from '../../../public/tutorials.json'

import TutorialSection from '@/components/tutoriales/TutorialSection'

import { ITutorial } from '@/interface/Content'

const Tutoriales = () => {
  return (
    <div className='container'>
      <div className="flex justify-center items-center flex-col w-full">
        <TutorialSection tutorialSection={tutorials.filter(t => t.category === "Lenguages de Programación") as ITutorial[]} title='Lenguages de Programación' />
        <TutorialSection tutorialSection={tutorials.filter(t => t.category === "Web Framework") as ITutorial[]} title='Frameworks Web' />
        <TutorialSection tutorialSection={tutorials.filter(t => t.category === "Data Sciense") as ITutorial[]} title='Analísis de Datos' />
      </div>
    </div>
  )
}

export default Tutoriales