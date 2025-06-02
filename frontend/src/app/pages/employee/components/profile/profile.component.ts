import { Component, Input, OnInit, OnChanges, SimpleChanges } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EmployeeService } from '../../services/employee.service';
import { Employee } from '../../../../shared/models/entity/employee';
import { ChartConfiguration, ChartData, ChartOptions } from 'chart.js';

@Component({
  selector: 'app-employee-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit, OnChanges {
  @Input() employee!: Employee;
  public editing: boolean = false;
  public originalEmployee!: Employee;

  // Линейный график (рейтинг)
  public lineChartData: ChartData<'line'> = {
    datasets: [
      {
        data: [],
        label: 'Рейтинг сотрудника',
        borderColor: '#007bff',
        fill: false,
        tension: 0.4,
        pointRadius: 4,
        pointBackgroundColor: '#007bff'
      }
    ],
    labels: []
  };
  public lineChartOptions: ChartOptions<'line'> = {
    responsive: true,
    scales: {
      x: {
        title: {
          display: true,
          text: 'Проект'
        }
      },
      y: {
        title: {
          display: true,
          text: 'Значение рейтинга'
        },
        suggestedMin: -10,
        suggestedMax: 10
      }
    }
  };

  // Круговая диаграмма (проекты)
  public pieChartData: ChartData<'pie'> = {
    labels: ['Открытые', 'Закрытые вовремя', 'Просроченные'],
    datasets: [{
      data: [],
      backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
      hoverOffset: 4
    }]
  };
  public pieChartOptions: ChartOptions<'pie'> = {
    responsive: true,
    plugins: {
      legend: { position: 'top' },
      tooltip: { enabled: true }
    }
  };

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: EmployeeService
  ) {}

  ngOnInit(): void {
    this.updateChartData();
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['employee'] && changes['employee'].currentValue) {
      this.employee = changes['employee'].currentValue;
      this.originalEmployee = { ...this.employee };
      this.updateChartData();
    }
  }

  private updateChartData(): void {
    console.log('Employee:', this.employee);
    if (this.employee && this.employee.rating && this.employee.rating.length > 0) {
      // Линейный график: исходные значения рейтинга
      const data = [...this.employee.rating];
      const labels = this.employee.rating.map((_, index) => `Проект ${index + 1}`);
      this.lineChartData = {
        labels: labels,
        datasets: [{ ...this.lineChartData.datasets[0], data: data }]
      };
    } else {
      console.log('Данные для линейного графика отсутствуют:', this.employee?.rating);
    }

    // Круговая диаграмма: соотношение проектов
    if (this.employee && this.employee.total_projects_count !== undefined) {
      const completedOnTime = this.employee.total_projects_count - this.employee.overdue_projects_count - this.employee.active_projects_count;
      const active = this.employee.active_projects_count || 0;
      this.pieChartData = {
        labels: ['Активные', 'Закрытые вовремя', 'Просроченные'],
        datasets: [{
          ...this.pieChartData.datasets[0],
          data: [active, completedOnTime, this.employee.overdue_projects_count]
        }]
      };
    } else {
      console.log('Данные для круговой диаграммы отсутствуют:', this.employee);
    }
  }

  toggleEditing(): void {
    if (this.editing) {
      this.employee = { ...this.originalEmployee };
    } else {
      this.originalEmployee = { ...this.employee };
    }
    this.editing = !this.editing;
  }

  saveEmployee(): void {
    if (!this.employee) return;
    this.api.updateEmployee(this.employee).subscribe({
      next: () => {
        this.editing = false;
        this.updateChartData();
      },
      error: (err) => console.error('Ошибка при сохранении:', err)
    });
  }

  deleteEmployee(): void {
    if (confirm('Вы уверены, что хотите удалить этого сотрудника?')) {
      this.api.deleteEmployee(this.employee.id).subscribe({
        next: () => this.router.navigate(['/employees']),
        error: (err) => console.error('Ошибка при удалении:', err)
      });
    }
  }
}
