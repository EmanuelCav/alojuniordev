export interface ITutorial {
    tool: string;
    image: string;
    category: string;
    shortDescription: string;
    description: string;
    tabs: ITab[];
}

export interface ITab {
    title: string;
    description: string;
    syntax: string[];
}