import { Variable } from './variable.model';

export interface Template {
    id: string;
    name: string;
    template: string;
    variables: Variable[];
}