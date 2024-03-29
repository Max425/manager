import { Component, Input } from '@angular/core';
import { Project } from "../../../shared/models/entity/project";

@Component({
  selector: 'app-project-card',
  templateUrl: './project-card.component.html',
  styleUrls: ['./project-card.component.scss']
})
export class ProjectCardComponent {
  @Input() project!: Project;
}
