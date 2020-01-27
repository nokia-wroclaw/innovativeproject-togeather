import { Component } from '@angular/core';
import { ApiService } from '../_services/api.service';
import { ToastrService } from 'ngx-toastr';
import { RedirectionService } from '../_services/redirection.service';
import { ActivatedRoute } from '@angular/router';
import { FormControl, Validators } from '@angular/forms';
import { pluck, switchMap, } from 'rxjs/operators';

@Component({
  selector: 'app-join-lobby-link',
  templateUrl: './join-lobby-link.component.html',
  styleUrls: ['./join-lobby-link.component.scss']
})
export class JoinLobbyLinkComponent {

  lobbyId;

  userName = new FormControl('', Validators.required);

  constructor(
    private api: ApiService,
    private toaster: ToastrService,
    private redirectionService: RedirectionService,
    private route: ActivatedRoute,
  ) {
  }

  onNoClick(): void {
    this.redirectionService.redirectToHomePage();
  }

  onJoinClick(): void {
    if (this.userName.valid) {
      this.route.params.pipe(
        pluck('lobbyId'),
        switchMap(lobbyId => {
          this.lobbyId = lobbyId;
          return this.api.joinLobby(lobbyId, this.userName.value);
        }),
      ).subscribe(
        () => this.redirectionService.redirectToSingleLobby(this.lobbyId),
        error => this.toaster.error(error),
      );
    }
  }

}
