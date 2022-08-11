import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './../../material.module';
import { TemplateResource } from './../../resources/template.resource';
import { TemplateResourceMock } from './../../resources/template.resource.mock';

import { TemplatesPage } from './templates.component';

describe('TemplatesPage', () => {
  let component: TemplatesPage;
  let fixture: ComponentFixture<TemplatesPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      declarations: [TemplatesPage],
      providers: [
        { provide: TemplateResource, useClass: TemplateResourceMock }
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

  it('should contain a table with the columns desc., template, variables, control', () => {
    const compiled = fixture.debugElement.nativeElement;
    const columns = compiled.querySelectorAll('.table-column-header');

    expect(columns.length).toBe(4);
    // TODO check each
  });
});
