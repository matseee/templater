import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './../material.module';
import { TemplateState } from './template.state';

import { TemplateFacade } from './template.facade';

describe('TemplateFacade', () => {
  let facade: TemplateFacade;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ]
    });
    facade = TestBed.inject(TemplateFacade);
  });

  it('should be created with initial state', () => {
    expect(facade).toBeTruthy();
    expect(facade.state$).toBeTruthy();

    facade.state$
      .subscribe((state: TemplateState) => {
        expect(state).toBeTruthy();
        expect(state.status).not.toBeUndefined();
        expect(state.templates).not.toBeUndefined();
      });
  });


});
