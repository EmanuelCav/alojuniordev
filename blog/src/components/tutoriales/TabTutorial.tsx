'use client'

import { useState } from "react";
import { MdOutlineKeyboardArrowRight, MdOutlineKeyboardArrowDown } from "react-icons/md";

import ContentTab from "./components/tabTutorial/ContentTab";

const TabTutorial = ({ name }: { name: string }) => {

    const [isClicked, setIsClicked] = useState<boolean>(false)

    const handleTab = () => {
        setIsClicked(!isClicked)
    }

    return (
        <div className="w-full">
            <div className="cursor-pointer border-2 border-emerald-200 border-solid mt-8 py-8 px-16 items-center justify-between flex-row flex shadow hover:shadow-md hover:shadow-emerald-200 active:shadow-none"
                onClick={handleTab}>
                <p className="text-2xl select-none text-slate-950 font-medium">{name}</p>
                {
                    isClicked ? <MdOutlineKeyboardArrowDown size={24} className="text-slate-950" />
                        : <MdOutlineKeyboardArrowRight size={24} className="text-slate-950" />
                }
            </div>
            {
                isClicked && <ContentTab />
            }
        </div>
    )
}

export default TabTutorial