import { TestBed } from '@angular/core/testing';
import { RedirectionService } from './redirection.service';
import { RouterTestingModule } from '@angular/router/testing';

describe('RedirectionService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [ RouterTestingModule ]
  }));

  it('should be created', () => {
    const service: RedirectionService = TestBed.get(RedirectionService);
    expect(service).toBeTruthy();
  });
});
