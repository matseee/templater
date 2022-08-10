import { Observable } from 'rxjs';
import { Status } from './../models/status.model';
import { Template } from './../models/template.model';

export interface TemplateResourceInterface {
    // templater application status
    readStatus(): Observable<Status>;

    updateActiveStatus(isActive: boolean): Observable<boolean>;

    // all templates
    readTemplates(): Observable<Template[]>;

    // single template
    createTemplate(template: Template): Observable<boolean>;
    updateTemplate(template: Template): Observable<boolean>;
    deleteTemplate(template: Template): Observable<boolean>;
}