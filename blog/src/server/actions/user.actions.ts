import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

import { ILogin, IMainUser} from "@/interface/User";

import * as userApi from "../api/user.api";

export const loginAction = async (userData: ILogin, login: (userData: IMainUser) => void, router: AppRouterInstance) => {

    const data = await userApi.loginApi(userData)
    login(data)
    router.push('/main')

}