import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {ApiService} from "../../../core/services/api.service";
import {Company} from "../../../shared/models/entity/company";
import {Project} from "../../../shared/models/entity/project";

@Injectable({
  providedIn: 'root'
})
export class CreateProjectService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }

  getCompany(): Observable<Company> {
    const url = `/api/companies`;
    return this.http.get<Company>(url, {});
  }

  createProject(project: Project): Observable<Project> {
    const url = `/api/projects`;
    return this.http.post<Project>(url, project);
  }
  // chooseEmployees(roles: string[]): Observable<Employee[]> {
  //   const url = `/api/projects`;
  //   return this.http.post<Employee[]>(url, roles);
  // }
}
