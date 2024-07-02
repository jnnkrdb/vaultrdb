import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailedKVSComponent } from './detailed-kvs.component';

describe('DetailedKVSComponent', () => {
  let component: DetailedKVSComponent;
  let fixture: ComponentFixture<DetailedKVSComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DetailedKVSComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DetailedKVSComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
