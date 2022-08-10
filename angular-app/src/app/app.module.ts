import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { TemplateResource } from './resources/template.resource';
import { TemplateResourceMock } from './resources/template.resource.mock';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MaterialModule } from './material.module';
import { TemplatesPage } from './pages/templates/templates.component';

@NgModule({
  declarations: [
    AppComponent,
    TemplatesPage
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MaterialModule,
  ],
  providers: [
    // TODO Remove mock
    { provide: TemplateResource, useClass: TemplateResourceMock }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
