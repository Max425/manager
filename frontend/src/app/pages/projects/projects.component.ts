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
    await this.getEmployees();
  }

  public async getEmployees() {
    const data = await firstValueFrom(this.api.getProjects());
    this._treeData.next(data);
    this.filteredProjects = data;
  }

  filterProjects(event: any) {
    const searchTerm = event.target.value;
    this.filteredProjects = this._treeData.value.filter(project => {
      return (
          project.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
          project.description.toLowerCase().includes(searchTerm.toLowerCase())
        // Add more fields as needed
      );
    });
  }

  protected readonly console = console;
}
