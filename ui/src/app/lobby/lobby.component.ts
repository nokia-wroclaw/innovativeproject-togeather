import { Component, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ActivatedRoute } from '@angular/router';
import { pluck, switchMap } from 'rxjs/operators';
import { ApiService } from '../_services/api.service';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-lobby',
    templateUrl: './lobby.component.html',
    styleUrls: ['./lobby.component.scss']
})
export class LobbyComponent implements OnInit {

    lobby$: Observable<Lobby>;

    constructor(
        private route: ActivatedRoute,
        private api: ApiService,
    ) { }

    ngOnInit() {
        this.lobby$ = this.route.params.pipe(
            pluck('lobbyId'),
            switchMap(lobbyId => this.api.getLobby(lobbyId)),
        );
    }

}
