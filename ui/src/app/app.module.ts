import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FooterComponent } from './footer/footer.component';
import { MatButtonModule } from '@angular/material/button';
import { HeaderComponent } from './header/header.component';
import {
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatMenuModule,
    MatCardModule,
    MatListModule,
    MatAutocompleteModule,
} from '@angular/material';
import { MatDialogModule } from '@angular/material/dialog';
import { MapComponent } from './map/map.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FormsModule } from '@angular/forms';
import { RestaurantsListComponent } from './restaurants-list/restaurants-list.component';
import { LobbiesService } from './_services/lobbies.service';
import { MapPopUpService } from './_services/map-pop-up.service';
import { RestaurantComponent } from './restaurant/restaurant.component';
import { HttpClientModule } from '@angular/common/http';
import { ApiService } from './_services/api.service';
import { RedirectionService } from './_services/redirection.service';
import { CreateLobbyComponent } from './create-lobby/create-lobby.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { BeautifyAddressPipe } from './_pipes/beautify-address.pipe';
import { CartComponent } from './cart/cart.component';
import { LobbiesListComponent } from './lobbies-list/lobbies-list.component';
import { BeautifyExpirationDatePipe } from './_pipes/beautify-expiration-date.pipe';
import { JoinLobbyComponent } from './join-lobby/join-lobby.component';
import { LobbyComponent } from './lobby/lobby.component';
import { RestaurantContainerComponent } from './restaurant/restaurant-container.component';
import { MapContainerComponent } from './map/map-container.component';
import { MapLobbyMarkerComponent } from './map-lobby-marker/map-lobby-marker.component';
import { ToastrModule } from 'ngx-toastr';
import { JoinLobbyLinkComponent } from './join-lobby-link/join-lobby-link.component';
import { NgxMaterialTimepickerModule } from 'ngx-material-timepicker';


@NgModule({
    declarations: [
        AppComponent,
        FooterComponent,
        HeaderComponent,
        MapComponent,
        MapContainerComponent,
        LandingPageComponent,
        RestaurantsListComponent,
        RestaurantComponent,
        CreateLobbyComponent,
        BeautifyAddressPipe,
        CartComponent,
        LobbiesListComponent,
        BeautifyExpirationDatePipe,
        JoinLobbyComponent,
        LobbyComponent,
        RestaurantContainerComponent,
        MapLobbyMarkerComponent,
        JoinLobbyLinkComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        MatButtonModule,
        MatMenuModule,
        MatIconModule,
        MatFormFieldModule,
        MatInputModule,
        MatCardModule,
        ReactiveFormsModule,
        FormsModule,
        HttpClientModule,
        FlexLayoutModule,
        MatAutocompleteModule,
        MatListModule,
        ToastrModule.forRoot({
            timeOut: 3500,
        }),
        MatDialogModule,
        NgxMaterialTimepickerModule.setLocale('pl-PL'),
    ],
    providers: [
        ApiService,
        RedirectionService,
        LobbiesService,
        MapPopUpService,
    ],
    bootstrap: [AppComponent],
    entryComponents: [JoinLobbyComponent],
})
export class AppModule {
}
