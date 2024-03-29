import {Component, ElementRef, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Chat, ChatQuestion} from "../../../../shared/models/chat/chat";
import {FlatTreeControl, NestedTreeControl} from "@angular/cdk/tree";
import {MatTreeFlatDataSource, MatTreeFlattener, MatTreeNestedDataSource} from "@angular/material/tree";
import {ChatService} from "../../services/chat.service";
import {Message} from "../../../../shared/models/chat/message";
import {BehaviorSubject, combineLatest, map, startWith} from "rxjs";
import {FormBuilder} from "@angular/forms";

@Component({
  selector: 'chat-tree-node',
  templateUrl: './chat-tree.component.html',
  styleUrls: ['./chat-tree.component.scss']
})

export class ChatTreeComponent {

  treeSearch= this.fb.group(
    {search: this.fb.nonNullable.control("")}
  )
  page = 1;
  count = 0;
  pageSize = 10;

  private _treeData = new BehaviorSubject<Chat[]>([]);

  public treeData$ = this._treeData.asObservable();

  public filterData$ = combineLatest([
    this.treeSearch.controls.search.valueChanges.pipe(startWith('')),
    this.treeData$
  ]).pipe(
    map(([searchString, data]) => {
      console.log(searchString, data);
      return this.filtrateTree(data, searchString);
    })
  )

  @Output('select')
  public readonly selectEmitter = new EventEmitter<Message>();

  @Input()
  public set dataSource(chats: Chat[]){
    this._treeData.next(chats);
  }


  constructor(
    private  readonly fb: FormBuilder,
    ) {

  }

  handlePageChange(event: number) {
    this.page = event;
  }

  convertMessage(chat: Chat, question: ChatQuestion) : Message {
    return {
      userId: chat.userId,
      userName: chat.userName,
      userTName: chat.userTName,
      questionId: question.questionId,
      requestTime: question.requestTime,
      text: question.text,
      answer: undefined
    };
  }

  updateView(chat: Chat){
    chat.isHidden = !chat.isHidden;
  }

  private filtrateTree(items: Chat[], searchText: string) {
    if (!items) return [];
    if (!searchText) return items;

    searchText = searchText.toLowerCase();
    return items.filter((it) => {
      return it.userName.toLowerCase().includes(searchText);
    });
  }
}
