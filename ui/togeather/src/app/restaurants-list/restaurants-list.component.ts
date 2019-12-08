import { Component, OnInit } from '@angular/core';
import {Restaurant} from "../_models/restaurant";
import {ApiService} from "../_services/api.service";

@Component({
  selector: 'app-restaurants-list',
  templateUrl: './restaurants-list.component.html',
  styleUrls: ['./restaurants-list.component.scss']
})
export class RestaurantsListComponent implements OnInit {

  restaurants: Restaurant[];

  constructor(
      private apiService: ApiService,
  ) { }

  ngOnInit() {
    this.apiService.getRestaurants().subscribe(
        restaurants => {
          this.restaurants = restaurants;
        }
    );
  }

}
