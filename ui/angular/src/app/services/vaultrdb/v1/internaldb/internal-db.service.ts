import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { APIENDPOINT_V1 } from '../endpoint';
import { Observable, map } from 'rxjs';

// ----------------------------------------------
// create the object bucketlist to request the
// current existing buckets in the backend storage
export class BucketList {
  buckets: string[] = []
}

// the actual bucket, containing info sets
export class Bucket {
  sink: BucketKeyValueSet[] = []
}

// Bucket Key Value Set
export class BucketKeyValueSet {
  key: string = ""
  value: string = ""
}

// ----------------------------------------------
// setting the default rout for the buckets endpoint
const baseRoute = APIENDPOINT_V1 + "/internaldb/buckets"

// ----------------------------------------------
// injectable for the internaldb service
@Injectable({
  providedIn: 'root'
})
export class InternalDBService {

  constructor(private http: HttpClient) { }

  Buckets_List(): Observable<BucketList> {
    return this.http.get<BucketList>(baseRoute)
      .pipe(map((bucketList: BucketList) => {
        console.log(bucketList)
        return bucketList
      }))
  }

  Buckets_Get(bucket: string): Observable<Bucket> {
    return this.http.get<Bucket>(baseRoute+"/"+bucket+"/sink")
      .pipe(map((b: Bucket) => {
        console.log(b)
        return b
      }))
  }
}
