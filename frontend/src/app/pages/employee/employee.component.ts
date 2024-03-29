import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {EmployeeService} from './services/employee.service';
import {Employee} from '../../shared/models/entity/employee';
import {firstValueFrom} from "rxjs";

@Component({
  selector: 'app-employee',
  templateUrl: './employee.component.html',
  styleUrls: ['./employee.component.scss']
})
export class EmployeeComponent implements OnInit {
  public employee!: Employee;
  public editing: boolean = false;
  originalEmployee!: Employee;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: EmployeeService
  ) {
  }

  async ngOnInit(): Promise<void> {
    await this.getEmployee();
  }

  public async getEmployee() {
    const idParam = this.route.snapshot.paramMap.get('id');
    if (idParam === null) return;
    const id = +idParam;
    this.employee = await firstValueFrom((this.api.getEmployeeById(id)));
  }

  toggleEditing(): void {
    if (this.editing) {
      // Если режим редактирования отключается, сбросить изменения
      this.employee = {...this.originalEmployee};
    } else {
      // Если режим редактирования включается, сохранить оригинальные данные
      this.originalEmployee = {...this.employee};
    }
    this.editing = !this.editing;
  }

  saveEmployee(): void {
    if (!this.employee) return;
    this.api.updateEmployee(this.employee).subscribe(() => {
      this.editing = false;
    });
  }

  deleteEmployee(): void {
    if (confirm('Are you sure you want to delete this employee?')) {
      this.api.deleteEmployee(this.employee.id)
        .subscribe(() => {
          this.router.navigate(['/employees']);
        });
    }
  }
}
