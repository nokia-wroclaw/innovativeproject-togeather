import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Restaurant } from '../_models/restaurant';
import { Lobby } from '../_models/lobby';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

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

    // TODO: Write test and check that api returns expected array of Lobby objects
    getLobbies(): Observable<Lobby[]> {
        return this.http.get(
            this.baseUrl + '/lobbies'
        ).pipe(
            map((objects: any[]): Lobby[] => {
                return objects.map(obj => {
                    return {
                        id: obj.id,
                        restaurant: obj.restaurant,
                        owner: obj.owner,
                        expires: new Date(obj.expires),
                        location: { lat: obj.lat, lon: obj.lon },
                    };
                });
            })
        );
    }
}
