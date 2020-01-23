import { Component, OnInit } from '@angular/core';
import { Coordinates } from './map.component';
import { LocationService } from '../_services/location.service';
import { Observable } from 'rxjs';
import { Lobby } from '../_models/lobby';
import { ApiService } from '../_services/api.service';

@Component({
    selector: 'map-container',
    templateUrl: './map-container.component.html'
})
export class MapContainerComponent implements OnInit {

    location: Coordinates;
    lobbies$: Observable<Lobby[]>;

    constructor(
        private locationService: LocationService,
        private api: ApiService,
    ) { }

    ngOnInit() {
        this.locationService.getLocation()
            .then((location: Coordinates) => {
                this.location = location;
            })
            .catch(error => {
                // TODO: Change this to toaster popup.
                // alert(error);
                console.error('Could not get your location');
            });

        this.lobbies$ = this.api.getLobbies();
    }
}
