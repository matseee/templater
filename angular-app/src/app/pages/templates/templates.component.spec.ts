import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TemplatesPage } from './templates.component';

describe('TemplatesPage', () => {
  let component: TemplatesPage;
  let fixture: ComponentFixture<TemplatesPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [TemplatesPage]
    })
      .compileComponents();

    fixture = TestBed.createComponent(TemplatesPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
