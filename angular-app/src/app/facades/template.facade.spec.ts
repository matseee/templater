import { TestBed } from '@angular/core/testing';

import { TemplateFacade } from './template.facade';

describe('TemplateService', () => {
  let service: TemplateFacade;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TemplateFacade);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
