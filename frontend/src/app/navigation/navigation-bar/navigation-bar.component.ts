import { Component } from '@angular/core';

interface Url {
  name: string;
  link: string;
  isActive: boolean;
}

@Component({
  selector: 'woodpecker-navigation-bar',
  templateUrl: './navigation-bar.component.html',
  styleUrls: ['./navigation-bar.component.scss']
})


export class NavigationBarComponent {

      public links: Url[] = [
          {
            name: "Сотрудники",
            link: "employees",
            isActive: false
          },
          {
            name: "Чаты",
            link: "chats",
            isActive: false
          },
          {
            name: "Статистика",
            link: "statistics",
            isActive: false
          },
          {
            name: "Логирование",
            link: "logging",
            isActive: false
          }
        ]

  onSelect(link: Url)
  {
    for (let l of this.links)
    {
      l.isActive = false;
    }
    link.isActive = true;
  }

  setActive(link: Url): string
  {
    return link.isActive ? "active" : ""
  }
}
