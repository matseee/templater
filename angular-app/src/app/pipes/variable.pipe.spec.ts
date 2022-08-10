import { Variable } from './../models/variable.model';
import { VariablePipe } from './variable.pipe';

describe('VariablePipe', () => {
  let pipe: VariablePipe;

  const variablesShort: Variable[] = [
    { id: 'var1', name: 'variable1' },
    { id: 'var2', name: 'variable2' },
    { id: 'var3', name: 'variable3' },
  ];

  const resultShort: string = 'variable1, variable2, variable3';

  const variablesLong: Variable[] = [
    { id: 'var1', name: 'variable1' },
    { id: 'var2', name: 'variable2' },
    { id: 'var3', name: 'variable3' },
    { id: 'var4', name: 'variable4' },
    { id: 'var5', name: 'variable5' },
    { id: 'var6', name: 'variable6' },
    { id: 'var7', name: 'variable7' },
    { id: 'var8', name: 'variable8' },
    { id: 'var9', name: 'variable9' },
  ];

  const resultLong: string = 'variable1, variable2, variable3...';

  it('create an instance', () => {
    pipe = new VariablePipe();
    expect(pipe).toBeTruthy();
  });

  it('should transform the short variable array to an string', () => {
    pipe = new VariablePipe();
    const result: string = pipe.transform(variablesShort);
    expect(result).toBe(resultShort);
  });

  it('should transform the long variable arry to an shorted string', () => {
    pipe = new VariablePipe();
    const result: string = pipe.transform(variablesLong, false);
    expect(result).toBe(resultLong);
  });
});
