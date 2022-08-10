import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './../material.module';
import { Template } from './../models/template.model';
import { TemplateResource } from './../resources/template.resource';
import { TemplateResourceMock } from './../resources/template.resource.mock';
import { TemplateState } from './template.state';

import { timer } from 'rxjs';
import { TemplateFacade } from './template.facade';

describe('TemplateFacade', () => {
  let facade: TemplateFacade;
  let spy: jasmine.SpyObj<TemplateResource>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      providers: [
        { provide: TemplateResource, useClass: TemplateResourceMock }
      ]
    });

    facade = TestBed.inject(TemplateFacade);
  });

  it('should be created with initial state', () => {
    expect(facade).toBeTruthy();
    expect(facade.state$).toBeTruthy();

    const state: TemplateState = facade.getStateSnapshot();
    expect(state).toBeTruthy();
    expect(state.status).not.toBeUndefined();
    expect(state.templates).not.toBeUndefined();
  });

  it('should create a new template', async () => {
    let state: TemplateState = facade.getStateSnapshot();
    let length = state.templates.length;

    const template: Template = {
      id: 'template3',
      description: 'template3',
      template: 'template3',
      variables: []
    };

    facade.createTemplate(template);

    await timer(300).toPromise();
    state = facade.getStateSnapshot();
    expect(state.templates.indexOf(template)).not.toBe(-1);
    expect(state.templates.length).toBe(length + 1);
  });

  it('should edit the first template', async () => {
    let state: TemplateState = facade.getStateSnapshot();

    const firstTemplate = {
      id: state.templates[0].id,
      description: 'template1 edited',
      template: 'template1 edited',
      variables: [],
    };

    facade.updateTemplate(firstTemplate);

    await timer(300).toPromise();
    state = facade.getStateSnapshot();
    expect(state.templates[0].description).toBe('template1 edited');
  });

  it('should delete the first template', async () => {
    let state: TemplateState = facade.getStateSnapshot();
    let templatesLength = state.templates.length;

    const firstTemplate = {
      id: state.templates[0].id,
      description: state.templates[0].description,
      template: state.templates[0].template,
      variables: state.templates[0].variables,
    };

    facade.deleteTemplate(firstTemplate);

    await timer(300).toPromise();
    state = facade.getStateSnapshot();
    expect(state.templates.length).toBe(templatesLength - 1);
    expect(state.templates.indexOf(firstTemplate)).toBe(-1);
  });

  it('should toggle the active status', async () => {
    let previous: TemplateState = facade.getStateSnapshot();
    let previousIsActive = previous.status.isActive;
    expect(previousIsActive).toBeFalse();

    facade.toggleActiveStatus();
    await timer(300).toPromise();

    let after: TemplateState = facade.getStateSnapshot();
    expect(after.status.isActive).toBeTrue();
    expect(previousIsActive).not.toBe(after.status.isActive);

    facade.toggleActiveStatus();
    await timer(300).toPromise();

    after = facade.getStateSnapshot();
    expect(after.status.isActive).toBeFalse();
  });
});
