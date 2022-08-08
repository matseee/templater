import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { Status } from './../models/status.model';
import { Template } from './../models/template.model';

@Injectable({ providedIn: 'root' })
export class TemplateResource {
    constructor(protected myHttpClient: HttpClient) { }

    getStatus(): Observable<Status> {
        const status: Status = {
            isActive: true,
        };

        return from([status]);
    }

    getTemplates(): Observable<Template[]> {
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