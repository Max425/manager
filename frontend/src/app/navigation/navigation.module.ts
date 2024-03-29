import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MaterialProxyModule} from "../material-proxy/material-proxy.module";
import {RouterModule} from "@angular/router";
import {FormsModule} from "@angular/forms";
import {AuthModule} from "../auth/auth.module";
import {CoreModule} from "../core/core.module";
import {NavigationBarComponent} from "./navigation-bar/navigation-bar.component";
import {StatusBarComponent} from "./status-bar/status-bar.component";
import {NavigateComponent} from "./navigate/navigate.component";

@NgModule({
  declarations: [
    NavigationBarComponent,
    StatusBarComponent,
    NavigateComponent
  ],
  exports: [
    NavigateComponent
  ],
  imports: [
    CommonModule,
    MaterialProxyModule,
    RouterModule,
    FormsModule,
    AuthModule,
    CoreModule,
  ]
})
export class NavigationModule { }
