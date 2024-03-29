import {Injectable} from '@angular/core';
import {ApiService} from "../../../core/services/api.service";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {ClientResponse} from "../../../shared/models/client-response";
import {Log} from "../../../shared/models/logging/log";

@Injectable({
  providedIn: 'root'
})
export class LoggingService extends ApiService {

  constructor(
    http: HttpClient,
  ) {
    super(http)
  }

  public getLogging(): Observable<ClientResponse<Log[]>> {
    const url = "/api/log/get";
    return this.http.get<ClientResponse<Log[]>>(url, {});
  }
}
