import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { AuthService, UserAuth } from '../../services/auth/auth.service';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [
    MatDialogModule,
    MatButtonModule,

    FormsModule, 
    MatFormFieldModule,

    MatInputModule,
    MatIconModule
  ],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.css'
})
export class AuthComponent {
  public username: string = ""
  public password: string = ""

  constructor(
    private authService: AuthService,
  ) { 
    const userLogin = authService.authValue;
    if (userLogin) {
      this.username = userLogin.username!
      this.password = userLogin.password!
    }
  }

  login() {
    this.authService.login(this.username, this.password)
    this.closeDialog()
  }

  closeDialog() {
    
  }
}
