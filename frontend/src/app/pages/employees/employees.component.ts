import {Component, OnInit} from '@angular/core';
import {EmployeesService} from "./services/employees.service";
import {BehaviorSubject, firstValueFrom} from "rxjs";
import {Employee} from "../../shared/models/entity/employee";

@Component({
  selector: 'woodpecker-admin',
  templateUrl: './employees.component.html',
  styleUrls: ['./employees.component.scss']
})
export class EmployeesComponent implements OnInit {
  public filteredEmployees: Employee[] = [];
  private _treeData = new BehaviorSubject<Employee[]>([]);

  constructor(private api: EmployeesService) {
  }

  async ngOnInit(): Promise<void> {
    await this.getEmployees();
  }

  public async getEmployees() {
    const data = await firstValueFrom(this.api.getEmployees());
    this._treeData.next(data);
    this.filteredEmployees = data;
  }

  filterEmployees(event: any) {
    const searchTerm = event.target.value;
    this.filteredEmployees = this._treeData.value.filter(employee => {
      return (
        employee.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        employee.position.toLowerCase().includes(searchTerm.toLowerCase()) ||
        // Add more fields as needed
        employee.mail.toLowerCase().includes(searchTerm.toLowerCase())
      );
    });
  }

    protected readonly console = console;
}
