import { Component, Input } from '@angular/core';
import {Employee} from "../../../shared/models/entity/employee";

@Component({
  selector: 'app-employee-card',
  templateUrl: './employee-card.component.html',
  styleUrls: ['./employee-card.component.scss']
})
export class EmployeeCardComponent {
  @Input() employee!: Employee;
}
