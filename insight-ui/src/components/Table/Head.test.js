import React from 'react'
import Head from './Head'

xit('should render child elements', () => {
  expect(shallow(<Head>
    <tr>
      <th>Table head</th>
    </tr>
  </Head>).contains(<tr>
    <th>Table head</th>
  </tr>)).toEqual(true)
});