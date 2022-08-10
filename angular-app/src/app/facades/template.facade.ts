import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Status } from './../models/status.model';
import { Template } from './../models/template.model';
import { TemplateResource } from './../resources/template.resource';
import { TemplateState } from './template.state';

let _state: TemplateState = {
  status: {
    isActive: false,
  },
  templates: [],
};

@Injectable({
  providedIn: 'root'
})
export class TemplateFacade {
  state$: Observable<TemplateState>;
  protected stateSubject: BehaviorSubject<TemplateState>;

  constructor(
    protected myTemplateResource: TemplateResource,
  ) {
    this.stateSubject = new BehaviorSubject(_state);
    this.state$ = this.stateSubject.asObservable();

    this.refresh();
  }

  createTemplate(template: Template) {
    this.myTemplateResource.createTemplate(template)
      .subscribe(() => this.refresh());
  }

  updateTemplate(template: Template) {
    this.myTemplateResource.updateTemplate(template)
      .subscribe(() => this.refresh());
  }

  deleteTemplate(template: Template) {
    this.myTemplateResource.deleteTemplate(template)
      .subscribe(() => this.refresh());
  }

  toggleActiveStatus() {
    this.myTemplateResource.updateActiveStatus(!_state.status.isActive)
      .subscribe(() => this.refresh());
  }

  getStateSnapshot() {
    return _state;
  }

  protected refresh() {
    this.refreshStatus();
    this.refreshTemplates();
  }

  protected refreshStatus() {
    this.myTemplateResource.readStatus()
      .subscribe((status: Status) => this.updateState({ ..._state, status }));
  }

  protected refreshTemplates() {
    this.myTemplateResource.readTemplates()
      .subscribe((templates: Template[]) => this.updateState({ ..._state, templates }));
  }

  protected updateState(state: TemplateState) {
    this.stateSubject.next(_state = state);
  }
}
