import {Component, OnInit} from '@angular/core';
import {FormBuilder, Validators} from "@angular/forms";
import {AuthService} from "../../services/auth.service";
import {ActivatedRoute, Router} from "@angular/router";
import {Login} from "../../../shared/models/auth/login-credential";

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent implements OnInit {
  private readonly SPACE_IN_INPUT = /^\S*$/;
  public errorText = '';
  public isPasswordHidden = true;
  public authForm = this.fb.group({
    login: ["",
      [
        Validators.required,
        Validators.pattern(this.SPACE_IN_INPUT)
      ]],
    password: ["",
      [
        Validators.required,
        Validators.pattern(this.SPACE_IN_INPUT)
      ]]
  });
  constructor(
      private router: Router,
      private route: ActivatedRoute,
      private auth: AuthService,
      private readonly fb: FormBuilder,) {
    console.log("AuthComponent");

  }

  public async onLogin(): Promise<void> {
    this.authForm.markAllAsTouched();
    if (!this.authForm.valid) {
      console.log('Введите логин и пароль');
      return;
    }

    try {
      console.log(this.authForm.value)
      let login : Login = {
        login: this.authForm.value.login,
        password: this.authForm.value.password}
      const session = await this.auth.login(login);
      console.log(session);
      if (session) {
        await this.redirectOnLogin();
      }
    } catch (e) {
      // @ts-ignore
      if (e.status === 401) {
        this.authForm.setErrors({ incorrect: true });
        this.errorText = 'Вы ввели неверный логин или пароль';
      } else {
        console.log('При авторизации произошла непредвиденная ошибка. Попробуйте еще раз или обратитесь к системному администратору.');
      }
    }
  }

  public async redirectOnLogin(): Promise<void> {
    const redirect = this.route.snapshot.queryParamMap.get('redirect') || '/employees';
    await this.router.navigateByUrl(redirect);
  }

  public onClear(): void {
    this.authForm.reset();
  }

  ngOnInit(): void {
    this.auth.onSessionRenew.subscribe(async session => {
      if (session) {
        await this.redirectOnLogin();
      } else {
        await this.router.navigateByUrl('/auth');
      }
    });
  }

}
