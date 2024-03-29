import {Component, Input} from '@angular/core';

@Component({
  selector: 'woodpecker-button',
  templateUrl: './woodpecker-button.component.html',
  styleUrls: ['./woodpecker-button.component.scss']
})
export class WoodpeckerButtonComponent {
  @Input() text!: string;
}
