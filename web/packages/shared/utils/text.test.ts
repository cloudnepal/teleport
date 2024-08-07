/**
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

import { pluralize, capitalizeFirstLetter } from './text';

test('pluralize', () => {
  expect(pluralize(0, 'apple')).toBe('apples');
  expect(pluralize(1, 'apple')).toBe('apple');
  expect(pluralize(2, 'apple')).toBe('apples');
});

test('capitalizeFirstLetter', () => {
  expect(capitalizeFirstLetter('hello')).toBe('Hello');
  expect(capitalizeFirstLetter('')).toBe('');
});
