import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BehaviorSubject, Observable } from 'rxjs';
import { AuthComponent } from '../pages/auth/auth.component';

export class UserAuth {
  username?: string
  password?: string
  b64?: string
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
