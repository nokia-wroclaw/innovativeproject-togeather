import { Component } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { ApiService } from '../_services/api.service';
import { ToastrService } from 'ngx-toastr';
import { User } from '../_models/user';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.scss', '../../styles/global/centered-form.scss']
})
export class RegistrationPageComponent {
  name = new FormControl('', Validators.required);
  disableButton: boolean = false;
  user: User;

  constructor(
      private api: ApiService,
      private toaster: ToastrService,
  ) { }

  register() {
    this.name.setValue(this.name.value.trim());

    if (this.name.valid) {
      this.disableButton = true;
      this.api.register(this.name.value.trim()).subscribe(
          (user: User) => {
            this.user = user;
            this.toaster.success('You can now log in', 'Successfully registered!');
            this.disableButton = false;
          },
          error => this.toaster.error(error, 'Could not register')
      );
    }
  }
}
