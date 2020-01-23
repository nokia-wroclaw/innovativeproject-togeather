import { Component, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ActivatedRoute } from '@angular/router';
import { catchError, pluck, switchMap } from 'rxjs/operators';
import { ApiService } from '../_services/api.service';
import { Observable, throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { RedirectionService } from '../_services/redirection.service';

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
        private toaster: ToastrService,
        private redirectionService: RedirectionService,
    ) { }

    ngOnInit() {
        this.lobby$ = this.route.params.pipe(
            pluck('lobbyId'),
            switchMap(lobbyId => this.api.getLobby(lobbyId)),
            catchError(error => {
                this.redirectionService.redirectToLobbies();
                this.toaster.error(error, 'Could not load this lobby');
                return throwError(error);
            }),
        );
    }

}
