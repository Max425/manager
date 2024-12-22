import { NgModule } from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {NotAuthGuard} from "../auth/guards/not-auth.guard";
import {AuthGuard} from "../auth/guards/auth.guard";
import {NotFoundComponent} from "../core/components/not-found/not-found.component";

const routes: Routes = [
  {
    path: 'auth',
    loadChildren: () => import('../auth/auth.module').then(m => m.AuthModule),
    canActivate: [NotAuthGuard],
    canActivateChild: [NotAuthGuard]
  },
  {
    path: '',
    loadChildren: () => import('../pages/pages.module').then(m => m.PagesModule),
    canActivate: [NotAuthGuard],
    canActivateChild: [NotAuthGuard]
  },
  {
    path: "**", component: NotFoundComponent
  }
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
