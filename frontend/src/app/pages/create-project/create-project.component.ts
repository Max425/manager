import {Component, OnInit} from '@angular/core';
import {Project} from "../../shared/models/entity/project";
import {CreateProjectService} from "./services/create-project.service";
import {firstValueFrom} from "rxjs";
import {Company} from "../../shared/models/entity/company";
import {AutoEmployee} from "../../shared/models/entity/auto-employee";
import {Employee} from "../../shared/models/entity/employee";
import {AutoEmployees} from "../../shared/models/entity/auto-employees";

@Component({
  selector: 'app-create-project',
  templateUrl: './create-project.component.html',
  styleUrls: ['./create-project.component.css']
})
export class CreateProjectComponent implements OnInit {
  project: Project = {
    company_id: 0, current_stage: 1, id: 0, image: "", status: 1,
    complexity: 0,
    deadline: '',
    description: '',
    name: '',
    stages: ['']
  };
  choosePositions: AutoEmployee[] = [{position: '', employee: undefined, pin: false}];
  company!: Company;
  employees: Employee[] = [];

  constructor(private api: CreateProjectService) {
  }

  async ngOnInit(): Promise<void> {
    await this.getCompany();
    this.employees = await firstValueFrom(this.api.getEmployees());
  }

  public async getCompany() {
    this.company = await firstValueFrom((this.api.getCompany()));
  }

  addStage(): void {
    this.project.stages.push('');
  }

  removeStage(index: number): void {
    this.project.stages.splice(index, 1);
  }

  addPositions(): void {
    this.choosePositions.push({position: '', employee: undefined, pin: false});
  }

  removePositions(index: number): void {
    this.choosePositions.splice(index, 1);
  }

  chooseEmployees(index: number) {
    return this.employees
      .filter(e =>
        this.choosePositions[index].position.length < 1
        || e.position === this.choosePositions[index].position);
  }

  autoChooseEmployees() {
    let auto: AutoEmployees = {
      project: this.project,
      auto_employee: this.choosePositions
    };
    this.api.autoChooseEmployees(auto).subscribe(data => {
      this.choosePositions = data;
    })
  }

  createProject(): void {
    console.log(this.project);
    this.api.createProject(this.project).subscribe(data => {
      if (data.status === 200) {
        console.log(`status: ${data.status}\nmessage: ${data}`);
      } else {
        console.log(`status: ${data.status}\nmessage: ${data}`);
      }
    });
  }
}
