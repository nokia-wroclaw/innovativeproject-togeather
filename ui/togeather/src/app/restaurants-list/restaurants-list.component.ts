import { Component, OnInit } from '@angular/core';
import {Restaurant} from "../_models/restaurant";
import {ApiService} from "../_services/api.service";
import { RedirectionService } from "../_services/redirection.service";

@Component({
  selector: 'app-restaurants-list',
  templateUrl: './restaurants-list.component.html',
  styleUrls: ['./restaurants-list.component.scss']
})
export class RestaurantsListComponent implements OnInit {

  restaurants: Restaurant[];

  constructor(
      private apiService: ApiService,
      private redirectionService: RedirectionService,
  ) { }

  ngOnInit() {
    this.apiService.getRestaurants().subscribe(
        restaurants => {
          this.restaurants = restaurants;
        }
    );
  }
  redirectToRestaurant(id: number): void{
    this.redirectionService.redirectToSingleRestaurant(id);
  }

}
