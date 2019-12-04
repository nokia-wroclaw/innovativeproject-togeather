import { Component, OnInit } from '@angular/core';
import { FormControl } from "@angular/forms";
import {LocationService} from "../_services/location.service";
import {Coordinates} from "../map/map.component";

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.scss']
})
export class LandingPageComponent implements OnInit {
  location = new FormControl('');
  coordinates: Coordinates = null;

  constructor(
      private locationService: LocationService
  ) { }

  ngOnInit() {

  }

}
