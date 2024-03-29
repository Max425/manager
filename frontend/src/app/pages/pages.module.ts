import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ChatsComponent} from "./chats/chats.component";
import {StatisticsComponent} from "./statistics/statistics.component";
import {LoggingComponent} from "./logging/logging.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {CoreModule} from "../core/core.module";
import {HomeComponent} from './home/home.component';
import {PageRoutingModule} from "./pages-routing.module";
import {AdminComponent} from "./admin/admin.component";
import {NavigationModule} from "../navigation/navigation.module";
import {MaterialProxyModule} from "../material-proxy/material-proxy.module";
import {MAT_DATE_LOCALE} from "@angular/material/core";
import {ChatTreeComponent} from './chats/components/chat-tree/chat-tree.component';
import {TreeModule} from "primeng/tree";
import {ButtonModule} from "primeng/button";
import {InputTextModule} from "primeng/inputtext";
import {CardModule} from "primeng/card";
import {InputTextareaModule} from "primeng/inputtextarea";
import {TreeSelectModule} from "primeng/treeselect";
import {CardComponent} from "./admin/components/card/card.component";
import {WoodpeckerButtonComponent} from "./admin/components/woodpecker-button/woodpecker-button.component";
import {PanelModule} from "primeng/panel";
import {PaginatorModule} from "primeng/paginator";
import {CalendarModule} from "primeng/calendar";
import {CheckboxModule} from "primeng/checkbox";
import {EmployeeComponent} from "./employee/employee.component";


@NgModule({
  declarations: [
    AdminComponent,
    EmployeeComponent,
    ChatsComponent,
    StatisticsComponent,
    LoggingComponent,
    HomeComponent,
    ChatTreeComponent,
    CardComponent,
    WoodpeckerButtonComponent
  ],
  exports: [
    ChatsComponent
  ],
  imports: [
    CommonModule,
    CoreModule,
    PageRoutingModule,
    NavigationModule,
    MaterialProxyModule,
    ReactiveFormsModule,
    TreeModule,
    ButtonModule,
    InputTextModule,
    FormsModule,
    CardModule,
    InputTextareaModule,
    TreeSelectModule,
    PanelModule,
    PaginatorModule,
    CalendarModule,
    CheckboxModule
  ],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: "ru-RU"},
  ]
})
export class PagesModule {
}
