import { Coordinates } from '../map/map.component';
import { Restaurant } from './restaurant';

export interface Lobby {
    ownerId: number;
    expirationDate: Date;
    location: Coordinates;
    restaurant: Restaurant;
    lobbyAddress: string;
}
