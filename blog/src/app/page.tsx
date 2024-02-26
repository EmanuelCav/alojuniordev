
import Presentation from "@/components/main/Presentation";
import ProjectsMain from "@/components/main/ProjectsMain";
import TutorialsMain from "@/components/main/TutorialsMain";

export default function Home() {
  return (
    <div className="container">
      <Presentation />
      <ProjectsMain />
      <TutorialsMain />
    </div>
  );
}
