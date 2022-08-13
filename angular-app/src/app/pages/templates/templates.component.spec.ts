import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { By } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { TemplateModalComponent, TemplateModalData } from './../../components/template-modal/template-modal.component';
import { MaterialModule } from './../../material.module';
import { TemplateResource } from './../../resources/template.resource';
import { mockTemplates, TemplateResourceMock } from './../../resources/template.resource.mock';

import { TemplatesPage } from './templates.component';

describe('TemplatesPage', () => {
  let component: TemplatesPage;
  let fixture: ComponentFixture<TemplatesPage>;
  let templateResourceMock: TemplateResourceMock;

  beforeEach(async () => {
    templateResourceMock = new TemplateResourceMock();

    await TestBed.configureTestingModule({
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      declarations: [TemplatesPage],
      providers: [
        { provide: TemplateResource, useValue: templateResourceMock },
        MatDialog,
      ]
    })
      .compileComponents();

    fixture = TestBed.createComponent(TemplatesPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should contain a search input', () => {
    const compiled = fixture.debugElement.nativeElement;
    const inputSearch = compiled.querySelector('#input-search');
    expect(inputSearch).toBeTruthy();
  });

  it('should contain an add button', () => {
    const compiled = fixture.debugElement.nativeElement;
    const buttonAdd = compiled.querySelector('#button-add');
    expect(buttonAdd).toBeTruthy();
  });

  it('should contain a table with the columns name, template, variables, control', () => {
    const compiled = fixture.debugElement.nativeElement;
    const columns = compiled.querySelectorAll('.table-column-header');

    expect(columns.length).toBe(3);
  });

  it('should open the template modal when the add button was clicked', () => {
    const spyObject = spyOn<MatDialog, any>(component.myMatDialog, 'open').and.callThrough();

    const buttonAdd = fixture.debugElement.query(By.css('#button-add'));
    buttonAdd.triggerEventHandler('click', null);

    expect(spyObject).toHaveBeenCalled();
  });

  it('should open the template modul when a template was selected in the table', () => {
    const spyObject = spyOn<MatDialog, any>(component.myMatDialog, 'open').and.callThrough();

    const firstTemplateRow = fixture.debugElement.queryAll(By.css('.template-table-row'))[0];
    firstTemplateRow.triggerEventHandler('click', null);

    expect(spyObject).toHaveBeenCalledWith(TemplateModalComponent, {
      data: { data: mockTemplates[0] } as TemplateModalData,
    });
  });
});