import { Injectable } from '@angular/core';
import { environment } from './../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LogService {

  constructor() { }

  debug(message: string) {
    if (environment.production) {
      return;
    }

    console.log(message);
  }

  info(message: string) {
    console.log(message);
  }

  error(message: string) {
    console.log(message);
  }
}
