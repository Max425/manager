import {Component, OnInit} from '@angular/core';
import {Project} from "../../shared/models/entity/project";
import {CreateProjectService} from "./services/create-project.service";
import {firstValueFrom} from "rxjs";
import {Company} from "../../shared/models/entity/company";

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
  choosePositions: string[] = [''];
  company!: Company;

  constructor(private api: CreateProjectService) {
  }

  async ngOnInit(): Promise<void> {
    await this.getCompany();
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
    this.choosePositions.push('');
  }

  removePositions(index: number): void {
    this.choosePositions.splice(index, 1);
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
