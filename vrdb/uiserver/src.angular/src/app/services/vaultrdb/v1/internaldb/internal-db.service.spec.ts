import { TestBed } from '@angular/core/testing';

import { InternalDBService } from './internal-db.service';

describe('InternalDBService', () => {
  let service: InternalDBService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InternalDBService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
