import { User } from './user';


export interface BaseSession {
    sessionToken: string,
    currentUserId: string,
}

export interface Session extends BaseSession {
    user: User;
}
