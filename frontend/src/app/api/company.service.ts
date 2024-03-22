import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {Company} from "../models/company.model";

@Injectable({
  providedIn: 'root'
})
export class CompanyService {

  constructor(private http: HttpClient) { }

  createCompany(company: Company): Observable<Company> {
    const url = 'http://localhost:8000/api/companies';
    return this.http.post<Company>(url, company);
  }

  getCompanyById(id: number): Observable<Company> {
    const url = `http://localhost:8000/api/companies/${id}`;
    return this.http.get<Company>(url);
  }

  updateCompany(id: number, company: Company): Observable<Company> {
    const url = `http://localhost:8000/api/companies/${id}`;
    return this.http.put<Company>(url, company);
  }

  deleteCompany(id: number): Observable<string> {
    const url = `http://localhost:8000/api/companies/${id}`;
    return this.http.delete<string>(url);
  }
}
