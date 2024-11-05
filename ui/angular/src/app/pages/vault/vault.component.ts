import { ChangeDetectionStrategy, Component } from '@angular/core';
import { KeyValueSet, StoreDBService } from '../../services/vaultrdb/v1/storedb/store-db.service';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { DetailedKVSComponent } from './detailed-kvs/detailed-kvs.component';
import { SlicePipe } from '@angular/common';

@Component({
  selector: 'app-vault',
  standalone: true,
  imports: [
    MatButtonModule,
    MatFormFieldModule,
    MatTableModule,
    DetailedKVSComponent,
    FormsModule,
    SlicePipe
  ],
  templateUrl: './vault.component.html',
  styleUrl: './vault.component.css',
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ])
  ],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class VaultComponent {

  // table content
  displayedColumns = ['key', 'description', 'function'];
  dataSource = new MatTableDataSource<KeyValueSet>([]);
  expandedElement!: KeyValueSet | null;

  // list of tags which are already in use
  possibleTags: string[] = [] // this has to be received from the backend, calculating all used tags for the autocomplete

  // constructor for the table
  // loads the kvs at startup
  constructor(
    private storedbService: StoreDBService
  ) {
    // load the buckets
    this.storedbService.KVS_List().subscribe(response => this.dataSource.data = response)
  }

  onChange(result: string, kvs: KeyValueSet) {
    console.log('result:',result, 'kvs:', kvs.key)
    if (result == 'deleted') {
      const index = this.dataSource.data.indexOf(kvs);
      this.dataSource.data.splice(index,1);
      this.dataSource._updateChangeSubscription()
    }
  }

  openCreateKVSForm() {
    this.storedbService.openDialog().subscribe(_ => 
      this.storedbService.KVS_List().subscribe(response => this.dataSource.data = response))
  }
}
