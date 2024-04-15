import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-manager-button',
  templateUrl: './manager-button.component.html',
  styleUrls: ['./manager-button.component.scss']
})
export class ManagerButtonComponent {
  @Input() text!: string;
}
