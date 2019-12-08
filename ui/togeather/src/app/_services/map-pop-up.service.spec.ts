import { TestBed } from '@angular/core/testing';

import { MapPopUpService } from './map-pop-up.service';

describe('MapPopUpService', () => {
  beforeEach(() => TestBed.configureTestingModule({ }));

  it('should be created', () => {
    const service: MapPopUpService = TestBed.get(MapPopUpService);
    expect(service).toBeTruthy();
  });
});
