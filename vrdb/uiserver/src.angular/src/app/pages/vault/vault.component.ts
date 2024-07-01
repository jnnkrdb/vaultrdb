import { Component } from '@angular/core';
import { KeyValueSet, StoreDBService } from '../../services/vaultrdb/v1/storedb/store-db.service';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { trigger, state, style, transition, animate } from '@angular/animations';

@Component({
  selector: 'app-vault',
  standalone: true,
  imports: [
    MatButtonModule,
    MatToolbarModule,
    MatFormFieldModule, 
    MatSelectModule, 
    MatInputModule,
    MatTableModule,
    MatIconModule,
    FormsModule
  ],
  templateUrl: './vault.component.html',
  styleUrl: './vault.component.css',
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ]
})
export class VaultComponent {

  // table content
  displayedColumns = ['key', 'description'];
  dataSource = new MatTableDataSource<KeyValueSet>([]);
  expandedElement!: KeyValueSet | null;

  // constructor for the table
  // loads the kvs at startup
  constructor(
    private storedbService: StoreDBService
  ) {
    // load the buckets
    this.storedbService.KVS_List().subscribe(response => this.dataSource.data = response)
  }

  openCreateKVSForm() {
    this.storedbService.openDialog().subscribe(_ => 
      this.storedbService.KVS_List().subscribe(response => this.dataSource.data = response))
  }

  saveKVS(kvs: KeyValueSet) {
    console.log(kvs)
  }

  deleteKVS(kvs: KeyValueSet) {
    console.log(kvs)
  }
}
