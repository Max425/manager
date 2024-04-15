import {NgModule} from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {HomeComponent} from "./home/home.component";
import {EmployeesComponent} from "./employees/employees.component";
import {EmployeeComponent} from "./employee/employee.component";
import {ProjectsComponent} from "./projects/projects.component";
import {CreateProjectComponent} from "./create-project/create-project.component";

const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
        children: [
            {
                path: 'employees',
                component: EmployeesComponent,
            },
            {
                path: 'employee/:id',
                component: EmployeeComponent,
            },
            {
                path: 'projects',
                component: ProjectsComponent,
            },
            {
                path: 'create-project',
                component: CreateProjectComponent,
            },
        ]
    }
];


@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})

export class PageRoutingModule {
}
