import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AuthComponent } from './components/auth/auth.component';
import { NotAuthGuard } from './guards/not-auth.guard';

const routes: Routes = [
    {
        path: 'sign-in',
        component: AuthComponent,
        canActivate: [NotAuthGuard],
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class AuthRoutingModule { }
