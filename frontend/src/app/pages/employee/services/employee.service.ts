import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {ApiService} from "../../../core/services/api.service";
import {Employee} from "../../../shared/models/entity/employee";
import {Project} from "../../../shared/models/entity/project";

@Injectable({
  providedIn: 'root'
})
export class EmployeeService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }

  getEmployeeById(id: number): Observable<Employee> {
    const url = `/api/employees/${id}`;
    return this.http.get<Employee>(url, {});
  }

  getEmployeeProjects(id: number): Observable<Project[]> {
    const url = `/api/employees/${id}/projects`;
    return this.http.get<Project[]>(url, {});
  }

  updateEmployee(employee: Employee): Observable<Employee> {
    const url = `/api/employees`;
    return this.http.put<Employee>(url, employee);
  }

  deleteEmployee(id: number): Observable<string> {
    const url = `/api/employees/${id}`;
    return this.http.delete<string>(url, {});
  }

}
