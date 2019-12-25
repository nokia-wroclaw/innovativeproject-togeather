import { TestBed } from '@angular/core/testing';

import { LobbiesService } from './lobbies.service';
import { HttpClientModule } from '@angular/common/http';

describe('LobbiesService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [ HttpClientModule ]
  }));

  it('should be created', () => {
    const service: LobbiesService = TestBed.get(LobbiesService);
    expect(service).toBeTruthy();
  });
});
