import { Pipe, PipeTransform } from '@angular/core';
import { Variable } from './../models/variable.model';

@Pipe({
  name: 'variable'
})
export class VariablePipe implements PipeTransform {
  transform(value: Variable[], displayAll: boolean = true): string {
    const strings: string[] = value.map((val: Variable) => val.name);
    let result: string = '';

    if (displayAll) {
      strings.forEach((str: string, index: number) => {
        result += `${str}${index + 1 != strings.length ? ', ' : ''}`;
      });
    } else {
      strings.forEach((str: string, index: number) => {
        if (index < 3) {
          result += `${str}${index == 2 ? '...' : ', '}`;
        }
      });
    }

    return result;
  }
}
