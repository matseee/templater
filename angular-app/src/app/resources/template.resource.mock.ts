import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { Status } from '../models/status.model';
import { Template } from '../models/template.model';
import { TemplateResourceInterface } from './template.resource.interface';

export let mockStatus: Status = {
    isActive: false,
};

export let mockTemplates: Template[] = [
    {
        id: 'template1',
        name: 'template1',
        template: 'template1'
    },
    {
        id: 'template2',
        name: 'template2',
        template: 'template2'
    }
];

@Injectable({ providedIn: 'root' })
export class TemplateResourceMock implements TemplateResourceInterface {
    myMockTemplates: Template[] = [
        {
            id: 'template1',
            name: 'template1',
            template: 'template1'
        },
        {
            id: 'template2',
            name: 'template2',
            template: 'template2'
        }
    ];

    constructor() { }

    readStatus(): Observable<Status> {
        return from([mockStatus]);
    }

    updateActiveStatus(isActive: boolean): Observable<boolean> {
        mockStatus.isActive = isActive;
        return from([true]);
    }

    readTemplates(): Observable<Template[]> {
        return from([this.myMockTemplates]);
    };

    createTemplate(template: Template): Observable<boolean> {
        this.myMockTemplates.push(template);
        return from([true]);
    }

    updateTemplate(template: Template): Observable<boolean> {
        const index = this.myMockTemplates.findIndex((tmp) => tmp.id === template.id);
        if (index > -1) {
            this.myMockTemplates[index] = template;
            return from([true]);
        }
        return from([false]);
    }

    deleteTemplate(template: Template): Observable<boolean> {
        const index = this.myMockTemplates.findIndex((tmp) => tmp.id === template.id);
        if (index > -1) {
            this.myMockTemplates.splice(index, 1);
            return from([true]);
        }

        return from([false]);
    }
}