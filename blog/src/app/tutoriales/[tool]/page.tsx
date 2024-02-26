import TabTutorial from "@/components/tutoriales/TabTutorial"
import TitleTutorialSection from "@/components/tutoriales/components/TutorialSection/TitleTutorialSection"

const Tool = ({ params }: { params: { tool: string } }) => {
    return (
        <div>
            <TitleTutorialSection title={params.tool} />
            <TabTutorial name="Variables" />
        </div>
    )
}

export default Tool