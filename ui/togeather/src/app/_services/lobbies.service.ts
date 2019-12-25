import { Injectable } from '@angular/core';
import * as LeafLet from 'leaflet';
import { MapPopUpService } from './map-pop-up.service';
import { Lobby } from '../_models/lobby';
import { ApiService } from './api.service';

@Injectable({
  providedIn: 'root'
})
export class LobbiesService {

  private markerIcon = LeafLet.icon({
    iconUrl: './assets/marker-icon.png',
    iconSize: [25, 45],
  });

  constructor(
      private lobbyPopupService: MapPopUpService,
      private apiService: ApiService,
  ) { }

  makeLobbiesMarkers(map: LeafLet.map): void {
    this.apiService.getLobbies().subscribe(lobbies => {
      for (const lobby of lobbies) {
        const marker = LeafLet.marker (
            [ lobby.location.lat, lobby.location.lon ],
            { icon: this.markerIcon}
        );

        marker.bindPopup(MapPopUpService.makeLobbyPopup(lobby));
        marker.addTo(map);
      }
    });
  }

}
