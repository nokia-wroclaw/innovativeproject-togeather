import { Injectable } from '@angular/core';
import * as LeafLet from 'leaflet';
import { MapPopUpService } from './map-pop-up.service';

@Injectable({
  providedIn: 'root'
})
export class LobbiesService {

  lobbies = [
    {address:'ul. H. Kołłątaja', lat:51.101417, lon: 17.036716},
    {address:'ul. J. Piłsudskiego', lat:51.100919, lon: 17.032328},
  ]
  constructor(private lobbyPopupService: MapPopUpService) { }

  makeLobbiesMarkers(map: LeafLet.map): void {
    for (const lobby of this.lobbies) {
      const lat_coord = lobby.lat;
      const lon_coord = lobby.lon;
      const marker = LeafLet.marker([lat_coord, lon_coord]);//.addTo(map);

      marker.bindPopup(this.lobbyPopupService.makeLobbyPopup(lobby));
      marker.addTo(map);
    }
  };

}
