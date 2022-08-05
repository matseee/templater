import { Variable } from './variable.model';

export interface Template {
    id: string;
    description: string;
    template: string;
    variables: Variable[];
}