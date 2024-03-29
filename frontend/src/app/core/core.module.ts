import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {FormsModule} from "@angular/forms";
import {RouterModule} from "@angular/router";
import {ProjectCardComponent} from "./components/project-card/project-card.component";
import {EmployeeCardComponent} from "./components/employee-card/employee-card.component";
import {ManagerButtonComponent} from "./components/manager-button/manager-button.component";


@NgModule({
  declarations: [
    ProjectCardComponent,
    EmployeeCardComponent,
    ManagerButtonComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    RouterModule,
  ],
  exports: [
    ProjectCardComponent,
    EmployeeCardComponent,
    ManagerButtonComponent,
  ]
})
export class CoreModule { }
