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

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: EmployeeService
  ) {}

  ngOnInit(): void {
    // Начальная инициализация, если employee уже доступен
    this.updateChartData();
  }

  ngOnChanges(changes: SimpleChanges): void {
    // Реагируем на изменения employee
    if (changes['employee'] && changes['employee'].currentValue) {
      this.employee = changes['employee'].currentValue;
      this.originalEmployee = { ...this.employee }; // Сохраняем оригинал для отмены
      this.updateChartData();
    }
  }

  private updateChartData(): void {
    if (this.employee && this.employee.rating && this.employee.rating.length > 0) {
      // Вычисляем кумулятивный рейтинг для оси Y
      const cumulativeRating = this.employee.rating.reduce((acc, curr, index) => {
        acc.push(curr);
        return acc;
      }, [] as number[]);

      // Создаём метки для оси X (равные промежутки)
      const labels = this.employee.rating.map((_, index) => `${index + 1}`);

      // Обновляем данные графика
      this.lineChartData = {
        labels: labels,
        datasets: [
          {
            ...this.lineChartData.datasets[0],
            data: cumulativeRating
          }
        ]
      };
    } else {
      console.log('Данные для графика отсутствуют:', this.employee?.rating);
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
