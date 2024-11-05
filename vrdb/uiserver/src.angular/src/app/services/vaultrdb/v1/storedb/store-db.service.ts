import { Injectable } from '@angular/core';
import { APIENDPOINT_V1 } from '../endpoint';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { CreateKvsFormComponent } from '../../../../pages/vault/create-kvs-form/create-kvs-form.component';

// ----------------------------------------------
// create the kvs object from the database
export class KeyValueSet {
  id: number = NaN
  key: string = ''
  value: string = ''
  tags: string[] = []
  description: string = ''
  created_at: string = ''
  updated_at: string = ''
}

// the *new* object cache
export class NewKeyValueSet {
  key: string = ''
  value: string = ''
  tags: string[] = []
  description: string = ''
}

// ----------------------------------------------
// setting the default rout for the buckets endpoint
const baseRoute = APIENDPOINT_V1 + "/storedb/kvs"

@Injectable({
  providedIn: 'root'
})
export class StoreDBService {

  constructor(
    private http: HttpClient,
    private dialog: MatDialog
  ) { }

  KVS_List(): Observable<KeyValueSet[]> {
    return this.http.get<KeyValueSet[]>(baseRoute)
      .pipe(map((kvsList: KeyValueSet[]) => {
        console.log('received multiple', kvsList)
        return kvsList
      }))
  }

  KVS_Get(key: string): Observable<KeyValueSet> {
    return this.http.get<KeyValueSet>(baseRoute+"/"+key)
      .pipe(map((kvs: KeyValueSet) => {
        console.log('received', kvs)
        return kvs
      }))
  }

  // : Observable<KeyValueSet>
  KVS_Create(kvs: NewKeyValueSet): Observable<KeyValueSet> {
    return this.http.post<KeyValueSet>(baseRoute, kvs)
      .pipe(map((resp) => {
        console.log('created', resp)
        return resp
      }))
  }

  KVS_Update(kvs: NewKeyValueSet): Observable<KeyValueSet> {
    return this.http.put<KeyValueSet>(baseRoute+"/"+kvs.key, kvs)
      .pipe(map((resp) => {
        console.log('updated', resp)
        return resp
      }))
  }

  KVS_Delete(key: string): Observable<any> {
    return this.http.delete<string>(baseRoute+"/"+key)
      .pipe((resp) => {
        console.log('deleted', resp)
        return resp
      })
  }

  // ----------------------------------------------------------------
  // open the dialog for creating a new kvs
  openDialog() {
    const dialogRef = this.dialog.open(CreateKvsFormComponent)
    return dialogRef.afterClosed()
  }
}
