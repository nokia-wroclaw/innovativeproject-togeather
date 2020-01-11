import { Pipe, PipeTransform } from '@angular/core';
import * as dateFns from 'date-fns';

@Pipe({
  name: 'beautifyExpirationDate'
})
export class BeautifyExpirationDatePipe implements PipeTransform {

  transform(date: Date): string {
    const dateUTC = new Date(date.toUTCString());

    return dateFns.formatDistanceToNow(dateUTC);
  }

}
