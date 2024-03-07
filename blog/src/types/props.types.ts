import { StaticImageData } from "next/image";

import { ITutorial } from "@/interface/Content";
import { CategoryTypes, InputTypes } from "./key.types";

export type NavigationTypeProps = {
    text: string;
    archor: string;
}

export type SearchTagTitlePropsType = {
    text: string;
    title: string;
}

export type SectionMainPropsType = {
    text: string;
    image: StaticImageData;
    isLeft: boolean;
}

export type TutorialSectionPropsType = {
    tutorialSection: ITutorial[];
    title: CategoryTypes;
}

export type InputPropsType = {
    type: InputTypes;
    text: string;
    autofocus: boolean;
    autocomplete: boolean;
}

export type InputPasswordPropsType = {
    text: string;
    name: string;
    isPassword: boolean;
    func: () => void;
}