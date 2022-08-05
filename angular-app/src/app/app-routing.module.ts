import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TemplatesPage } from './pages/templates/templates.component';

const routes: Routes = [
  {
    path: 'templates',
    component: TemplatesPage,
  },
  {
    path: '**',
    redirectTo: 'templates',
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
