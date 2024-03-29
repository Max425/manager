import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {Login} from "../../shared/models/auth/login-credential";
import {Session} from "../../shared/models/auth/session";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private _url = `/api/auth`
  constructor(
    protected readonly http: HttpClient,
    ) { }

  public getSessionUser(login: Login): Observable<Session> {
    const url = `${this._url}/login`;

    return this.http.post<Session>(url, login);
  }

  public deleteSession(session: Session): Observable<Session> {
    const url = `${this._url}/logout`;
    const headers = new HttpHeaders({
      Authorization: `Bearer ${session.sessionToken}`,
      UserId: `${session.currentUserId}`
    });
    return this.http.delete<Session>(url, { headers });
  }

}
