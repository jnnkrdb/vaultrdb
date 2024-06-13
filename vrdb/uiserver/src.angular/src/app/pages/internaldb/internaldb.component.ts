import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatFormFieldModule } from '@angular/material/form-field';
import { Bucket, BucketList, InternalDBService } from '../../services/vaultrdb/v1/internaldb/internal-db.service';

@Component({
  selector: 'app-internaldb',
  standalone: true,
  imports: [
    MatFormFieldModule, 
    MatSelectModule, 
    MatInputModule, 
    FormsModule
  ],
  templateUrl: './internaldb.component.html',
  styleUrl: './internaldb.component.css'
})
export class InternaldbComponent {

  buckets: string[] = []
  selectedBucket: string = ""

  constructor(private idbService: InternalDBService) {
    this.idbService.Buckets_List().subscribe(response => this.buckets = response.buckets)
  }
}
