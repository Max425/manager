import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {ApiService} from "../../../core/services/api.service";
import {Project} from "../../../shared/models/entity/project";

@Injectable({
  providedIn: 'root'
})
export class ProjectsService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }

  public getProjects(): Observable<Project[]> {
    const url = "/api/companies/projects"
    return this.http.get<Project[]>(url, {});
  }

}
