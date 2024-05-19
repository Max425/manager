import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {ApiService} from "../../../core/services/api.service";
import {Company} from "../../../shared/models/entity/company";
import {Project} from "../../../shared/models/entity/project";
import {Employee} from "../../../shared/models/entity/employee";
import {AutoEmployee} from "../../../shared/models/entity/auto-employee";
import {AutoEmployees} from "../../../shared/models/entity/auto-employees";

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
  autoChooseEmployees(autoEmployees: AutoEmployees): Observable<AutoEmployee[]> {
    const url = `/api/employees/auto`;
    return this.http.post<AutoEmployee[]>(url, autoEmployees);
  }

  public getEmployees(): Observable<Employee[]> {
    const url = "/api/companies/employees"
    return this.http.get<Employee[]>(url, {});
  }
}
