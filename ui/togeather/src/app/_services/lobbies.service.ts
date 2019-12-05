import { Injectable } from '@angular/core';
import * as LeafLet from 'leaflet';
import { MapPopUpService } from './map-pop-up.service';
import { Lobby } from '../_models/lobby';

@Injectable({
  providedIn: 'root'
})
export class LobbiesService {

  lobbies: Lobby[] = [
    {
      ownerId: null,
      expirationDate: null,
      location: {
        lat:51.101417,
        lon:17.036716
      },
      restaurant: null,
      addressLobby: null,
    },
    {
      ownerId: null,
      expirationDate: null,
      location: {
        lat: 51.100919,
        lon: 17.032328
      },
      restaurant: null,
      addressLobby: null,
    },
  ]
  constructor(private lobbyPopupService: MapPopUpService) { }

  makeLobbiesMarkers(map: LeafLet.map): void {
    for (const lobby of this.lobbies) {
      const lat_coord = lobby.location.lat;
      const lon_coord = lobby.location.lon;
      const marker = LeafLet.marker([lat_coord, lon_coord]);

      marker.bindPopup(this.lobbyPopupService.makeLobbyPopup(lobby));
      marker.addTo(map);
    }
  };

}
