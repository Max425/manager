import { Component } from '@angular/core';

@Component({
  selector: 'woodpecker-not-found',
  template: `<h3>Страница не найдена</h3>`
})
export class NotFoundComponent {
  constructor() {
    console.log("NotFoundComponent")
  }
}
