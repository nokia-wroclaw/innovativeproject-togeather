import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { Restaurant } from '../_models/restaurant';
import { Lobby } from '../_models/lobby';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError, map, pluck } from 'rxjs/operators';
import { PostLobbyDto } from '../_models/post-lobby-dto';
import { environment } from '../../environments/environment';
import { User } from '../_models/user';
import { UserDto } from '../_models/user-dto';

@Injectable({
    providedIn: 'root'
})
export class ApiService {

    readonly baseUrl = environment.apiUrl;

    constructor(
        private http: HttpClient,
    ) { }

    logIn(id: number): Observable<User> {
        return this.http.post<UserDto>(
            this.baseUrl + '/auth/login',
            { id: id },
            { withCredentials: true }
        ).pipe(
            map(userDto => {
                return {
                    id: userDto.user_id,
                    name: userDto.user_name
                };
            }),
            catchError(this.handleError)
        );
    }

    logOut(): Observable<void> {
        return this.http.delete<void>(this.baseUrl + '/auth/logout')
            .pipe(catchError(this.handleError));
    }

    register(name: string): Observable<User> {
        return this.http.post<UserDto>(
            this.baseUrl + '/auth/register',
            { name: name }
        ).pipe(
            map(userDto => {
                return {
                    id: userDto.user_id,
                    name: userDto.user_name,
                };
            }),
            catchError(this.handleError)
        );
    }

    checkUserLogin(): Observable<boolean> {
        return this.http.get(
            this.baseUrl + '/ping',
            { withCredentials: true }
        ).pipe(
            map(response => {
                return response === 'pong';
            }),
        );
    }

    getRestaurants(): Observable<Restaurant[]> {
        return this.http.get<Restaurant[]>(
            this.baseUrl + '/restaurants'
        ).pipe(
            catchError(this.handleError)
        );
    }

    getRestaurant(id: string | number): Observable<Restaurant> {
        return this.http.get<Restaurant>(
            this.baseUrl + '/restaurants/' + id
        ).pipe(
            catchError(this.handleError)
        );
    }

    getLobbies(): Observable<Lobby[]> {
        return this.http.get(
            this.baseUrl + '/lobbies'
        ).pipe(
            map((objects: any[]): Lobby[] => {
                return objects.map(obj => {
                    return {
                        id: obj.id,
                        restaurant: obj.restaurant,
                        expires: new Date(obj.expires),
                        location: { lat: obj.location.lat, lon: obj.location.lon },
                        address: obj.location.lobby_address,
                    };
                });
            }),
            catchError(this.handleError)
        );
    }

    getLobby(lobbyId: number): Observable<Lobby> {
        return this.http.get(
            `${this.baseUrl}/lobbies/${lobbyId}`,
            { withCredentials: true },
        ).pipe(
            map((data: any): Lobby => {
                return {
                    id: data.id,
                    restaurant: data.restaurant,
                    expires: new Date(data.expires),
                    location: { lat: data.location.lat, lon: data.location.lon },
                    address: data.location.lobby_address
                };
            }),
            catchError(this.handleError)
        );
    }

    postLobby(lobby: PostLobbyDto): Observable<Lobby> {
        lobby.restaurant_id = Number(lobby.restaurant_id);

        return this.http.post<Lobby>(
            this.baseUrl + '/lobbies',
            lobby,
            { withCredentials: true },
        ).pipe(
            catchError(this.handleError)
        );
    }

    joinLobby(lobbyId: number, userName?: string) {
        return this.http.post<{ id: number }>(
            this.baseUrl + `/lobbies/${lobbyId}`,
            { user_name: userName },
            { withCredentials: true },
        ).pipe(
            pluck('id'),
            catchError(this.handleError)
        );
    }

    private handleError(error: HttpErrorResponse) {
        if (error.error instanceof ErrorEvent) {
            // A client-side or network error occurred. Handle it accordingly.
            console.error('A client-side or network error occurred:', error.error.message);
        } else {
            // The backend returned an unsuccessful response code.
            // The response body may contain clues as to what went wrong,
            console.error('Backend returned code' + error.status);
            console.error('body was: ', error.error);
        }
        // return an observable with a user-facing error message
        return throwError(`Error ${error.status}: ${error.error.error}`);
    }
}
