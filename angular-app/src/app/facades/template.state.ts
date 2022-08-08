import { Status } from './../models/status.model';
import { Template } from './../models/template.model';

export interface TemplateState {
    status: Status;
    templates: Template[];
}