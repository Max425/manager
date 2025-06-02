import {Component, OnInit} from '@angular/core';
import {ProjectsService} from "./services/projects.service";
import {BehaviorSubject, firstValueFrom} from "rxjs";
import {Project} from "../../shared/models/entity/project";

@Component({
  selector: 'woodpecker-admin',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.scss']
})
export class ProjectsComponent implements OnInit {
  public filteredProjects: Project[] = [];
  private _treeData = new BehaviorSubject<Project[]>([]);

  constructor(private api: ProjectsService) {
  }

  async ngOnInit(): Promise<void> {
    await this.getProjects();
  }

  public async getProjects() {
    const data = await firstValueFrom(this.api.getProjects());
    // Filter active projects and sort by deadline in descending order
    const activeSortedProjects = data
      .filter(project => project.status === 1)
      .sort((a, b) => new Date(b.deadline).getTime() - new Date(a.deadline).getTime());
    this._treeData.next(activeSortedProjects);
    this.filteredProjects = activeSortedProjects;
  }

  filterProjects(event: any) {
    const searchTerm = event.target.value;
    this.filteredProjects = this._treeData.value.filter(project => {
      return (
        project.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        project.description.toLowerCase().includes(searchTerm.toLowerCase())
      );
    });
  }

  protected readonly console = console;
}
