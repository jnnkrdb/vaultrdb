import { Component } from '@angular/core';
import { KeyValueSet, StoreDBService } from '../../services/vaultrdb/v1/storedb/store-db.service';
import { CreateKvsFormComponent } from './create-kvs-form/create-kvs-form.component';

@Component({
  selector: 'app-vault',
  standalone: true,
  imports: [
    CreateKvsFormComponent
  ],
  templateUrl: './vault.component.html',
  styleUrl: './vault.component.css'
})
export class VaultComponent {

  kvsList: KeyValueSet[] = []

  // constructor for the table
  // loads the kvs at startup
  constructor(
    private storedbService: StoreDBService
  ) {
    // load the buckets
    this.storedbService.KVS_List().subscribe(response => this.kvsList = response)
  }

}
