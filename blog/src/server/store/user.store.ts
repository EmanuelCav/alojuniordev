import { create } from "zustand";

import { IReducerUser, IMainUser } from "@/interface/User";

export const userStore = create<IReducerUser>()((set) => ({
    user: {},
    isLoggedIn: false,
    login: (userData: IMainUser) => set(() => ({
        user: userData,
        isLoggedIn: true
    }))
}))