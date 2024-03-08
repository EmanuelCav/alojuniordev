import { StaticImageData } from "next/image";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { FieldError, UseFormRegister } from "react-hook-form";

import { ITutorial } from "@/interface/Content";
import { ILogin, IMainUser } from "@/interface/User";
import { CategoryTypes, InputLoginTextTypes, InputTypes } from "./key.types";

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
    text: InputLoginTextTypes;
    autofocus: boolean;
    autocomplete: boolean;
    register: UseFormRegister<ILogin>;
    errors: FieldError | undefined;
}

export type InputPasswordPropsType = {
    text: string;
    name: string;
    isPassword: boolean;
    func: () => void;
}

export type LoginPropsType = {
    login: (userData: IMainUser) => void;
    router: AppRouterInstance;
}