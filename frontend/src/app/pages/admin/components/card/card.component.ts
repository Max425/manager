import { Component, Input } from '@angular/core';
import {Employee} from "../../../../shared/models/entity/employee";

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss']
})
export class CardComponent {
  @Input() employee!: Employee;
}
