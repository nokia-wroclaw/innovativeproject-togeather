import { TestBed } from '@angular/core/testing';

import { RedirectionService } from './redirection.service';

describe('RedirectionService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: RedirectionService = TestBed.get(RedirectionService);
    expect(service).toBeTruthy();
  });
});
