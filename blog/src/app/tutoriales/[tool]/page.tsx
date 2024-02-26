import TabTutorial from "@/components/tutoriales/TabTutorial"
import TitleTutorialSection from "@/components/tutoriales/components/TutorialSection/TitleTutorialSection"

const Tool = ({ params }: { params: { tool: string } }) => {
    return (
        <div>
            <TitleTutorialSection title={params.tool} />
            <TabTutorial name="Variables" />
            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae esse, quam dolore repellendus sequi eius neque eum dolorum maxime tempora illum fuga quidem consequuntur quae facere vitae dolores tenetur sed.</p>
        </div>
    )
}

export default Tool