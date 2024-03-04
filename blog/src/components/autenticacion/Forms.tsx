import Login from "./components/forms/Login"
import Register from "./components/forms/Register"

import { FormsPropsType } from "@/types/props.types"

const Forms = ({ isRegister, handleViewForm }: FormsPropsType) => {

    return (
        <div className="h-full w-1/2 items-center justify-center flex">
            {
                isRegister ? <Register /> 
                : <Login />
            }
        </div>
    )
}

export default Forms