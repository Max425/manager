import {Injectable} from '@angular/core';
import {ApiService} from "../../../core/services/api.service";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Grafana} from "../../../shared/models/statictics/grafana";
import {ClientResponse} from "../../../shared/models/client-response";

@Injectable({
  providedIn: 'root'
})
export class StatisticsService extends ApiService {

  constructor(
    http: HttpClient
  ) {
    super(http)
  }

  public getAddressGrafana(): Observable<ClientResponse<Grafana>> {
    const url = "/api/statistics/address";
    return this.http.get<ClientResponse<Grafana>>(url, {});
  }

  public setFilterGrafana(filterValues: any): Observable<ClientResponse<Grafana>> {
    const url = "/api/statistics/filter";
    return this.http.post<ClientResponse<Grafana>>(url,
      filterValues);
  }
}
