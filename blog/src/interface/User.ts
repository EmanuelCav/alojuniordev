export interface IReducerUser {
    user: IMainUser;
    isLoggedIn: boolean;
    login: (userData: IMainUser) => void;
}

export interface IMainUser {
    token?: string;
    user?: IUser;
}

export interface IUser {
    _id: number;
    username: string;
    password: string
    email: string;
    role: string;
    phone: string;
    reputation: number;
    status: boolean;
    createdAt: string;
    updatedAt: string;
}

export interface ILogin {
    email: string;
    password: string;
}