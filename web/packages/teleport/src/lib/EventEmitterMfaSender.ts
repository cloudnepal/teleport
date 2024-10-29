/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { EventEmitter } from 'events';

import {
  MfaChallengeResponse,
  WebauthnAssertionResponse,
} from 'teleport/services/auth';

class EventEmitterMfaSender extends EventEmitter {
  constructor() {
    super();
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  sendChallengeResponse(data: MfaChallengeResponse) {
    throw new Error('Not implemented');
  }

  // TODO (avatus) DELETE IN 18
  /**
   * @deprecated Use sendChallengeResponse instead.
   */
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  sendWebAuthn(data: WebauthnAssertionResponse) {
    throw new Error('Not implemented');
  }
}

export { EventEmitterMfaSender };