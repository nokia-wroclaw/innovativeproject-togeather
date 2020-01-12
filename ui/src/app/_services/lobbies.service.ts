import { Injectable } from '@angular/core';
import * as LeafLet from 'leaflet';
import { MapPopUpService } from './map-pop-up.service';
import { Lobby } from '../_models/lobby';

@Injectable({
    providedIn: 'root'
})
export class LobbiesService {

    private markerIcon = LeafLet.icon({
        iconUrl: './assets/marker-icon.png',
        iconSize: [25, 45],
    });

    static removeMarker(map: ReturnType<LeafLet.map>, marker: ReturnType<LeafLet.marker>) {
        map.removeLayer(marker);
    }

    makeLobbyMarker(map: LeafLet.map, lobby: Lobby): ReturnType<LeafLet.marker> {
        const marker = LeafLet.marker(
            [lobby.location.lat, lobby.location.lon],
            { icon: this.markerIcon }
        );

        marker.bindPopup(MapPopUpService.makeLobbyPopup(lobby));
        marker.addTo(map);

        return marker;
    }

}
