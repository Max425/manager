import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { interceptorProviders } from './interceptors/app-interceptors';
import { AppComponent } from './components/app.component';
import {AppRoutingModule} from "./app-routing.module";
import {HttpClientModule} from "@angular/common/http";
import {CoreModule} from "../core/core.module";
import {PagesModule} from "../pages/pages.module";
import {NotFoundComponent} from "../core/components/not-found/not-found.component";
import {RouterModule} from "@angular/router";
import {AuthModule} from "../auth/auth.module";
import {CommonModule} from "@angular/common";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent,
    NotFoundComponent,
  ],
  imports: [
    AuthModule,
    BrowserModule,
    RouterModule,
    CommonModule,
    CoreModule,
    AppRoutingModule,
    PagesModule,
    HttpClientModule,
    BrowserAnimationsModule,
  ],
  exports: [
  ],
  bootstrap: [AppComponent],
  providers: [
    ...interceptorProviders
  ]
})
export class AppModule { }
