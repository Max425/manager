import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {firstValueFrom} from "rxjs";
import {Employee} from "../../shared/models/entity/employee";
import {EmployeeService} from "./services/employee.service";
import {Project} from "../../shared/models/entity/project";

@Component({
  selector: 'app-employee',
  templateUrl: './employee.component.html',
  styleUrls: ['./employee.component.scss']
})
export class EmployeeComponent implements OnInit {
  public employee!: Employee;
  public projects!: Project[];

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: EmployeeService
  ) {
  }

  async ngOnInit(): Promise<void> {
    await this.getEmployee();
    await this.getEmployeeProjects();
  }

  public async getEmployee() {
    const idParam = this.route.snapshot.paramMap.get('id');
    if (idParam === null) return;
    const id = +idParam;
    this.employee = await firstValueFrom((this.api.getEmployeeById(id)));
  }


  public async getEmployeeProjects() {
    const idParam = this.route.snapshot.paramMap.get('id');
    if (idParam === null) return;
    const id = +idParam;
    this.projects = await firstValueFrom((this.api.getEmployeeProjects(id)));
  }
}
