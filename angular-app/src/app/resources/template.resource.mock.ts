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
        description: 'template1',
        template: 'template1',
        variables: [
            {
                id: 'template1-variable1',
                name: 'template1-variable1',
            },
            {
                id: 'template1-variable2',
                name: 'template1-variable2',
            }
        ]
    },
    {
        id: 'template2',
        description: 'template2',
        template: 'template2',
        variables: [
            {
                id: 'template2-variable1',
                name: 'template2-variable1',
            },
            {
                id: 'template2-variable2',
                name: 'template2-variable2',
            }
        ]
    }
];

@Injectable({ providedIn: 'root' })
export class TemplateResourceMock implements TemplateResourceInterface {
    constructor() { }

    readStatus(): Observable<Status> {
        return from([mockStatus]);
    }

    updateActiveStatus(isActive: boolean): Observable<boolean> {
        mockStatus.isActive = isActive;
        return from([true]);
    }

    readTemplates(): Observable<Template[]> {
        return from([mockTemplates]);
    };

    createTemplate(template: Template): Observable<boolean> {
        mockTemplates.push(template);
        return from([true]);
    }

    updateTemplate(template: Template): Observable<boolean> {
        const index = mockTemplates.findIndex((tmp) => tmp.id === template.id);
        if (index > -1) {
            mockTemplates[index] = template;
            return from([true]);
        }
        return from([false]);
    }

    deleteTemplate(template: Template): Observable<boolean> {
        const index = mockTemplates.findIndex((tmp) => tmp.id === template.id);
        if (index > -1) {
            mockTemplates.splice(index, 1);
            return from([true]);
        }

        return from([false]);
    }
}