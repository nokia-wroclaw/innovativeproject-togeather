import { Component } from '@angular/core';
import { RedirectionService } from '../_services/redirection.service';
import { UserService } from '../_services/user.service';
import { ApiService } from '../_services/api.service';
import { ToastrService } from 'ngx-toastr';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss']
})
export class HeaderComponent {

    constructor(
        private redirectionService: RedirectionService,
        private api: ApiService,
        private toaster: ToastrService,
        public user: UserService,
    ) {
    }

    redirectToHomePage(): void {
        this.redirectionService.redirectToHomePage();
    }

    redirectToLogin() {
        this.redirectionService.redirectToLoginPage();
    }

    redirectToRegistration() {
        this.redirectionService.redirectToRegistrationPage();
    }

    logOut() {
        this.user.clearLoggedInUser();
        this.api.logOut().subscribe(() => {
            this.toaster.info('You\'ve been logged out');
        });
    }
}
