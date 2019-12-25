import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Restaurant } from '../_models/restaurant';
import { Lobby } from '../_models/lobby';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class ApiService {

    readonly baseUrl = 'http://localhost:8000/api';

    constructor(
        private http: HttpClient,
    ) { }

    getRestaurants(): Observable<Restaurant[]> {
        return this.http.get<Restaurant[]>(
            this.baseUrl + '/restaurants'
        );
    }

    getSingleRestaurant(id: string | number): Observable<Restaurant> {
        return this.http.get<Restaurant>(
            this.baseUrl + '/restaurants/' + id
        );
    }

    getLobbies(): Observable<Lobby[]> {
        return of([
            {
                ownerId: 123,
                expirationDate: new Date('2019-12-24'),
                location: { lat: 51.1264, lon: 16.9918 },
                restaurant: {
                    id: 3,
                    name: 'KFC',
                    address: 'Kromera, Wrocław',
                    menu: [
                        'Burger',
                        'Frytki',
                        'Cola',
                    ]
                },
                lobbyAddress: 'Lobby address',
            },
            {
                ownerId: 123,
                expirationDate: new Date('2019-12-24'),
                location: { lat: 51.1292, lon: 16.9708 },
                restaurant: {
                    id: 2,
                    name: 'Burger King',
                    address: 'Plac Dominikański, Wrocław',
                    menu: [
                        'Burger',
                        'Frytki',
                        'Cola',
                    ]
                },
                lobbyAddress: 'Nokia West link',
            },
        ]);
    }
}
