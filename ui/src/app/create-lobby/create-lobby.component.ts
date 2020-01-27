import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Restaurant } from '../_models/restaurant';
import { ApiService } from '../_services/api.service';
import { Observable, of } from 'rxjs';
import { PostLobbyDto } from '../_models/post-lobby-dto';
import { Lobby } from '../_models/lobby';
import { RedirectionService } from '../_services/redirection.service';
import { ToastrService } from 'ngx-toastr';
import { NgxMaterialTimepickerTheme } from 'ngx-material-timepicker';

@Component({
  selector: 'app-create-lobby',
  templateUrl: './create-lobby.component.html',
  styleUrls: ['../../styles/global/centered-form.scss']
})
export class CreateLobbyComponent implements OnInit {

  disableCreateButton = false;
  restaurants$: Observable<Restaurant[]> = of([]);

  lobbyForm = this.fb.group({
    street: [ '', Validators.required ],
    nr: [ '', Validators.required ],
    city: [ '', Validators.required ],
    restaurantId: [ null, Validators.required ],
    expirationHour: ['', [Validators.required, Validators.pattern(/(\d{1,2}):(\d{2})/)]],
  });

  timePickerTheme: NgxMaterialTimepickerTheme = {
    container: {
      buttonColor: '#3f51b5'
    },

    dial: {
      dialBackgroundColor: '#3f51b5',
      dialEditableActiveColor: '#3f51b5',
    },

    clockFace: {
      clockHandColor: '#3f51b5',
    },
  };

  static sanitize (field: string) {
    return field.trim().replace(',', ' ');
  }

  constructor(
      private api: ApiService,
      private fb: FormBuilder,
      private redirectionService: RedirectionService,
      private toaster: ToastrService,
  ) { }

  ngOnInit() {
    const today = new Date();
    this.lobbyForm.controls['expirationHour'].setValue(`${ today.getHours() }:${ today.getMinutes() + 10 }`);

    this.restaurants$ = this.api.getRestaurants();
  }

  createNewLobby(): void {
    if (this.lobbyForm.valid) {
      this.disableCreateButton = true;

      const expirationDate = new Date();
      const expirationTime = this.lobbyForm.controls['expirationHour'].value.split(':', 2);
      expirationDate.setHours(expirationTime[0], expirationTime[1]);

      const address = CreateLobbyComponent.sanitize(this.lobbyForm.get('nr').value)
          + ','
          + CreateLobbyComponent.sanitize(this.lobbyForm.get('street').value)
          + ','
          + CreateLobbyComponent.sanitize(this.lobbyForm.get('city').value);

      const newLobby: PostLobbyDto = {
        restaurant_id: this.lobbyForm.controls['restaurantId'].value,
        expires: expirationDate.toISOString(),
        address: address
      };


      this.api.postLobby(newLobby).subscribe((lobby: Lobby) => {
        this.disableCreateButton = false;
        this.redirectionService.redirectToSingleLobby(lobby.id);
      }, error => {
        this.toaster.error(error, 'Could not create lobby');
        this.disableCreateButton = false;
      });
    } else {
      this.toaster.error('Form is not valid! I won\'t send the request >.<');
    }
  }
}
