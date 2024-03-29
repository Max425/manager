import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule} from "@angular/router";
import {FormsModule} from "@angular/forms";
import {AuthModule} from "../auth/auth.module";
import {CoreModule} from "../core/core.module";
import {NavigationBarComponent} from "./navigation-bar/navigation-bar.component";
import {StatusBarComponent} from "./status-bar/status-bar.component";
import {NavigateComponent} from "./navigate/navigate.component";
import {MatMenuModule} from "@angular/material/menu";
import {MatIconModule} from '@angular/material/icon';

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
        RouterModule,
        FormsModule,
        AuthModule,
        CoreModule,
        MatIconModule,
        MatMenuModule,
    ]
})
export class NavigationModule {
}
