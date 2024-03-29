import {HTTP_INTERCEPTORS} from "@angular/common/http";
import {JWTTokenInterceptor} from "../../auth/interceptors/jwttoken.interceptor";
import {ApiOriginInterceptor} from "../../core/interceptors/api-origin.interceptor";

export const interceptorProviders = [{
  provide: HTTP_INTERCEPTORS,
  useClass: JWTTokenInterceptor,
  multi: true
},
  { provide: HTTP_INTERCEPTORS, useClass: ApiOriginInterceptor, multi: true },]
