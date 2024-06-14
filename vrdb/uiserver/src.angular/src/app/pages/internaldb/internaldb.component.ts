import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatFormFieldModule } from '@angular/material/form-field';
import { BucketKeyValueSet, InternalDBService } from '../../services/vaultrdb/v1/internaldb/internal-db.service';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

const TEST_ITEMS: BucketKeyValueSet[] = [
  {key: "hello", value: "world"},
  {key: "hello1", value: "world"},
  {key: "hello2", value: "world"},
  {key: "hello3", value: "world"},
  {key: "hello4", value: "world"},
  {key: "hello", value: "world"},
  {key: "hello1", value: "world"},
  {key: "hello2", value: "world"},
  {key: "hello3", value: "world"},
  {key: "hello4", value: "world"},
  {key: "hello", value: "world"},
  {key: "hello1", value: "world"},
  {key: "hello2", value: "world"},
  {key: "hello3", value: "world"},
  {key: "hello4", value: "world"},
  {key: "hello", value: "world"},
  {key: "hello1", value: "world"},
  {key: "hello2", value: "world"},
  {key: "hello3", value: "world"},
  {key: "hello4", value: "world"},
  {key: "hello", value: "world"},
  {key: "hello1", value: "world"},
  {key: "hello2", value: "world"},
  {key: "hello3", value: "world"},
  {key: "hello4", value: "world"},
]

@Component({
  selector: 'app-internaldb',
  standalone: true,
  imports: [
    MatFormFieldModule, 
    MatSelectModule, 
    MatInputModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    FormsModule
  ],
  templateUrl: './internaldb.component.html',
  styleUrl: './internaldb.component.css'
})
export class InternaldbComponent {

  buckets: string[] = [];
  selectedBucket: string = "";

  // the data source with all its configs, plus the
  // used table columns
  displayedColumns: string[] = ['key', 'value'];
  dataSource = new MatTableDataSource<BucketKeyValueSet>([]);

  // creating the filtering
  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  // constructor for the table
  // loads the buckets at startup
  constructor(
    private idbService: InternalDBService
  ) {
    // load the buckets
    this.idbService.Buckets_List().subscribe(response => this.buckets = response.buckets)

    // apply the filters
   // this.dataSource.filterPredicate = (data: BucketKeyValueSet, filter: string) => data.key.trim().toLowerCase().indexOf(filter) != -1;
  }

  // this function is required to reload the actual bucket content on 
  // selection change event
  reloadSelection() {
    this.idbService.Buckets_Get(this.selectedBucket).subscribe(
      response => {
        console.log(this.selectedBucket)
        console.log(response)
        this.dataSource.data = response.sink
      }
    )    
  }
}
