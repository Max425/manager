import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
import {DomSanitizer, SafeResourceUrl} from "@angular/platform-browser";
import {StatisticsService} from "./service/statistics.service";
import {UntypedFormControl} from "@angular/forms";
import {ReplaySubject, Subject, take, takeUntil} from "rxjs";
import {MatSelect} from "@angular/material/select";

@Component({
  selector: 'woodpecker-statistics',
  templateUrl: './statistics.component.html',
  styleUrls: ['./statistics.component.scss']
})
export class StatisticsComponent implements OnInit, AfterViewInit {

  private statisticsDashboard: SafeResourceUrl | undefined;
  private userIds: number[] | undefined;
  private dateStart: string | undefined;
  private dateEnd: string | undefined;

  public userIdCtrl: UntypedFormControl = new UntypedFormControl();
  public userIdFilterCtrl: UntypedFormControl = new UntypedFormControl();
  public filteredUserIds: ReplaySubject<number[]> = new ReplaySubject<number[]>(1);

  @ViewChild('singleSelect', { static: true })
  singleSelect!: MatSelect;

  isIndeterminate = false;
  isChecked = false;

  constructor(
    private sanitizer: DomSanitizer,
    private api: StatisticsService,
    ) {
  }

  ngOnInit(): void {
      this.api.getAddressGrafana().subscribe(data => {
          if(data.status == 200) {
              console.log(data)
              this.userIds = data.payload?.userIds;
              this.userIdCtrl.setValue(this.userIds);
              this.filteredUserIds.next(this.userIds!.slice());

              // listen for search field value changes
              this.userIdFilterCtrl.valueChanges
                  .subscribe(() => {
                      this.filterUserIds();
                  });

              this.statisticsDashboard = this.sanitizer
                  .bypassSecurityTrustResourceUrl(
                      `${data.payload?.url}`)
              this.userIdCtrl.valueChanges.subscribe(() => this.setUserIdFilter())
          }
          else {
              console.log(`status: ${data.status}\nmessage: ${data.message}`);
          }
      });
  }

  ngAfterViewInit() {
    this.setInitialValue();
  }

  setUserIdFilter() {
      let filterValue = {
          userIds: this.userIdCtrl.getRawValue() as number[],
          dateStart: this.dateStart,
          dateEnd: this.dateEnd
      };

      this.api.setFilterGrafana(filterValue).subscribe(data => {
          this.statisticsDashboard = this.sanitizer
              .bypassSecurityTrustResourceUrl(
                  `${data.payload?.url}`);
      })
  }

  getDashboard() {
    return this.statisticsDashboard;
  }

  toggleSelectAll(selectAllValue: boolean) {
    this.filteredUserIds
      .pipe(take(1))
      .subscribe((val) => {
        if (selectAllValue) {
          this.userIdCtrl.patchValue(val);
        } else {
          this.userIdCtrl.patchValue([]);
        }
      });
  }

  private filterUserIds() {
        if (!this.userIds) {
            return;
        }
        // get the search keyword
        let search = this.userIdFilterCtrl.value;

        if (!search) {
            this.filteredUserIds.next(this.userIds.slice());
            return;
        } else {
            search = search.toLowerCase();
        }
        this.filteredUserIds.next(
            this.userIds.filter((e) => e.toString().indexOf(search) > -1)
        );
    }
  private setInitialValue() {
      this.filteredUserIds
          .pipe(take(1))
          .subscribe(() => {
              this.singleSelect.compareWith = (a, b) => a && b && a === b;
          });
  }

  click(dateRangeStart: HTMLInputElement, dateRangeEnd: HTMLInputElement) {
      this.dateStart = dateRangeStart.value;
      this.dateEnd = dateRangeEnd.value;
      let filterValue = {dateStart: dateRangeStart.value, dateEnd: dateRangeEnd.value};
      this.setUserIdFilter();
  }
}
