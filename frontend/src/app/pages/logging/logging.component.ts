import {Component, OnInit, ViewChild} from '@angular/core';
import {MatPaginator} from "@angular/material/paginator";
import {MatTableDataSource} from "@angular/material/table";
import {Log} from "../../shared/models/logging/log";
import {LoggingService} from "./services/logging.service";
import {MatSort} from "@angular/material/sort";

@Component({
  selector: 'woodpecker-logging',
  templateUrl: './logging.component.html',
  styleUrls: ['./logging.component.scss']
})
export class LoggingComponent implements OnInit {

  columnObjects = [
    { columnId: 'timeRequest', propertyName: 'Время обращения' },
    { columnId: 'userId', propertyName: 'Telegram Id' },
    { columnId: 'userName', propertyName: 'Имя' },
    { columnId: 'command', propertyName: 'Сообщение' },
    { columnId: 'message', propertyName: 'Сообщение' },
    { columnId: 'answer', propertyName: 'Ответ' }
  ];

  columnIds = this.columnObjects.map(c => c.columnId);
  dataSource!: MatTableDataSource<Log[]>;
  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;
  constructor(private api: LoggingService) {
  }

  ngOnInit(): void {
    this.api.getLogging().subscribe(data => {
      if(data.status == 200) {
        let logs = data.payload!;
        console.log(logs);
        // @ts-ignore
        this.dataSource = new MatTableDataSource<Log[]>(logs);
        this.dataSource.paginator = this.paginator;
        this.dataSource.sort = this.sort;
      } else {
      console.log(`status: ${data.status}\nmessage: ${data.message}`);
      }
    })
  }
}
