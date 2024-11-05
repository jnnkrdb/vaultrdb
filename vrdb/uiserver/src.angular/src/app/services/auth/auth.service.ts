import { Injectable, inject } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BehaviorSubject, Observable } from 'rxjs';
import { AuthComponent } from '../../pages/auth/auth.component';
import { HttpEvent, HttpHandlerFn, HttpRequest } from '@angular/common/http';

export class UserAuth {
  username?: string
  password?: string
  b64?: string
}

// http interceptor
export function BasicAuth(request: HttpRequest<unknown>, next: HttpHandlerFn): Observable<HttpEvent<unknown>> {
  // add header with basic auth credentials if user is logged in and request is to the api url
  const user = inject(AuthService).authValue;
  if (user?.username && user?.password) {
      request = request.clone({
          setHeaders: { 
              Authorization: `Basic ${user.b64}`
          }
      });
  }
  return next(request);
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private authSubject: BehaviorSubject<UserAuth | null>;
  public auth: Observable<UserAuth | null>;

  constructor(
    private dialog: MatDialog
  ) { 
    this.authSubject = new BehaviorSubject(JSON.parse(localStorage.getItem('auth')!));
    this.auth = this.authSubject.asObservable();
  }
  
  public get authValue() {
    return this.authSubject.value;
  }

  login(_username: string, _password: string) {
    var user: UserAuth = {
      username: _username,
      password: _password,
      b64: window.btoa(_username+":"+_password)
    }
    localStorage.setItem('auth', JSON.stringify(user))
    this.authSubject.next(user)
    console.log("logged in with: "+user.username+", "+user.password+", "+user.b64)

  }

  logout() {    
    localStorage.removeItem('auth');
    this.authSubject.next(null);
  }

  openPopup() {
    this.dialog.open(AuthComponent);
  }
}
