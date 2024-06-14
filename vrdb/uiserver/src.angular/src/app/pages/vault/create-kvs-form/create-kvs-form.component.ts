import { Component } from '@angular/core';
import { StoreDBService } from '../../../services/vaultrdb/v1/storedb/store-db.service';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-create-kvs-form',
  standalone: true,
  imports: [
    MatFormFieldModule,
    MatButtonModule,
    MatInputModule,
    MatIconModule, 
    FormsModule
  ],
  templateUrl: './create-kvs-form.component.html',
  styleUrl: './create-kvs-form.component.css'
})
export class CreateKvsFormComponent {

  _key: string = ""
  _value: string = ""
  _description: string = ""

  constructor(
    private storedbService: StoreDBService
  ) { }
  
  submit() {

  }
}
