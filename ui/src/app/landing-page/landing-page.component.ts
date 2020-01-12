import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Coordinates } from '../map/map.component';
import { RedirectionService } from '../_services/redirection.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.scss']
})
export class LandingPageComponent implements OnInit {
  location = new FormControl('');
  coordinates: Coordinates = null;

  constructor(
      private redirectionService: RedirectionService,
  ) { }

  ngOnInit() {

  }

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
