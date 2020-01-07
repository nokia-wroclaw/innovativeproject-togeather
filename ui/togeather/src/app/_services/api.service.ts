import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { Restaurant } from '../_models/restaurant';
import { Lobby } from '../_models/lobby';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { catchError, map } from 'rxjs/operators';
import { PostLobbyDto } from '../_models/post-lobby-dto';

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
                        location: { lat: obj.location.lat, lon: obj.location.lon },
                        address: obj.location.lobby_address,
                    };
                });
            })
        );
    }

    postLobby(lobby: PostLobbyDto): Observable<Lobby> {
        lobby.restaurant_id = Number(lobby.restaurant_id);

        return this.http.post<Lobby>(
            this.baseUrl + '/lobbies',
            lobby,
        ).pipe(
            catchError(this.handleError)
        );
    }

    private handleError(error: HttpErrorResponse) {
        if (error.error instanceof ErrorEvent) {
            // A client-side or network error occurred. Handle it accordingly.
            console.error('An error occurred:', error.error.message);
        } else {
            // The backend returned an unsuccessful response code.
            // The response body may contain clues as to what went wrong,
            console.error(
                `Backend returned code ${error.status}, ` +
                `body was: ${error.error}`);
        }
        // return an observable with a user-facing error message
        return throwError('Something bad happened; please try again later.');
    }
}
