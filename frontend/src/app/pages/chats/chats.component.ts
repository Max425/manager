import {Component, OnInit} from '@angular/core';
import {Chat} from "../../shared/models/chat/chat";
import {Message} from "../../shared/models/chat/message";
import {AuthService} from "../../auth/services/auth.service";
import {Session} from "../../shared/models/auth/session";
import {BehaviorSubject, firstValueFrom} from "rxjs";
import {ChatService} from "./services/chat.service";

@Component({
  selector: 'woodpecker-chats',
  templateUrl: './chats.component.html',
  styleUrls: ['./chats.component.scss']
})

export class ChatsComponent implements OnInit {

  message?: Message = undefined;
  protected session?: Session;
  protected answerDate?: number;
  constructor(
      public readonly auth: AuthService,
      public readonly api: ChatService,
    ) {
  }
  private _treeData = new BehaviorSubject<Chat[]>([])

  public treeDate$ = this._treeData.asObservable();
  async ngOnInit(): Promise<void> {
    this.session = await firstValueFrom(this.auth.sessionStarted);
    await this.chats();
  }

  private async chats(){
    const data = await firstValueFrom(this.api.getChatTree())
    if (data.status == 200) {
      if (data.payload) {
        this._treeData.next(data.payload);
      }
    } else {
      console.log(`status: ${data.status}\nmessage: ${data.message}`);
    }
  }

  selectNode(message: Message){
    if (message.text != undefined){
      this.message = message;
    }
  }

  public async onSubmit(value: string){
    if(value.length! > 0) {
      this.message!.answer = value
      this.answerDate = Date.now()
      const data = await firstValueFrom(this.api.sendMessage(this.message!))
      if (data.status == 200) {
        if (data.payload) {
          this._treeData.next(data.payload);
        }
      } else {
        console.log(`status: ${data.status}\nmessage: ${data.message}`);
      }
      this.message = undefined;
    }
  }

  public clear(){
    this.message = undefined;
  }

}
