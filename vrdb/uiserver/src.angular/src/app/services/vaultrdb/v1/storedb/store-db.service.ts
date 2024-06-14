import { Injectable } from '@angular/core';
import { APIENDPOINT_V1 } from '../endpoint';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';

// ----------------------------------------------
// create the kvs object from the database
export class KeyValueSet {
  id: number = NaN
  key: string = ""
  value: string = ""
  description: string = ""
  created_at: string = ""
  updated_at: string = ""
}

// the *new* object cache
export class NewKeyValueSet {
  key: string = ""
  value: string = ""
  description: string = ""
}

// ----------------------------------------------
// setting the default rout for the buckets endpoint
const baseRoute = APIENDPOINT_V1 + "/storedb/kvs"

@Injectable({
  providedIn: 'root'
})
export class StoreDBService {

  constructor(private http: HttpClient) { }

  KVS_List(): Observable<KeyValueSet[]> {
    return this.http.get<KeyValueSet[]>(baseRoute)
      .pipe(map((kvsList: KeyValueSet[]) => {
        console.log(kvsList)
        return kvsList
      }))
  }

  KVS_Get(key: string): Observable<KeyValueSet> {
    return this.http.get<KeyValueSet>(baseRoute+"/"+key)
      .pipe(map((kvs: KeyValueSet) => {
        console.log(kvs)
        return kvs
      }))
  }

  // : Observable<KeyValueSet>
  KVS_Create(kvs: NewKeyValueSet) {
    return this.http.post<NewKeyValueSet>(baseRoute, kvs)
      .pipe(map((resp) => {
        console.log(resp)
        return resp
      }))
  }

  KVS_Update(kvs: KeyValueSet): Observable<KeyValueSet> {

    return new Observable<KeyValueSet>
  }

  KVS_Delete(key: string): Observable<KeyValueSet> {

    return new Observable<KeyValueSet>
  }
}
