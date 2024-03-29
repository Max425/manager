import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {ApiService} from "../../../core/services/api.service";
import {Employee} from "../../../shared/models/entity/employee";

@Injectable({
  providedIn: 'root'
})
export class EmployeesService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }

  public getEmployees(): Observable<Employee[]> {
    const url = "/api/companies/employees"
    return this.http.get<Employee[]>(url, {});
  }

}
