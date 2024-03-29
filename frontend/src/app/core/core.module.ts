import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {InputComponent} from "./components/input/input.component";
import {FormsModule} from "@angular/forms";
import {RouterModule} from "@angular/router";
import {MaterialProxyModule} from "../material-proxy/material-proxy.module";


@NgModule({
  declarations: [
    InputComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    RouterModule,
    MaterialProxyModule
  ],
  exports: [
    InputComponent
  ]
})
export class CoreModule { }
