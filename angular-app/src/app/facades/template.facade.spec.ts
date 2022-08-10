import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './../material.module';

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

  it('should be created', () => {
    expect(facade).toBeTruthy();
  });
});
