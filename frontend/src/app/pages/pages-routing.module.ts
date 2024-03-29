import {NgModule} from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {ChatsComponent} from "./chats/chats.component";
import {AuthGuard} from "../auth/guards/auth.guard";
import {LoggingComponent} from "./logging/logging.component";
import {StatisticsComponent} from "./statistics/statistics.component";
import {HomeComponent} from "./home/home.component";
import {AdminComponent} from "./admin/admin.component";
import {EmployeeComponent} from "./employee/employee.component";

const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    children: [
      {
        path: 'chats',
        component: ChatsComponent,
      },
      {
        path: 'employees',
        component: AdminComponent,
      },
      {
        path: 'employee/:id',
        component: EmployeeComponent,
      },
      {
        path: 'logging',
        component: LoggingComponent,
      },
      {
        path: 'statistics',
        component: StatisticsComponent,
      }
    ]
  }
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})

export class PageRoutingModule {}
