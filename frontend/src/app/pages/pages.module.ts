import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {CoreModule} from "../core/core.module";
import {HomeComponent} from './home/home.component';
import {PageRoutingModule} from "./pages-routing.module";
import {EmployeesComponent} from "./employees/employees.component";
import {NavigationModule} from "../navigation/navigation.module";
import {MAT_DATE_LOCALE} from "@angular/material/core";
import {ButtonModule} from "primeng/button";
import {InputTextModule} from "primeng/inputtext";
import {CardModule} from "primeng/card";
import {InputTextareaModule} from "primeng/inputtextarea";
import {CalendarModule} from "primeng/calendar";
import {CheckboxModule} from "primeng/checkbox";
import {ProfileComponent} from "./employee/components/profile/profile.component";
import {EmployeeComponent} from "./employee/employee.component";
import {ProjectsComponent} from "./projects/projects.component";
import {CreateProjectComponent} from './create-project/create-project.component';
import {InputNumberModule} from "primeng/inputnumber";
import {DropdownModule} from 'primeng/dropdown';
import { NgChartsModule } from 'ng2-charts';

@NgModule({
    declarations: [
        ProjectsComponent,
        EmployeesComponent,
        ProfileComponent,
        EmployeeComponent,
        HomeComponent,
        CreateProjectComponent,
    ],
    exports: [],
    imports: [
        CommonModule,
        CoreModule,
        PageRoutingModule,
        NavigationModule,
        ReactiveFormsModule,
        ButtonModule,
        InputTextModule,
        FormsModule,
        CardModule,
        InputTextareaModule,
        CalendarModule,
        CheckboxModule,
        InputNumberModule,
        DropdownModule,
        NgChartsModule,
    ],
    providers: [
        {provide: MAT_DATE_LOCALE, useValue: "ru-RU"},
    ]
})
export class PagesModule {
}
