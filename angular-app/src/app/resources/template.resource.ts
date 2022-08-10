import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { Status } from './../models/status.model';
import { Template } from './../models/template.model';

export interface TemplateResourceInterface {
    // templater application status
    readStatus(): Observable<Status>;

    // all templates
    readTemplates(): Observable<Template[]>;

    // single template
    createTemplate(template: Template): Observable<boolean>;
    readTemplate(template: Template): Observable<Template>;
    updateTemplate(template: Template): Observable<boolean>;
    deleteTemplate(template: Template): Observable<boolean>;
}

@Injectable({ providedIn: 'root' })
export class TemplateResource implements TemplateResourceInterface {
    constructor(protected myHttpClient: HttpClient) { }

    readStatus(): Observable<Status> {
        const status: Status = {
            isActive: true,
        };

        return from([status]);
    }

    createTemplate(template: Template): Observable<boolean> {
        // TODO
        return from([true]);
    }

    readTemplate(template: Template): Observable<Template> {
        // TODO
        return from([template]);
    }

    updateTemplate(template: Template): Observable<boolean> {
        // TODO
        return from([true]);
    }

    deleteTemplate(template: Template): Observable<boolean> {
        // TODO
        return from([true]);
    }

    readTemplates(): Observable<Template[]> {
        const templates: Template[] = [
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            },
            {
                id: 'template1',
                description: 'template1',
                template: 'template1',
                variables: [
                    {
                        id: 'variable1',
                        name: 'variable1',
                    },
                    {
                        id: 'variable2',
                        name: 'variable2',
                    }
                ]
            }
        ];

        return from([templates]);
    }
}