import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TemplateModalComponent } from './template-modal.component';

describe('TemplateModalComponent', () => {
  let component: TemplateModalComponent;
  let fixture: ComponentFixture<TemplateModalComponent>;
  let compiled: any;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [TemplateModalComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(TemplateModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    compiled = fixture.debugElement.nativeElement;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have a title', () => {
    component.templateName = 'Test123';
    fixture.detectChanges();

    expect(component.templateName).toBe('Test123');

    const title = compiled.querySelector('.title');
    expect(title).toBeTruthy();
    expect(title.innerHTML).toBe('Template Test123');
  });

  it('should have a form-group', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('form.form-group')).toBeTruthy();
  });

  it('should contain the input "name"', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#input-name')).toBeTruthy();
  });

  it('should contain the textarea "template"', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#textarea-template')).toBeTruthy();
  });

  it('should contain the button save', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#button-save')).toBeTruthy();
  });

  it('should contain the button cancle', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#button-cancle')).toBeTruthy();
  });
});
