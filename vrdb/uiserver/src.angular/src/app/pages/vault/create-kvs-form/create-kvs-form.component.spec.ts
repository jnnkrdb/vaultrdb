import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateKvsFormComponent } from './create-kvs-form.component';

describe('CreateKvsFormComponent', () => {
  let component: CreateKvsFormComponent;
  let fixture: ComponentFixture<CreateKvsFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateKvsFormComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CreateKvsFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
