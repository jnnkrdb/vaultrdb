import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { MatListModule } from '@angular/material/list';
import { FormsModule } from '@angular/forms';
import { SlicePipe } from '@angular/common';

@NgModule({
  declarations: [
    AppComponent,
    SlicePipe
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatIconModule,
    MatButtonModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    FormsModule,
  ],
  providers: [
    SlicePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }