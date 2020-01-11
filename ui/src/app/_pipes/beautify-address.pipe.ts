import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'beautifyAddress'
})
export class BeautifyAddressPipe implements PipeTransform {

  transform(address: string): string {
    const addressArray = address.split(',');

    return `${ addressArray[1] } ${ addressArray[0] }, ${ addressArray[2] }`;
  }

}
