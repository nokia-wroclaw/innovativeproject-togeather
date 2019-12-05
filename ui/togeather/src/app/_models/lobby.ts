import { Coordinates } from '../map/map.component';
import { Restaurant } from './restaurant';

export class Lobby {
    ownerId: number;
    expirationDate: Date;
    location: Coordinates;
    restaurant: Restaurant;
    addressLobby: string;
}
