import { Component } from '@angular/core';
import { RedirectionService } from '../_services/redirection.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {

  constructor(
    private redirectionService: RedirectionService,
  ) { }

  redirectToHomePage(): void {
    this.redirectionService.redirectToHomePage();
  }
}
