import { TestBed } from '@angular/core/testing';

import { StoreDBService } from './store-db.service';

describe('StoreDBService', () => {
  let service: StoreDBService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(StoreDBService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
