import { api } from "./api"

import { ILogin, IMainUser } from "@/interface/User"

export const loginApi = async (userData: ILogin): Promise<IMainUser> => {

    const response = await fetch(api + "/users/login", {
        method: 'POST',
        body: JSON.stringify(userData),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    })

    const data = await response.json()

    return data

}