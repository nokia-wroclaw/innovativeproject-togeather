import { Component } from '@angular/core';
import { RedirectionService } from '../_services/redirection.service';
import { UserService } from '../_services/user.service';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss']
})
export class HeaderComponent {

    constructor(
        private redirectionService: RedirectionService,
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
}
