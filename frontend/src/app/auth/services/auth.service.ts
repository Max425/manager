import {Injectable} from '@angular/core';
import * as store from 'store';
import {BehaviorSubject, filter, firstValueFrom, fromEvent, map} from "rxjs";
import {Session} from "../../shared/models/auth/session";
import {Login} from "../../shared/models/auth/login-credential";
import {ApiService} from "../../core/services/api.service";
import {User} from "../../shared/models/auth/user";

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private readonly SESSION_STORAGE_KEY = 'currentSession';

    private currentSession = new BehaviorSubject<Session | null>(
        store.get(this.SESSION_STORAGE_KEY) ?? null
    );

    private sessionChannel = new BroadcastChannel('auth');

    public onSessionRenew = fromEvent<MessageEvent>(this.sessionChannel, 'message').pipe(
        map(msg => JSON.parse(msg.data))
    );

    public currentSessionObservable = this.currentSession.asObservable();

    constructor(
        private api: ApiService,
    ) {
        this.onSessionRenew.subscribe(session => {

            if (session) {
                store.set(this.SESSION_STORAGE_KEY, session);
                this.currentSession.next(session);

            } else {
                store.remove(this.SESSION_STORAGE_KEY);
                this.currentSession.next(null);
            }
        });
    }

    public async login(login: Login): Promise<Session | undefined> {
        const clientResponse = await firstValueFrom(this.api.getSessionUser(login));
        const session = clientResponse! as Session;
        console.log(session)
        store.set(this.SESSION_STORAGE_KEY, session);
        this.currentSession.next(session);
        this.sessionChannel.postMessage(JSON.stringify(session));

        return session;
    }

    public clearToken() {
        store.remove(this.SESSION_STORAGE_KEY);
        this.currentSession.next(null);
        this.sessionChannel.postMessage(JSON.stringify(null));
    }

    public async logout(): Promise<void> {
        if (this.currentSession.value) {
            try {
                await firstValueFrom(this.api.deleteSession(this.currentSession.value))
            } finally {
                this.clearToken();
                window.location.reload();
            }
        }
    }

    public sessionStarted: User = {
      userName: 'Максим'
};

    public get isAuthed(): boolean {
        return !!this.currentSession.value;
    }

    public get JWTToken(): string | null {
        return this.currentSession.value?.sessionToken ?? null;
    }
}
