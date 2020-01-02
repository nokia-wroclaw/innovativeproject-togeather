import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
    providedIn: 'root'
})
export class RedirectionService {

    constructor(
        private router: Router,
    ) { }

    redirectToSingleRestaurant(id: number): void {
        this.router.navigateByUrl(`/restaurants/${id}`);
    }

    redirectToHomePage(): void {
        this.router.navigateByUrl('/');
    }

    redirectToLobbyCreation(): void {
        this.router.navigateByUrl('/create-lobby');
    }

    redirectToRestaurants(): void {
        this.router.navigateByUrl('/restaurants');
    }
}
