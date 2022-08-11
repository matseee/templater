import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { By } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { TemplateFacade } from './../../facades/template.facade';
import { MaterialModule } from './../../material.module';
import { Template } from './../../models/template.model';

import { TemplateModalComponent } from './template-modal.component';

class TemplateFacadeMock {
  template: Template;

  constructor() {
    this.template = {
      id: 'DUMMY',
      name: 'DUMMY',
      template: 'DUMMY',
    };
  }

  createTemplate(template: Template) {
    this.template = template;
    this.template.id = 'Test123';
  }

  updateTemplate(template: Template) {
    this.template = template;
  }
}

class MatDialogRefMock {
  isClosed: boolean = false;
  constructor() { }

  reset() {
    this.isClosed = false;
  }

  close() {
    this.isClosed = true;
  }
}

describe('TemplateModalComponent', () => {
  let templateFacade: TemplateFacadeMock;
  let matDialogRef: MatDialogRefMock;
  let component: TemplateModalComponent;
  let fixture: ComponentFixture<TemplateModalComponent>;
  let compiled: any;

  beforeEach(async () => {
    templateFacade = new TemplateFacadeMock();
    matDialogRef = new MatDialogRefMock();

    await TestBed.configureTestingModule({
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      declarations: [TemplateModalComponent],
      providers: [
        { provide: MatDialogRef, useValue: matDialogRef },
        { provide: MAT_DIALOG_DATA, useValue: {} },
        { provide: TemplateFacade, useValue: templateFacade }
      ]
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

  it('should be a new creation', () => {
    expect(component).toBeTruthy();
    expect(component.isEditMode).toBeFalse();
  });

  it('should have the title "New template"', () => {
    fixture.detectChanges();
    expect(component.templateName).toBe('');

    const title = compiled.querySelector('.title');
    expect(title).toBeTruthy();
    expect(title.innerHTML).toBe('New template');
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

  it('should only enable the save button when the formgroup is valid', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#button-save').disabled).toBeTrue();

    component.formGroup.setValue({
      name: 'Test123',
      template: 'Test123'
    });

    fixture.detectChanges();
    expect(compiled.querySelector('#button-save').disabled).toBeFalse();
  });

  it('should create the template after clicking on button save', async () => {
    matDialogRef.reset();

    fixture.detectChanges();

    component.formGroup.patchValue({
      name: 'Test123',
      template: 'Test1234'
    });

    fixture.detectChanges();
    expect(compiled.querySelector('#button-save').disabled).toBeFalse();

    const buttonSave = fixture.debugElement.query(By.css('#button-save'));
    buttonSave.triggerEventHandler('click', null);

    fixture.whenStable().then(() => {
      expect(templateFacade.template.id).toBe('Test123');
      expect(templateFacade.template.name).toBe('Test123');
      expect(templateFacade.template.template).toBe('Test1234');
      expect(matDialogRef.isClosed).toBe(true);
    });
  });

  it('should close the modal after clicking on button cancle', async () => {
    matDialogRef.reset();

    fixture.detectChanges();

    const buttonCancle = fixture.debugElement.query(By.css('#button-cancle'));
    buttonCancle.triggerEventHandler('click', null);

    fixture.whenStable().then(() => {
      expect(matDialogRef.isClosed).toBe(true);
    });
  });
});

describe('TemplateModalComponent-WithData', () => {
  let templateFacade: TemplateFacadeMock;
  let matDialogRef: MatDialogRefMock;
  let component: TemplateModalComponent;
  let fixture: ComponentFixture<TemplateModalComponent>;
  let compiled: any;

  beforeEach(async () => {
    templateFacade = new TemplateFacadeMock();
    matDialogRef = new MatDialogRefMock();

    await TestBed.configureTestingModule({
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      declarations: [TemplateModalComponent],
      providers: [
        { provide: MatDialogRef, useValue: matDialogRef },
        {
          provide: MAT_DIALOG_DATA, useValue: {
            data: {
              id: 'test123',
              name: 'test123',
              template: 'test123',
            }
          }
        },
        { provide: TemplateFacade, useValue: templateFacade }
      ]
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
    fixture.detectChanges();

    expect(component.templateName).toBe('test123');

    const title = compiled.querySelector('.title');
    expect(title).toBeTruthy();
    expect(title.innerHTML).toBe('Template test123');
  });

  it('should have a form-group', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('form.form-group')).toBeTruthy();
  });

  it('should contain the input "name" with the value "test123"', () => {
    fixture.detectChanges();
    const input = compiled.querySelector('#input-name');
    expect(input).toBeTruthy();
    expect(input.value).toBe('test123');
  });

  it('should contain the textarea "template" with the vlaue "test123"', () => {
    fixture.detectChanges();
    const input = compiled.querySelector('#textarea-template');
    expect(input).toBeTruthy();
    expect(input.value).toBe('test123');
  });

  it('should contain the button save', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#button-save')).toBeTruthy();
  });

  it('should contain the button cancle', () => {
    fixture.detectChanges();
    expect(compiled.querySelector('#button-cancle')).toBeTruthy();
  });

  it('should update the template after clicking on button save', async () => {
    matDialogRef.reset();

    fixture.detectChanges();

    component.formGroup.patchValue({
      template: 'edited',
    });

    const buttonSave = fixture.debugElement.query(By.css('#button-save'));
    buttonSave.triggerEventHandler('click', null);

    fixture.whenStable().then(() => {
      expect(templateFacade.template.template).toBe('edited');
      expect(matDialogRef.isClosed).toBe(true);
    });
  });

  it('should close the modal after clicking on button cancle', async () => {
    matDialogRef.reset();

    fixture.detectChanges();

    const buttonCancle = fixture.debugElement.query(By.css('#button-cancle'));
    buttonCancle.triggerEventHandler('click', null);

    fixture.whenStable().then(() => {
      expect(matDialogRef.isClosed).toBe(true);
    });
  });
});

