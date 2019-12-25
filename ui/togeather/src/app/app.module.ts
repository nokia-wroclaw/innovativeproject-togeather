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
    MatAutocompleteModule
} from '@angular/material';
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
import { MaterialTimePickerModule } from '@candidosales/material-time-picker';
import { FlexLayoutModule } from '@angular/flex-layout';
import { BeautifyAddressPipe } from './_pipes/beautify-address.pipe';
import { CartComponent } from './cart/cart.component';


@NgModule({
    declarations: [
        AppComponent,
        FooterComponent,
        HeaderComponent,
        MapComponent,
        LandingPageComponent,
        RestaurantsListComponent,
        RestaurantComponent,
        CreateLobbyComponent,
        BeautifyAddressPipe,
        CartComponent,
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
        MaterialTimePickerModule,
        FlexLayoutModule,
        MatAutocompleteModule,
        MatListModule,
    ],
    providers: [
        ApiService,
        RedirectionService,
        LobbiesService,
        MapPopUpService,
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
