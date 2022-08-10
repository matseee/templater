import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Status } from '../models/status.model';
import { Template } from '../models/template.model';
import { TemplateResourceInterface } from './template.resource.interface';

@Injectable({ providedIn: 'root' })
export class TemplateResource implements TemplateResourceInterface {
    constructor(protected myHttpClient: HttpClient) { }

    readStatus(): Observable<Status> {
        throw new Error('Method not implemented.');
    }

    updateActiveStatus(isActive: boolean): Observable<boolean> {
        throw new Error('Method not implemented.');
    }

    readTemplates(): Observable<Template[]> {
        throw new Error('Method not implemented.');
    }

    createTemplate(template: Template): Observable<boolean> {
        throw new Error('Method not implemented.');
    }

    updateTemplate(template: Template): Observable<boolean> {
        throw new Error('Method not implemented.');
    }

    deleteTemplate(template: Template): Observable<boolean> {
        throw new Error('Method not implemented.');
    }
}