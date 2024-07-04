import { Component } from '@angular/core';
import { NewKeyValueSet, StoreDBService } from '../../../services/vaultrdb/v1/storedb/store-db.service';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogActions, MatDialogClose, MatDialogContent, MatDialogModule, MatDialogTitle } from '@angular/material/dialog';

@Component({
  selector: 'app-create-kvs-form',
  standalone: true,
  imports: [
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,
    MatDialogModule,
    MatButtonModule,

    FormsModule, 
    MatFormFieldModule,

    MatInputModule,
    MatIconModule
  ],
  templateUrl: './create-kvs-form.component.html',
  styleUrl: './create-kvs-form.component.css'
})
export class CreateKvsFormComponent {

  public newKVS: NewKeyValueSet = {
    key: '',
    value: '',
    tags: [],
    description: ''
  }

  constructor(
    private storedbService: StoreDBService
  ) {

  }

  replaceUnwantedCharacters() {
    this.newKVS.key = this.newKVS.key
      .replaceAll('\\', ' ')
      .replaceAll(' ', '_')
  }
  
  submit() {
    console.log("submitted", this.newKVS)
    if (this.newKVS.key == "") {
      console.log("newKVS.key shouldn't be empty")
      return
    }

    this.storedbService.KVS_Create(this.newKVS).subscribe()
  }
}
