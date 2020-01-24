import { Component, OnInit } from '@angular/core';
import { ApiService } from '../_services/api.service';
import { ToastrService } from 'ngx-toastr';
import { RedirectionService } from '../_services/redirection.service';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-join-lobby-link',
  templateUrl: './join-lobby-link.component.html',
  styleUrls: ['./join-lobby-link.component.scss']
})
export class JoinLobbyLinkComponent implements OnInit {

  lobbyId;

  joinLobbyForm = this.fb.group({ userName: ['', Validators.required ],});

  constructor(
    private api: ApiService,
    private toaster: ToastrService,
    private redirectionService: RedirectionService,
    private route: ActivatedRoute,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
  }

  onNoClick(): void {
    this.redirectionService.redirectToHomePage();
  }

  onJoinClick(): void {
    this.route.paramMap.subscribe(params => {
        this.lobbyId = params.get('lobbyId');
      });
    if (this.joinLobbyForm.valid) {
      const userName = this.joinLobbyForm.controls['userName'].value;
      this.api.joinLobby(this.lobbyId, userName).subscribe(
        () => {
          this.redirectionService.redirectToSingleLobby(this.lobbyId);
        },
        error => {
          this.toaster.error(error);
        },
      ); 
    }
  }
  
}
