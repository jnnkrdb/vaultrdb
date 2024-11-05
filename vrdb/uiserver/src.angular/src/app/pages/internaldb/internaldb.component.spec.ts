import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InternaldbComponent } from './internaldb.component';

describe('InternaldbComponent', () => {
  let component: InternaldbComponent;
  let fixture: ComponentFixture<InternaldbComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InternaldbComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(InternaldbComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
