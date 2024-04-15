import { Component, OnInit } from '@angular/core';
import {Project} from "../../shared/models/entity/project";

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

  constructor() { }

  ngOnInit(): void {
  }

  addStage(): void {
    this.project.stages.push('');
  }

  removeStage(index: number): void {
    this.project.stages.splice(index, 1);
  }

  createProject(): void {
    console.log(this.project);
  }
}
