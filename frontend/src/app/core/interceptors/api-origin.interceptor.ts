import { Injectable } from '@angular/core';
import {
    HttpRequest,
    HttpHandler,
    HttpInterceptor,
    HttpErrorResponse
} from '@angular/common/http';
import { throwError } from 'rxjs';
import { AuthService } from '../../auth/services/auth.service';
import { catchError } from 'rxjs/operators';
import {environment} from "../../../environments/environment";


@Injectable()
export class ApiOriginInterceptor implements HttpInterceptor {
    constructor(
        private readonly auth: AuthService,
    ) { }

    public intercept(request: HttpRequest<unknown>, next: HttpHandler) {
      if(environment.api){
        let url = `${environment.api}${request.url}`
        console.log(url);
        request = request.clone({url});
      }
      return next.handle(request).pipe(
          catchError((error: HttpErrorResponse) => {
              if (error.status === 401 && this.auth.JWTToken) {
                  this.auth.clearToken();
              }
              return throwError(error);
          })
      );
    }
}
