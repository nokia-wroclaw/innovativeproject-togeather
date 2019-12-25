import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from '../_services/api.service';
import { Restaurant } from '../_models/restaurant';

@Component({
    selector: 'app-restaurant',
    templateUrl: './restaurant.component.html',
    styleUrls: ['./restaurant.component.scss']
})
export class RestaurantComponent implements OnInit {

    restaurant: Restaurant;

    constructor(
        private route: ActivatedRoute,
        private apiService: ApiService,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe(params => {
            this.apiService.getSingleRestaurant(params.get('restaurantId')).subscribe(
                restaurant => {
                    this.restaurant = restaurant;
                }
            );
        });
    }

}
