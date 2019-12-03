import { TestBed } from '@angular/core/testing';

import { LobbiesService } from './lobbies.service';

describe('LobbiesService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: LobbiesService = TestBed.get(LobbiesService);
    expect(service).toBeTruthy();
  });
});
