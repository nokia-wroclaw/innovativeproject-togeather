import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { RedirectionService } from '../_services/redirection.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.scss']
})
export class LandingPageComponent {
  location = new FormControl('');

  constructor(
      private redirectionService: RedirectionService,
  ) { }

  redirectToLobbyCreation() {
    this.redirectionService.redirectToLobbyCreation();
  }

  redirectToRestaurants() {
    this.redirectionService.redirectToRestaurants();
  }

  redirectToLobbies() {
    this.redirectionService.redirectToLobbies();
  }
}
