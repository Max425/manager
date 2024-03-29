import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ApiService} from "../../../core/services/api.service";
import {Observable} from "rxjs";
import {ClientResponse} from "../../../shared/models/client-response";
import {Chat} from "../../../shared/models/chat/chat";
import {Message} from "../../../shared/models/chat/message";

@Injectable({
  providedIn: 'root'
})
export class ChatService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }


  public getChatTree(): Observable<ClientResponse<Chat[]>> {
    const url = "/api/chat/get-tree"
    return this.http.get<ClientResponse<Chat[]>>(url, {});
  }

  public sendMessage(message: Message) : Observable<ClientResponse<Chat[]>> {
    const url = "/api/chat/send-message"
    return this.http.post<ClientResponse<Chat[]>>(url, message);
  }
}
